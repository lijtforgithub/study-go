package main

import (
	"fmt"
	"time"
)

func main() {
	num := 10
	intChan := make(chan int, num)

	for i := 0; i < num; i++ {
		intChan <- i
	}

	read := func() {
		for v := range intChan {
			fmt.Println(v)
		}
	}

	// 不执行关闭 for range会一直读报错 使用协程不会报错
	//close(intChan)
	//read()

	go read()
	time.Sleep(5 * time.Second)

}
