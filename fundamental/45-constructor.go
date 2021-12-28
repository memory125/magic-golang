package main

import "fmt"

// 构造函数

type Student struct {
	name string
	age int
	gender string
	grade int
}

// 构造函数：约定成俗new开头
// 返回的是结构体还是结构体指针
// 当结构体比较大的时候尽量返回结构体指针，减少程序的内容开销
func newStudent(name string, age int, gender string, grade int) *Student {
	return &Student{
		name,
		age,
		gender,
		grade,
	}
}

func main()  {
	s1 := newStudent("David", 18, "Male", 5 )
	s2 := newStudent("Jane", 16, "Female", 4 )
	/*
		&{David 18 Male 5} &{Jane 16 Female 4}
		s1 type: *main.Student, value: &{David 18 Male 5}
		s2 type: *main.Student, value: &{Jane 16 Female 4}
	*/
	fmt.Println(s1, s2)
	fmt.Printf("s1 type: %T, value: %v \n", s1, s1)
	fmt.Printf("s2 type: %T, value: %v \n", s2, s2)
}
