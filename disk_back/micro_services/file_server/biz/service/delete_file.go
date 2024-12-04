package service

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/enum"
	errno2 "github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/mysql"
	"github.com/cutejiuges/disk_back/micro_services/file_server/pojo/param"
	"os"
	"sync"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/27 下午10:55
 * @FilePath: delete_file
 * @Description: 执行文件删除操作
 */

func ProcessDeleteFile(ctx context.Context, req *file_server.DeleteFileRequest) (*file_server.DeleteFileData, error) {
	data := file_server.NewDeleteFileData()
	//1. 根据id查询所有待删除的文件信息
	fileList, err := dao.QueryFileMetaList(ctx, &param.QueryFileMetaParam{IdList: req.GetId()})
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessDeleteFile -> dao.QueryFileMetaList error: %v", err)
		data.SetStatus(thrift.Int8Ptr(enum.OperateFileStatusFailed))
		data.SetStatusName(thrift.StringPtr(enum.OperateFileStatusMap[enum.OperateFileStatusFailed]))
		return data, &errno2.BizError{Code: errno2.ErrCodeDbUnknownError, Msg: err.Error()}
	}
	//2. 根据文件的引用次数，判断本次操作是删除文件记录还是减少引用次数
	dropFileIds := make([]int64, 0)
	dropFilePaths := make([]string, 0)
	editFileIds := make([]int64, 0)
	for _, file := range fileList {
		if file.RefNum > 1 {
			editFileIds = append(editFileIds, file.ID)
		} else {
			dropFileIds = append(dropFileIds, file.ID)
			dropFilePaths = append(dropFilePaths, file.FileAddr)
		}
	}

	success, failed := util.NewSafeMap[int64, *file_server.OperateFileRes](), util.NewSafeMap[int64, *file_server.OperateFileRes]()
	//3. 处理引用次数-1
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := decreaseFileRefNum(ctx, editFileIds); err != nil {
			for _, id := range editFileIds {
				failed.Put(id, &file_server.OperateFileRes{Id: id, Msg: "文件删除失败, " + err.Error()})
			}
		} else {
			for _, id := range editFileIds {
				success.Put(id, &file_server.OperateFileRes{Id: id, Msg: "文件删除成功"})
			}
		}
	}()

	//4. 处理文件删除
	for _, path := range dropFilePaths {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			err := os.Remove(s)
			klog.CtxWarnf(ctx, "remove file error: %v", err)
		}(path)
	}
	qry := query.Use(mysql.DB())
	if err := dao.DeleteFile(ctx, qry, &param.EditFileMetaParam{IdList: dropFileIds}); err != nil {
		for _, id := range dropFileIds {
			failed.Put(id, &file_server.OperateFileRes{Id: id, Msg: "文件删除失败, " + err.Error()})
		}
	} else {
		for _, id := range dropFileIds {
			success.Put(id, &file_server.OperateFileRes{Id: id, Msg: "文件删除成功"})
		}
	}
	wg.Wait()

	status := judgeOperateStatus(success, failed)
	data.SetStatus(thrift.Int8Ptr(status))
	data.SetStatusName(thrift.StringPtr(enum.OperateFileStatusMap[status]))
	data.SetSuccessMap(success.GetData())
	data.SetFailedMap(failed.GetData())
	return data, nil
}

func decreaseFileRefNum(ctx context.Context, fileIds []int64) error {
	qry := query.Use(mysql.DB())
	return dao.ModifyFileRef(ctx, qry, &param.EditFileMetaParam{
		IdList:   fileIds,
		RefDealt: int64(-1),
	})
}
