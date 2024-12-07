package cache

import (
	"context"
	"fmt"
	"github.com/cutejiuges/disk_back/micro_services/user_server/biz/cache/redis"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/6 上午12:32
 * @FilePath: user_cache
 * @Description:
 */

// GenVerifyCodeKey 生成用户注册验证码的key
func GenVerifyCodeKey(email string) string {
	return fmt.Sprintf("{verify_code}:%s", email)
}

// SaveVerifyCode 存储用户验证码
func SaveVerifyCode(ctx context.Context, email string, code int64) error {
	key := GenVerifyCodeKey(email)
	return redis.DB().SetEx(ctx, key, code, time.Second*60).Err()
}

// GetVerifyCode 获取用户验证码
func GetVerifyCode(ctx context.Context, email string) (int64, error) {
	key := GenVerifyCodeKey(email)
	return redis.DB().Get(ctx, key).Int64()
}
