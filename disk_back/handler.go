package main

import (
	"context"
	"github.com/cutejiuges/disk_back/biz/handler"
	"github.com/cutejiuges/disk_back/biz/handler/file_handler"
	disk_back "github.com/cutejiuges/disk_back/kitex_gen/disk_back"
	disk_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

// DiskBackServiceImpl implements the last service interface defined in the IDL.
type DiskBackServiceImpl struct{}

// UploadFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) UploadFile(ctx context.Context, req *disk_back.UploadFileRequest) (resp *disk_back.UploadFileResponse, err error) {
	return file_handler.NewUploadFileHandler(ctx, req).Handle()
}

// Ping implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) Ping(ctx context.Context, req *disk_common.PingRequest) (resp *disk_common.PingResponse, err error) {
	return handler.NewPingHandler(ctx, req).Handle()
}

// QueryFileInfo implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) QueryFileInfo(ctx context.Context, req *disk_common.QueryFileInfoRequest) (resp *disk_back.QueryFileInfoResponse, err error) {
	return file_handler.NewQueryFileInfoHandler(ctx, req).Handle()
}

// DownloadFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) DownloadFile(ctx context.Context, req *disk_common.DownloadFileRequest) (resp *disk_back.DownloadFileResponse, err error) {
	return file_handler.NewDownloadFileHandler(ctx, req).Handle()
}

// EditFileInfo implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) EditFileInfo(ctx context.Context, req *disk_common.EditFileInfoRequest) (resp *disk_back.EditFileInfoResponse, err error) {
	return file_handler.NewEditFileInfoHandler(ctx, req).Handle()
}

// DeleteFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) DeleteFile(ctx context.Context, req *disk_common.DeleteFileRequest) (resp *disk_back.DeleteFileResponse, err error) {
	// TODO: Your code here...
	return
}
