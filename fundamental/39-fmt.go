package main

import "fmt"

// fmt demo

func main()  {
	/*
	------fmt.Print------------fmt.Printf------
	------fmt.Println------
	 */
	fmt.Print("------fmt.Print------")
	/*
	fmt.Printf()：格式化字符串，值
	  %T：查看类型
	  %d：十进制数
	  %b：二进制数
	  %o：八进制数
	  %x：十六进制数
	  %c：字符
	  %s：字符串
	  %p：指针
	  %v：值
	  %f：浮点数
	  %t：布尔值
	 */
	fmt.Printf("------fmt.Printf------\n")
	fmt.Println("------fmt.Println------")

	// %v %#v
	var m = make(map[string]int, 1)
	m["David"] = 36
	/*
	m(v) = map[David:36]
	m(#v) = map[string]int{"David":36}
	 */
	fmt.Printf("m(v) = %v\n", m)
	fmt.Printf("m(#v) = %#v\n", m)

	/*
		96%
	 */
	printPercentage(96)

	// 整数 -> 字符
	// 'a'
	fmt.Printf("%q\n", 97)
	// 浮点数和复数
	// 7074237808027956p-51
	fmt.Printf("%b\n", 3.141592678458574)
}

// 打印百分比
func printPercentage(num int)  {
	fmt.Printf("%d%%\n", num)
}
