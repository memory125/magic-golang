package main

import (
	"flag"
	"fmt"
)

// flag示例
/*
   flag: 获取命令行参数
*/

func main() {
	// 创建标志位参数
	// 这里name返回的是*string, age返回的是*int
	/*
			使用时：exe -name value -age value
		    查询使用方法：exe -help
	*/
	name := flag.String("name", "Jack", "Please input the name")
	age := flag.Int("age", 26, "Please input the age")
	flag.Parse()
	/*
		name is Jack.
	*/
	fmt.Printf("name is %v, age is %v \n", *name, *age)

	fmt.Println("flag.Args() ", flag.Args())   // 返回命令行参数的其他参数，[]string类型
	fmt.Println("flag.NArg() ", flag.NArg())   // 返回命令行参数后的其他参数个数
	fmt.Println("flag.NFlag() ", flag.NFlag()) // 返回使用命令行参数个数
}
