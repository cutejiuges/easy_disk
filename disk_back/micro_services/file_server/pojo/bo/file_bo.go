package bo

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/17 上午12:01
 * @FilePath: file_bo
 * @Description:
 */

type SimpleFile struct {
	ID   int64  //文件id
	Addr string //文件存储地址
	Key  string //文件key
	Msg  string //补充说明信息
}
