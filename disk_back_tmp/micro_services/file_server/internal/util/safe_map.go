package util

import "sync"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/26 下午7:12
 * @FilePath: safe_map
 * @Description: 并发安全的map结构
 */

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex //读写锁
	data map[K]V      //泛型map结构
}

// NewSafeMap 创建一个并发安全的map
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

// Put 在安全map中写入一对数据
func (m *SafeMap[K, V]) Put(key K, val V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

// Get 从安全map中读一组值
func (m *SafeMap[K, V]) Get(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.data[key]
	return val, ok
}

// Len 获取安全map的大小
func (m *SafeMap[K, V]) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Clear 清空安全map
func (m *SafeMap[K, V]) Clear() {
	m.mu.Lock()
	m.mu.Unlock()
	m.data = make(map[K]V)
}

// GetData 获取安全map中的所有键值对
func (m *SafeMap[K, V]) GetData() map[K]V {
	m.mu.RLock()
	m.mu.RUnlock()
	return m.data
}
