package main

import "fmt"

func main() {
	// 10进制变量
	var a = 10

	/*
		a = 10
		a = 0xa
		a = 12
		a = 1010
		Type a = int
	 */
	fmt.Println("============= 输出 -> a ===========")
	// 10进制
	fmt.Printf("a = %d\n", a)
	// 16进制
	fmt.Printf("a = 0x%x\n", a)
	// 8进制
	fmt.Printf("a = %o\n", a)
	// 2进制
	fmt.Printf("a = %b\n", a)
	// 原始类型
	fmt.Printf("Type a = %T\n", a)

	// 8进制变量
	b := 077
	/*
		b = 63
		b = 0x3f
		b = 77
		b = 111111
		Type b = int
	 */
	fmt.Println("============= 输出 -> b ===========")
	// 10进制
	fmt.Printf("b = %d\n", b)
	// 16进制
	fmt.Printf("b = 0x%x\n", b)
	// 8进制
	fmt.Printf("b = %o\n", b)
	// 2进制
	fmt.Printf("b = %b\n", b)
	// 原始类型
	fmt.Printf("Type b = %T\n", b)

	// 16进制变量
	 c := 0x1234567
	/*
		c = 19088743
		c = 0x1234567
		c = 110642547
		c = 1001000110100010101100111
		Type c = int
	 */
	fmt.Println("============= 输出 -> c ===========")
	// 10进制
	fmt.Printf("c = %d\n", c)
	// 16进制c
	fmt.Printf("c = 0x%x\n", c)
	// 8进制
	fmt.Printf("c = %o\n", c)
	// 2进制
	fmt.Printf("c = %b\n", c)
	// 原始类型
	fmt.Printf("Type c = %T\n", c)
}
