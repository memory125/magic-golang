package main

import "fmt"

func main()  {
	/*
		0 - a
		0 - b
		0 - c
	 */
	forDemo()
}

// 跳出多层for循环
func forDemo()  {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'a'; j <= 'z'; j++ {
			//if i == 5 && j == 'd' {
			if j == 'd' {
				// 跳出内层循环
				flag = true
				break
			}
			fmt.Printf("%v - %c\n", i, j)
		}

		if flag {
			// 跳出外层循环
			break
		}
	}
}