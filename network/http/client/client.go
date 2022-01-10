package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http - client

func main() {
	resp, err := http.Get("http://127.0.0.1:6688/req")
	if err != nil {
		fmt.Printf("Http Get URL failed, error is %v.\n", err)
		return
	}

	defer resp.Body.Close()

	bytesRead, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read response body failed, error is %v.\n", err)
		return
	}

	fmt.Printf("Resp Body, %v.\n", string(bytesRead))
}
