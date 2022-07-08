package main

import (
	"fmt"
)

var s = "S"

func main() {

	fmt.Println("Hello Word!")

	var a int
	a = 1

	var b int = 2
	var c = 3
	// 初始化声明 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误 只能在函数体中出现
	d := true
	e := "xxoo"

	fmt.Println(a + b + c)
	fmt.Println(d, e)

	e = "XXOO"
	fmt.Println(e)

	var f, g = "F", "G"
	h, i := "H", "I"
	fmt.Println(f, g, h, i)

	var (
		j int
		k bool
	)

	fmt.Println(s, j, k)

	// 内存地址
	fmt.Println(&s)

	// 交换
	s, h = h, s
	fmt.Println(s, h)

	// 舍弃 5
	_, b = 5, 7
	fmt.Println(b)

	ch := 'c'
	fmt.Printf("%d\t%c", ch, ch)

}
