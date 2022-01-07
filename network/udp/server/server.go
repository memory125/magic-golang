package main

import (
	"fmt"
	"net"
	"strings"
)

// udp - 服务端

func main() {
	// 监听udp
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 6688,
	})
	if err != nil {
		fmt.Printf("Server UDP ========> listen failed! error is %v.\n", err)
		return
	}

	defer conn.Close()

	// 读取数据
	var data [1024]byte
	for {
		bytes, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("Server UDP ========> read failed! error is %v.\n", err)
			return
		}

		// 输出读到的数据
		fmt.Printf("read %d bytes from UDP, data is %v.\n", bytes, string(data[:bytes]))

		// 将读取的数据转成大写
		msg := strings.ToUpper(string(data[:bytes]))

		// 发送数据
		_, err = conn.WriteToUDP([]byte(msg), addr)
		if err != nil {
			fmt.Printf("Server UDP ========> write data to UDP failed! error is %v.\n", err)
			return
		}
	}

}
