package util

import (
	"math/rand"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/6 上午12:20
 * @FilePath: random_util
 * @Description: 随机数生成
 */

func GenerateVerifyCode() int64 {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	return r.Int63n(900000) + 100000
}
