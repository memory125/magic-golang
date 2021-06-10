package main

import "fmt"

// 切片：引用类型
func main()  {
	// 定义切片
	var s1 [] int
	var s2 [] string

	// 打印切片的值
	/*
		s1 = [], s2 = []
	 */
	fmt.Printf("s1 = %v, s2 = %v\n", s1, s2)

	// 判断是否为空
	/*
		true
		true
	*/
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 切片初始化
	/*
		s1 = [1 2 3 4 5 6 7 8 9], s2 = [Jack David Bruce Ham Joe Lily]
	 */
	s1 = []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 = []string {"Jack", "David", "Bruce", "Ham", "Joe", "Lily"}
	// 打印切片的值
	fmt.Printf("s1 = %v, s2 = %v\n", s1, s2)

	// 切片的长度
	/*
		len(s1) = 9, len(s2) = 6
	 */
	fmt.Printf("len(s1) = %v, len(s2) = %v\n", len(s1), len(s2))

	// 判断是否为空
	/*
		false
		false
	 */
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 切片的容量
	/*
		cap(s1) = 9, cap(s2) = 6
	 */
	fmt.Printf("cap(s1) = %v, cap(s2) = %v\n", cap(s1), cap(s2))

	// 从数组得到切片
	/*
		[1 2 3 4 5 6 7 8 9 10]
	 */
	a1 := [...]int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	s3 := a1[0:10]           // 表示从数组a1[0] ~ a[9]的元素赋值给切片 s3
	fmt.Println(s3)
	s4 := a1[:10]           // 表示从数组a1[0] ~ a[9]的元素赋值给切片 s4 ==> s3
	s5 := a1[3:]            // ==> [3:len(a1)]
	s6 := a1[:]             // ==> [0:len(a1)]
	/*
		[1 2 3 4 5 6 7 8 9 10] [4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20] [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	 */
	fmt.Println(s4, s5, s6)
	/*
	   len(s4) = 10, cap(s4) = 20
	   len(s5) = 17, cap(s5) = 17
	   len(s6) = 20, cap(s6) = 20
	 */
	fmt.Printf("len(s4) = %d, cap(s4) = %d\n", len(s4), cap(s4))
	fmt.Printf("len(s5) = %d, cap(s5) = %d\n", len(s5), cap(s5))
	fmt.Printf("len(s6) = %d, cap(s6) = %d\n", len(s6), cap(s6))

	s7 := s5[4:]
	/*
		[8 9 10 11 12 13 14 15 16 17 18 19 20]
		len(s7) = 13, cap(s7) = 13
	 */
	fmt.Println(s7)
	fmt.Printf("len(s7) = %d, cap(s7) = %d\n", len(s7), cap(s7))

	/*
		s3 =  [1 2 3 4 5 6 7 8 9 10]
		s3 =  [1 2 3 4 5 6 7 8 9 100]
	 */
	// 切片是引用类型，都指向了底层的一个具体的数组
	fmt.Println("s3 = ", s3)
	a1[9] = 100
	fmt.Println("s3 = ", s3)
}
