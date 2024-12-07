package user_service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cutejiuges/disk_api/biz/model/user_server"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	user_back "github.com/cutejiuges/disk_back/kitex_gen/user_server"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 下午12:34
 * @FilePath: user_sign_up
 * @Description: 用户注册服务
 */

func ProcessUserSignUp(ctx context.Context, req *user_server.UserSignUpRequest) (*user_back.UserSignUpResponse, error) {
	var rpcReq user_back.UserSignUpRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "user_service.ProcessUserSignUp -> localutils.Converter error: %v", err)
		return nil, err
	}

	resp, err := rpc.GetUserServerClient().UserSignUp(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "user_service.ProcessUserSignUp -> rpc.UserSignUp error: %v", err)
		return nil, err
	}
	return resp, nil
}
