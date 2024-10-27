package file_service

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/dao/file_meta_dao"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/util/localutil"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"golang.org/x/net/context"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/26 下午7:05
 * @FilePath: query_file_info
 * @Description: 通过查询参数列表查询文件信息
 */

func ProcessQueryFileInfo(ctx context.Context, req *disk_common.QueryFileInfoRequest) ([]*disk_back.QueryFileInfoData, int64, error) {
	data := make([]*disk_back.QueryFileInfoData, 0)
	fileParam := &param.QueryFileMetaParam{
		IdList:   req.GetID(),
		FileName: req.GetFileName(),
		Status:   req.GetStatus(),
		Page:     int(req.GetPage()),
		Size:     int(req.GetSize()),
	}
	if len(req.GetUploadTimeInterval()) > 0 {
		fileParam.MinCreateTime = req.GetUploadTimeInterval()[0]
		fileParam.MaxCreateTime = req.GetUploadTimeInterval()[1]
	}
	if len(req.GetFileSizeInterval()) > 0 {
		fileParam.MinFileSize = req.GetFileSizeInterval()[0]
		fileParam.MaxFileSize = req.GetFileSizeInterval()[1]
	}
	klog.CtxInfof(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaList, param: %v", fileParam)
	list, total, err := file_meta_dao.QueryFileMetaList(ctx, fileParam)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaList error: %v", err)
		return data, total, err
	}
	klog.CtxInfof(ctx, "file_service.ProcessQueryFileInfo -> file_meta_dao.QueryFileMetaList, resp: %v, total: %d", list, total)
	for _, meta := range list {
		info := &disk_back.QueryFileInfoData{
			ID:             meta.ID,
			FileName:       meta.FileName,
			FileAddress:    meta.FileAddr,
			FileSize:       meta.FileSize,
			FileSizeString: thrift.StringPtr(FormatByteCount(meta.FileSize, int(enum.FileSizeUnitIEC))),
			UploadTime:     localutil.FormatTime(meta.CreateAt, string(enum.TimeLayoutCompleteMinus)),
			Status:         meta.Status,
			StatusName:     thrift.StringPtr(enum.FileMetaStatusNameMap[enum.FileMetaStatus(meta.Status)]),
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
