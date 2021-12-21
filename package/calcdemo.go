package main

import (
	calc "code.wing.com/package/calc"
	"fmt"
)

func main()  {
	ret := calc.Add(50, 80)
	/*
		ret:  130
	 */
	fmt.Println("ret: ", ret)
}
