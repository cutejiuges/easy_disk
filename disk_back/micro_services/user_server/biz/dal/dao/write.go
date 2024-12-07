package dao

import (
	"context"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/model"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/dal/model/query"
	"github.com/cutejiuges/disk_back/micro_services/user_server/infra/pojo"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午9:37
 * @FilePath: write
 * @Description:
 */

func CreateUser(ctx context.Context, q *query.Query, user *model.User) error {
	dao := q.WithContext(ctx).User
	return dao.Create(user)
}

func CreateUserInBatch(ctx context.Context, q *query.Query, users []*model.User, batch int) error {
	dao := q.WithContext(ctx).User
	return dao.CreateInBatches(users, batch)
}

func UpdateUserInfo(ctx context.Context, q *query.Query, param *pojo.UserEditParam) error {
	if len(param.IdList) <= 0 && param.Id <= 0 && param.Email == "" && param.QueryStatus <= int8(0) {
		return nil
	}
	qry := q.User
	dao := qry.WithContext(ctx)
	if param.QueryStatus != int8(0) {
		dao = dao.Where(qry.Status.Eq(param.QueryStatus))
	}
	if param.Id != 0 {
		dao = dao.Where(qry.ID.Eq(param.Id))
	}
	if param.Email != "" {
		dao = dao.Where(qry.Email.Eq(param.Email))
	}
	if len(param.IdList) > 0 {
		dao = dao.Where(qry.ID.In(param.IdList...))
	}
	if param.Password != "" {
		dao = dao.Select(qry.Password)
	}
	if param.Phone != "" {
		dao = dao.Select(qry.Phone)
	}
	if param.Profile != "" {
		dao = dao.Select(qry.Profile)
	}
	if param.EditStatus != int8(0) {
		dao = dao.Select(qry.Status)
	}
	if param.UserName != "" {
		dao = dao.Select(qry.UserName)
	}
	_, err := dao.Updates(model.User{
		UserName: param.UserName,
		Password: param.Password,
		Phone:    param.Email,
		Profile:  param.Profile,
		Status:   param.EditStatus,
	})
	return err
}
