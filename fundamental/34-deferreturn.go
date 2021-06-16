package main

import "fmt"

func main()  {
	fmt.Println(func1())
	fmt.Println(func2())
	fmt.Println(func3())
	fmt.Println(func4())
}

/*
	Go语言中函数的return不是原子操作，在底层分为两步来执行
	1. 返回值赋值
	defer
	2. 真正的RET返回
	函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
 */

// defer带有return的示例
func func1() int {
	x := 5
	defer func() {
		x++
	}()
	return x           // 5
}

func func2() (x int) {
	defer func() {
		x++
	}()
	return 5        // 6
}

func func3() (y int) {
	x := 5
	defer func() {
		x++       // 6
	}()
	return x       // 5
}
func func4() (x int) {
	defer func(x int) {
		x++          // 这里是副本
	}(x)
	return 5      // 5
}