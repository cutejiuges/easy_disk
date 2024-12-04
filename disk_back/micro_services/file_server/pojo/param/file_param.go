package param

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/17 上午12:07
 * @FilePath: file_param
 * @Description: 文件元信息相关参数
 */

// QueryFileMetaParam 查询文件元信息参数
type QueryFileMetaParam struct {
	ID            int64
	FileKey       string
	IdList        []int64
	MinFileSize   int64
	MaxFileSize   int64
	MinCreateTime string
	MaxCreateTime string
	Status        []int8
	Page          int
	Size          int
}

// EditFileMetaParam 编辑文件信息参数
type EditFileMetaParam struct {
	ID       int64
	FileKey  string
	IdList   []int64
	FileAddr string
	RefNum   int64
	RefDealt int64
	Status   int8
}
