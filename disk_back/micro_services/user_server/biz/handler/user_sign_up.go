package handler

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/service"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午8:51
 * @FilePath: user_sign_up
 * @Description: 用户注册功能
 */

type UserSignUpHandler struct {
	ctx  context.Context
	req  *user_server.UserSignUpRequest
	resp *user_server.UserSignUpResponse
	data *user_server.UserSignUpData
	err  error
}

func NewUserSignUpHandler(ctx context.Context, req *user_server.UserSignUpRequest) *UserSignUpHandler {
	return &UserSignUpHandler{
		ctx:  ctx,
		req:  req,
		resp: user_server.NewUserSignUpResponse(),
	}
}

func (h *UserSignUpHandler) checkParam() error {
	if h.req.GetVerifyCode() <= 0 {
		return &errno.BizError{
			Code: errno.ErrCodeParamIllegal,
			Msg:  errno.ErrMsgMap[errno.ErrCodeParamIllegal] + ", 未输入正确的验证码",
		}
	}
	if h.req.GetUserInfo() == nil {
		return &errno.BizError{
			Code: errno.ErrCodeParamIllegal,
			Msg:  errno.ErrMsgMap[errno.ErrCodeParamIllegal] + ", 用户身份信息未完善",
		}
	}
	userInfo := h.req.GetUserInfo()
	if !util.CheckEmail(userInfo.GetEmail()) {
		return &errno.BizError{
			Code: errno.ErrCodeParamIllegal,
			Msg:  "邮箱地址不合法",
		}
	}
	if userInfo.GetUserName() == "" {
		userInfo.SetUserName(thrift.StringPtr(fmt.Sprintf("游客%d", util.GenUserSuffix())))
	}
	return nil
}

func (h *UserSignUpHandler) checkPermission() error {
	return nil
}

func (h *UserSignUpHandler) processBusiness() error {
	data, err := service.ProcessUserSignUp(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *UserSignUpHandler) packResp() {
	h.resp.SetBaseResp(errno.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "GetEmailVerifyCodeHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *UserSignUpHandler) Handle() (*user_server.UserSignUpResponse, error) {
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
