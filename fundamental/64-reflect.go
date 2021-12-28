package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Printf("Type %v is %v.\n ", x, v)
	fmt.Printf("Type %v, Name is %v, Kind is %v.\n", x, v.Name(), v.Kind())
}

func main()  {
	var a float32 = 3.14265
	/*
	   Type 3.14265 is float32.
	   Type 3.14265, Name is float32, Kind is float32.
	 */
	reflectType(a)

	var b int64 = 66666
	/*
		Type 66666 is int64.
	 	Type 66666, Name is int64, Kind is int64.
	 */
	reflectType(b)
}
