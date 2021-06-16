package main

import "fmt"

func main()  {
	deferDemo()
}

// defer: 延时执行
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