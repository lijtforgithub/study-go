package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("/Users/lijingtang/Downloads/go/test.txt")
	if err != nil {
		fmt.Println("打开文件出错", err)
	} else {
		fmt.Printf("打开文件%v\n", file)
		defer func() {
			err = file.Close()
			if err != nil {
				fmt.Println("关闭文件异常")
			}
		}()
	}
}
