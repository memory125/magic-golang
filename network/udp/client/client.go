package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// udp - 客户端

func main() {
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 6688,
	})
	if err != nil {
		fmt.Printf("Client UDP ========> dail failed! error is %v.\n", err)
		return
	}

	defer conn.Close()

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Please input message: ")
		msg, _ := reader.ReadString('\n')
		_, _ = conn.Write([]byte(msg))

		// 收回回复的数据
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Printf("Client UDP ========> read from UDP failed! error is %v.\n", err)
			return
		}

		// 输出收到的数据
		fmt.Printf("Client UDP =========> received replied message: %v.\n", string(reply[:n]))
	}
}
