package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * context 对于派生goroutine有更强的控制力
 * timerCtx继承自cancelCtx
 */
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, _ := context.WithDeadline(context.Background(), 5 * time.Second)
	go handleRequest(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println(ctx.Err())
}

func writeRedis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeDatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}

func handleRequest(ctx context.Context) {
	go writeRedis(ctx)
	go writeDatabase(ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running")
			time.Sleep(2 * time.Second)
		}
	}
}
