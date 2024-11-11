package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type path struct {
	ID string `uri:"id" binding:"required"`
}

func main() {
	server := gin.Default()

	server.GET("/bind/path/:id", func(c *gin.Context) {
		var p path
		err := c.ShouldBindUri(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ShouldBindUri": p,
		})
	})

	// header参数绑定
	server.GET("/bind/header", func(c *gin.Context) {
		var p path
		err := c.ShouldBindHeader(&p)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ShouldBindHeader": p,
		})
	})

	server.Run()
}
