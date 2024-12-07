package dao

import (
	"context"
	"github.com/cutejiuges/disk_back/internal/enum"
	errno2 "github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/file_server/pojo"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 上午1:01
 * @FilePath: write
 * @Description: file_meta的写相关操作
 */

func CreateFile(ctx context.Context, q *query.Query, fileParam *model.FileMeta) error {
	dao := q.WithContext(ctx).FileMeta
	fileMeta := &model.FileMeta{
		ID:       fileParam.ID,
		FileKey:  fileParam.FileKey,
		FileSize: fileParam.FileSize,
		FileAddr: fileParam.FileAddr,
		RefNum:   fileParam.RefNum,
		Status:   fileParam.Status,
	}
	return dao.Create(fileMeta)
}

func CreateFilesInBatch(ctx context.Context, q *query.Query, fileList []*model.FileMeta, batchSize int) error {
	dao := q.WithContext(ctx).FileMeta
	fileMetaList := make([]*model.FileMeta, 0)
	for _, file := range fileList {
		meta := &model.FileMeta{
			ID:       file.ID,
			FileKey:  file.FileKey,
			FileSize: file.FileSize,
			FileAddr: file.FileAddr,
			RefNum:   file.RefNum,
			Status:   file.Status,
		}
		fileMetaList = append(fileMetaList, meta)
	}
	return dao.CreateInBatches(fileMetaList, batchSize)
}

func DeleteFile(ctx context.Context, q *query.Query, params *pojo.EditFileMetaParam) error {
	if params.ID <= 0 && params.FileKey == "" && len(params.IdList) <= 0 {
		return nil
	}
	meta := q.FileMeta
	dao := q.WithContext(ctx).FileMeta.Where(meta.Status.Neq(enum.FileMetaStatusDeleted))
	if params.ID > 0 {
		dao = dao.Where(meta.ID.Eq(params.ID))
	}
	if params.FileKey != "" {
		dao = dao.Where(meta.FileKey.Eq(params.FileKey))
	}
	if len(params.IdList) > 0 {
		dao = dao.Where(meta.ID.In(params.IdList...))
	}
	_, err := dao.UpdateSimple(meta.Status.Value(enum.FileMetaStatusDeleted))
	return err
}

func ModifyFileRef(ctx context.Context, q *query.Query, param *pojo.EditFileMetaParam) error {
	if len(param.IdList) <= 0 {
		return &errno2.BizError{
			Code: errno2.ErrCodeDbUnknownError,
			Msg:  "检索条件不正确，id为必传参数",
		}
	}
	meta := q.FileMeta
	dao := q.WithContext(ctx).FileMeta.Where(meta.Status.Neq(enum.FileMetaStatusDeleted))
	dao = dao.Where(meta.ID.In(param.IdList...))
	_, err := dao.UpdateSimple(meta.RefNum.Add(param.RefDealt))
	return err
}
