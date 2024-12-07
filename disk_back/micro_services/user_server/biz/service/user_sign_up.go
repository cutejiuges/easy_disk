package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/enum"
	"github.com/cutejiuges/disk_back/internal/errno"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/cache"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/mysql"
	"github.com/cutejiuges/disk_back/micro_services/user_server/infra/pojo"
	"gorm.io/gorm"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午8:51
 * @FilePath: user_sign_up
 * @Description: 执行用户注册行为
 */

func ProcessUserSignUp(ctx context.Context, req *user_server.UserSignUpRequest) (*user_server.UserSignUpData, error) {
	data := user_server.NewUserSignUpData()

	//检查验证码是否正确
	code, err := cache.GetVerifyCode(ctx, req.GetUserInfo().GetEmail())
	if err != nil {
		return data, fmt.Errorf("请检查验证码有效性")
	}
	if code != req.GetVerifyCode() {
		return data, &errno.BizError{
			Code: errno.ErrCodeCommon,
			Msg:  "验证码错误",
		}
	}

	//检查用户是否已经注册过
	userInfo, err := dao.QuerySingleUser(ctx, &pojo.UserQueryParam{Email: req.GetUserInfo().GetEmail()})
	//如果出现了记录不存在之外的其他错误，需要报错
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		klog.CtxErrorf(ctx, "service.ProcessUserSignUp -> dao.QuerySingleUser error: %v", err)
		return data, err
	}
	//如果记录存在
	if userInfo != nil {
		data.SetAccountId(thrift.Int64Ptr(userInfo.ID))
		//未注销，返回账号信息
		if userInfo.Status != enum.UserStatusDeleted {
			return data, fmt.Errorf("邮箱已注册，请使用邮箱或账号登陆")
		} else { //记录存在但已注销，恢复记录
			qry := query.Use(mysql.DB())
			err := dao.UpdateUserInfo(ctx, qry, &pojo.UserEditParam{Id: userInfo.ID, EditStatus: enum.UserStatusEnable})
			if err != nil {
				klog.CtxErrorf(ctx, "service.ProcessUserSignUp -> dao.UpdateUserInfo error: %v", err)
				return data, err
			}
			return data, nil
		}
	}
	//记录不存在需要注册
	qry := query.Use(mysql.DB())
	user := &model.User{
		UserName: req.GetUserInfo().GetUserName(),
		Password: req.GetUserInfo().GetPassword(),
		Email:    req.GetUserInfo().GetEmail(),
		Profile:  req.GetUserInfo().GetProfile(),
		Status:   enum.UserStatusEnable,
	}
	err = dao.CreateUser(ctx, qry, user)
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessUserSignUp -> dao.CreateUser error: %v", err)
		return data, err
	}

	data.SetAccountId(thrift.Int64Ptr(user.ID))
	return data, nil
}
