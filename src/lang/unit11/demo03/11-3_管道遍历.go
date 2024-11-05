package main

import "fmt"

func main() {
	num := 10
	intChan := make(chan int, num)

	for i := 0; i < num; i++ {
		intChan <- i
	}

	// 不执行关闭 for range会一直读 报错
	close(intChan)

	for v := range intChan {
		fmt.Println(v)
	}
}
