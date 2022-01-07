package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp服务端

func processConnection(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Printf("server ======> listen failed, error is %v.\n", err)
		return
	}

	// 循环接收消息
	for {
		// 建立连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("server ======> accept failed, error is %v.\n", err)
			return
		}

		go processConnection(conn)
	}
}
