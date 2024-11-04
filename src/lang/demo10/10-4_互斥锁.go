package main

import (
	"fmt"
	"sync"
)

var global = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add1w() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		global++
		lock.Unlock()
	}
	wg.Done()
}

func sub1w() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		global--
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add1w()
	go sub1w()
	wg.Wait()
	fmt.Println(global)
}
