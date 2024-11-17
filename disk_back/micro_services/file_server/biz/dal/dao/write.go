package dao

import (
	"context"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/enum"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/pojo/param"
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
		FileName: fileParam.FileName,
		FileSize: fileParam.FileSize,
		FileAddr: fileParam.FileAddr,
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
			FileName: file.FileName,
			FileSize: file.FileSize,
			FileAddr: file.FileAddr,
			Status:   file.Status,
		}
		fileMetaList = append(fileMetaList, meta)
	}
	return dao.CreateInBatches(fileMetaList, batchSize)
}

func EditFileInfo(ctx context.Context, q *query.Query, param *param.EditFileMetaParam) error {
	meta := q.FileMeta
	dao := q.WithContext(ctx).FileMeta.Where(meta.Status.Neq(enum.FileMetaStatusDeleted))
	if param.ID > 0 {
		dao = dao.Where(meta.ID.Eq(param.ID))
	}
	if param.FileKey != "" {
		dao = dao.Where(meta.FileKey.Eq(param.FileKey))
	}
	_, err := dao.UpdateSimple(meta.FileName.Value(param.FileName), meta.FileAddr.Value(param.FileAddr))
	return err
}

func DeleteFile(ctx context.Context, q *query.Query, param *param.EditFileMetaParam) error {
	meta := q.FileMeta
	dao := q.WithContext(ctx).FileMeta.Where(meta.Status.Neq(enum.FileMetaStatusDeleted))
	if param.ID > 0 {
		dao = dao.Where(meta.ID.Eq(param.ID))
	}
	if param.FileKey != "" {
		dao = dao.Where(meta.FileKey.Eq(param.FileKey))
	}
	if len(param.IdList) > 0 {
		dao = dao.Where(meta.ID.In(param.IdList...))
	}
	_, err := dao.UpdateSimple(meta.Status.Value(enum.FileMetaStatusDeleted))
	return err
}
