package main

import "fmt"

func main()  {
	/*
		不带参数，没有返回值
		f1
		带参数，没有返回值
		f2 7
		带参数，有返回值
		函数不同格式
		f3
		11
		f4
		11
		f5
		11
		变长参数
		f6
		10
		[1st 2nd 3rd 4th 5th]
		f7
		10
		Jack
		10 Jack
	 */
	fmt.Println("不带参数，没有返回值")
	f1()
	fmt.Println("带参数，没有返回值")
	f2(3, 4)
	fmt.Println("带参数，有返回值")
	fmt.Println("函数不同格式")
	ret1 := f3(5, 6)
	fmt.Println(ret1)
	ret2 := f4(5, 6)
	fmt.Println(ret2)
	ret3 := f5(5, 6)
	fmt.Println(ret3)
	fmt.Println("变长参数")
	f6(10, "1st", "2nd", "3rd", "4th", "5th")

	fmt.Println("多个返回值")
	v1, v2 := f7(10, "Jack")
	fmt.Println(v1, v2)

}

// 不带参数，且没有返回值
func f1()  {
	fmt.Println("f1")
}

// 带参数，但没有返回值
func f2(x int, y int)  {
	fmt.Println("f2", x + y)
}

// 带参数，且有返回值，返回值为int类型,f3的第一种方式
func f3(x int, y int) int {
	fmt.Println("f3")
	ret := x + y
	return ret
}

// 带参数，且有返回值，返回值为int类型，f3的第二种方式
func f4(x int, y int) (ret int) {
	fmt.Println("f4")
	ret = x + y
	return
}

// 带参数，且有返回值，返回值为int类型，f3的第三种方式
func f5(x , y int) (ret int) {
	fmt.Println("f5")
	ret = x + y
	return
}

// 变长参数
func f6(x int, y...string)  {
	fmt.Println("f6")

	fmt.Println(x)
	fmt.Println(y)
}

// 多个返回值
func f7(x int, y string) (int, string)  {
	fmt.Println("f7")

	fmt.Println(x)
	fmt.Println(y)

	return x, y
}
