package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(2)
	intChan := make(chan int, num)

	go func(len int) {
		for i := 0; i < len; i++ {
			fmt.Println("写入 ", i)
			intChan <- i
			time.Sleep(time.Second)
		}
		close(intChan)
		wg.Done()
	}(num)

	go func() {
		for v := range intChan {
			fmt.Println("读出 ", v)
			//time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
}
