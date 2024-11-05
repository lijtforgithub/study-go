package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	path := "/Users/lijingtang/Downloads/go/test.txt"

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("打开文件出错", err)
	} else {
		fmt.Println("文件内容：\n", string(content))
	}

	fmt.Println("\n开始缓存读取")
	file, err := os.Open(path)
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
}
