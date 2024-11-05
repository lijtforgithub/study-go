package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	path := "/Users/lijingtang/Downloads/go/test.txt"
	path2 := "/Users/lijingtang/Downloads/go/test2.txt"

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("打开文件出错", err)
	} else {
		err := os.WriteFile(path2, content, fs.ModePerm)
		if err != nil {
			fmt.Println("文件复制出错", err)
		}
	}
}
