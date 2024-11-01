package file_meta_dao

import (
	"context"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/model/model"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/model/query"
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
	dao := q.WithContext(ctx).FileMeta.Where(meta.Status.Neq(enum.FileMetaStatusUnknown), meta.Status.Neq(enum.FileMetaStatusDeleted))
	if param.ID > 0 {
		dao = dao.Where(meta.ID.Eq(param.ID))
	}
	if param.FileKey != "" {
		dao = dao.Where(meta.FileKey.Eq(param.FileKey))
	}
	_, err := dao.UpdateSimple(meta.FileName.Value(param.FileName), meta.FileAddr.Value(param.FileAddr))
	return err
}
