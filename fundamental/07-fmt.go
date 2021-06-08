package main

import "fmt"

func main()  {
	var x = 100
	// 输出
	/*
		x = 1100100
		x = 0144
		x = 100
		x = 0x64
		x = 100
		Type x = int
	 */
	// 二进制
	fmt.Printf("x = %b\n", x)
	// 八进制
	fmt.Printf("x = 0%o\n", x)
	// 十进制
	fmt.Printf("x = %d\n", x)
	// 十六进制
	fmt.Printf("x = 0x%x\n", x)
	// 值
	fmt.Printf("x = %v\n", x)
	// 类型
	fmt.Printf("Type x = %T\n", x)

	s := "你好Go！"
	/*
		s = 你好Go！
		s = 你好Go！
		s = "你好Go！"
	 */
	fmt.Printf("s = %s\n", s)
	fmt.Printf("s = %v\n", s)
	fmt.Printf("s = %#v\n", s)
}
