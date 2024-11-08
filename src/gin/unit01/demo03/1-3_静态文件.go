package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

// go build 之后执行
func main() {
	server := gin.Default()

	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	fmt.Println(execPath)
	// 获取静态文件目录的绝对路径
	staticPath := filepath.Join(filepath.Dir(execPath), "static")
	fmt.Println(staticPath)

	server.Static("/static", staticPath)

	server.StaticFile("/index", filepath.Join(filepath.Dir(execPath), "index.html"))

	server.Run()
}
