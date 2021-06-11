package main

import "fmt"

func main()  {
	sliceCopy()
}

func sliceCopy()  {
	a1 := []int {1, 3, 5}
	a2 := a1
	a3 := make([]int, 3)
	// 复制
	copy(a3, a1)

	/*
		[1 3 5] [1 3 5] [1 3 5]
	 */
	fmt.Println(a1, a2, a3)

	/*
		[20 3 5] [20 3 5] [1 3 5]
	 */
	a1[0] = 20
	fmt.Println(a1, a2, a3)

	// 将a1中的索引为1的3这个元素删掉
	/*
		[20 5 5] [20 5] [1 3 5]
	 */
	// 修改了底层数组
	a2 = append(a2[:1], a2[2:]...)
	fmt.Println(a1, a2, a3)
}
