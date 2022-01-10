package main

import "net/http"

// http - server

// http请求处理函数
func handle1(w http.ResponseWriter, r *http.Request) {
	str := "<h1>Hello Http Web!!!</h1>"
	_, _ = w.Write([]byte(str))
}

func main() {
	http.HandleFunc("/wing", handle1)
	http.ListenAndServe("127.0.0.1:6688", nil)
}
