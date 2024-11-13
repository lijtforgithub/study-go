package main

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Birthday string

// UnmarshalParam 方法用于将请求中的参数（如查询参数、表单参数或路径参数）解码并绑定到结构体中。这在处理复杂的请求参数时非常有用，可以避免手动解析每个参数。
func (b *Birthday) UnmarshalParam(param string) error {
	*b = Birthday(strings.Replace(param, "-", "/", -1))
	return nil
}

func main() {
	server := gin.Default()

	var request struct {
		Birthday Birthday `form:"birthday"`
	}
	server.GET("/test", func(c *gin.Context) {
		_ = c.BindQuery(&request)
		c.JSON(200, request.Birthday)
	})

	server.Run()
}
