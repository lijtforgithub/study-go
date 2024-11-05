package main

import "fmt"

func main() {
	rwChan := make(chan string, 1)
	rwChan <- "默认可读可写"
	fmt.Println(<-rwChan)

	// 只读
	rChan := make(<-chan string, 1)
	//rChan <- "只读"
	fmt.Println(rChan)

	// 只写
	wChan := make(chan<- string, 1)
	wChan <- "只写"
	//fmt.Println(<- wChan)
}
