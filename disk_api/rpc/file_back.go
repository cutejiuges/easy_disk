package rpc

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cutejiuges/disk_api/conf"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server/fileservice"
	consul "github.com/kitex-contrib/registry-consul"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 上午12:00
 * @FilePath: disk_back
 * @Description: 生成file_server的客户端，以发起rpc调用
 */

var fileCli fileservice.Client

const file_server_name string = "cutejiuges.server.file"

func initFileServerClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Resolver.ResolverAddress[file_server_name])
	if err != nil {
		hlog.Fatalf("initFileServerClient error: %v", err)
		return
	}
	fileCli = fileservice.MustNewClient(
		file_server_name,
		client.WithResolver(r),
		client.WithRPCTimeout(time.Second*5),
	)
}

func GetFileServerClient() fileservice.Client {
	return fileCli
}
