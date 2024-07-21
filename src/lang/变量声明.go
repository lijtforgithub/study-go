package main

import (
	"fmt"
)

// 全局变量是允许声明但不使用的；局部变量声明必须使用，否则编译错误
var s = "S"

func main() {

	var bl bool
	bl = false
	fmt.Println("bl = ", bl)

	// 一行代表一个语句结束，不需要像 C 家族中的其它语言一样以分号 ; 结尾。如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分
	fmt.Println("Hello Word!")

	// 指定变量类型 先声明后赋值 有默认值
	var a int
	fmt.Println("a = ", a)
	a = 1
	fmt.Println("a = ", a)

	// 声明和赋值一起
	var b int = 2
	// 根据值自行判定变量类型
	var c = 3
	// 初始化声明 省略var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误；这种不带声明格式的只能在函数体中出现
	d := true
	e := "xxoo"
	fmt.Println(a + b + c)
	fmt.Println(d, e)

	e = "XXOO"
	fmt.Println(e)

	// 多变量声明
	var f, g = "F", "G"
	h, i := "H", "I"
	fmt.Println(f, g, h, i)

	// 这种因式分解关键字的写法一般用于声明全局变量
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

	// 空白标识符 _ 也被用于抛弃值 舍弃 5 _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
	_, b = 5, 7
	fmt.Println(b)

	var ss string = `很长的字符串
包含换行的
可以用这个符号`
	fmt.Println(ss)

	ss = "加号拼接的字符串" +
		"；加号要放在每行的末尾" +
		"；不然效果你试试"
	fmt.Println(ss)

}
