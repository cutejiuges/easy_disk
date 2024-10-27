package service

import (
	"context"
	"fmt"
	"github.com/cutejiuges/disk_back/kitex_gen/disk_common"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/1 下午7:22
 * @FilePath: ping
 * @Description:
 */

func ProcessPing(ctx context.Context, req *disk_common.PingRequest) (string, error) {
	reply := fmt.Sprintf("hello, %s!", req.GetMsg())
	return reply, nil
}
