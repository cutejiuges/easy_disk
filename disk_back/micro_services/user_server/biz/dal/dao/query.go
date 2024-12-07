package dao

import (
	"context"
	"fmt"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/mysql"
	"github.com/cutejiuges/disk_back/micro_services/user_server/infra/pojo"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午9:37
 * @FilePath: query
 * @Description: 用户的查询操作
 */

func QuerySingleUser(ctx context.Context, param *pojo.UserQueryParam) (*model.User, error) {
	if param.UserName == "" && param.Id <= 0 && param.Email == "" && param.Phone == "" {
		return nil, fmt.Errorf("无有效查询条件")
	}
	q := query.Use(mysql.DB()).User
	dao := q.WithContext(ctx)
	if param.UserName != "" {
		dao = dao.Where(q.UserName.Eq(param.UserName))
	}
	if param.Id > 0 {
		dao = dao.Where(q.ID.Eq(param.Id))
	}
	if param.Email != "" {
		dao = dao.Where(q.Email.Eq(param.Email))
	}
	if param.Phone != "" {
		dao = dao.Where(q.Phone.Eq(param.Phone))
	}
	return dao.First()
}
