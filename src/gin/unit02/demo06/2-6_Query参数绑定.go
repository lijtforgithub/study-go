package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Id   int
	Name string
}

type student struct {
	No   int    `form:"no" binding:"required"`
	Name string `form:"name"`
}

func main() {
	server := gin.Default()

	// ?Id=1&Name=xxoo
	server.GET("/bind/query", func(c *gin.Context) {
		var u user

		err := c.BindQuery(&u)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"BindQuery": u,
		})
	})

	server.GET("/bind/query1", func(c *gin.Context) {
		var s student

		// 在绑定失败时，会自动返回一个 400 Bad Request 响应。
		err := c.BindQuery(&s)
		fmt.Println(err)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"BindQuery": s,
		})
	})

	server.GET("/bind/query2", func(c *gin.Context) {
		var s student

		// 在绑定失败时，不会自动返回错误响应，需要手动处理绑定失败的情况。
		err := c.ShouldBindQuery(&s)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ShouldBindQuery": s,
		})
	})

	server.Run()
}
