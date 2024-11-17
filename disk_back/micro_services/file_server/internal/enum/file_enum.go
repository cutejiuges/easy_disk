package enum

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午4:33
 * @FilePath: upload_file
 * @Description:
 */

type OperateFileStatus = int8

const (
	OperateFileStatusSuccess             OperateFileStatus = 1
	OperateFileStatusFailed              OperateFileStatus = 2
	OperateFileStatusPartiallySuccessful OperateFileStatus = 3
)

var OperateFileStatusMap = map[int8]string{
	OperateFileStatusSuccess:             "文件操作成功",
	OperateFileStatusFailed:              "文件操作失败",
	OperateFileStatusPartiallySuccessful: "文件操作部分成功",
}

type FileMetaStatus = int8

const (
	FileMetaStatusEnable  FileMetaStatus = 1
	FileMetaStatusDeleted FileMetaStatus = 2
)

var FileMetaStatusNameMap = map[FileMetaStatus]string{
	FileMetaStatusEnable:  "生效中",
	FileMetaStatusDeleted: "已删除",
}

type FileSizeUnit int

const (
	FileSizeUnitSI  FileSizeUnit = 1000
	FileSizeUnitIEC FileSizeUnit = 1024
)
