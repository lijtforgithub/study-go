package main

import "fmt"

func main() {
	// 基本类型和数组都是值传递
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	arrayParam(arr)
	fmt.Println(arr)

	x, y := swap("Mahesh", "Kumar")
	fmt.Println(x, y)

	/* 定义局部变量 */
	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a = %d b = %d\n", a, b)
	// 值传递
	fmt.Println("函数返回值", swapValue(a, b))
	fmt.Printf("交换后（值传递） a = %d b = %d\n", a, b)

	// 引用传递
	sr := swapRef
	fmt.Printf("sr 类型 %T\n", sr)
	sr(&a, &b)
	fmt.Printf("交换后（引用传递） a = %d b = %d\n", a, b)

	// 别名
	type myInt int
	var mi myInt = 20

	sum, sub := returnValue(int(mi), 10)
	fmt.Println(sum, sub)
}

func arrayParam(arr [3]int) {
	arr[0] = 100
}

func swap(x, y string) (string, string) {
	return y, x
}

/* 值传递 */
func swapValue(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

/* 引用传递 */
func swapRef(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

func returnValue(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return
}
