package main

import "fmt"

func main() {
	fmt.Println("函数返回值 = ", testDefer(10, 20))
}

func testDefer(num1 int, num2 int) int {
	defer fmt.Println("num1 = ", num1)
	defer fmt.Println("num2 = ", num2)
	num1++
	num2++
	result := num1 + num2
	fmt.Println("result = ", result)
	// 发生异常 defer也会执行 执行之后抛出异常
	//i := 0
	//fmt.Println(num2 / i)

	return result
}
