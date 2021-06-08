package main

import "fmt"

func main()  {
	operationDemo()
	relationDemo()
	logicDemo()
	bitDemo()
}

func operationDemo()  {
	var (
		x = 6
		y = 3
	)

	/*
		9
		3
		18
		2
		0
		7 2
	 */
	// 加
	fmt.Println(x + y)
	// 减
	fmt.Println(x - y)
	// 乘
	fmt.Println(x * y)
	// 除
	fmt.Println(x / y)
	// 求模
	fmt.Println(x % y)
	x++            // x = x + 1
	y--            // y = y - 1
	fmt.Println(x, y)
}

func relationDemo()  {
	var (
		x = 8
		y = 6
	)

	/*
		false
		true
		true
		true
		false
		false
	 */
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x >= y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x < y)
}

func logicDemo()  {
	var age = 30

	// && 逻辑与
	/*
		Working...
	 */
	if age > 18 && age < 65 {
		fmt.Println("Working...")
	} else {
		fmt.Println("Happy...")
	}

	// || 逻辑或
	/*
		Working...
	*/
	if age < 18 || age > 65 {
		fmt.Println("Happy...")
	} else {
		fmt.Println("Working...")
	}

	// ! 逻辑非
	isOk := true
	if !isOk{
		fmt.Println("Something is not ok...")
	} else {
		fmt.Println("Something is ok...")
	}
}

func bitDemo()  {
	var (
		x = 5
		y = 2
	)

	/*
		0
		7
		7
		20
		1
	 */
	// 与
	fmt.Println(x & y)
	// 或
	fmt.Println(x | y)
	// 抑或
	fmt.Println(x ^ y)
	// 左移
	fmt.Println(x << y)
	// 右移
	fmt.Println(x >> y)
}