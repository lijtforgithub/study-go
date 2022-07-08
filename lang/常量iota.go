package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 常量
	const x string = "X"
	const y string = "Y"

	fmt.Println(x + y)

	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

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
