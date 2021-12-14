package main

import "fmt"

// 自定义类型方法

// 自定义类型
type myInt int

func (m myInt) add1(x, y int) {
	fmt.Println("myInt1-------", x + y)
}

func (m myInt) add2(x, y int) myInt {
	fmt.Println("myInt2-------", x + y)
	return myInt(x + y)
}

func main()  {
	m1 := myInt(10)
	/*
		myInt1------- 30
	 */
	m1.add1(10, 20)

	m2 := myInt(20)
	/*
		myInt2------- 70
		70
	 */
	m3 := m2.add2(30, 40)
	fmt.Println(m3)
}