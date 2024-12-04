package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	errno2 "github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/service"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午10:25
 * @FilePath: upload_file_single
 * @Description: 文件批量上传功能
 */

type UploadFileBatchHandler struct {
	ctx  context.Context
	req  *file_server.UploadFileRequest
	resp *file_server.UploadFileResponse
	data *file_server.UploadFileData
	err  error
}

func NewUploadFileBatchHandler(ctx context.Context, req *file_server.UploadFileRequest) *UploadFileBatchHandler {
	return &UploadFileBatchHandler{
		ctx:  ctx,
		req:  req,
		resp: file_server.NewUploadFileResponse(),
	}
}

func (h *UploadFileBatchHandler) checkParam() error {
	if len(h.req.GetFiles()) <= 0 {
		return &errno2.BizError{
			Code: errno2.ErrCodeParamIllegal,
			Msg:  errno2.ErrMsgMap[errno2.ErrCodeParamIllegal] + ", 未上传有效文件",
		}
	}
	return nil
}

func (h *UploadFileBatchHandler) checkPermission() error {
	return nil
}

func (h *UploadFileBatchHandler) processBusiness() error {
	data, err := service.ProcessUploadFileBatch(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *UploadFileBatchHandler) packResp() {
	h.resp.SetBaseResp(errno2.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "UploadFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno2.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *UploadFileBatchHandler) Handle() (*file_server.UploadFileResponse, error) {
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
