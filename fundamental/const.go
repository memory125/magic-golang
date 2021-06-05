package main

import "fmt"

/*
	定义常量：const
	1. 定义之后不能修改
	2. 程序运行期间不能更改
 */

// 单独声明
const statusOk = 200

// 批量声明
const (
	pi = 3.1415926
	logFile = "log.log"
)

// 批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	s1 = 100
	s2
	s3
	s4
)

// iota
const (
	k1 = iota           // 默认是0
	k2                  // 1
	k3
	k4
)

// 匿名变量
const (
	b1 = iota           // 默认是0
	b2                  // 1
	_                   // 2
	b3                  // 3
)

// 插队
const (
	c1 = iota           // 默认是0
	c2 = 100            // 100
	c3 = iota           // 2
	c4                  // 3
)

// 多个常量声明一行,iota在常量声明时，每增加一行常量声明，iota的值加1
const (
	d1, d2 = iota + 1, iota + 2          // d1: 1 = 0 + 1, d2: 2 = 0 + 2
	d3, d4 = iota + 1, iota + 2          // d3: 2 = 1 + 1, d2: 3 = 1 + 2
)

// 定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)             // 2的10 * 1次方 = 1024
	MB = 1 << (10 * iota)             // 2的10 * 2次方 = 1024 * 1024
	GB = 1 << (10 * iota)             // 2的10 * 3次方 = 1024 * 1024 * 1024
	TB = 1 << (10 * iota)             // 2的10 * 4次方 = 1024 * 1024 * 1024 * 1024
	PB = 1 << (10 * iota)             // 2的10 * 5次方 = 1024 * 1024 * 1024 * 1024 * 1024
)

func main()  {
	fmt.Println("============ const 1 ===============")
	fmt.Println(statusOk)
	fmt.Println(pi)
	fmt.Println(logFile)

	/*
		s1 100
		s2 100
		s3 100
		s4 100
	 */
	fmt.Println("============ const 2 ===============")
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)

	/*
		k1 0
		k2 1
		k3 2
		k4 3
	 */
	fmt.Println("============ const 3 ===============")
	fmt.Println("k1", k1)
	fmt.Println("k2", k2)
	fmt.Println("k3", k3)
	fmt.Println("k4", k4)

	/*
		b1 0
		b2 1
		b3 3
	 */
	fmt.Println("============ const 4 ===============")
	fmt.Println("b1", b1)
	fmt.Println("b2", b2)
	fmt.Println("b3", b3)

	/*
		c1 0
		c2 100
		c3 2
		c4 3
	 */
	fmt.Println("============ const 5 ===============")
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)
	fmt.Println("c4", c4)

	/*
		d1 1
		d2 2
		d3 2
		d4 3
	 */
	fmt.Println("============ const 6 ===============")
	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)


	/*
		KB 1024
		MB 1048576
		GB 1073741824
		TB 1099511627776
		PB 1125899906842624
	 */
	fmt.Println("============ const 7 ===============")
	fmt.Println("KB", KB)
	fmt.Println("MB", MB)
	fmt.Println("GB", GB)
	fmt.Println("TB", TB)
	fmt.Println("PB", PB)
}
