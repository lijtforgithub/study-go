package main

import "fmt"

/**
 * 管道-数据结构本质队列
 */
func main() {
	intChan := make(chan int, 3)
	intChan <- 10
	intChan <- 20
	fmt.Printf("管道的地址=%v 容量=%d 长度=%d\n", intChan, cap(intChan), len(intChan))

	fmt.Println(<-intChan)
	intChan <- 30
	intChan <- 40
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

	// 没有使用协程的情况下 取出超过长度会报错
	//fmt.Println(<-intChan)
}
