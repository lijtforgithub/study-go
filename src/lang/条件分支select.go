package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
	select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
	如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。
	*/
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	c4 := make(chan string)
	c5 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c4 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c5 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c4:
			fmt.Println("received", msg1)
		case msg2 := <-c5:
			fmt.Println("received", msg2)
		}
	}

}
