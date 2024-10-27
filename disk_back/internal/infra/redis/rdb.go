package redis

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/conf"
	"github.com/redis/go-redis/v9"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 上午12:50
 * @FilePath: rdb
 * @Description: 初始化redis
 */

var rdb *redis.Client

func Init() {
	var err error
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	klog.Info("connect redis success: ", rdb)
}

func DB() *redis.Client {
	return rdb
}
