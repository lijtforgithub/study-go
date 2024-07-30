package import1

import "fmt"

var Pkg_1 string = "大写字母开头的变量和函数可以被其他包使用"
var pkg_1 string = "小写开头的变量和函数不能被其他包使用"

func UpCase() {
	fmt.Println("大写字母开头的函数")
}

func loCase() {
	fmt.Println("小写字母开头的函数")
}
