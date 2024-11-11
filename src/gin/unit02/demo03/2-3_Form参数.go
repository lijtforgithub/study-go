package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.POST("/post/form", func(c *gin.Context) {
		fmt.Println(c.Params)

		p1 := c.PostForm("p1")
		p2 := c.DefaultPostForm("p2", "P2默认值")
		c.JSON(http.StatusOK, gin.H{
			"PostForm":        p1,
			"DefaultPostForm": p2,
		})
	})

	server.POST("/post/form/array", func(c *gin.Context) {
		p := c.PostFormArray("p")
		c.JSON(http.StatusOK, gin.H{
			"PostFormArray": p,
		})
	})

	server.POST("/post/form/map", func(c *gin.Context) {
		p := c.PostFormMap("p")
		c.JSON(http.StatusOK, gin.H{
			"PostFormMap": p,
		})
	})

	server.Run()
}
