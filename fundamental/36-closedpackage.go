package main

import (
	"fmt"
)

// 闭包
/*
   闭包：闭包是一个函数，这个函数包含了它外部作用域的一个变量
   底层原理：
	1. 函数可以作为返回值
    2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找
 */

func test1(f func())  {
	fmt.Println("This is function test1()")
	f()
}

func test2(x, y int)  {
	fmt.Println("This is function test2()")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func test3(f func(int, int), m, n int) func() {
	// 将ret定义为跟test1函数参数形式一致，然后返回一个函数
	ret := func() {
		// ret是不带参数的函数，然后这个函数中调用func(int, int)格式的函数
		f(m, n)
	}

	return ret
}

// 另外一个示例
func addr(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}


func main()  {
	/*
		This is function test1()
		This is function test2()
		300
	 */
	tmp := test3(test2, 100, 200)
	test1(tmp)

	/*
		Ret2 is 300
	*/
	ret := addr(100)
	ret2 := ret(200)
	fmt.Printf("Ret2 is %v \n", ret2)
}
