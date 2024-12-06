package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cutejiuges/disk_api/conf"
	"github.com/cutejiuges/disk_back/kitex_gen/user_server/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/6 上午12:59
 * @FilePath: user_back
 * @Description: 生成user_server的客户端，发起rpc调用
 */

var userCli userservice.Client

const user_server_name = "cutejiuges.server.user"

func initUserServerClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Resolver.ResolverAddress[user_server_name])
	if err != nil {
		hlog.Fatalf("initUserServerClient error: %v", err)
		return
	}

	userCli = userservice.MustNewClient(
		user_server_name,
		client.WithResolver(r),
		client.WithRPCTimeout(5*time.Second),
	)
}

func GetUserServerClient() userservice.Client {
	return userCli
}
