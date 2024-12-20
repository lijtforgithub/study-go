package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

// go get -u github.com/go-redis/redis/v8

func main() {
	rdb := createConnect()
	ctx := rdb.Context()

	// 设置一个键值对
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键值对
	val1 := rdb.Get(ctx, "key").Val()
	fmt.Println("key:", val1)

	// 获取不存在的键值对
	val2, err := rdb.Get(ctx, "key2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2:", val2)
	}

	// 删除一个键值对
	err = rdb.Del(ctx, "key").Err()
	if err != nil {
		panic(err)
	}

	// 检查键是否存在
	exists, err := rdb.Exists(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	if exists == 0 {
		fmt.Println("key does not exist")
	} else {
		fmt.Println("key exists")
	}

	// 关闭连接
	if err := rdb.Close(); err != nil {
		panic(err)
	}
}

func createConnect() *redis.Client {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // 密码（如果需要）
		DB:       0,                // 默认数据库
		PoolSize: 10,
	})

	s, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis连接成功", s)
	return rdb
}
