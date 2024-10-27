package enum

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午4:33
 * @FilePath: upload_file
 * @Description:
 */

type UploadFileStatus string

const (
	UploadFileStatusSuccess             UploadFileStatus = "Success"
	UploadFileStatusFailed              UploadFileStatus = "Failed"
	UploadFileStatusPartiallySuccessful UploadFileStatus = "PartiallySuccessful"
)

type FileMetaStatus = int8

const (
	FileMetaStatusUnknown FileMetaStatus = 0
	FileMetaStatusEnable  FileMetaStatus = 1
	FileMetaStatusDeleted FileMetaStatus = 2
)

var FileMetaStatusNameMap = map[FileMetaStatus]string{
	FileMetaStatusUnknown: "未知",
	FileMetaStatusEnable:  "生效中",
	FileMetaStatusDeleted: "已删除",
}

type FileSizeUnit int

const (
	FileSizeUnitSI  FileSizeUnit = 1000
	FileSizeUnitIEC FileSizeUnit = 1024
)
