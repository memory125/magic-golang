package main

import (
	"fmt"
	"net"
	"strconv"
)

// tcp - 黏包

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6688")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := "Hello, Hello. How are you? + " + strconv.Itoa(i)
		conn.Write([]byte(msg))
	}
}
