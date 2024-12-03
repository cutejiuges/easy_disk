package main

import (
	"context"
	file_server "github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/handler"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct{}

// UploadFileBatch implements the FileServiceImpl interface.
func (s *FileServiceImpl) UploadFileBatch(ctx context.Context, req *file_server.UploadFileRequest) (resp *file_server.UploadFileResponse, err error) {
	return handler.NewUploadFileBatchHandler(ctx, req).Handle()
}

// QueryFileInfo implements the FileServiceImpl interface.
func (s *FileServiceImpl) QueryFileInfo(ctx context.Context, req *file_server.QueryFileInfoRequest) (resp *file_server.QueryFileInfoResponse, err error) {
	return handler.NewQueryFileInfoHandler(ctx, req).Handle()
}

// DownloadFile implements the FileServiceImpl interface.
func (s *FileServiceImpl) DownloadFile(ctx context.Context, req *file_server.DownloadFileRequest) (resp *file_server.DownloadFileResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteFile implements the FileServiceImpl interface.
func (s *FileServiceImpl) DeleteFile(ctx context.Context, req *file_server.DeleteFileRequest) (resp *file_server.DeleteFileResponse, err error) {
	return handler.NewDeleteFileHandler(ctx, req).Handle()
}
