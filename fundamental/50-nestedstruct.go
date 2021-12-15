package main

import "fmt"

// 结构体嵌套

// person info
type personInfo struct {
	name string
	age int
	addr addrInfo // 结构体嵌套
}

// addr info
type addrInfo struct {
	province string
	city string
}

// company info
type companyInfo struct {
	name string
	addr addrInfo       // 结构体嵌套
}

//
type homeInfo struct {
	name string
	id int64
	addrInfo           // 匿名结构体嵌套
}

func main()  {
	// person结构体嵌套的初始化
	p1 := personInfo{
		name: "Jack",
		age: 30,
		addr: addrInfo{
			province: "广东省",
			city: "深圳",
		},
	}
	// 输出p1信息
	/*
		person info ===============> {Jack 30 {广东省 深圳}}
	 */
	fmt.Println("person info ===============>", p1)

	// 访问结构体嵌套的成员
	fmt.Println("p1 info", p1.name, p1.addr.city)

	// 匿名结构体嵌套示例
	// 匿名结构体嵌套初始化
	h1 := homeInfo{
		name: "幸福港湾",
		id: 10000000,
		addrInfo: addrInfo{
			province: "四川省",
			city: "成都市",
		},
	}
	// 访问结构体匿名嵌套的成员
	/*
		h1 info 幸福港湾 10000000 成都市
		h1 info 幸福港湾 10000000 成都市
	 */
	fmt.Println("h1 info", h1.name, h1.id, h1.addrInfo.city)
	// h1.city：先在自己结构体自己安排这个成员，找不到就去匿名嵌套的结构体中查找该成员
	fmt.Println("h1 info", h1.name, h1.id, h1.city)
}
