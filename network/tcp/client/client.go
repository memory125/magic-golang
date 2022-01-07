package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp客户端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("client ==========> dial failed, error is %v.\n", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("请输入要发送的信息：")
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			fmt.Printf("client ==========> write failed, error is %v.\n", err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("recv failed, error is %v.\n", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
