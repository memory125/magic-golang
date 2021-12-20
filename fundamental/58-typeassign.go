package main

import (
	"fmt"
)

// 类型断言

func assign1(a interface{})  {
	fmt.Printf("type: %T.\n", a)
	v, ok := a.(string)
	if !ok {
		fmt.Printf("a is not the instance of string, but the instance of %T.\n", a)
	} else {
		fmt.Printf("a is the instance of string, and value is %v.\n", v)
	}
}

func assign2(a interface{})  {
	fmt.Printf("type: %T.\n", a)
	switch a.(type) {
	case string:
		fmt.Println("a is string, the value is ", a.(string))
	case int:
		fmt.Println("a is int, the value is ", a.(int))
	case int8:
		fmt.Println("a is int8, the value is ", a.(int8))
	case int16:
		fmt.Println("a is int8, the value is ", a.(int16))
	case int32:
		fmt.Println("a is int32, the value is ", a.(int32))
	case int64:
		fmt.Println("a is int64, the value is ", a.(int64))
	case uint:
		fmt.Println("a is uint, the value is ", a.(uint))
	case uint8:
		fmt.Println("a is uint8, the value is ", a.(uint8))
	case uint16:
		fmt.Println("a is uint16, the value is ", a.(uint16))
	case uint32:
		fmt.Println("a is uint32, the value is ", a.(uint32))
	case uint64:
		fmt.Println("a is uint64, the value is ", a.(uint64))
	case float32:
		fmt.Println("a is float32, the value is ", a.(float32))
	case float64:
		fmt.Println("a is float64, the value is ", a.(float64))
	case bool:
		fmt.Println("a is bool, the value is ", a.(bool))
	default:
		fmt.Println("Unsupported type.")
	}
}

func main()  {
	/*
		type: string.
		a is the instance of string, and value is GoLang.
		type: int.
		a is not the instance of string, but the instance of int.
		type: <nil>.
		a is not the instance of string, but the instance of <nil>.
		type: float64.
		a is not the instance of string, but the instance of float64.
	 */
	assign1("GoLang")
	assign1(100)
	assign1(nil)
	assign1(3.02112)

	//
	fmt.Println("=======================================")
	/*
		type: string.
		a is string, the value is  Jacky
		type: bool.
		a is bool, the value is  true
		type: int.
		a is int, the value is  100
		type: int8.
		a is int8, the value is  10
		type: int16.
		a is int8, the value is  300
		type: int32.
		a is int32, the value is  400
		type: int64.
		a is int64, the value is  500
		type: uint.
		a is uint, the value is  100
		type: uint8.
		a is uint8, the value is  10
		type: uint16.
		a is uint16, the value is  300
		type: uint32.
		a is uint32, the value is  400
		type: uint64.
		a is uint64, the value is  500
		type: float32.
		a is float32, the value is  2.23232
		type: float64.
		a is float64, the value is  3.6e+60
	 */
	assign2("Jacky")
	assign2(true)
	assign2(int(100))
	assign2(int8(10))
	assign2(int16(300))
	assign2(int32(400))
	assign2(int64(500))
	assign2(uint(100))
	assign2(uint8(10))
	assign2(uint16(300))
	assign2(uint32(400))
	assign2(uint64(500))
	assign2(float32(2.23232))
	assign2(float64(3.6E60))
}