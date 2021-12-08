package main

import (
	"fmt"
	"strings"
)

// 闭包示例

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasPrefix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main()  {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	/*
		test.jpg
		test.jpg.jpg
		test.txt.jpg
		test.txt
		text.txt.txt
		text.jpg.txt
	 */
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(jpgFunc("test.txt"))
	fmt.Println(txtFunc("test"))
	fmt.Println(txtFunc("text.txt"))
	fmt.Println(txtFunc("text.jpg"))
}
