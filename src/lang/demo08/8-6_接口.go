package main

import "fmt"

/**
接口是隐式关系 必须实现接口定义的全部方法
结构可以嵌套结构体或者接口，接口只能嵌套接口
*/

type Hello interface {
	SayHello()
}

type NewHello interface {
	Hello
	newMethod()
}

type anyType interface {
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

// 断言
func (c Chinese) niuYangGe() {
	fmt.Println("扭秧歌")
}

func (a American) disco() {
	fmt.Println("迪斯科")
}

func greet2(h Hello) {
	h.SayHello()

	//if ch, flag := h.(Chinese); flag {
	//	ch.niuYangGe()
	//} else {
	//	fmt.Println("不是中国人")
	//}

	switch h.(type) {
	case Chinese:
		c := h.(Chinese)
		c.niuYangGe()
	case American:
		a := h.(American)
		a.disco()
	}
}

func main() {
	c := Chinese{}
	a := American{}
	greet(c)
	greet(a)

	fmt.Println()

	greet2(a)
	greet2(c)
}
