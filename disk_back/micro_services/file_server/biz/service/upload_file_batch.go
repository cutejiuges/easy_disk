package service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/cache"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/mysql"
	"github.com/cutejiuges/disk_back/micro_services/file_server/pojo"
	"os"
	"sync"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午11:09
 * @FilePath: upload_file_single
 * @Description: 实现文件批量上传
 */

func ProcessUploadFileBatch(ctx context.Context, req *file_server.UploadFileRequest) (*file_server.UploadFileData, error) {
	data := file_server.NewUploadFileData()
	//处理文件上传，使用文件总数的1/5个协程执行文件上传，但最多只开10个协程
	wg := sync.WaitGroup{}
	success, failed := util.NewSafeMap[string, *file_server.OperateFileRes](), util.NewSafeMap[string, *file_server.OperateFileRes]()
	processNum := min((len(req.GetFiles())+enum.NumberOfGoroutineProcessed-1)/enum.NumberOfGoroutineProcessed, enum.MaxNumberOfGoroutines)
	groupLength := len(req.GetFiles()) / processNum
	for i := 0; i < processNum; i++ {
		wg.Add(1)
		go func(processId int) {
			defer wg.Done()
			//计算当前协程要处理的分片起止位置
			startIdx, endIdx := processId*groupLength, (processId+1)*groupLength
			if processId == processNum-1 {
				endIdx = len(req.GetFiles())
			}

			//处理文件
			for j := startIdx; j < endIdx; j++ {
				func(idx int) {
					f := req.GetFiles()[idx]
					fileKey := f.GetFileKey()
					// 1. 获取全局锁，取锁成功进入上传流程，取锁失败说明在上传中
					fileLock := cache.NewDistributedLock(cache.UploadFileLockKey(fileKey))
					if ok := fileLock.TryLock(); !ok {
						//文件正在上传
						failed.Put(fileKey, &file_server.OperateFileRes{Id: 0, FileKey: fileKey, Msg: "文件正在上传中"})
						return
					}
					//记得释放锁
					defer fileLock.Unlock()

					// 2. 非正在上传，校验是否已经上传过
					if simpleFile, exits := fileExistInDB(ctx, fileKey); exits {
						//已经上传过，增加文件引用数量，直接返回历史信息
						err := increaseFileRefNum(ctx, []int64{simpleFile.ID})
						if err != nil {
							failed.Put(fileKey, &file_server.OperateFileRes{Id: simpleFile.ID, FileKey: fileKey, Msg: err.Error()})
							return
						}
						success.Put(fileKey, &file_server.OperateFileRes{Id: simpleFile.ID, FileKey: simpleFile.Key, Msg: simpleFile.Msg})
						return
					}

					// 3. 没有上传过，执行上传动作
					path := fmt.Sprintf("%s%s", enum.LocationOfUploadFiles, fileKey)
					simpleFile, err := saveFileInfo(ctx, path, f.GetFileData())
					if err != nil {
						failed.Put(fileKey, &file_server.OperateFileRes{Id: simpleFile.ID, FileKey: fileKey, Msg: simpleFile.Msg + err.Error()})
						return
					}
					success.Put(fileKey, &file_server.OperateFileRes{Id: simpleFile.ID, FileKey: simpleFile.Key, Msg: simpleFile.Msg})
					return
				}(j)
			}
		}(i)
	}
	wg.Wait()
	status := judgeOperateStatus(success, failed)
	data.Status = thrift.Int8Ptr(status)
	data.StatusName = thrift.StringPtr(enum.OperateFileStatusMap[status])
	data.SetSuccessRes(success.GetData())
	data.SetFailedRes(failed.GetData())
	return data, nil
}

// 存储文件实体
func writeFileEntity(path string, content []byte) error {
	//先判断文件是否存在，如果存在不用写入，如果不存在执行写入
	_, err := os.Stat(path)
	if !os.IsNotExist(err) { //已经存在
		klog.Warnf("service.writeFileEntity warning: %v", err)
		return nil
	}
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		klog.Errorf("service.writeFileEntity error: %v", err)
		return err
	}
	return nil
}

// 根据key在mysql中查询文件是否上传过
func fileExistInDB(ctx context.Context, fileKey string) (pojo.SimpleFile, bool) {
	simpleFile := pojo.SimpleFile{}
	f, err := dao.QuerySingleFileMeta(ctx, &pojo.QueryFileMetaParam{
		FileKey: fileKey,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "service.fileExistInDB error: %v", err)
		return simpleFile, false
	}
	simpleFile.ID = f.ID
	simpleFile.Msg = "文件已上传过"
	simpleFile.Key = f.FileKey
	return simpleFile, true
}

// 执行文件保存动作
func saveFileInfo(ctx context.Context, path string, content []byte) (simpleFile pojo.SimpleFile, err error) {
	simpleFile = pojo.SimpleFile{ID: 0, Msg: "文件上传失败"}
	//对未上传的文件执行文件写入
	id, err := cache.GetNextUniqID(ctx)
	if err != nil {
		return simpleFile, err
	}
	//构建写入sql的fileMeta
	uniqKey := util.GetSha256Key(content)
	fileMeta := &model.FileMeta{
		ID:       id,
		FileKey:  uniqKey,
		FileAddr: path,
		FileSize: int64(len(content)),
		Status:   enum.FileMetaStatusEnable,
	}
	klog.CtxInfof(ctx, "service.saveFileInfo create fileMeta: %v", fileMeta)
	//写数据库，开启事务
	qry := query.Use(mysql.DB()).Begin()
	defer func() {
		if e := recover(); e != nil || err != nil {
			_ = qry.Rollback()
		}
	}()
	//先写入mysql，如果写数据库失败，回滚后直接返回
	if err = dao.CreateFile(ctx, qry.Query, fileMeta); err != nil {
		klog.CtxErrorf(ctx, "service.saveFileInfo -> dao.CreateFile error: %v", err)
		return simpleFile, err
	}
	//再写入本地存储，如果写文件失败，也可以把数据库信息回滚掉，确保信息是干净的
	if err = writeFileEntity(path, content); err != nil {
		klog.CtxErrorf(ctx, "service.saveFileInfo -> service.writeFileEntity error: %v", err)
		return simpleFile, err
	}
	_ = qry.Commit()
	simpleFile.ID = id
	simpleFile.Msg = "文件上传成功"
	simpleFile.Key = uniqKey
	return simpleFile, nil
}

func increaseFileRefNum(ctx context.Context, fileIds []int64) error {
	qry := query.Use(mysql.DB())
	return dao.ModifyFileRef(ctx, qry, &pojo.EditFileMetaParam{
		IdList:   fileIds,
		RefDealt: int64(1),
	})
}

func judgeOperateStatus[K comparable, V any](success, failed *util.SafeMap[K, V]) int8 {
	res := enum.OperateFileStatusPartiallySuccessful
	if success.Len() == 0 {
		res = enum.OperateFileStatusFailed
	}
	if failed.Len() == 0 {
		res = enum.OperateFileStatusSuccess
	}
	return res
}
