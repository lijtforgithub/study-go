package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port int    `toml:"port"`
		Host string `toml:"host"`
	} `toml:"server"`

	Database struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`
}

func main() {
	tomlFile, err := os.ReadFile("resources/config/app.toml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 解析 TOML 文件
	var config Config
	_, err = toml.Decode(string(tomlFile), &config)
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
