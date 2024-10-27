package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cutejiuges/disk_api/biz/model/disk_common"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	back_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 上午1:01
 * @FilePath: ping
 * @Description:
 */

func ProcessPing(ctx context.Context, req *disk_common.PingRequest) (*disk_common.PingResponse, error) {
	resp := disk_common.NewPingResponse()
	var rpcReq back_common.PingRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "ProcessPing convert req error: %v", err)
		return resp, err
	}
	rpcResp, err := rpc.GetDiskBackClient().Ping(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "ProcessPing -> rpc Ping error: %v", err)
		return resp, err
	}

	err = localutils.Converter(rpcResp, resp)
	if err != nil {
		hlog.CtxErrorf(ctx, "ProcessPing convert resp error: %v", err)
		return resp, err
	}
	return resp, nil
}
