package main

import "fmt"

// 结构体嵌套

// personinfo
type personinfo struct {
	name string
	age int
	addr addrinfo // 结构体嵌套
}

// addrinfo
type addrinfo struct {
	province string
	city string
}

// companyinfo
type companyinfo struct {
	name string
	addr addrinfo       // 结构体嵌套
}

func main()  {
	// person结构体嵌套的初始化
	p1 := personinfo{
		name: "Jack",
		age: 30,
		addr: addrinfo{
			province: "广东省",
			city: "深圳",
		},
	}
	// 输出p1信息
	/*
		person info ===============> {Jack 30 {广东省 深圳}}
	 */
	fmt.Println("person info ===============>", p1)
}
