package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	//router := gin.Default()
	//http.ListenAndServe(":8080", router)

	router := gin.Default()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("服务启动失败", err.Error())
	}
}
