package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex = &sync.Mutex{}

	go syncMethod(0, mutex)
	go syncMethod(1, mutex)
	go syncMethod(2, mutex)

	time.Sleep(time.Second * 30)
}

func syncMethod(no int, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("开始执行逻辑", no)
	time.Sleep(time.Second * 3)
	fmt.Println("结束执行逻辑", no)
}
