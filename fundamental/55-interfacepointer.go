package main

import "fmt"

// 接口 - interface
/*
   接口(interface)是一种类型，是一种特殊的类型，它规定了变量有哪些方法
*/

type barker1 interface {
	bark()         // 只要实现了bark方法的变量都是barker类型，方法签名
}

type cat2 struct {
	name string
}

type dog2 struct {
	name string
}

type person2 struct {
	name string
}

// 使用指针接收者实现了接口的所有方法
func (c *cat2) bark()  {
	fmt.Println("喵喵喵~~~~~")
}

func (d *dog2) bark()  {
	fmt.Println("汪汪汪~~~~")
}

func (p *person2) bark()  {
	fmt.Println("啊啊啊~~~~")
}

/*
    总结：使用值接收者实现接口与使用指针接收者实现接口的区别？
	- 使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存储。
    - 使用指针接收者实现接口，只能存储结构体指针类型的变量。
 */
func roar1(x barker1)  {
	x.bark()
}

func main()  {
	var b1 barker1
	c1 := cat2{
		name: "Tommy",
	}

	d1 := dog2{
		name: "Dalk",
	}

	p1 := person2{
		name: "Kim",
	}

	b1 = &c1
	/*
		b1==========1=====> &{Tommy}
	 */
	fmt.Println("b1==========1=====>", b1)
	b1 = &d1
	/*
		b1==========2=====> &{Dalk}
	 */
	fmt.Println("b1==========2=====>", b1)
	b1 = &p1
	/*
		b1==========3=====> &{Kim}
	 */
	fmt.Println("b1==========3=====>", b1)

	/*
		喵喵喵~~~~~
		汪汪汪~~~~
		啊啊啊~~~~
	 */
	roar1(&c1)
	roar1(&d1)
	roar1(&p1)

}