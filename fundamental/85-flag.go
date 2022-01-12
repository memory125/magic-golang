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
}
