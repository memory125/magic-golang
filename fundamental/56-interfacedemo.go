package main

import "fmt"

// 同一个结构体可以实现多个接口
// 接口还可以嵌套

// 接口的嵌套
type animal interface {
	mover
	eater
}

// mover接口
type mover interface {
	move()
}

// eater接口
type eater interface {
	eat(string)
}

type cat3 struct {
	name string
	feet uint16
}

// cat实现了mover接口
func (c *cat3) move()  {
	fmt.Printf("%s is moving!!!!!\n", c.name)
}

// cat实现了eater接口
func (c *cat3) eat(food string)  {
	fmt.Printf("%s is eating %v!!!!!\n", c.name, food)
}

func main()  {
	var a animal
	c1 := &cat3{
		name: "Timmy",
		feet: 4,
	}
	a = c1
	/*
		Timmy is moving!!!!!
		Timmy is eating fish!!!!!
	 */
	a.move()
	a.eat("fish")
}