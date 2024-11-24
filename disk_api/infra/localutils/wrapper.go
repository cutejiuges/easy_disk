package localutils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cutejiuges/disk_back/kitex_gen/file_server"
	"net/url"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 上午1:29
 * @FilePath: wrapper
 * @Description:
 */

type commonResp struct {
	StatusCode int32
	StatusMsg  string
}

func Wrapper(c *app.RequestContext, data any, err error) {
	if err != nil {
		c.JSON(consts.StatusOK, commonResp{
			StatusCode: int32(-100),
			StatusMsg:  "服务调用异常: " + err.Error(),
		})
		return
	}
	switch data.(type) {
	case *file_server.DownloadFileResponse:
		file := data.(*file_server.DownloadFileResponse).Data
		c.Response.Header.Set("Content-Disposition", "attachment; filename="+url.QueryEscape(file.GetFileName()))
		c.Response.Header.Set("Content-Type", "application/octet-stream")
		c.Response.SetBodyRaw(file.Content)
		return
	}
	c.JSON(consts.StatusOK, data)
}
