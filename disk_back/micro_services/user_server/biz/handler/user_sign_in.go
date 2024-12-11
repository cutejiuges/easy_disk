package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/service"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/10 上午8:36
 * @FilePath: user_sign_in
 * @Description: 用户登录接口
 */

type UserSignInHandler struct {
	ctx  context.Context
	req  *user_server.UserSignInRequest
	resp *user_server.UserSignInResponse
	data *user_server.UserSignInData
	err  error
}

func NewUserSignInHandler(ctx context.Context, req *user_server.UserSignInRequest) *UserSignInHandler {
	return &UserSignInHandler{
		ctx:  ctx,
		req:  req,
		resp: user_server.NewUserSignInResponse(),
	}
}

func (h *UserSignInHandler) checkParam() error {
	err := &errno.BizError{Code: errno.ErrCodeParamIllegal, Msg: errno.ErrMsgMap[errno.ErrCodeInternal]}
	if len(h.req.GetPassword()) <= 0 {
		err.Msg += ", 未输入有效密码"
		return err
	}
	if len(h.req.GetEmail()) != 0 && !util.CheckEmail(h.req.GetEmail()) {
		err.Msg += ", 为输入有效的邮箱"
		return err
	}
	if len(h.req.GetEmail()) <= 0 && h.req.GetAccountId() <= 0 {
		err.Msg += ", 请使用email或账户id进行登陆"
		return err
	}
	return nil
}

func (h *UserSignInHandler) checkPermission() error {
	return nil
}

func (h *UserSignInHandler) processBusiness() error {
	data, err := service.ProcessUserSignIn(h.ctx, h.req)
	if err != nil {
		klog.CtxErrorf(h.ctx, "UserSignInHandler.processBusiness -> service.ProcessUserSignIn error: %v", err)
		return err
	}
	h.data = data
	return nil
}

func (h *UserSignInHandler) packResp() {
	h.resp.SetBaseResp(errno.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "GetEmailVerifyCodeHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *UserSignInHandler) Handle() (*user_server.UserSignInResponse, error) {
	defer func() {
		h.packResp()
	}()
	if h.err = h.checkParam(); h.err != nil {
		goto end
	}
	if h.err = h.checkPermission(); h.err != nil {
		goto end
	}
	if h.err = h.processBusiness(); h.err != nil {
		goto end
	}
end:
	return h.resp, nil
}
