package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(msg int) {
			fmt.Println(msg)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
