package main

import "fmt"

func main()  {
	pointerDemo()
}

/*
	1. &：取地址
	2. *：根据地址取值
 */
func pointerDemo()  {
	// 定义int型变量
	n := 20
	// 获取变量n的地址
	p := &n
	// 输出地址
	/*
		p =  0xc00000a088
	 */
	fmt.Println("p = ", p)
	// 输出指针类型
	/*
		Type p = *int    // 表示int型的指针
	 */
	fmt.Printf(" Type p = %T\n", p)

	// 取值
	m := *p
	/*
	 m = 20
	 */
	fmt.Printf("m = %v\n", m)
}
