package main

import "fmt"

/*
*
变量是一种使用方便的占位符，用于引用计算机内存地址。
Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
*/
func main() {
	a := 4
	var ptr *int

	fmt.Printf("a 的值 %d\n", a)
	fmt.Printf("a 的类型 %T\n", a)
	fmt.Println()

	/*  指针 & 和 * */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("指针 ptr 的类型 %T\n", ptr)
	fmt.Printf("指针 ptr 的内存值 %d\n", ptr)
	fmt.Printf("指针 *ptr 的值 %d\n", *ptr)
	fmt.Println()

	/*  指针的指针 */
	var pptr **int
	pptr = &ptr /* 'pptr' 包含了 'ptr' 指针的地址 */
	fmt.Printf("指针 pptr 的类型 %T\n", pptr)
	fmt.Printf("指针 pptr 的内存值 %d\n", pptr)
	fmt.Printf("指针 *pptr 的值 %d\n", *pptr)
	fmt.Println()

	// 指针接收控制台输入
	var str string
	fmt.Printf("请输入一个字符串：")
	fmt.Scan(&str)
	fmt.Printf("输入的字符串 %v\n", str)
	fmt.Println()

	var num int8
	fmt.Printf("请输入一个数字：")
	fmt.Scan(&num)
	fmt.Printf("输入的数字（类型不匹配获取不到） %d\n", num)
}
