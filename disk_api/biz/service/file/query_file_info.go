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
 * @Date: 2024/10/27 上午12:41
 * @FilePath: query_file_info
 * @Description: 查询文件信息
 */

func ProcessQueryFileInfo(ctx context.Context, req *file_server.QueryFileInfoRequest) (*file_back.QueryFileInfoResponse, error) {
	var rpcReq file_back.QueryFileInfoRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.ProcessQueryFileInfo -> localutils.Converter, err: %v", err)
		return nil, err
	}
	resp, err := rpc.GetDiskBackClient().QueryFileInfo(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.ProcessQueryFileInfo -> rpc.QueryFileInfo, err: %v", err)
		return nil, err
	}
	return resp, nil
}
