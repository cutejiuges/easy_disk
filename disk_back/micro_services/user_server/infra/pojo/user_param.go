package pojo

import "github.com/cutejiuges/disk_back/internal/enum"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午9:34
 * @FilePath: user_param
 * @Description: user查询的参数
 */

type UserQueryParam struct {
	Id       int64
	IdList   []int64
	UserName string
	Password string
	Email    string
	Phone    string
	Status   enum.UserStatus
}

type UserEditParam struct {
	//检索条件
	Id          int64
	IdList      []int64
	Email       string
	QueryStatus enum.UserStatus
	//更新内容
	Phone      string
	Profile    string
	Password   string
	UserName   string
	EditStatus enum.UserStatus
}
