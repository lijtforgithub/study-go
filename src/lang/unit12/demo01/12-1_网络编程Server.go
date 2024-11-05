package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("读取客户端数据：", string(buf[0:n]))
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("服务端启动失败: ", err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("连接客户端失败: ", err)
			continue
		}

		fmt.Println("连接客户端成功: ", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}
