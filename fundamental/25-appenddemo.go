package main

import "fmt"

func main()  {
	appendDemo()
	sliceDemo()
}

func appendDemo()  {
	a := [...]int {1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29}
	s := a[:]
	fmt.Println("================= 1 =================")
	/*
		[1 3 5 7 9 11 13 15 17 19 21 23 25 27 29]
	 */
	fmt.Println(s)

	// 修改元素
	// 将切片中第2和第3个元素删掉,即将数组中的3，5删掉
	s = append(s[0:1], s[3:]...)
	fmt.Println("================= 2 =================")
	/*
		[1 7 9 11 13 15 17 19 21 23 25 27 29]
	 */
	fmt.Println(s)
	/*
		s = [1 7 9 11 13 15 17 19 21 23 25 27 29], len(s) = 13, cap(s) = 15
	 */
	fmt.Printf("s = %v, len(s) = %d, cap(s) = %d\n", s, len(s), cap(s))

	// 输出原数组
	/*
		1. 切片不保存具体的值
		2. 切片对应一个底层数组
		3. 底层数组都是一块连续的内存
	 */
	fmt.Println("================= 3 =================")
	/*
		a =  [1 7 9 11 13 15 17 19 21 23 25 27 29 27 29]
	 */
	fmt.Println("a = ", a)
}

func sliceDemo()  {
	fmt.Println("================= sliceDemo =================")
	var a = make([]int, 5, 10)   // 创建切片，长度为5，容量为10
	/*
		before a = [0 0 0 0 0]
	 */
	fmt.Printf("before a = %v\n", a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	/*
		after a = [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	 */
	fmt.Printf("afetr a = %v\n", a)
}
