package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	server.POST("/route", func(c *gin.Context) {
		c.String(http.StatusOK, "POST")
	})

	server.PUT("/route", func(c *gin.Context) {
		c.String(http.StatusOK, "PUT")
	})

	server.PATCH("/route", func(c *gin.Context) {
		c.String(http.StatusOK, "PATCH")
	})

	server.DELETE("/route", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE")
	})

	server.Any("/any", func(c *gin.Context) {
		c.String(http.StatusOK, "Any")
	})

	server.Run()
}
