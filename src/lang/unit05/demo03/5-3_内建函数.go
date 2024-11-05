package main

import "fmt"

func main() {
	// new 分配内存 new函数的实参是一个类型而不是具体值 new函数返回值是对应类型的指针
	ptr := new(int)
	fmt.Printf("类型=%T 值=%v 地址=%v 指针指向的值=%v", ptr, ptr, &ptr, *ptr)

}
