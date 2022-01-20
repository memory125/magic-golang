package main

import (
	"fmt"
	"net"
	"strings"
)

// ip示例

// 获取对外网络访问的ip地址
func getOutboundIPAddr() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("Dial dns server failed, error is %v.\n", err)
		return
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Printf("local address is %v.\n", localAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() {
	ip, err := getOutboundIPAddr()
	if err != nil {
		fmt.Printf("getOutboundIPAddr failed, error is %v.\n", err)
		return
	}
	fmt.Println("ip address: ", ip)
}
