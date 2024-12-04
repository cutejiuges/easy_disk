package main

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/13 下午11:11
 * @FilePath: generate
 * @Description:
 */

import (
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"strings"
)

func main() {
	mysql.Init()
	db := mysql.DB()

	g := gen.NewGenerator(gen.Config{
		OutPath:          "micro_services/user_server/biz/dal/model/query",
		Mode:             gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithTypeTag: true,
	})
	g.UseDB(db)

	dataMap := map[string]func(columnType gorm.ColumnType) (dateType string){
		"int": func(columnType gorm.ColumnType) (dateType string) {
			if n, ok := columnType.Nullable(); ok && n {
				return "*int64"
			}
			return "int64"
		},
		"tinyint": func(columnType gorm.ColumnType) (dateType string) {
			ct, _ := columnType.ColumnType()
			if strings.HasPrefix(ct, "tinyint(1)") {
				return "bool"
			}
			return "int8"
		},
	}
	g.WithDataTypeMap(dataMap)

	g.ApplyBasic(
		g.GenerateModelAs("tb_user", "User"),
	)
	g.ApplyInterface(func() {})
	g.Execute()
}
