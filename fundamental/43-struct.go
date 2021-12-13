package main

import "fmt"

// 结构体

type Person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {
	// 声明一个Person类型的变量p
	var p Person
	// 通过字段赋值
	p.name = "David"
	p.age = 18
	p.gender = "Male"
	p.hobby = []string{"Swimming", "Music", "Basketball", "Coding"}

	/*
		{David 18 Male [Swimming Music Basketball Coding]}
		Type of p is main.Person
		name:  David
	 */
	fmt.Println(p)
	fmt.Printf("Type of p is %T \n", p)
	fmt.Println("name: ", p.name)

	var p2 Person
	// 通过字段赋值
	p2.name = "Jack"
	p2.age = 18
	/*
	   未赋值的结构体成员，默认是空值，如int型，默认0，string类型，默认空
		{Jack 18  []}
	 */
	fmt.Println(p2)
}