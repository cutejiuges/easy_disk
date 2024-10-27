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
 * @Date: 2024/10/27 下午3:09
 * @FilePath: download_file
 * @Description: 文件下载接口
 */

type DownloadFileHandler struct {
	ctx  context.Context
	req  *disk_common.DownloadFileRequest
	resp *disk_back.DownloadFileResponse
	data *disk_back.DownloadFileData
	err  error
}

func NewDownloadFileHandler(ctx context.Context, req *disk_common.DownloadFileRequest) *DownloadFileHandler {
	return &DownloadFileHandler{
		ctx:  ctx,
		req:  req,
		resp: disk_back.NewDownloadFileResponse(),
	}
}

func (h *DownloadFileHandler) checkParam() error {
	if h.req.Id <= 0 {
		return &errutil.BizError{
			Code: errutil.ErrCodeParamIllegal,
			Msg:  errutil.ErrMsgMap[errutil.ErrCodeParamIllegal] + ",文件id非法",
		}
	}
	return nil
}

func (h *DownloadFileHandler) checkPermission() error {
	return nil
}

func (h *DownloadFileHandler) processBusiness() error {
	data, err := file_service.ProcessDownloadFile(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.data = data
	return nil
}

func (h *DownloadFileHandler) packResp() {
	h.resp.SetBaseResp(errutil.NewBaseRespWithOK())
	h.resp.SetData(h.data)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "DownloadFileHandler exec error: %v", h.err)
		h.resp.SetBaseResp(errutil.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *DownloadFileHandler) Handle() (*disk_back.DownloadFileResponse, error) {
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
