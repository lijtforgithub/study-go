package main

import "fmt"

/*
	Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。
*/

type Student struct {
	name string
}

// 值传递
func (c Student) method1(newName string) {
	fmt.Println("值传递", newName)
	c.name = newName
}

// 指针传递
func (c *Student) method2(newName string) {
	fmt.Println("指针传递", newName)
	c.name = newName
}

func (c Student) String() string {
	return "String方法 " + c.name
}

func main() {

	stu := Student{"小小"}
	fmt.Println(stu.name)

	stu.method1("小值")
	fmt.Println(stu.name)

	stu.method2("小指")
	fmt.Println(stu)

}
