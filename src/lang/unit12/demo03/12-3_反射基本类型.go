package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 100
	reType := reflect.TypeOf(num)
	fmt.Printf("%T = %v\n", reType, reType)
	reValue := reflect.ValueOf(num)
	fmt.Printf("%T = %v\n", reValue, reValue)

	val := reValue.Int()
	fmt.Println(val)

	newType := reValue.Interface()
	fmt.Printf("%T = ", newType)
	fmt.Println(newType.(int))
}
