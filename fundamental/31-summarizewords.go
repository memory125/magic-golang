package main

import (
	"fmt"
	"strings"
)

func main()  {
	summarizeWords("how do you do and do you know do you")
}

// 统计单词的个数
func summarizeWords(words string) {
	// 把字符串按照空格切割得到切片
	str := strings.Split(words, " ")
	// 定义map，遍历切片存储到map
	m := make(map[string]int, 20)
	for _, word := range str {
		// 获取map中的key信息
		_, ok := m[word]
		if !ok {
			// 1. 如果原来map中不存在这个key，那么出现次数为1
			m[word] = 1
		} else {
			// 2. 如果原来map中存在这个key，那么出现次数+1
			m[word]++
		}
	}

	// 输出结果
	fmt.Println(m)
	// 或者
	for key, value := range m{
		fmt.Printf("The '%v' -  %d times.\n", key, value)
	}
}
