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
	i := 10
	stu2 := student{"李四", 4}
	t1 := reflect.TypeOf(i)
	fmt.Println(t1.Kind(), t1)
	t2 := reflect.TypeOf(stu2)
	fmt.Println(t2.Kind(), t2)
}
