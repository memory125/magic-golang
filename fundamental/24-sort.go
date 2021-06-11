package main

import (
	"fmt"
	"sort"
)

func main()  {
	// 声明数组
	var a = [...]int {3, 96, 5, 612, 44, 15, 144, 71, 5, 15, -8, 890, -12}
	// 对数组切片排序
	sort.Ints(a[:])
	/*
		[-12 -8 3 5 5 15 15 44 71 96 144 612 890]
	 */
	fmt.Println(a)
}
