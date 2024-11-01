package main

import (
	"fmt"
	"time"
)

func loopPrint(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, "-", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go loopPrint("go")

	// 主死从随
	for i := 0; i < 5; i++ {
		fmt.Println("main-", i)
		time.Sleep(time.Second)
	}

}
