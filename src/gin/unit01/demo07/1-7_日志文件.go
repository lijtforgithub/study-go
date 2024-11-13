package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {
	//gin.DisableConsoleColor()
	//gin.ForceConsoleColor()

	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	server.Run()
}
