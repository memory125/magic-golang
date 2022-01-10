package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http - server

// http请求处理函数
func handle1(w http.ResponseWriter, r *http.Request) {
	str := "<h1>Hello Http Web!!!</h1>"
	_, _ = w.Write([]byte(str))
}

func handle2(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	fmt.Println(queryParams)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/hello", handle1)
	http.HandleFunc("/req", handle2)
	_ = http.ListenAndServe("127.0.0.1:6688", nil)
}
