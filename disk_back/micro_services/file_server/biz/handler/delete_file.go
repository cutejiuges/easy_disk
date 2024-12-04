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
 * @Date: 2024/11/27 下午10:54
 * @FilePath: delete_file
 * @Description: 执行文件删除操作
 */

type DeleteFileHandler struct {
	ctx  context.Context
	req  *file_server.DeleteFileRequest
	data *file_server.DeleteFileData
	resp *file_server.DeleteFileResponse
	err  error
}

func NewDeleteFileHandler(ctx context.Context, req *file_server.DeleteFileRequest) *DeleteFileHandler {
	return &DeleteFileHandler{
		ctx:  ctx,
		req:  req,
		resp: file_server.NewDeleteFileResponse(),
	}
}

func (h *DeleteFileHandler) checkParam() error {
	if len(h.req.GetId()) <= 0 {
		return &errno2.BizError{
			Code: errno2.ErrCodeParamIllegal,
			Msg:  errno2.ErrMsgMap[errno2.ErrCodeParamIllegal] + ", 无有效文件id",
		}
	}
	return nil
}

func (h *DeleteFileHandler) checkPermission() error {
	return nil
}

func (h *DeleteFileHandler) processBusiness() error {
	data, err := service.ProcessDeleteFile(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *DeleteFileHandler) packResp() {
	h.resp.SetBaseResp(errno2.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "UploadFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errno2.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *DeleteFileHandler) Handle() (*file_server.DeleteFileResponse, error) {
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
