package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/test1/:action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test1": c.Param("action"),
		})
	})

	// 前缀匹配
	server.GET("/test2/*action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"test2": c.Param("action"),
		})
	})

	server.Run()
}
