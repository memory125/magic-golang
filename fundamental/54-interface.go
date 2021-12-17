package main

import "fmt"

// 接口 - interface
/*
   接口(interface)是一种类型，是一种特殊的类型，它规定了变量有哪些方法
 */

type barker interface {
	bark()         // 只要实现了bark方法的变量都是barker类型，方法签名
}

type cat1 struct {

}

type dog1 struct {

}

type person1 struct {

}

func (a cat1) bark()  {
	fmt.Println("喵喵喵~~~~~")
}

func (d dog1) bark()  {
	fmt.Println("汪汪汪~~~~")
}

func (p person1) bark()  {
	fmt.Println("啊啊啊~~~~")
}

func roar(x barker)  {
	x.bark()
}

func main()  {
	var c1 cat1
	var d1 dog1
	var p1 person1

	/*
		=======================
		喵喵喵~~~~~
		汪汪汪~~~~
		啊啊啊~~~~
	 */
	fmt.Println("=======================")
	roar(c1)
	roar(d1)
	roar(p1)
}