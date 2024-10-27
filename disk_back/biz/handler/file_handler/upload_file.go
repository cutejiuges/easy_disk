package file_handler

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/service/file_service"
	"github.com/cutejiuges/disk_back/internal/util/errutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 上午1:33
 * @FilePath: upload_file
 * @Description:
 */

type UploadFileHandler struct {
	ctx  context.Context
	req  *disk_back.UploadFileRequest
	resp *disk_back.UploadFileResponse
	data *disk_back.UploadFileData
	err  error
}

func NewUploadFileHandler(ctx context.Context, req *disk_back.UploadFileRequest) *UploadFileHandler {
	return &UploadFileHandler{
		ctx:  ctx,
		req:  req,
		resp: disk_back.NewUploadFileResponse(),
	}
}

func (h *UploadFileHandler) checkParam() error {
	if len(h.req.GetFiles()) <= 0 {
		return errors.New("未上传有效文件")
	}
	return nil
}

func (h *UploadFileHandler) checkPermission() error {
	return nil
}

func (h *UploadFileHandler) processBusiness() error {
	data, err := file_service.ProcessUploadFile(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *UploadFileHandler) packResp() {
	h.resp.SetBaseResp(errutil.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "UploadFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errutil.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *UploadFileHandler) Handle() (*disk_back.UploadFileResponse, error) {
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
