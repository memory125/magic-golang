package main

import "fmt"

func main()  {
	// 声明数组
	var a1[3] int
	var a2[5] int

	// 打印数组类型和初始值
	/*
		a1: [3]int, a2: [5]int
		a1: [0 0 0], a2: [0 0 0 0 0]
	 */
	fmt.Printf("a1: %T, a2: %T\n", a1, a2)
	fmt.Printf("a1: %v, a2: %v\n", a1, a2)

	// 1. 数组初始化
	/*
		a1: [1 3 5], a2: [2 4 6 8 10]
	 */
	a1 = [3]int{1, 3, 5}
	a2 = [5]int{2, 4, 6, 8, 10}
	fmt.Printf("a1: %v, a2: %v\n", a1, a2)

	// 2. 根据初始值自动推断数组的长度是多少
	/*
		len a3: 10, a3: [0 1 2 3 4 5 6 7 8 9]
	 */
	a3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len a3: %d, a3: %v\n", len(a3), a3)
	// 数组遍历
	for _, v := range a3{
		fmt.Println(v)
	}

	// 3. 根据索引来初始化
	/*
		len a4: 5, a4: [1 0 0 0 2]
	 */
	a4 := [5]int{0 : 1, 4 : 2}
	fmt.Printf("len a4: %d, a4: %v\n", len(a4), a4)

	// 数组遍历
	for _, v := range a4{
		fmt.Println(v)
	}
}
