package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/internal/util"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/cache"
	"github.com/cutejiuges/disk_back/micro_services/user_server/infra/email"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/5 下午11:37
 * @FilePath: get_email_verify_code
 * @Description: 向用户输入的邮箱中发送限时验证码
 */

func ProcessGetEmailVerifyCode(ctx context.Context, req *user_server.GetEmailVerifyCodeRequest) error {
	//先生成一个随机的6位整数作为验证码的内容
	code := util.GenerateVerifyCode()

	var err error
	//调用邮箱服务发送验证码
	err = email.SendVerifyCode(req.GetEmail(), code)
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessGetEmailVerifyCode -> email.SendVerifyCode error: %v", err)
		return err
	}

	//将code存储到redis中，时长1min
	err = cache.SaveVerifyCode(req.GetEmail(), code)
	if err != nil {
		klog.CtxErrorf(ctx, "service.ProcessGetEmailVerifyCode -> cache.SaveVerifyCode error: %v", err)
		return err
	}
	return nil
}
