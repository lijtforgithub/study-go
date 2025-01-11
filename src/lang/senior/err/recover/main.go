package main

import "fmt"

func main() {
	recoverDemo1()
}

func recoverDemo1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("A", err)
		}
	}()

	panic(nil)
	fmt.Println("B")
}
