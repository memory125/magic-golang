package main

import "fmt"

func main()  {
	sliceDemo()
}

func sliceDemo()  {
	// 声明长度为5，容量为10的整形切片a
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	/*
		[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	 */
	fmt.Println(a)
	/*
		20
	 */
	fmt.Println(cap(a))
}