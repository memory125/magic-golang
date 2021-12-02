package main

import "fmt"

func main()  {
	sliceAppend()
}

func sliceAppend()  {
	s1 := []string {"Java", "C", "C++", "Go", "Rust"}

	/*
		[Java C C++ Go Rust]
		len(s1) = 5, cap(s1) = 5
	 */
	fmt.Println(s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))


	// 添加数据
	// 调用append函数必须用原来的切片变量接收返回值
	/*
		[Java C C++ Go Rust C#]
		len(s1) = 6, cap(s1) = 10
	 */
	// append追加元素，原来的底层数组放不下的时候，Go底层就会把底层数组换一个
	s1 = append(s1, "C#")
	fmt.Println(s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	/*
		[Java C C++ Go Rust C# JS Ruby PHP VB Kotlin Swift Object-C]
		len(s1) = 13, cap(s1) = 20
	 */
	s2 := []string {"JS", "Ruby", "PHP", "VB", "Kotlin", "Swift", "Object-C"}
	s1 = append(s1, s2...)         // ...表示将s2拆开
	fmt.Println(s1)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))
}
