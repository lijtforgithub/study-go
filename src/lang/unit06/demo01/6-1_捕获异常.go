package main

import "fmt"

func main() {
	//defaultFun()

	recoverFun()

	fmt.Println("异常后面的代码")
}

func recoverFun() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常", err)
		}
	}()
	num1 := 1
	num2 := 0
	return num1 / num2
}

func defaultFun() int {
	num1 := 1
	num2 := 0
	return num1 / num2
}
