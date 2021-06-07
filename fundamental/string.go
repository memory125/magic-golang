package main

import (
	"fmt"
	"strings"
)

func main()  {
	// 声明字符串：文件路径
	path := "D:\\project\\go\\gopath\\src\\code.wing.com\\fundamental"

	// 获取字符串的长度
	/*
		len = 50
	 */
	fmt.Printf("len = %d\n", len(path))


	// 字符串拼接
	name := "Jack"
	gender := "Male"
	info := name + gender
	/*
		info = JackMale
	 */
	fmt.Printf("info = %v\n", info)

	// 使用函数拼接
	/*
		s0 = info is Jack Male
	 */
	s0 := fmt.Sprintf("info is %s %s", name, gender)
	fmt.Printf("s0 = %v\n", s0)

	// 字符串截取
	/*
		s1 = [D: project go gopath src code.wing.com fundamental]
	 */
	s1 := strings.Split(path, "\\")
	fmt.Printf("s1 = %s\n", s1)


	// 循环输出
	/*
		D:
		project
		go
		gopath
		src
		code.wing.com
		fundamental
	 */
	for _, str := range s1 {
		fmt.Println(str)
	}

	// 可以使用反引号``声明字符串，如下
	s2 := `
               春晓
			春眠不觉晓，
			处处闻啼鸟。
			夜来风雨声，
			花落知多少。
		`
	fmt.Println(s2)

	// 包含
	/*
		Contains 'info' -------- false
		Contains 'go' -------- true
	 */
	fmt.Printf("Contains 'info' -------- %v\n", strings.Contains(path,"info"))
	fmt.Printf("Contains 'go' -------- %v\n", strings.Contains(path,"go"))

	/*
		HasPrefix 'info' -------- false
		HasSuffix 'go' -------- false
	 */
	// 前缀
	fmt.Printf("HasPrefix 'info' -------- %v\n", strings.HasPrefix(path,"info"))
	// 后缀
	fmt.Printf("HasSuffix 'go' -------- %v\n", strings.HasSuffix(path,"go"))

	// 获取字符出线的位置
	s3 := "Hello World"
	/*
		The first position of 'e' is -------- 1
		The last position of 'l' is -------- 9
	 */
	fmt.Printf("The first position of 'e' is -------- %v\n", strings.Index(s3, "e"))
	fmt.Printf("The last position of 'l' is -------- %v\n", strings.LastIndex(s3, "l"))

}
