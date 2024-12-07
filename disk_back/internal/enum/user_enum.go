package enum

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午9:49
 * @FilePath: user_enum
 * @Description: 用于用户行为的枚举
 */

type UserStatus = int8

const (
	UserStatusEnable  UserStatus = 1 //生效中
	UserStatusDeleted UserStatus = 2 //已注销
)

var UserStatusNameMap = map[UserStatus]string{
	UserStatusEnable:  "生效中",
	UserStatusDeleted: "已注销",
}
