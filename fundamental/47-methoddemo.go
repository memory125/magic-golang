package main

import "fmt"

// 自定义类型方法

// 自定义类型
type myInt int

type cat struct {
	name string
	age int
}

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

	// 结构体初始化
	// 方式1
	var c1 cat
	c1.name = "Loo"
	c1.age = 1
	/*
		c1=======> {Loo 1}
	 */
	fmt.Println("c1=======>", c1)

	// 方式2
	var c2 cat
	c2 = cat{
		name: "Poo",
		age : 2,
	}
	/*
		c2=======> {Poo 2}
	 */
	fmt.Println("c2=======>", c2)

	// 方式3
	var c3 cat
	c3 = cat{
		"Too",
		 3,
	}
	/*
		c3=======> {Too 3}
	 */
	fmt.Println("c3=======>", c3)

	// 方式4
	c4 := cat{
		"Noo",
		4,
	}
	/*
		c4=======> {Noo 4}
	 */
	fmt.Println("c4=======>", c4)
}