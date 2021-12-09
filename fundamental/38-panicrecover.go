package main

import "fmt"

func funcA()  {
	fmt.Println("A")
}

func funcB()  {
	// 假设这里打开了数据库的连接
	/*
		1. recover必须搭配defer使用
		2. defer一定要在可能引发panic的语句之前定义
	 */
	defer func() {
		// 修复程序
		err := recover()
		if err != nil {
			// fmt.Println(err)
			fmt.Printf("Error is %v\n", err)
		}
		fmt.Println("关闭数据库连接！")
	}()
	// 程序crash退出
	panic("There has a severe error!!!")
	fmt.Println("B")
}

func funcC()  {
	fmt.Println("C")
}

func main()  {
	/*
	未添加recover修复程序之前：
		A
		关闭数据库连接！
		panic: There has a severe error!!!

		goroutine 1 [running]:
		main.funcB(...)
			/38-panicrecover.go:11
		main.main()
			/38-panicrecover.go:21 +0x66

	添加了recover修复程序之后：
		A
		There has a severe error!!!
		关闭数据库连接！
		C
	 */
	funcA()
	funcB()
	funcC()
}
