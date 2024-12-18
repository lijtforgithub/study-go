package main

import (
	"fmt"
	"unsafe"
)

// iota代表了const声明块的行索引（下标从0开始）
// const声明有一个特点：如果常量指定了一个表达式，但后续的常量没有表达式；则继承上面的表达式
func main() {
	// 常量
	const x string = "X"
	const y string = "Y"

	fmt.Println(x + y)

	const (
		Unknown = 0
		// 不赋值和上一个常量一样
		XXOO
		Female = 1
		XXOO1
		Male = 2
		XXOO2
	)

	fmt.Println(XXOO, XXOO1, XXOO2)

	const (
		r = "abc"
		s = len(r)
		t = unsafe.Sizeof(Unknown)
	)

	fmt.Println(r, s, t)

	const (
		a = iota // 0
		b        // 1
		c        // 2
		d = "ha" // 独立值，iota += 1
		e        // "ha"   iota += 1
		f = 100  // iota +=1
		g        // 100  iota +=1
		h = iota // 7,恢复计数
		i        // 8
	)

	fmt.Println(a, b, c, d, e, f, g, h, i)

	const (
		j = "xx"
		k
	)

	fmt.Println(j, k)

	const (
		l = 1 << iota
		m = 3 << iota
		n
		o
	)

	fmt.Println(l, m, n, o)
}
