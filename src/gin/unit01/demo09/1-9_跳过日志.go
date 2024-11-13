package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.New()

	// skip logging for desired paths by setting SkipPaths in LoggerConfig
	loggerConfig := gin.LoggerConfig{SkipPaths: []string{"/metrics"}}

	// skip logging based on your logic by setting Skip func in LoggerConfig
	loggerConfig.Skip = func(c *gin.Context) bool {
		// as an example skip non server side errors
		return c.Writer.Status() < http.StatusInternalServerError
	}

	server.Use(gin.LoggerWithConfig(loggerConfig))
	server.Use(gin.Recovery())

	// skipped
	server.GET("/metrics", func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	})

	// skipped
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// not skipped
	server.GET("/data", func(c *gin.Context) {
		c.Status(http.StatusNotImplemented)
	})

	server.Run()
}
