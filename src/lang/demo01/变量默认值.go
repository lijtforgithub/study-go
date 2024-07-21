package main

import (
	"fmt"
)

func main() {

	var a *int
	var b []int
	var c map[string]int
	var d chan int
	var e func(string) int
	var f error // error 是接口
	var g bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}
