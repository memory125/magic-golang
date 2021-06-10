package main

import "fmt"

func main()  {
	// 使用make函数定义切片
	// 下面的例子定义了一个长度为5， 容量为10的int切片
	a := make([]int, 5, 10)
	/*
		a = [0 0 0 0 0], len(a) = 5, cap(a) = 10
	 */
	fmt.Printf("a = %d, len(a) = %d, cap(a) = %d\n", a, len(a), cap(a))
}
