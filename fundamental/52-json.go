package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1. 序列化：把Go语言中的结构体变量  ----> json格式的字符串
// 2. 反序列化：json格式的字符串  ----> Go语言中能够识别的结构体变量

/*
   注意事项：
    1. 结构体内部的字段首字母要大写！！！不大写外部访问不到
    2. 反序列化时要传递指针！！！
*/
type personInfo1 struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age" db:"age" ini:"age"`
}

func main() {
	p1 := personInfo1{
		Name: "Gavin",
		Age:  25,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("p1 Marshal failed! Error is %v.\n", err)
	}

	/*
		{"name":"Gavin","age":25}
	*/
	fmt.Println(string(b))

	// 反序列化
	str := `{"name": "Lee", "age": 36}`
	var p2 personInfo1
	json.Unmarshal([]byte(str), &p2) // 传指针是为了在json.Unmarshal内部修改p2的值
	/*
		p2: main.personInfo1{Name:"Lee", Age:36}
	*/
	fmt.Printf("p2: %#v\n", p2)
}
