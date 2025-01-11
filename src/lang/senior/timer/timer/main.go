package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := time.NewTimer(time.Second * 5)
	fmt.Println("5s倒计时", time.Now().Format("2006-01-02 15:04:05"))
	<-t1.C
	fmt.Println("5s时间到", time.Now().Format("2006-01-02 15:04:05"))

	var ch chan string
	waitChannel(ch)
	delayFunction()

	stop()
	reset()
	after()
	afterFun()
}

// waitChannel 设定超时时间
func waitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)
	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C:
		log.Println("等待超时")
		return false
	}
}

// delayFunction 延迟执行某个方法
func delayFunction() {
	log.Println("延迟方法开始")
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		log.Println("延迟5s开始执行")
	}
}

func stop() {
	t1 := time.NewTimer(1 * time.Second)
	t2 := time.NewTimer(2 * time.Second)
	time.Sleep(time.Second)
	fmt.Println(t1.Stop())
	fmt.Println(t2.Stop())
}

func reset() {
	log.Println("reset开始")
	t1 := time.NewTimer(1 * time.Second)
	t1.Stop()
	t1.Reset(5 * time.Second)
	<-t1.C
	log.Println("reset结束")
}

func after() {
	log.Println("after开始")
	<-time.After(time.Second)
	log.Println("after结束")
}

func afterFun() {
	log.Println("afterFun开始")
	time.AfterFunc(time.Second, func() {
		log.Println("afterFun结束-异步")
	})
	time.Sleep(2 * time.Second)
}
