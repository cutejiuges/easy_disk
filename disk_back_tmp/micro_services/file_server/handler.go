package main

import (
	"context"
	disk_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	file_server "github.com/cutejiuges/disk_back/kitex_gen/file_server"
)

// DiskBackServiceImpl implements the last service interface defined in the IDL.
type DiskBackServiceImpl struct{}

// UploadFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) UploadFile(ctx context.Context, req *file_server.UploadFileRequest) (resp *file_server.UploadFileResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFileInfo implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) QueryFileInfo(ctx context.Context, req *disk_common.QueryFileInfoRequest) (resp *file_server.QueryFileInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DownloadFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) DownloadFile(ctx context.Context, req *disk_common.DownloadFileRequest) (resp *file_server.DownloadFileResponse, err error) {
	// TODO: Your code here...
	return
}

// EditFileInfo implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) EditFileInfo(ctx context.Context, req *disk_common.EditFileInfoRequest) (resp *file_server.EditFileInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteFile implements the DiskBackServiceImpl interface.
func (s *DiskBackServiceImpl) DeleteFile(ctx context.Context, req *disk_common.DeleteFileRequest) (resp *file_server.DeleteFileResponse, err error) {
	// TODO: Your code here...
	return
}
