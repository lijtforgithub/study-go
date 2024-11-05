package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwLock sync.RWMutex

func read() {
	defer wg.Done()
	rwLock.RLock()
	fmt.Println("读取数据开始")
	time.Sleep(time.Second)
	fmt.Println("读取数据结束")
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	rwLock.Lock()
	fmt.Println("写数据开始")
	time.Sleep(time.Second * 3)
	fmt.Println("写数据结束")
	rwLock.Unlock()
}

func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}

	go write()

	wg.Wait()
}
