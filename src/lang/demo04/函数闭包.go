package main

import "fmt"

/*
*
闭包就是一个函数和与其相关的引用环境组合的一个整体
*/
func main() {
	// 匿名函数
	sum := func(num1 int, num2 int) int {
		return num1 + num2
	}(10, 20)
	fmt.Println("匿名函数", sum)

	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
}

/** 闭包 */
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
