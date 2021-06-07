package main

import "fmt"

func main()  {
	multiplicationTable(9)
}

// 打印乘法表
func multiplicationTable(num int)  {
	for i := 1; i <= num; i++ {
		for j := 1; i <= i; j++ {
			fmt.Printf("%d * %d = %d", i, j, i * j)
			fmt.Printf("  ")
			if j == i {
				fmt.Println()
				break
			}
		}
	}
}