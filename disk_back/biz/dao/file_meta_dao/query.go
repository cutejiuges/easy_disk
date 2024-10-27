package file_meta_dao

import (
	"context"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/infra/mysql"
	"github.com/cutejiuges/disk_back/internal/model/model"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/model/query"
	"github.com/cutejiuges/disk_back/internal/util/localutil"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/17 上午12:06
 * @FilePath: query
 * @Description: 文件元信息查询接口
 */

func QuerySingleFileMeta(ctx context.Context, param *param.QueryFileMetaParam) (*model.FileMeta, error) {
	q := query.Use(mysql.DB()).FileMeta
	dao := q.WithContext(ctx).Where(q.Status.Neq(enum.FileMetaStatusUnknown))
	if param.ID != 0 {
		dao = dao.Where(q.ID.Eq(param.ID))
	}
	if param.FileKey != "" {
		dao = dao.Where(q.FileKey.Eq(param.FileKey))
	}
	dao = dao.Where(q.Status.Neq(enum.FileMetaStatusDeleted))
	return dao.First()
}

func QueryFileMetaList(ctx context.Context, param *param.QueryFileMetaParam) ([]*model.FileMeta, int64, error) {
	q := query.Use(mysql.DB()).FileMeta
	dao := q.WithContext(ctx).Where(q.Status.Neq(enum.FileMetaStatusUnknown))
	if len(param.IdList) > 0 {
		dao.Where(q.ID.In(param.IdList...))
	}
	if len(param.FileName) > 0 {
		dao.Where(q.FileName.Like(param.FileName))
	}
	if len(param.Status) > 0 {
		statusDao := q.WithContext(ctx)
		statusCond := statusDao.Where(q.Status.In(param.Status...))
		dao.Where(statusCond)
	}
	if param.MinCreateTime != "" && param.MaxCreateTime != "" {
		minTime, _ := localutil.ParseTime(param.MinCreateTime, string(enum.TimeLayoutCompleteMinus))
		maxTime, _ := localutil.ParseTime(param.MaxCreateTime, string(enum.TimeLayoutCompleteMinus))
		dao.Where(q.CreateAt.Gte(minTime), q.CreateAt.Lte(maxTime))
	}
	if param.MinFileSize > 0 && param.MaxFileSize > 0 {
		dao.Where(q.FileSize.Gte(param.MinFileSize), q.FileSize.Lte(param.MaxFileSize))
	}

	return dao.FindByPage((param.Page-1)*param.Size, param.Size)
}
