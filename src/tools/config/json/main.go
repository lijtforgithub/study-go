package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int    `json:"port"`
		Host string `json:"host"`
	} `json:"server"`

	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"database"`
}

func main() {
	jsonFile, err := os.ReadFile("resources/config/app.json")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析 JSON 文件
	var config Config
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// 打印配置信息
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("Database User: %s\n", config.Database.User)
	fmt.Printf("Database Password: %s\n", config.Database.Password)
	fmt.Printf("Database Name: %s\n", config.Database.Name)
}
