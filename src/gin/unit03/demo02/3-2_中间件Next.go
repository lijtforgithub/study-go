package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func middleware1(c *gin.Context) {
	log.Println("中间件-1-进入")

	c.Next()

	log.Println("中间件-1-结束")
}

func middleware2() func(c *gin.Context) {
	return func(c *gin.Context) {
		log.Println("中间件-2-进入")

		c.Next()

		log.Println("中间件-2-结束")
		c.JSON(http.StatusOK, "middleware2")
	}
}

func main() {
	server := gin.Default()

	server.Use(middleware1)
	server.Use(middleware2())

	server.GET("/ping", func(c *gin.Context) {
		log.Println("请求-进入")
		c.String(http.StatusOK, c.GetString("pong"))
		log.Println("请求-结束")
	})
	server.Run()
}
