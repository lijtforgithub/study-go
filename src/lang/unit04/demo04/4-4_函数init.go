package main

import (
	"fmt"
	"study-go/src/lang/pkg/init"
)

var global int = initGlobalValue()

func initGlobalValue() int {
	fmt.Println("initGlobalValue 执行了")
	return 10
}

/*
*
每个源文件都可以有一个init函数
*/
func init() {
	fmt.Println("init 执行了")
}

func main() {
	pkgInit.TestInit()
	fmt.Println("main 执行了")
}
