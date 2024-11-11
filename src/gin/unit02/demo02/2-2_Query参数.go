package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	// ?q1=Q1&q2=Q2
	server.GET("/query", func(c *gin.Context) {
		fmt.Println(c.Params)

		q1 := c.Query("q1")
		q2 := c.DefaultQuery("q2", "Q2默认值")
		c.JSON(http.StatusOK, gin.H{
			"Query":        q1,
			"DefaultQuery": q2,
		})
	})

	// ?q=1&q=2&q=3
	server.GET("/query/array", func(c *gin.Context) {
		fmt.Println(c.Query("q"))
		q := c.QueryArray("q")
		c.JSON(http.StatusOK, gin.H{
			"QueryArray": q,
		})
	})

	// ?q[m1]=1&q[m2]=2&q[m3]=3
	server.GET("/query/map", func(c *gin.Context) {
		q := c.QueryMap("q")
		c.JSON(http.StatusOK, gin.H{
			"QueryMap": q,
		})
	})

	server.Run()
}
