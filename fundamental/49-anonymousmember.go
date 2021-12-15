package main

import "fmt"

// 结构体匿名成员

// 下例中的string和int就是匿名成员
type anonymember struct {
	string
	int
}

func main()  {
	a := anonymember{
		"Java",
		30,
	}

	/*
		a===========> {Java 30}
		Java
		30
	 */
	fmt.Println("a===========>", a)
	// 访问成员
	fmt.Println(a.string)
	fmt.Println(a.int)
}
