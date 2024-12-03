package file_service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	file_server "github.com/cutejiuges/disk_api/biz/model/file_server"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	file_back "github.com/cutejiuges/disk_back/kitex_gen/file_server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/12 上午12:13
 * @FilePath: delete_file
 * @Description:
 */

func ProcessDeleteFile(ctx context.Context, req *file_server.DeleteFileRequest) (*file_back.DeleteFileResponse, error) {
	var rpcReq file_back.DeleteFileRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "converter req error: %v", err)
		return nil, err
	}
	resp, err := rpc.GetDiskBackClient().DeleteFile(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.DeleteFile -> rpc.DeleteFile error: %v", err)
		return nil, err
	}
	return resp, nil
}
