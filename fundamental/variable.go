package main

import "fmt"

/*
	声明变量:
	格式: var 变量名 类型
*/
// 这里是全局变量
var name1 string
var age1 int
var isOk1 bool

/*
	批量声明
 */
// 全局变量
var (
	name2 string             // ”“
	age2 int                 // 0
	isOk2 bool               // false
)

/*
格式化打印占位符：
            %v,原样输出
            %T，打印类型
            %t,bool类型
            %s，字符串
            %f，浮点
            %d，10进制的整数
            %b，2进制的整数
            %o，8进制
            %x，%X，16进制
                %x：0-9，a-f
                %X：0-9，A-F
            %c，打印字符
            %p，打印地址
 */

// Go语言中推荐使用驼峰命名法: 可参考下述方式
// 推荐第二种方式，和Java中的类似
var student_name string
var studentName string
var StudentName string

func main()  {
	// 全局变量赋值
	name1 = "David"
	age1 = 12
	isOk1 = true

	// Go语言中变量声明后必须使用
	fmt.Println("================Name1======================")
	fmt.Printf("name: %s\n", name1)
	fmt.Printf("age: %d\n", age1)
	fmt.Printf("isOk: %v\n", isOk1)

	// 全局变量赋值
	name2= "Jack"
	age2 = 18
	isOk2 = false
	fmt.Println("================Name2======================")
	fmt.Printf("name: %s\n", name2)
	fmt.Printf("age: %d\n", age2)
	fmt.Printf("isOk: %v\n", isOk2)

	// 局部变量
	// 声明同时并赋值
	var s1 string = "Joke"
	fmt.Println(s1)

	// 类型推导
	var a = 1
	var b = 2
	fmt.Println(a + b)

	// 简短变量声明
	s2 := "Kevin"
	fmt.Println(s2)
}
