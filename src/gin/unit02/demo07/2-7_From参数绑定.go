package main

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type form struct {
	FormKey  string                `form:"key"`
	FormFile *multipart.FileHeader `form:"file"`
}

type json struct {
	JsonKey string `json:"key"`
}

func main() {
	server := gin.Default()

	server.POST("/bind/form", func(c *gin.Context) {
		var f form

		err := c.ShouldBind(&f)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ShouldBind": f,
		})
	})

	server.POST("/bind/json", func(c *gin.Context) {
		var j json

		err := c.ShouldBindJSON(&j)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ShouldBindJSON": j,
		})
	})

	server.Run()
}
