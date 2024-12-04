package cache

import (
	"context"
	"fmt"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/cache/redis"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 上午1:31
 * @FilePath: id_gen
 * @Description: 生成全局唯一id
 */

const globalFileID = "globalFileID"

func GetCurUniqID(ctx context.Context) (int64, error) {
	return redis.DB().Get(ctx, globalFileID).Int64()
}

func GetNextUniqID(ctx context.Context) (int64, error) {
	return redis.DB().Incr(ctx, globalFileID).Result()
}

func UploadFileLockKey(key string) string {
	return fmt.Sprintf("{%s}:%s", "uploadFileSha1Key", key)
}

func DeleteFileLockKey(id int64) string {
	return fmt.Sprintf("{%s}:%d", "deleteFileId", id)
}
