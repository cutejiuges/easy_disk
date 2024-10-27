package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cutejiuges/disk_api/conf"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_back/diskbackservice"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 上午12:00
 * @FilePath: disk_back
 * @Description: 生成disk_back的客户端，以发起rpc调用
 */

var cli diskbackservice.Client

const diskbackservice_name string = "cutejiuges.practice.disk_back"

func initDiskBackClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Resolver.ResolverAddress[diskbackservice_name])
	if err != nil {
		hlog.Fatalf("initDiskBackClient error: %v", err)
		return
	}
	cli = diskbackservice.MustNewClient(
		diskbackservice_name,
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithRPCTimeout(time.Second*5),
	)
}

func GetDiskBackClient() diskbackservice.Client {
	return cli
}
