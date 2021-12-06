package main

import "fmt"

func main()  {
	deferDemo()
}

// defer: 延时执行
/*
   Go语言中的return不是原子操作，在底层是分为两步来执行
	1. 返回值赋值
	2. defer语句
	3. 真正的RET返回
   函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间
 */
func deferDemo()  {
	/*
		============== Begin ================
		============== End ================
		GaGaGa
		AhAhAh
		HoHoHo

	*/
	// defer把它后面的语句延迟到函数即将返回的时候在执行
	// 多个defer语句按照先进后出(后进先出)的顺序延迟执行
	fmt.Println("============== Begin ================")
	defer fmt.Println("HoHoHo")
	defer fmt.Println("AhAhAh")
	defer fmt.Println("GaGaGa")
	fmt.Println("============== End ================")
}