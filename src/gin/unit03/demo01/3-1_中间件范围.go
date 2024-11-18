package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const KEY = "md"

func globalMiddleware(c *gin.Context) {
	log.Println("全局中间件")
	c.Set(KEY, c.GetString(KEY)+"/global")
}

func groupMiddleware(c *gin.Context) {
	log.Println("请求组中间件")
	c.Set(KEY, c.GetString(KEY)+"/group")
}

func reqMiddleware(c *gin.Context) {
	log.Println("单个请求中间件")
	c.Set(KEY, c.GetString(KEY)+"/single")
}

func main() {
	server := gin.Default()

	server.Use(globalMiddleware)

	server.GET("/ping", func(c *gin.Context) {
		log.Println(c.Request.URL)
		c.String(http.StatusOK, c.GetString(KEY))
	})

	{
		v1 := server.Group("/v1", groupMiddleware)

		v1.GET("/group", func(c *gin.Context) {
			log.Println(c.Request.URL)
			c.String(http.StatusOK, c.GetString(KEY))
		})

		v1.GET("/single", reqMiddleware, func(c *gin.Context) {
			log.Println(c.Request.URL)
			c.String(http.StatusOK, c.GetString(KEY))
		})
	}

	server.Run()
}
