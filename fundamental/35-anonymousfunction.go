package main

import "fmt"

func main()  {
	// 定义匿名函数
	f1 := func(x int, y int) int {
		fmt.Println("匿名函数1 =====> Sum")
		return x + y
	}

	/*
		匿名函数1 =====> Sum
		300
	 */
	fmt.Println(f1(100, 200))

	/*
		匿名函数2 ====> Sub
		100
	 */
	fmt.Println(func(a int, b int) int {
		fmt.Println("匿名函数2 ====> Sub")
		return  a - b
	}(200, 100))
}
