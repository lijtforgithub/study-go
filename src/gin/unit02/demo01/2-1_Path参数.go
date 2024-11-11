package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/path/:param", func(c *gin.Context) {
		fmt.Println(c.Params)
		p := c.Param("param")
		c.JSON(http.StatusOK, gin.H{
			"Path参数": p,
		})
	})

	server.Run()
}
