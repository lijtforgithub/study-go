package main

import (
	"context"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"time"
)

func main() {
	// 配置选项
	config := bigcache.Config{
		Shards:             1024,            // 分片数量，默认值为 1024
		LifeWindow:         5 * time.Minute, // 缓存条目的生命周期，默认值为永不过期
		CleanWindow:        5 * time.Minute, // 清理过期条目的间隔时间，默认值为 5 分钟
		Verbose:            false,           // 是否启用详细日志，默认值为 false
		MaxEntrySize:       500,             // 单个条目最大大小（字节），默认值为 500 字节
		HardMaxCacheSize:   8192,            // 最大缓存分片大小，默认值为 8192
		OnRemoveWithReason: nil,             // 条目移除时的回调函数，默认值为 nil
	}

	config.OnRemoveWithReason = func(key string, entry []byte, reason bigcache.RemoveReason) {
		fmt.Printf("Removed key %s due to %v\n", key, reason)
	}

	ctx := context.Background()

	// 创建一个新的 bigcache 实例
	cache, err := bigcache.New(ctx, config)
	if err != nil {
		panic(err)
	}
	defer cache.Close()

	// 设置缓存条目
	key := "user-123"
	value := []byte("John Doe")
	err = cache.Set(key, value)
	if err != nil {
		fmt.Println("Set error:", err)
	}

	// 获取缓存条目
	result, err := cache.Get(key)
	if err == nil {
		fmt.Printf("Get %s: %s\n", key, string(result))
	} else {
		fmt.Println("Get error:", err)
	}

	// 删除缓存条目
	err = cache.Delete(key)
	if err == nil {
		fmt.Printf("Deleted %s from cache\n", key)
	} else {
		fmt.Println("Delete error:", err)
	}
}
