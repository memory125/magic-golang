package main

import "fmt"

// 空接口
// interface：关键字
// interface{}：空接口类型

// 空接口作为函数参数
func showAll(a... interface{})  {
	fmt.Printf("type: %T, value: %v.\n", a, a)
}

func main()  {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 20)

	m1["name"] = "Jane"
	m1["age"] = 26
	m1["married"] = false
	m1["hobby"] = [...]string{"Coding", "Music", "Dancing","Jogging"}
	/*
		m1=================> map[age:26 hobby:[Coding Music Dancing Jogging] married:false name:Jane]
	 */
	fmt.Println("m1=================>", m1)

	/*
		type: []interface {}, value: [false].
		type: []interface {}, value: [<nil>].
		type: []interface {}, value: [map[age:26 hobby:[Coding Music Dancing Jogging] married:false name:Jane]].
		type: []interface {}, value: [See...].
	 */
	showAll(false)
	showAll(nil)
	showAll(m1)
	showAll("See...")
}
