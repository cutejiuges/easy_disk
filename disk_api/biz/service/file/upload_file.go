package file_service

import (
	"bytes"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cutejiuges/disk_api/biz/model/disk_api"
	"github.com/cutejiuges/disk_api/infra/localutils"
	"github.com/cutejiuges/disk_api/rpc"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	"golang.org/x/net/context"
	"io"
	"mime/multipart"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午6:30
 * @FilePath: upload_file
 * @Description:
 */

func ProcessUploadFile(ctx context.Context, c *app.RequestContext, req *disk_api.UploadFileRequest) (*disk_back.UploadFileResponse, error) {
	var rpcReq disk_back.UploadFileRequest
	err := localutils.Converter(req, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "ProcessUploadFile convert req error: %v", err)
		return nil, err
	}

	var fileList []*disk_back.UploadFileMeta
	form, _ := c.MultipartForm()
	fileHeaders := form.File["file"]
	for _, fileHeader := range fileHeaders {
		err := func(f *multipart.FileHeader) error {
			file, err := f.Open()
			if err != nil {
				return err
			}
			defer file.Close()

			buf := bytes.Buffer{}
			_, _ = io.Copy(&buf, file)

			fileList = append(fileList, &disk_back.UploadFileMeta{
				FileName: f.Filename,
				FileData: buf.Bytes(),
			})
			return nil
		}(fileHeader)

		if err != nil {
			hlog.CtxErrorf(ctx, "load file error: %v", err)
			return nil, err
		}
	}
	rpcReq.SetFiles(fileList)

	rpcResp, err := rpc.GetDiskBackClient().UploadFile(ctx, &rpcReq)
	if err != nil {
		hlog.CtxErrorf(ctx, "ProcessUploadFile -> rpc UploadFile error: %v", err)
		return nil, err
	}
	return rpcResp, nil
}
