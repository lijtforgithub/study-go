package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func main() {
	yamlFile, err := os.ReadFile("resources/config/app.yml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析 YAML 文件
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
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
