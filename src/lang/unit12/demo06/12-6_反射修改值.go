package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 10
	// 传指针
	reValue := reflect.ValueOf(&i)
	fmt.Println(reValue.Kind())
	reValue.Elem().SetInt(100)
	fmt.Println(i)
}
