package main

import (
	calc "code.wing.com/package/calc"
	"fmt"
)

// 全局变量
var x int = 20

const pi = 3.1415926

func init()  {
	fmt.Printf("x = %v, pi = %v.\n", x, pi)
}

func main()  {
	ret := calc.Add(50, 80)
	/*
		-------------------------Module cals---------------------
		x = 20, pi = 3.1415926.
		ret:  130
	 */
	fmt.Println("ret: ", ret)
}
