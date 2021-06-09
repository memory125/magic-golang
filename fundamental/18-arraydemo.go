package main

import (
	"fmt"
)

func main()  {
	arraySum()
	arrayIndex()
}

// 数组求和
func arraySum()  {
	a1 := [5]int {1, 3, 5, 7, 8}
	sum := 0

	for _, v := range a1 {
		sum += v
	}

	fmt.Println(sum)
}

// 输出数组中两个数的和为8的数组下标
func arrayIndex()  {
	a := [5]int {1, 3, 5, 7, 8}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] + a[j] == 8{
				fmt.Printf("[%d %d]\n",i, j)
			}
		}
	}
}
