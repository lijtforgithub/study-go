package main

import (
	"fmt"
)

func main() {
	var a int
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值 %d\n", a)
	}

	fmt.Printf("a = %d\n", a)

	var b int = 5
	for a < b {
		a++
		fmt.Printf("a 的值 %d\n", a)
	}

	var s string = "hello 你好"
	for i := 0; i < len(s); i++ {
		fmt.Printf("第 %d 位 %c\n", i, s[i])
	}
	fmt.Println()
	for i, c := range s {
		fmt.Printf("第 %d 位 %c\n", i, c)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	//for {
	//	fmt.Println("死循环")
	//}
}
