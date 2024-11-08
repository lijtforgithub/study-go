package main

import (
	"fmt"
	"study-go/src/lang/pkg/import1"
	// 这里使用源文件目录
	"study-go/src/lang/pkg/import"
	// 使用别名区分
	p1 "study-go/src/lang/pkg/alias/pkg1"
	p2 "study-go/src/lang/pkg/alias/pkg2"
)

func main() {
	fmt.Println(import1.Pkg_1)
	import1.UpCase()
	fmt.Println(import1.Pkg_2)

	// 这里使用包名
	fmt.Println(unsame.Pkg_3)

	p1.SamePkg()
	p2.SamePkg()
}
