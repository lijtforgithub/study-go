package main

import (
	"fmt"
	"time"
)

func main() {
	testWaitForStopOrTimeout()

	select {}
}

func testWaitForStopOrTimeout() {
	stopChan := make(chan struct{}, 1)
	ch := waitForStopOrTimeout(stopChan, time.Second*10)

	go func() {
		time.Sleep(time.Second * 30)
		stopChan <- struct{}{}
	}()

	<-ch
}

func waitForStopOrTimeout(stopCh <-chan struct{}, timeout time.Duration) <-chan struct{} {
	stopChWithTimeout := make(chan struct{})
	go func() {
		select {
		case <-stopCh:
			fmt.Println("收到信号")
		case <-time.After(timeout):
			fmt.Println("时间到了")
		}
		close(stopChWithTimeout)
	}()
	return stopChWithTimeout
}
