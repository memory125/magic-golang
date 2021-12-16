package main

import "fmt"

// 继承 - 结构体模拟实现其它语言中的“继承”

type animalInfo struct {
	name string
}

type dogInfo struct {
	feet uint8
	animalInfo          // 匿名结构体嵌套，animal拥有的方法，dog此时也有了，实现了“继承”
}

// 给animal实现的一个move的方法
func (a animalInfo) move()  {
	fmt.Printf("%s is moving!!!\n", a.name)
}

// dog实现一个bark的方法
func (d dogInfo) bark()  {
	fmt.Printf("%s is barking!!!!!\n", d.name)
}


func main()  {
	d1 := dogInfo{
		feet: 4,
		animalInfo:animalInfo{
			name: "WangCai",
		},
	}

	/*
		{4 {WangCai}}
		WangCai is barking!!!!!
		WangCai is moving!!!
	 */
	fmt.Println(d1)
	d1.bark()
	d1.move()
}



