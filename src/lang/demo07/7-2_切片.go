package main

import "fmt"

func main() {
	arr := [6]int{1, 4, 7, 2, 5, 8}
	slice := arr[1:]
	fmt.Println(slice)
	// 长度
	fmt.Println(len(slice))
	// 容量
	fmt.Println(cap(slice))
	// 切片是引用类型 底层是结构体
	slice[0] = 10
	fmt.Println(arr)

	slice1 := make([]int, 2, 10)
	slice1[1] = 2
	fmt.Println(slice1)

	slice2 := []int{2, 5, 8}
	fmt.Println(cap(slice2))
	// 切片可以再切片 同一个数组
	slice3 := slice[1:5]
	fmt.Println(slice3)
	// 切片追加 新的数据
	slice4 := append(slice2, 4, 8)
	slice4[0] = 10
	fmt.Println(slice2)
	fmt.Println(slice4)

	// 切片复制 新的数组
	var b []int = make([]int, 10)
	copy(b, slice)
	fmt.Println(b)

}
