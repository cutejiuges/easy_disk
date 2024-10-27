package ping_handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cutejiuges/disk_api/biz/model/disk_common"
	"github.com/cutejiuges/disk_api/biz/service"
	"github.com/cutejiuges/disk_api/infra/localutils"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/8 下午11:01
 * @FilePath: ping
 * @Description:
 */

// Ping .
// @router /ping [GET]
func Ping(ctx context.Context, c *app.RequestContext) {
	var err error
	var req disk_common.PingRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := service.ProcessPing(ctx, &req)
	localutils.Wrapper(c, resp, err)
}
