package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	go tickerDemo()
	go tickerLaunch()

	ch := time.Tick(time.Second)
	for t := range ch {
		fmt.Println(t)
	}

	time.Sleep(10 * time.Second)
}

// 简单定时任务
func tickerDemo() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick")
	}
}

// 定时聚合任务：公交车每隔5分钟发一班，不管是否已坐满乘客；已坐满乘客的情况下，不足5分钟也发车。
func tickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 10
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := getNewPassenger()
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}

		select {
		case <-ticker.C:
			launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger {
				launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func launch(passengers []string) {
	log.Println("发车", passengers)
}

func getNewPassenger() string {
	time.Sleep(1 * time.Second)
	p := rand.Int()
	log.Println("新来一个乘客", p)
	return strconv.Itoa(p)
}
