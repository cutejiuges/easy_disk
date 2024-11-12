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
 * @Date: 2024/10/27 下午5:25
 * @FilePath: edit_file_info
 * @Description: 编辑文件信息
 */

type EditFileInfoHandler struct {
	ctx  context.Context
	req  *disk_common.EditFileInfoRequest
	resp *disk_back.EditFileInfoResponse
	data *disk_back.OperateFileRes
	err  error
}

func NewEditFileInfoHandler(ctx context.Context, req *disk_common.EditFileInfoRequest) *EditFileInfoHandler {
	return &EditFileInfoHandler{
		ctx:  ctx,
		req:  req,
		resp: disk_back.NewEditFileInfoResponse(),
	}
}

func (h *EditFileInfoHandler) checkParam() error {
	if h.req.GetId() <= 0 {
		return &errutil.BizError{
			Code: errutil.ErrCodeParamIllegal,
			Msg:  errutil.ErrMsgMap[errutil.ErrCodeParamIllegal] + ", 文件id不合法",
		}
	}
	return nil
}

func (h *EditFileInfoHandler) checkPermission() error {
	return nil
}

func (h *EditFileInfoHandler) processBusiness() error {
	data, err := file_service.ProcessEditFileInfo(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *EditFileInfoHandler) packResp() {
	h.resp.SetBaseResp(errutil.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "file_handler.EditFileInfoHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errutil.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *EditFileInfoHandler) Handle() (*disk_back.EditFileInfoResponse, error) {
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
