package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		fmt.Println("main-defer")
	}()

	go routine1()
	go routine2()

	fmt.Println("main")
	time.Sleep(time.Second * 5)
}

func routine1() {
	defer func() {
		fmt.Println("routine1-defer")
	}()
	fmt.Println("routine1")
	time.Sleep(time.Second * 5)
	foo1()
}

func foo1() {
	defer func() {
		fmt.Println("foo1-defer")
	}()
	fmt.Println("foo1")
}

func routine2() {
	defer func() {
		fmt.Println("routine2-defer")
	}()
	fmt.Println("routine2")
	foo2()
}

func foo2() {
	defer func() {
		fmt.Println("foo2-defer")
	}()
	fmt.Println("foo2")
	panic("foo2-panic")
}
