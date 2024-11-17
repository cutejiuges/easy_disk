// Code generated by Kitex v0.10.0. DO NOT EDIT.

package fileservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	disk_common "github.com/cutejiuges/disk_back/kitex_gen/disk_common"
	file_server "github.com/cutejiuges/disk_back/kitex_gen/file_server"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"UploadFile": kitex.NewMethodInfo(
		uploadFileHandler,
		newFileServiceUploadFileArgs,
		newFileServiceUploadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"QueryFileInfo": kitex.NewMethodInfo(
		queryFileInfoHandler,
		newFileServiceQueryFileInfoArgs,
		newFileServiceQueryFileInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DownloadFile": kitex.NewMethodInfo(
		downloadFileHandler,
		newFileServiceDownloadFileArgs,
		newFileServiceDownloadFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"EditFileInfo": kitex.NewMethodInfo(
		editFileInfoHandler,
		newFileServiceEditFileInfoArgs,
		newFileServiceEditFileInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"DeleteFile": kitex.NewMethodInfo(
		deleteFileHandler,
		newFileServiceDeleteFileArgs,
		newFileServiceDeleteFileResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	fileServiceServiceInfo                = NewServiceInfo()
	fileServiceServiceInfoForClient       = NewServiceInfoForClient()
	fileServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return fileServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return fileServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return fileServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "FileService"
	handlerType := (*file_server.FileService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "file_server",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.10.0",
		Extra:           extra,
	}
	return svcInfo
}

func uploadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file_server.FileServiceUploadFileArgs)
	realResult := result.(*file_server.FileServiceUploadFileResult)
	success, err := handler.(file_server.FileService).UploadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceUploadFileArgs() interface{} {
	return file_server.NewFileServiceUploadFileArgs()
}

func newFileServiceUploadFileResult() interface{} {
	return file_server.NewFileServiceUploadFileResult()
}

func queryFileInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file_server.FileServiceQueryFileInfoArgs)
	realResult := result.(*file_server.FileServiceQueryFileInfoResult)
	success, err := handler.(file_server.FileService).QueryFileInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceQueryFileInfoArgs() interface{} {
	return file_server.NewFileServiceQueryFileInfoArgs()
}

func newFileServiceQueryFileInfoResult() interface{} {
	return file_server.NewFileServiceQueryFileInfoResult()
}

func downloadFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file_server.FileServiceDownloadFileArgs)
	realResult := result.(*file_server.FileServiceDownloadFileResult)
	success, err := handler.(file_server.FileService).DownloadFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceDownloadFileArgs() interface{} {
	return file_server.NewFileServiceDownloadFileArgs()
}

func newFileServiceDownloadFileResult() interface{} {
	return file_server.NewFileServiceDownloadFileResult()
}

func editFileInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file_server.FileServiceEditFileInfoArgs)
	realResult := result.(*file_server.FileServiceEditFileInfoResult)
	success, err := handler.(file_server.FileService).EditFileInfo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceEditFileInfoArgs() interface{} {
	return file_server.NewFileServiceEditFileInfoArgs()
}

func newFileServiceEditFileInfoResult() interface{} {
	return file_server.NewFileServiceEditFileInfoResult()
}

func deleteFileHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*file_server.FileServiceDeleteFileArgs)
	realResult := result.(*file_server.FileServiceDeleteFileResult)
	success, err := handler.(file_server.FileService).DeleteFile(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFileServiceDeleteFileArgs() interface{} {
	return file_server.NewFileServiceDeleteFileArgs()
}

func newFileServiceDeleteFileResult() interface{} {
	return file_server.NewFileServiceDeleteFileResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) UploadFile(ctx context.Context, req *file_server.UploadFileRequest) (r *file_server.UploadFileResponse, err error) {
	var _args file_server.FileServiceUploadFileArgs
	_args.Req = req
	var _result file_server.FileServiceUploadFileResult
	if err = p.c.Call(ctx, "UploadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryFileInfo(ctx context.Context, req *disk_common.QueryFileInfoRequest) (r *file_server.QueryFileInfoResponse, err error) {
	var _args file_server.FileServiceQueryFileInfoArgs
	_args.Req = req
	var _result file_server.FileServiceQueryFileInfoResult
	if err = p.c.Call(ctx, "QueryFileInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DownloadFile(ctx context.Context, req *disk_common.DownloadFileRequest) (r *file_server.DownloadFileResponse, err error) {
	var _args file_server.FileServiceDownloadFileArgs
	_args.Req = req
	var _result file_server.FileServiceDownloadFileResult
	if err = p.c.Call(ctx, "DownloadFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EditFileInfo(ctx context.Context, req *disk_common.EditFileInfoRequest) (r *file_server.EditFileInfoResponse, err error) {
	var _args file_server.FileServiceEditFileInfoArgs
	_args.Req = req
	var _result file_server.FileServiceEditFileInfoResult
	if err = p.c.Call(ctx, "EditFileInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFile(ctx context.Context, req *disk_common.DeleteFileRequest) (r *file_server.DeleteFileResponse, err error) {
	var _args file_server.FileServiceDeleteFileArgs
	_args.Req = req
	var _result file_server.FileServiceDeleteFileResult
	if err = p.c.Call(ctx, "DeleteFile", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
