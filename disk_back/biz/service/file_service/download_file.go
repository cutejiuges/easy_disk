package file_service

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/dao/file_meta_dao"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"os"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 下午3:32
 * @FilePath: down_file
 * @Description: 文件下载接口
 */

func ProcessDownloadFile(ctx context.Context, req *disk_common.DownloadFileRequest) (*disk_back.DownloadFileData, error) {
	data := disk_back.NewDownloadFileData()
	//通过文件id查询文件信息
	fileMeta, err := file_meta_dao.QuerySingleFileMeta(ctx, &param.QueryFileMetaParam{
		ID: req.GetId(),
	})
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessDownloadFile -> file_meta_dao.QuerySingleFileMeta error: %v", err)
		return data, err
	}
	content, err := loadFileContent(fileMeta.FileAddr)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.loadFileContent error: %v", err)
		return data, err
	}
	data.FileName = thrift.StringPtr(fileMeta.FileName)
	data.Id = thrift.Int64Ptr(fileMeta.ID)
	data.Content = content
	return data, nil
}

func loadFileContent(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
