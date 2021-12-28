package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime caller
// 主要用来获取函数调用信息：函数名，文件(名)，行号等
func main()  {
	// 调用当前文件
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller failed！")
		return
	}

	/*
		main.main
		63-runtimecaller.go
		13
	 */
	// 调用的函数名
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	// 文件名
	fmt.Println(path.Base(file))
	// 行号
	fmt.Println(line)
}
