package file_service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/cache"
	"github.com/cutejiuges/disk_back/biz/dao/file_meta_dao"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/infra/mysql"
	"github.com/cutejiuges/disk_back/internal/model/bo"
	"github.com/cutejiuges/disk_back/internal/model/model"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/model/query"
	"github.com/cutejiuges/disk_back/internal/util/localutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"os"
	"sync"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 上午1:53
 * @FilePath: upload_file
 * @Description: 执行文件上传
 */

const prefix = "uploads/"

func ProcessUploadFile(ctx context.Context, req *disk_back.UploadFileRequest) (*disk_back.UploadFileData, error) {
	data := disk_back.NewUploadFileData()

	//异步写文件
	wg := sync.WaitGroup{}
	successMap, failedMap := localutil.NewSafeMap[string, *disk_back.OperateFileRes](), localutil.NewSafeMap[string, *disk_back.OperateFileRes]()
	wg.Add(len(req.GetFiles()))
	for _, file := range req.GetFiles() {
		go func(f *disk_back.UploadFileMeta) {
			defer wg.Done()
			//检查文件是否正在上传，获取文件锁失败说明在上传中
			sha1Key := localutil.GetSha1Key(f.GetFileData())
			fileLock := cache.NewDistributedLock(cache.UploadFileLockKey(sha1Key))
			if ok := fileLock.TryLock(); !ok {
				//文件正在上传中
				failedMap.Put(f.GetFileName(), &disk_back.OperateFileRes{
					Id:       0,
					FileName: f.GetFileName(),
					Msg:      "文件正在上传中",
				})
				return
			}
			//延迟释放文件锁
			defer fileLock.Unlock()

			//无正在上传中的文件，再校验是否已上传
			if simpleFile, exits := isExistFileByKey(ctx, sha1Key); exits {
				//如果已经上传过了，用已存在的文件信息返回
				successMap.Put(f.GetFileName(), &disk_back.OperateFileRes{
					Id:       simpleFile.ID,
					FileName: simpleFile.Name,
					Msg:      simpleFile.Msg,
				})
				return
			}

			//在mysql和实际的存储中写入
			path := fmt.Sprintf("%s%s", prefix, f.GetFileName())
			simpleFile, err := saveFileInfo(ctx, path, f.GetFileName(), f.GetFileData())
			if err != nil {
				failedMap.Put(f.GetFileName(), &disk_back.OperateFileRes{
					Id:       simpleFile.ID,
					FileName: simpleFile.Name,
					Msg:      simpleFile.Msg + err.Error(),
				})
				return
			}

			//写成功
			successMap.Put(f.GetFileName(), &disk_back.OperateFileRes{
				Id:       simpleFile.ID,
				FileName: simpleFile.Name,
				Msg:      simpleFile.Msg,
			})
			return
		}(file)
	}
	wg.Wait()

	data.SetFailedRes(failedMap.GetData())
	data.SetSuccessRes(successMap.GetData())
	status := judgeOperateStatus(successMap, failedMap)
	data.SetStatus(thrift.Int8Ptr(status))
	data.SetStatusName(thrift.StringPtr(enum.OperateFileStatusMap[status]))
	return data, nil
}

func writeFile(ctx context.Context, path, fileName string, fileData []byte) error {
	err := os.WriteFile(path, fileData, 0644)
	if err != nil {
		klog.CtxErrorf(ctx, "WriteFile error: %v", err)
		return err
	}
	return nil
}

func isExistFileByKey(ctx context.Context, sha1Key string) (bo.SimpleFile, bool) {
	simpleFile := bo.SimpleFile{}
	// 查询文件是否上传过
	f, err := file_meta_dao.QuerySingleFileMeta(ctx, &param.QueryFileMetaParam{
		FileKey: sha1Key,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.saveFileInfo -> QueryFileByKey error: %v", err)
		return simpleFile, false
	}
	//如果已经上传过，返回文件已上传的信息
	simpleFile.Name = f.FileName
	simpleFile.ID = f.ID
	simpleFile.Msg = "文件已经上传过"
	return simpleFile, true
}

// 保存文件信息
func saveFileInfo(ctx context.Context, path, fileName string, fileData []byte) (bo.SimpleFile, error) {
	simpleFile := bo.SimpleFile{
		ID:   0,
		Name: fileName,
		Msg:  "文件上传失败",
	}
	//未上传文件，执行文件写入
	id, err := cache.GetNextUniqID(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.saveFileInfo -> GetNextUniqID failed, error: %v", err)
		return simpleFile, err
	}

	//先写入mysql，再写本地文件
	fileMeta := &model.FileMeta{
		ID:       id,
		FileKey:  localutil.GetSha1Key(fileData),
		FileName: fileName,
		FileAddr: path,
		FileSize: int64(len(fileData)),
		Status:   enum.FileMetaStatusEnable,
	}
	klog.CtxInfof(ctx, "file_service.saveFileInfo create FileMeta: %v", fileMeta)
	qry := query.Use(mysql.DB()).Begin()
	defer func() {
		if e := recover(); e != nil {
			klog.CtxErrorf(ctx, "file_service uploadFile paniced: %v", e)
			_ = qry.Rollback()
		}
	}()
	err = file_meta_dao.CreateFile(ctx, qry.Query, fileMeta)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.saveFileInfo -> CreateFile failed, error: %v", err)
		//写失败回滚
		_ = qry.Rollback()
		return simpleFile, err
	}
	err = writeFile(ctx, path, fileName, fileData)
	if err != nil {
		_ = qry.Rollback()
		return simpleFile, err
	}
	_ = qry.Commit()

	//mysql和文件存储都写成功了
	simpleFile.Name = fileName
	simpleFile.ID = id
	simpleFile.Msg = "文件上传成功"
	return simpleFile, nil
}

func judgeOperateStatus[K comparable, V any](success, failed *localutil.SafeMap[K, V]) int8 {
	res := enum.OperateFileStatusPartiallySuccessful
	if success.Len() == 0 {
		res = enum.OperateFileStatusFailed
	}
	if failed.Len() == 0 {
		res = enum.OperateFileStatusSuccess
	}
	return res
}
