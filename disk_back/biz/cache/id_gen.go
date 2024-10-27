package cache

import (
	"context"
	"github.com/cutejiuges/disk_back/internal/infra/redis"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 上午1:31
 * @FilePath: id_gen
 * @Description: 生成全局唯一id
 */

const globalUniqID = "globalUniqID"

func GetCurUniqID(ctx context.Context) (int64, error) {
	return redis.DB().Get(ctx, globalUniqID).Int64()
}

func GetNextUniqID(ctx context.Context) (int64, error) {
	return redis.DB().Incr(ctx, globalUniqID).Result()
}
