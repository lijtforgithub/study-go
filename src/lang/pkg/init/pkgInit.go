package pkgInit

import "fmt"

var global int = initGlobalValue()

func initGlobalValue() int {
	fmt.Println("pkg initGlobalValue 执行了")
	return 0
}

func init() {
	fmt.Println("pkg init 执行了")
}

func TestInit() {
	fmt.Println("pkg TestInit 执行了 测试多个init执行顺序")
}
