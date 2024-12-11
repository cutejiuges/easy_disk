package service

import (
	"context"
	"encoding/json"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/dao"
	"github.com/cutejiuges/disk_back/micro_services/user_server/infra/pojo"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/10 上午8:51
 * @FilePath: user_sign_in
 * @Description: 用户登陆执行
 */

func ProcessUserSignIn(ctx context.Context, req *user_server.UserSignInRequest) (*user_server.UserSignInData, error) {
	data := user_server.NewUserSignInData()

	//直接使用用户身份进行查询，能查询到用户就返回身份信息，查不到说明无此用户
	user, err := dao.QuerySingleUser(ctx, &pojo.UserQueryParam{
		Id:       req.GetAccountId(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessUserSignIn -> dao.QuerySingleUser error: %v", err)
		return data, err
	}

	var profile user_server.UserProfile
	err = json.Unmarshal([]byte(user.Profile), &profile)
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessUserSignIn -> json.Unmarshal error: %v", err)
		return data, err
	}

	data.SetUserName(thrift.StringPtr(user.UserName))
	data.SetAdmin(profile.Admin)
	data.SetEmail(thrift.StringPtr(user.Email))
	data.SetAccountId(thrift.Int64Ptr(user.ID))
	return data, nil
}
