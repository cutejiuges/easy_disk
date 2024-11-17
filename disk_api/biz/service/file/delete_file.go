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
 * @Date: 2024/11/12 上午12:13
 * @FilePath: delete_file
 * @Description:
 */

func ProcessDeleteFile(ctx context.Context, req *disk_common.DeleteFileRequest) (*file_server.DeleteFileResponse, error) {
	var rpcReq disk_back_common.DeleteFileRequest
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
