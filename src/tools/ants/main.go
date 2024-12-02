package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func task(i int) {
	fmt.Printf("Running task %d\n", i)
	time.Sleep(1 * time.Second) // 模拟任务处理时间
}

func main() {
	//newPool()
	newPoolWithFunc()
}

func newPool() {
	// 创建一个新的线程池，最大并发数为 10
	pool, err := ants.NewPool(10)
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	// 提交任务到线程池
	for i := 0; i < 20; i++ {
		if err := pool.Submit(func() {
			task(i)
		}); err != nil {
			fmt.Println(err)
		}
	}

	// 等待所有任务完成
	time.Sleep(3 * time.Second)
}

func newPoolWithFunc() {
	// 创建一个新的线程池，最大并发数为 10，任务超时时间为 5 秒
	pool, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		task(i.(int))
	}, ants.WithExpiryDuration(5*time.Second))
	if err != nil {
		panic(err)
	}
	defer pool.Release()

	// 提交任务到线程池
	for i := 0; i < 20; i++ {
		if err := pool.Invoke(i); err != nil {
			fmt.Println(err)
		}
	}

	// 等待所有任务完成
	time.Sleep(3 * time.Second)
}
