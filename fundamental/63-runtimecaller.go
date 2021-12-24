package main

import (
	"fmt"
	"runtime"
)

// runtime caller

func main()  {
	// 调用当前文件
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller failed！")
		return
	}

	/*
		main.main
		D:/project/go/gopath/src/code.wing.com/fundamental/63-runtimecaller.go
		12
	 */
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(line)
}
