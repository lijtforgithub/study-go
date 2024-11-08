package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	u := server.Group("/user")
	u.GET("", getHandler)
	u.POST("/v1", postHandler)

	g1 := server.Group("g1")
	{
		g1.GET("", getHandler)

		g2 := g1.Group("/g2")
		g2.POST("", postHandler)
	}

	server.Run()
}

func getHandler(c *gin.Context) {
	c.String(http.StatusOK, "GET 请求")
}

func postHandler(c *gin.Context) {
	c.String(http.StatusOK, "POST 请求")
}
