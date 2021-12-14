package main

import "fmt"

// 方法

type dog struct {
	name string
}

type staff struct {
	name string
	age int
	salary float32
}

// 构造函数
func newDog(name string) dog  {
	return dog{name}
}

func newStaff(name string, age int, salary float32) staff {
	return staff{
		name,
		age,
		salary,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) bark()  {
	fmt.Println("Woo, Woo, Woo....", d.name)
}

func (s *staff) addAge()  {
	s.age++
}

func (s *staff) addSalary(salary float32)  {
	s.salary += salary
}

func main()  {
	d1 := newDog("Wangcai")
	/*
		Woo, Woo, Woo.... Wangcai
	 */
	d1.bark()

	s1 := newStaff("Kevin", 32, 6866.66)
	/*
		{Kevin 32 6866.66}
	 */
	fmt.Println(s1)

	s1.addAge()
	/*
		{Kevin 33 6866.66}
	 */
	fmt.Println(s1)

	s1.addSalary(2000.22)
	/*
		{Kevin 33 8866.88}
	 */
	fmt.Println(s1)
}
