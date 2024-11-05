package main

import "fmt"

func main() {
	// 只声明map内存是没有分配空间的 数组在声明的时候就分配内存
	var a map[int]string
	// make之后分配内存
	a = make(map[int]string, 10)
	a[101] = "张三"
	a[102] = "李四"

	fmt.Println(a)

	// 使用字面量创建 Map
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3, // 最后一个也有逗号
	}

	// 获取键值对
	v1 := m["apple"]
	v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值

	fmt.Println(v1, v2, ok)

	// 遍历 Map
	for k, v := range m {
		fmt.Printf("key=%s, value=%d\n", k, v)
	}

	// 删除键值对
	delete(m, "banana")
	fmt.Println(m)
}
