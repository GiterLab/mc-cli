package main

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var mc *memcache.Client

// 初始化连接
func MemcacheInit(url string) {
	// init mc
	mc = memcache.New(url)
}

// 获取cache信息
func MemcacheGet(key string) ([]byte, error) {
	it, err := mc.Get(key)
	if err != nil {
		return nil, err
	}

	return it.Value, err
}

// 获取cache更多的信息
func MemcacheGetMore(key string) (*memcache.Item, error) {
	it, err := mc.Get(key)
	if err != nil {
		return nil, err
	}

	return it, err
}

// 设置cache信息
func MemcacheSet(key string, value []byte) error {
	item := &memcache.Item{
		Key:   key,
		Value: value,
	}
	return mc.Set(item)
}

// 设置cache信息，带时间过期
func MemcacheSetByExpired(key string, value []byte, expiration int32) error {
	item := &memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	return mc.Set(item)
}

// 删除key
func MemcacheDel(key string) error {
	return mc.Delete(key)
}
