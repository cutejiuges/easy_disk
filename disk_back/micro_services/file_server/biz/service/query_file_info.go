package service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/enum"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/pojo/param"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/util"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/17 上午12:00
 * @FilePath: query_file_info
 * @Description:
 */

func ProcessQueryFileInfo(ctx context.Context, req *disk_common.QueryFileInfoRequest) ([]*file_server.QueryFileInfoData, int64, error) {
	data := make([]*file_server.QueryFileInfoData, 0)
	fileParam := &param.QueryFileMetaParam{
		IdList: req.GetID(),
		Status: req.GetStatus(),
		Page:   int(req.GetPage()),
		Size:   int(req.GetSize()),
	}
	if len(req.GetUploadTimeInterval()) > 0 {
		fileParam.MinCreateTime = req.GetUploadTimeInterval()[0]
		fileParam.MaxCreateTime = req.GetUploadTimeInterval()[1]
	}
	if len(req.GetFileSizeInterval()) > 0 {
		fileParam.MinFileSize = req.GetFileSizeInterval()[0]
		fileParam.MaxFileSize = req.GetFileSizeInterval()[1]
	}
	klog.CtxInfof(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaListByPage, param: %v", fileParam)
	list, total, err := dao.QueryFileMetaListByPage(ctx, fileParam)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaListByPage error: %v", err)
		return data, total, err
	}
	klog.CtxInfof(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaListByPage, resp: %v, total: %d", list, total)
	for _, meta := range list {
		info := &file_server.QueryFileInfoData{
			ID:             meta.ID,
			FileAddress:    meta.FileAddr,
			FileSize:       meta.FileSize,
			FileSizeString: thrift.StringPtr(FormatByteCount(meta.FileSize, int(enum.FileSizeUnitIEC))),
			UploadTime:     util.FormatTime(meta.CreateAt, string(enum.TimeLayoutCompleteMinus)),
			Status:         meta.Status,
			StatusName:     thrift.StringPtr(enum.FileMetaStatusNameMap[meta.Status]),
		}
		data = append(data, info)
	}
	return data, total, err
}

// FormatByteCount 根据文件的字节大小,计算格式化的大小字符串
func FormatByteCount(b int64, base int) string {
	unit := int64(base)
	if b < int64(unit) {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTEP"[exp])
}
