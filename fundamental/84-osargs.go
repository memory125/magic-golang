package main

import (
	"fmt"
	"os"
)

// 获取命令行参数

/*
   运行时：可以带参数，如：fundamental.exe -p a b c 等
*/
func main() {
	/*
		Value is []string{"fundamental.exe"}
		Type is []string
	*/
	fmt.Printf("Value is %#v\n", os.Args)
	fmt.Printf("Type is %T\n", os.Args)
}
