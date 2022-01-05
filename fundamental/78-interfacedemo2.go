package main

import "fmt"

// interface示例2 - 可变参数

func testInterface(a ...interface{}) {
	fmt.Printf("Type: %T, Value: %#v.\n", a, a)
}

func main() {
	var s = []interface{}{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	/*
			Type: []interface {}, Value: []interface {}{[]interface {}{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}}.
		    备注：s是一个整体，s的第一个元素是一个interface类型的切片
	*/
	testInterface(s)
	/*
		Type: []interface {}, Value: []interface {}{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}.
		备注：s是一个切片
	*/
	testInterface(s...)
}
