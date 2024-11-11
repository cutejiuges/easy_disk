package file_handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/service/file_service"
	"github.com/cutejiuges/disk_back/internal/util/errutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/28 下午11:22
 * @FilePath: delete_file
 * @Description: 删除文件
 */

type DeleteFileHandler struct {
	ctx  context.Context
	req  *disk_common.DeleteFileRequest
	resp *disk_back.DeleteFileResponse
	data *disk_back.DeleteFileData
	err  error
}

func NewDeleteFileHandler(ctx context.Context, req *disk_common.DeleteFileRequest) *DeleteFileHandler {
	return &DeleteFileHandler{
		ctx:  ctx,
		req:  req,
		resp: disk_back.NewDeleteFileResponse(),
	}
}

func (h *DeleteFileHandler) checkParam() error {
	if len(h.req.GetId()) <= 0 {
		return &errutil.BizError{
			Code: errutil.ErrCodeParamIllegal,
			Msg:  errutil.ErrMsgMap[errutil.ErrCodeParamIllegal] + ", 无有效文件id",
		}
	}
	for _, id := range h.req.GetId() {
		if id <= 0 {
			return &errutil.BizError{
				Code: errutil.ErrCodeParamIllegal,
				Msg:  errutil.ErrMsgMap[errutil.ErrCodeParamIllegal] + ", 存在非法文件id",
			}
		}
	}
	return nil
}

func (h *DeleteFileHandler) checkPermission() error {
	return nil
}

func (h *DeleteFileHandler) processBusiness() error {
	data, err := file_service.ProcessDeleteFile(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *DeleteFileHandler) packResp() {
	h.resp.SetBaseResp(errutil.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "file_handler.DeleteFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errutil.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *DeleteFileHandler) Handle() (*disk_back.DeleteFileResponse, error) {
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
