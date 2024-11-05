package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type student struct {
	name string
	age  int
}

type Person struct {
	Name string
	age  int
}

// 非指针传递 不会修改原对象属性
func (stu student) SetName(newName string) (oldName string) {
	oldName = stu.name
	stu.name = newName
	fmt.Printf("姓名由%v变成%v\n", oldName, newName)
	return oldName
}

func (stu student) SetAge(newAge int) (oldAge int) {
	oldAge = stu.age
	stu.age = newAge
	fmt.Printf("年龄由%v变成%v\n", oldAge, newAge)
	return oldAge
}

// 小写字母开头的方法 反射获取不到
func (stu student) getInfo() string {
	return "name: " + stu.name + " age: " + strconv.Itoa(stu.age)
}

func main() {
	stu := student{"张三", 3}

	fmt.Println(stu.SetName("李四"))
	fmt.Println(stu.SetAge(4))
	fmt.Println(stu.getInfo())
	fmt.Println()

	reType := reflect.TypeOf(stu)
	reValue := reflect.ValueOf(stu)
	n1 := reValue.NumField()
	fmt.Println("属性个数 ", n1)
	for i := 0; i < n1; i++ {
		fmt.Println(i, reType.Field(i).Name, reValue.Field(i))
	}

	n2 := reValue.NumMethod()
	fmt.Println("方法个数 ", n2)
	for i := 0; i < n1; i++ {
		fmt.Println(i, reType.Method(i).Name, reValue.Method(i))
	}

	var params []reflect.Value
	params = append(params, reflect.ValueOf("反射"))
	reValue.Method(1).Call(params)

	fmt.Println()
	p := Person{"指针+大写字母才能CanSet", 0}
	pValue := reflect.ValueOf(&p)
	pValue = pValue.Elem()

	val1 := pValue.FieldByName("Name")
	val2 := pValue.FieldByName("age")
	fmt.Println("Name", val1.IsValid(), val1.CanSet())
	fmt.Println("age", val2.IsValid(), val2.CanSet())
}
