package main

import "fmt"

func main() {
	strChan := make(chan string, 5)
	strChan <- "A"
	strChan <- "B"
	strChan <- "C"

	// 管道关闭之后只能读不能写
	close(strChan)
	//strChan <- "D"

	fmt.Println(<-strChan)
	fmt.Println(<-strChan)
}
