package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/service"
	"regexp"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/5 上午12:32
 * @FilePath: get_email_verify_code
 * @Description: 通过邮箱发送验证码
 */

type GetEmailVerifyCodeHandler struct {
	ctx  context.Context
	req  *user_server.GetEmailVerifyCodeRequest
	resp *user_server.GetEmailVerifyCodeResponse
	err  error
}

func NewGetEmailVerifyCodeHandler(ctx context.Context, req *user_server.GetEmailVerifyCodeRequest) *GetEmailVerifyCodeHandler {
	return &GetEmailVerifyCodeHandler{
		ctx:  ctx,
		req:  req,
		resp: user_server.NewGetEmailVerifyCodeResponse(),
	}
}

func (h *GetEmailVerifyCodeHandler) checkParam() error {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(h.req.GetEmail()) {
		return &errno.BizError{
			Code: errno.ErrCodeParamIllegal,
			Msg:  "邮箱地址不合法",
		}
	}
	return nil
}

func (h *GetEmailVerifyCodeHandler) checkPermission() error {
	return nil
}

func (h *GetEmailVerifyCodeHandler) processBusiness() error {
	return service.ProcessGetEmailVerifyCode(h.ctx, h.req)
}

func (h *GetEmailVerifyCodeHandler) packResp() {
	h.resp.SetBaseResp(errno.NewBaseRespWithOK())
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "GetEmailVerifyCodeHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *GetEmailVerifyCodeHandler) Handle() (*user_server.GetEmailVerifyCodeResponse, error) {
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
