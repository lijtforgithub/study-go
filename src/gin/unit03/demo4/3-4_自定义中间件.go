package main

import "fmt"

func main() {
	var arr [3]func(s string)
	arr[0] = func(s string) {
		fmt.Println(s)
	}
	arr[0]("xxoo")
}
