package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "GO Gin",
			"data":  "GO Web框架",
		})
	})

	server.GET("/ping", ping)
	server.GET("/map", mapFun)

	err := server.Run("localhost:1234")
	if err != nil {
		fmt.Println("服务启动失败：", err)
	}
}

func mapFun(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"m": "Map",
	})
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
