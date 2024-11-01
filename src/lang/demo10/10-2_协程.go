package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go func(msg int) {
			fmt.Println(msg)
			time.Sleep(time.Second)
		}(i)
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 5)
}
