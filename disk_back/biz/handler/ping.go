package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/service"
	"github.com/cutejiuges/disk_back/internal/util/errutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 下午7:10
 * @FilePath: ping
 * @Description:
 */

type PingHandler struct {
	ctx   context.Context
	req   *disk_common.PingRequest
	resp  *disk_common.PingResponse
	reply string
	err   error
}

func NewPingHandler(ctx context.Context, req *disk_common.PingRequest) *PingHandler {
	return &PingHandler{
		ctx:  ctx,
		req:  req,
		resp: disk_common.NewPingResponse(),
	}
}

func (h *PingHandler) checkParam() error {
	return nil
}

func (h *PingHandler) checkPermission() error {
	return nil
}

func (h *PingHandler) processBusiness() error {
	reply, err := service.ProcessPing(h.ctx, h.req)
	if err != nil {
		klog.CtxErrorf(h.ctx, "processBusiness -> ProcessPing error: %v", err)
		return err
	}
	h.reply = reply
	return nil
}

func (h *PingHandler) packResp() {
	h.resp.SetBaseResp(errutil.NewBaseRespWithOK())
	h.resp.SetReply(h.reply)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "PingHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errutil.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *PingHandler) Handle() (*disk_common.PingResponse, error) {
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
