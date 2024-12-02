package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func worker(ctx context.Context, id int) error {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: canceled\n", id)
			return ctx.Err()
		default:
			fmt.Printf("Worker %d: working\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var g errgroup.Group

	// 创建一个带有取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动多个工作协程
	for i := 0; i < 5; i++ {
		i := i // 捕获循环变量
		g.Go(func() error {
			return worker(ctx, i)
		})
	}

	// 模拟一个错误发生的情况
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Main: canceling workers")
		cancel()
	}()

	// 等待所有工作协程完成或出现错误
	if err := g.Wait(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}
}
