package mysql

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/13 下午10:52
 * @FilePath: db
 * @Description:
 */

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/micro_services/file_server/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := conf.GetConf().MySQL.DSN
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	klog.Info("db connect success! db: ", db)
}

func DB() *gorm.DB {
	return db
}
