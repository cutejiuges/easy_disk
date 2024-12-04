package handler

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/service"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午11:59
 * @FilePath: query_file_info
 * @Description: 查询文件信息
 */

type QueryFileInfoHandler struct {
	ctx      context.Context
	req      *file_server.QueryFileInfoRequest
	resp     *file_server.QueryFileInfoResponse
	dataList []*file_server.QueryFileInfoData
	total    int64
	err      error
}

func NewQueryFileInfoHandler(ctx context.Context, req *file_server.QueryFileInfoRequest) *QueryFileInfoHandler {
	return &QueryFileInfoHandler{
		ctx:  ctx,
		req:  req,
		resp: file_server.NewQueryFileInfoResponse(),
	}
}

func (h *QueryFileInfoHandler) checkParam() error {
	param := h.req
	if param.Size <= 0 || param.Page <= 0 {
		return fmt.Errorf("分页参数设置有误, 请检查")
	}
	if len(param.UploadTimeInterval) > 0 {
		if len(param.UploadTimeInterval) < 2 {
			return fmt.Errorf("上传时间设置有误,请检查开始时间和截止时间")
		}
		minTime, err := util.ParseTime(param.UploadTimeInterval[0], string(enum.TimeLayoutCompleteMinus))
		if err != nil {
			return fmt.Errorf("上传时间设置有误,请检查开始时间, err: %v", err)
		}
		maxTime, err := util.ParseTime(param.UploadTimeInterval[1], string(enum.TimeLayoutCompleteMinus))
		if err != nil {
			return fmt.Errorf("上传时间设置有误,请检查截止时间, err: %v", err)
		}
		if minTime.Sub(maxTime) > 0 {
			return fmt.Errorf("上传时间设置有误,开始时间大于截至时间")
		}
	}
	if len(param.FileSizeInterval) > 0 {
		if len(param.FileSizeInterval) < 2 {
			return fmt.Errorf("大小区间设置有误,请检查大小上下限")
		}
		if param.FileSizeInterval[0] > param.FileSizeInterval[1] {
			return fmt.Errorf("大小区间设置有误,size下限大于上限")
		}
	}
	return nil
}

func (h *QueryFileInfoHandler) checkPermission() error {
	return nil
}

func (h *QueryFileInfoHandler) processBusiness() error {
	list, total, err := service.ProcessQueryFileInfo(h.ctx, h.req)
	if err != nil {
		return err
	}
	h.dataList = list
	h.total = total
	return nil
}

func (h *QueryFileInfoHandler) packResp() {
	h.resp.SetBaseResp(errno.NewBaseRespWithOK())
	h.resp.SetData(h.dataList)
	h.resp.SetTotal(&h.total)
	if h.err != nil {
		klog.CtxErrorf(h.ctx, "QueryFileInfo exec failed, error: %v", h.err)
		h.resp.SetBaseResp(errno.NewBaseResp(h.err))
		h.err = nil
	}
}

func (h *QueryFileInfoHandler) Handle() (*file_server.QueryFileInfoResponse, error) {
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
