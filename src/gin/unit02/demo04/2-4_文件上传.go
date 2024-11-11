package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	server := gin.Default()
	server.MaxMultipartMemory = 8 << 20 // 每次处理文件所占用的最大内存 8 MiB

	server.POST("/file", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		fmt.Println(file.Size)

		c.String(http.StatusOK, file.Filename+" 上传成功")
	})

	server.POST("/file/multi", func(c *gin.Context) {
		form, err := c.MultipartForm()
		fmt.Println(err)
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename, file.Size)
		}

		c.String(http.StatusOK, fmt.Sprintf("%d 个文件上传成功", len(files)))
	})

	server.Run()
}
