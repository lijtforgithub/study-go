package main

import "fmt"

/**
接口是隐式关系 必须实现接口定义的全部方法
*/

type Hello interface {
	SayHello()
}

type Chinese struct {
}

func (c Chinese) SayHello() {
	fmt.Println("你好")
}

type American struct {
}

func (a American) SayHello() {
	fmt.Println("hi")
}

func greet(h Hello) {
	h.SayHello()
}

func main() {
	c := Chinese{}
	a := American{}
	greet(c)
	greet(a)
}
