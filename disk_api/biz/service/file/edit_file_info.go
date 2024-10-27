package file_service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cutejiuges/disk_api/biz/model/disk_common"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	disk_back_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 下午11:22
 * @FilePath: edit_file_info
 * @Description: 修改文件信息,目前仅支持修改远端的文件名
 */

func ProcessEditFileInfo(ctx context.Context, req *disk_common.EditFileInfoRequest) (*disk_back.EditFileInfoResponse, error) {
	var rpcReq disk_back_common.EditFileInfoRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.ProcessEditFileInfo request converter error: %v", err)
		return nil, err
	}
	resp, err := rpc.GetDiskBackClient().EditFileInfo(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "file_service.ProcessEditFileInfo -> rpc EditFileInfo error: %v", err)
		return nil, err
	}
	return resp, nil
}
