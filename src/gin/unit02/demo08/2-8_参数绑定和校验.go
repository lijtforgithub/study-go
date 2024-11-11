package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type person struct {
	ID        string   `form:"id" binding:"required,uuid"`
	Name      string   `form:"name" binding:"required,min=2"`
	Age       int      `form:"age,default=1"`
	Friends   []string `form:"friends,default=Will;Bill"`
	Addresses []string `form:"addresses,default=foo bar" collection_format:"ssv"`
	LapTimes  []int    `form:"lap_times,default=1,2,3" collection_format:"csv"`
}

func main() {
	server := gin.Default()

	server.POST("/valid", func(c *gin.Context) {
		var p person
		err := c.ShouldBind(&p)

		fmt.Println("xxooo--------")

		if err != nil {
			fmt.Println("错误信息", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"person": p,
		})
	})

	server.Run()
}
