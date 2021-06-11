package main

import "fmt"

func main()  {
	appendDemo()
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
}
