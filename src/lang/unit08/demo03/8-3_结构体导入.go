package main

import "fmt"
import "study-go/src/lang/pkg/struct"

func main() {
	s1 := pckStruct.Student{"大写字母开头的结构体其他包可以访问"}
	fmt.Println(s1)

	s2 := pckStruct.NewStu("工厂模式访问小写字母开头的结构体")
	fmt.Println(*s2)

}
