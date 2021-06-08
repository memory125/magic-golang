package main

import "fmt"

func main()  {
	b1 := true
	var b2 bool

	// 类型
	/*
		Type b1 = bool
		Type b2 = bool
	 */
	fmt.Printf("Type b1 = %T\n", b1)
	fmt.Printf("Type b2 = %T\n", b2)

	// 值
	/*
		Value b1 = true
		Value b2 = false
	 */
	fmt.Printf("Value b1 = %v\n", b1)
	fmt.Printf("Value b2 = %v\n", b2)
}
