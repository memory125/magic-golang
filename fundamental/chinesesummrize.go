package main

import (
	"fmt"
	"unicode"
)

func main()  {
	summarizeChinese("Hello,我的世界！")
}
// 统计字符串中汉字的数目
func summarizeChinese(str string)  {
	//str := "Hello, 世界"
	fmt.Println(str)
	var chineseCount int

	for _, r := range str {
		// 判断中文的关键点：unicode.Han
		if unicode.Is(unicode.Han, r) {
			chineseCount++
			fmt.Print(string(r))
			fmt.Println()
		}
	}

	fmt.Printf("The Chinese Count is totally %d\n.", chineseCount)
}
