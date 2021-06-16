package main

import "fmt"

func main()  {
	stringDemo("上海自来水来自海上")
	stringDemo("山西煤运车运煤西山")
	stringDemo("黄山落叶松叶落山黄")
	stringDemo("abcdefg")
}

// 判断字符串是否为回文，如“上海自来水来自海上”，“山西煤运车运煤西山”，“黄山落叶松叶落山黄”
func stringDemo(str string)  {
	fmt.Println("str: ", str)
	// 把字符串总的字符存档到一个[]rune中，通过uint8值来判断
	r := make([]rune, 0, len(str))
	for _, c := range str{
		r = append(r, c)
	}

	// 输出原始内容
	fmt.Println("[]rune: ", r)

	// 判断首位是否一致
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r) - 1 - i] {
			fmt.Println("不是回文！")
		}
	}
	fmt.Println("是回文！")
}