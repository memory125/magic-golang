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

	// new：申请内存
	/*
	  make和new的区别
	  1. make和new都是用来申请内存的
	  2. new很少用，一般用来给基本数据类型申请内存，如`string`，`int`，new返回的是对应类型的指针(*string, *int)
	  3. make是用来给`slice`，`map`，`chan`申请内存的，make函数返回的是对应的这三个类型本身
	 */
	var a = new(int)
	*a = 100
	/*
		&a = 0xc00000a0a8, *a = 100
	 */
	fmt.Printf("&a = %p, *a = %d\n", a, *a)
}
