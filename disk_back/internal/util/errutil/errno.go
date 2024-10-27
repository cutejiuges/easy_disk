package errutil

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 上午12:35
 * @FilePath: errno
 * @Description: 定义错误信息
 */

type ErrCode = int32

const (
	ErrCodeOK           ErrCode = 0
	ErrCodeCommon       ErrCode = -1
	ErrCodeInternal     ErrCode = 1000
	ErrCodeParamIllegal ErrCode = 1001

	ErrCodeNotLogin    ErrCode = 2000
	ErCodeNoPermission ErrCode = 2001
	ErrCodeNotAllowed  ErrCode = 2002

	ErrCodeDbUnknownError      ErrCode = 3000
	ErrCodeDbTimeout           ErrCode = 3001
	ErrCodeDbSessionCommitFail ErrCode = 3002
)

var ErrMsgMap = map[int32]string{
	ErrCodeOK:                  "OK",
	ErrCodeCommon:              "系统错误",
	ErrCodeInternal:            "系统内部异常",
	ErrCodeParamIllegal:        "参数错误",
	ErrCodeNotLogin:            "用户未登录",
	ErCodeNoPermission:         "用户无权限",
	ErrCodeNotAllowed:          "不允许执行该操作",
	ErrCodeDbUnknownError:      "数据库未知异常",
	ErrCodeDbTimeout:           "数据库超时",
	ErrCodeDbSessionCommitFail: "数据库连接失败",
}
