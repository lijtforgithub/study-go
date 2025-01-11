package main

import (
	"fmt"
	"sync"
)

func main() {
	releaseResource()

	processControl()

	catchError()

	fmt.Println(deferFuncReturn())
}

/**
 * 资源释放
 */
func releaseResource() {
	l := sync.Mutex{}
	l.Lock()
	defer l.Unlock()
	fmt.Println("defer常用于关闭文件句柄、数据库连接、停止定时器Ticker及关闭管道等资源清理场景")
}

func processControl() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("协程执行完毕")
		wg.Done()
	}()
	defer func() {
		wg.Wait()
		fmt.Println("defer也常用于控制函数执行顺序，比如配合WaitGroup实现等待协程退出")
	}()

	fmt.Println("processControl 函数逻辑")
}

func catchError() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("defer也常用于处理异常，与recover配合可以消除panic。另外，recover只能用于defer函数中", err)
		}
	}()
	num1 := 1
	num2 := 0
	return num1 / num2
}

// 修改函数具名返回值
func deferFuncReturn() (result int) {
	i := 1
	defer func() {
		result++
	}()

	return i
}
