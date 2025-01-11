package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * context 对于派生goroutine有更强的控制力
 */
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go handleRequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("It's time to stop all sub goroutines!")
	cancel()
	fmt.Println("ERROR = ", ctx.Err())

	time.Sleep(5 * time.Second)
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
