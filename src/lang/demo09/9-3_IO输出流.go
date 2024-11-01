package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	path := "/Users/lijingtang/Downloads/go/test.txt"
	file, err := os.OpenFile(path, os.O_RDWR|os.O_TRUNC, fs.ModePerm)
	defer file.Close()
	writer := bufio.NewWriter(file)

	for i := 1; i < 6; i++ {
		ok, _ := writer.WriteString(fmt.Sprintf("第%d行\n", i))
		fmt.Print(ok)
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}

}
