package main

import (
	"fmt"
	"time"
)

func process(ch chan int) {
	time.Sleep(time.Second)
	// 在管道中写入一个元素表示当前协程已结束
	ch <- 1
}

/**
 * 使用channel控制字协程的有点事实现简单 缺点是当需要大量创建协程时就需要有相同数量的channel，而且对于子协程继续派生出来的协程不方便控制
 */
func main() {
	// 创建一个包含10个元素的切边 元素类型为channel
	channels := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go process(channels[i])
	}

	for i, ch := range channels {
		<-ch
		fmt.Println("Routine ", i, " quit!")
	}
}
