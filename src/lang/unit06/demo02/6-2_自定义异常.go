package main

import (
	"errors"
	"fmt"
)

func main() {
	err := newError()
	if err != nil {
		fmt.Println("自定义异常", err)

		// 抛出异常 中断程序代码
		panic(err)
	}

	fmt.Println("异常后面的代码")
}

func newError() error {
	num1 := 1
	num2 := 0
	if num2 == 0 {
		err := errors.New("除数不能为0")
		return err
	}

	result := num1 / num2
	fmt.Println(result)
	return nil
}
