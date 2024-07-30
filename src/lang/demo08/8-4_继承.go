package main

import "fmt"

type Animal struct {
	Age    int
	Weight float32
}

func (a Animal) show() {
	fmt.Println("Animal ", a.Age)
}

type House struct {
	Color string
}

type Cat struct {
	// 匿名字段 表示继承
	Animal
	// 也可以是基本类型
	string
	// 正常字段
	Age int
	// 组合
	h House
}

func (c Cat) show() {
	fmt.Println("Cat ", c.Age)
}

func main() {
	c := Cat{Animal{1, 2.0}, "xxoo", 2, House{"红色"}}
	fmt.Println(c)
	c.show()
	c.Animal.Age = 10
	c.Animal.show()
}
