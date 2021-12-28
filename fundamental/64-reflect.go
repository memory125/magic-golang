package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Println("type x is ", v)
}

func main()  {
	var a float32 = 3.14265
	/*
	   type x is  float32
	 */
	reflectType(a)

	var b int64 = 66666
	/*
		type x is  int64
	 */
	reflectType(b)
}
