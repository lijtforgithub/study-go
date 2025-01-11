package main

import (
	"context"
	"fmt"
	"time"
)

/**
 * context 对于派生goroutine有更强的控制力
 * context查找不到key时，会向父节点查找，如果查询不到最终返回interface{}
 */
func main() {
	ctx := context.WithValue(context.Background(), "param1", "v1")
	go handleRequest(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println(ctx.Err())
}

func handleRequest(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("param1"))
			time.Sleep(2 * time.Second)
		}
	}
}
