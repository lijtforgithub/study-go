package main

import (
	"fmt"
	"strconv"
)

/*
*
GO 语言没有隐式类型转换 必须显示转换
*/
func main() {

	var a int8 = 8
	var b int16 = 8
	fmt.Println(a == int8(b))

	var s = string(b)
	fmt.Println("s = ", s)

	s = fmt.Sprintf("%d", b)
	fmt.Println("s = ", s)

	var str string = "10"
	var num int
	num, _ = strconv.Atoi(str)
	fmt.Println("num = ", num)

	num++
	str = strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：%q\n", num, str)

	c := 'c'
	var sp = fmt.Sprintf("%d \t %c", c, c)
	fmt.Println("字符 " + sp)

	var bt byte = 97
	fmt.Printf("数值 bt = %d\n", bt)
	fmt.Printf("字符 bt = %c\n", bt)

}
