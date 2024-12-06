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
 * @Date: 2024/12/6 上午12:52
 * @FilePath: get_email_verify_code
 * @Description: 获取邮箱验证码
 */

func ProcessGetEmailVerifyCode(ctx context.Context, req *user_server.GetEmailVerifyCodeRequest) (*user_back.GetEmailVerifyCodeResponse, error) {
	var rpcReq user_back.GetEmailVerifyCodeRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "user.ProcessGetEmailVerifyCode -> localutils.Converter error: %v", err)
		return nil, err
	}
	userCli := rpc.GetUserServerClient()
	resp, err := userCli.GetEmailVerifyCode(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "user.ProcessGetEmailVerifyCode -> rpc.GetEmailVerifyCode error: %v", err)
		return nil, err
	}
	return resp, nil
}
