package file_service

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/cache"
	"github.com/cutejiuges/disk_back/biz/dao/file_meta_dao"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/infra/mysql"
	"github.com/cutejiuges/disk_back/internal/model/bo"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/model/query"
	"github.com/cutejiuges/disk_back/internal/util/localutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"os"
	"sync"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/28 下午11:33
 * @FilePath: delete_file
 * @Description: 删除文件
 */

func ProcessDeleteFile(ctx context.Context, req *disk_common.DeleteFileRequest) (*disk_back.DeleteFileData, error) {
	data := disk_back.NewDeleteFileData()

	//异步删除文件
	wg := sync.WaitGroup{}
	successMap, failedMap := localutil.NewSafeMap[int64, *disk_back.OperateFileRes](), localutil.NewSafeMap[int64, *disk_back.OperateFileRes]()
	wg.Add(len(req.GetId()))

	for _, id := range req.GetId() {
		go func(fileId int64) {
			defer wg.Done()
			//获取删除文件锁，如果能获取到，继续删除，否则说明正在删除中
			fileLock := cache.NewDistributedLock(cache.DeleteFileLockKey(fileId))
			if ok := fileLock.TryLock(); !ok {
				//文件正在删除
				failedMap.Put(fileId, &disk_back.OperateFileRes{
					Id:  fileId,
					Msg: "文件正在删除中",
				})
				return
			}
			//释放文件锁
			defer fileLock.Unlock()

			//判断文件是否存在
			simpleFile, ok := isExitsFileById(ctx, fileId)
			if !ok {
				//文件不存在
				failedMap.Put(fileId, &disk_back.OperateFileRes{
					Id:  fileId,
					Msg: simpleFile.Msg,
				})
				return
			}

			//删除文件
			simpleFile, err := deleteFile(ctx, fileId, simpleFile.Addr)
			if err != nil {
				failedMap.Put(fileId, &disk_back.OperateFileRes{
					Id:  fileId,
					Msg: simpleFile.Msg + ", " + err.Error(),
				})
			}

			//删除成功
			successMap.Put(fileId, &disk_back.OperateFileRes{
				Id:  fileId,
				Msg: simpleFile.Msg,
			})
		}(id)
	}
	wg.Wait()
	data.SetSuccessMap(successMap.GetData())
	data.SetFailedMap(failedMap.GetData())
	status := judgeOperateStatus(successMap, failedMap)
	data.SetStatus(thrift.Int8Ptr(status))
	data.SetStatusName(thrift.StringPtr(enum.OperateFileStatusMap[status]))
	return data, nil
}

// 文件是否存在
func isExitsFileById(ctx context.Context, id int64) (bo.SimpleFile, bool) {
	simpleFile := bo.SimpleFile{ID: id}
	file, err := file_meta_dao.QuerySingleFileMeta(ctx, &param.QueryFileMetaParam{ID: id})
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.isExitsFile -> file_meta_dao.QuerySingleFileMeta error: %v", err)
		simpleFile.Msg = "文件不存在"
		return simpleFile, false
	}
	//存在的话返回相关信息
	simpleFile.Addr = file.FileAddr
	simpleFile.Name = file.FileName
	return simpleFile, true
}

func deleteFile(ctx context.Context, id int64, path string) (bo.SimpleFile, error) {
	simpleFile := bo.SimpleFile{
		ID:  id,
		Msg: "文件删除失败",
	}
	//事务执行文件删除，先删除db数据，再删除实际存储
	qry := query.Use(mysql.DB()).Begin()
	//如果异常中断了，也需要回滚
	defer func() {
		if e := recover(); e != nil {
			klog.CtxErrorf(ctx, "file_service deleteFile paniced: %v", e)
			_ = qry.Rollback()
		}
	}()
	err := file_meta_dao.DeleteFile(ctx, qry.Query, &param.EditFileMetaParam{ID: id})
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.deleteFile -> file_meta_dao.DeleteFile error: %v", err)
		_ = qry.Rollback()
		return simpleFile, err
	}
	err = os.Remove(path)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.deleteFile -> os.Remove error: %v", err)
		_ = qry.Rollback()
		return simpleFile, err
	}
	_ = qry.Commit()
	simpleFile.Msg = "文件删除成功"
	return simpleFile, nil
}
