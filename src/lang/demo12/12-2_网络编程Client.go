package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("连接服务端失败：", err)
		return
	}

	fmt.Println("连接服务端成功")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取终端输入失败：", err)
	} else {
		_, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("发送数据失败：", err)
		}
	}
}
