package file_service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/biz/dao/file_meta_dao"
	"github.com/cutejiuges/disk_back/internal/infra/mysql"
	"github.com/cutejiuges/disk_back/internal/model/param"
	"github.com/cutejiuges/disk_back/internal/model/query"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	"os"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 下午5:18
 * @FilePath: edit_file_info
 * @Description: 编辑文件信息
 */

func ProcessEditFileInfo(ctx context.Context, req *disk_common.EditFileInfoRequest) (*disk_back.SaveFileRes, error) {
	data := disk_back.NewSaveFileRes()

	//查询文件信息
	meta, err := file_meta_dao.QuerySingleFileMeta(ctx, &param.QueryFileMetaParam{
		ID: req.GetId(),
	})
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessEditFileInfo -> file_meta_dao.QuerySingleFileMeta error: %v", err)
		return data, err
	}

	//执行文件的编辑操作
	newPath := getNewPath(meta.FileAddr, req.GetFileName())
	qry := query.Use(mysql.DB()).Begin()
	//先更改db存储的信息
	//如果先改文件存储，db更改有问题的话，这一步骤无法回滚
	err = file_meta_dao.EditFileInfo(ctx, qry.Query, &param.EditFileMetaParam{
		ID:       meta.ID,
		FileAddr: newPath,
		FileName: req.GetFileName(),
	})
	//有问题需要回滚
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessEditFileInfo -> file_meta_dao.EditFileInfo error: %v", err)
		_ = qry.Rollback()
		return data, err
	}
	//再更改实际的文件存储,确保实际更改有问题，也能回滚掉db的更改
	err = os.Rename(meta.FileAddr, newPath)
	if err != nil {
		klog.CtxErrorf(ctx, "file_service.ProcessEditFileInfo -> os.Rename error: %v", err)
		_ = qry.Rollback()
		return data, err
	}
	_ = qry.Commit()

	//更改成功，返回信息
	data.SetFileName(req.GetFileName())
	data.SetId(meta.ID)
	data.SetMsg("文件更改成功")
	return data, nil
}

func getNewPath(path, newName string) string {
	dirIdx := strings.LastIndex(path, "/")
	fileIdx := strings.LastIndex(path, ".")
	return path[:dirIdx+1] + newName + path[fileIdx:]
}
