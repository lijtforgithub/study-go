package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/test1/:name/:action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			c.Param("name"): c.Param("action"),
		})
	})

	// 前缀匹配
	server.GET("/test2/:name/*action", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			c.Param("name"): c.Param("action"),
		})
	})

	server.Run()
}
