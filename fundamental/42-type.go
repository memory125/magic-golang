package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型
type myType int  // 自定义类型
type ownType = int  // 类型别名

func main()  {
	var m myType
	m = 100
	/*
		100
		type of m is main.myType
	*/
	fmt.Println(m)
	fmt.Printf("type of m is %T \n", m)

	var n ownType
	n = 100
	/*
		100
		type of n is int
	*/
	fmt.Println(n)
	fmt.Printf("type of n is %T \n", n)
}