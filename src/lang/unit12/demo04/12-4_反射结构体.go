package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
	age  int
}

func main() {
	stu := student{"张三", 15}
	reType := reflect.TypeOf(stu)
	fmt.Printf("%T = %v\n", reType, reType)
	reValue := reflect.ValueOf(stu)
	fmt.Printf("%T = %v\n", reValue, reValue)

	newType := reValue.Interface()

	if stu1, ok := newType.(student); ok {
		fmt.Printf("姓名: %v 年龄: %v", stu1.name, stu1.age)
	}
}
