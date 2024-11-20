package cache

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cutejiuges/disk_back/micro_services/file_server/biz/cache/redis"
	"math/rand"
	"sync"
	"time"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/17 下午11:02
 * @FilePath: distributed_lock
 * @Description: 基于redis实现全局分布式锁，带watch_dog机制
 */

const (
	expireTime = time.Second * 30
	charset    = "abcdefghijklmnopqrstuvwxyz0123456789"
)

type DistributedLock struct {
	mutex sync.Mutex
	key   string
	value string
}

func NewDistributedLock(key string) *DistributedLock {
	return &DistributedLock{
		key:   key,
		value: generateRandomValue(),
	}
}

func generateRandomValue() string {
	bytes := make([]byte, 16)
	for i := range bytes {
		bytes[i] = charset[rand.Intn(len(charset))]
	}
	return string(bytes)
}

func (dl *DistributedLock) startWatchDog() {
	ticker := time.NewTicker(expireTime / 3)
	defer ticker.Stop()

	for range ticker.C {
		err := redis.DB().Expire(context.Background(), dl.key, expireTime).Err()
		if err != nil {
			klog.Error("锁续期失败， error: ", err)
			break
		}
	}
}

func (dl *DistributedLock) tryLock() bool {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	res, err := redis.DB().SetNX(context.Background(), dl.key, dl.value, expireTime).Result()
	if err != nil {
		return false
	}
	if res {
		go dl.startWatchDog()
	}
	return res
}

func (dl *DistributedLock) lock() bool {
	if !dl.tryLock() {
		time.Sleep(100 * time.Millisecond)
	}
	return true
}

func (dl *DistributedLock) unlock() bool {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	val, err := redis.DB().Get(context.Background(), dl.key).Result()
	if err != nil || val != dl.value {
		return false
	}

	_, err = redis.DB().Del(context.Background(), dl.key).Result()
	if err != nil {
		return false
	}
	return true
}

func (dl *DistributedLock) TryLock() bool {
	return dl.tryLock()
}

func (dl *DistributedLock) Lock() bool {
	return dl.lock()
}

func (dl *DistributedLock) Unlock() bool {
	return dl.unlock()
}
