package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	file_server "github.com/cutejiuges/disk_back/kitex_gen/file_server/diskbackservice"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/cache/redis"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/dal/mysql"
	"github.com/cutejiuges/disk_back/micro_services/file_server/conf"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/enum"
	"github.com/cutejiuges/disk_back/micro_services/file_server/internal/util/localutil"
	"github.com/cutejiuges/disk_back/middleware"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	opts := kitexInit()
	svr := file_server.NewServer(new(DiskBackServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	//服务地址
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	//服务名
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	//微服务注册
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	//thrift meta handler类型
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	//日志收集
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName + localutil.FormatTime(time.Now(), string(enum.TimeLayoutDateYYMMDD)),
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackup,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
	}
	//同时输出到日志文件和屏幕
	multiWriter := io.MultiWriter(asyncWriter, os.Stdout)
	klog.SetOutput(multiWriter)
	server.RegisterShutdownHook(func() {
		_ = asyncWriter.Sync()
	})

	//日志中间件
	opts = append(opts, server.WithMiddleware(middleware.LogMW))

	//连接数据库
	mysql.Init()
	redis.Init()
	return
}
