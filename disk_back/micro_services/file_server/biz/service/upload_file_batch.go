package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/cache"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/enum"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/pojo/bo"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/pojo/param"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/util"
	"os"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午11:09
 * @FilePath: upload_file_single
 * @Description: 实现文件批量上传
 */

func ProcessUploadFileBatch(ctx context.Context, req *file_server.UploadFileRequest) (*file_server.UploadFileData, error) {
	data := file_server.NewUploadFileData()
	// 先查询文件是否已经上传过，抽取文件的sha256进行校验
	return data, nil
}

// 存储文件实体
func writeFileEntity(path string, content []byte) error {
	//先判断文件是否存在，如果存在不用写入，如果不存在执行写入
	_, err := os.Stat(path)
	if !os.IsNotExist(err) { //已经存在
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
func fileExistInDB(ctx context.Context, fileKey string) (bo.SimpleFile, bool) {
	simpleFile := bo.SimpleFile{}
	f, err := dao.QuerySingleFileMeta(ctx, &param.QueryFileMetaParam{
		FileKey: fileKey,
	})
	if err != nil {
		klog.CtxErrorf(ctx, "service.fileExistInDB error: %v", err)
		return simpleFile, false
	}
	simpleFile.ID = f.ID
	simpleFile.Msg = "文件已上传过"
	return simpleFile, true
}

// 执行文件保存动作
func saveFileInfo(ctx context.Context, path string, content []byte) (bo.SimpleFile, error) {
	simpleFile := bo.SimpleFile{ID: 0, Msg: "文件上传失败"}
	//对未上传的文件执行文件写入
	id, err := cache.GetNextUniqID(ctx)
	if err != nil {
		klog.CtxErrorf(ctx, "service.saveFileInfo -> cache.GetNextUniqID error: %v", err)
		return simpleFile, err
	}
	//先写入mysql，再进行本地存储
	uniqKey := util.GetSha256Key(content)
	fileMeta := &model.FileMeta{
		ID:       id,
		FileKey:  uniqKey,
		FileName: uniqKey,
		FileAddr: path,
		FileSize: int64(len(content)),
		Status:   enum.FileMetaStatusEnable,
	}
	klog.CtxInfof(ctx, "service.saveFileInfo create fileMeta: %v", fileMeta)
	return simpleFile, nil
}
