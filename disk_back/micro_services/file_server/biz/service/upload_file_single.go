package service

import (
	"context"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午11:09
 * @FilePath: upload_file_single
 * @Description: 实现单文件上传
 */

func ProcessUploadFileSingle(ctx context.Context, req *file_server.UploadFileRequest) (*file_server.UploadFileData, error) {
	data := file_server.NewUploadFileData()

	return data, nil
}

// 存储文件实体
func writeSingleFile(path string) {

}
