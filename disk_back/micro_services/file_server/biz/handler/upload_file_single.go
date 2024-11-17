package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/errno"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/service"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午10:25
 * @FilePath: upload_file_single
 * @Description: 单独文件上传功能
 */

type UploadFileSingleHandler struct {
	ctx  context.Context
	req  *file_server.UploadFileRequest
	resp *file_server.UploadFileResponse
	data *file_server.UploadFileData
	err  error
}

func NewUploadFileHandler(ctx context.Context, req *file_server.UploadFileRequest) *UploadFileSingleHandler {
	return &UploadFileSingleHandler{
		ctx:  ctx,
		req:  req,
		resp: file_server.NewUploadFileResponse(),
	}
}

func (h *UploadFileSingleHandler) checkParam() error {
	if len(h.req.GetFiles()) <= 0 {
		return &errno.BizError{
			Code: errno.ErrCodeParamIllegal,
			Msg:  errno.ErrMsgMap[errno.ErrCodeParamIllegal] + ", 未上传有效文件",
		}
	}
	return nil
}

func (h *UploadFileSingleHandler) checkPermission() error {
	return nil
}

func (h *UploadFileSingleHandler) processBusiness() error {
	data, err := service.ProcessUploadFileSingle(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *UploadFileSingleHandler) packResp() {
	h.resp.SetBaseResp(errno.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "UploadFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *UploadFileSingleHandler) Handle() (*file_server.UploadFileResponse, error) {
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
