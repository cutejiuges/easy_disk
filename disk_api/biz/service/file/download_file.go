package file_service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cutejiuges/disk_api/biz/model/disk_common"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	disk_back_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 下午3:56
 * @FilePath: download_file
 * @Description: 下载文件
 */

func ProcessDownloadFile(ctx context.Context, request *disk_common.DownloadFileRequest) (*file_server.DownloadFileResponse, error) {
	var rpcReq disk_back_common.DownloadFileRequest
	err := localutils.Converter(request, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "converter req error: %v", err)
		return nil, err
	}
	resp, err := rpc.GetDiskBackClient().DownloadFile(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.ProcessDownloadFile -> rpc DownloadFile error: %v", err)
		return nil, err
	}
	return resp, nil
}
