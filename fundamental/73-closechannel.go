package main

import "fmt"

// 关闭通道

func main() {
	// 初始化通道
	ch1 := make(chan int, 2)

	// 发送
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	// 接收
	<-ch1
	<-ch1

	x1, ok := <-ch1
	if !ok {
		/*
			Get data from channel failed! 0 false
		*/
		fmt.Println("Get data from channel failed!", x1, ok)
	}

	x2, ok := <-ch1
	/*
		0 false
	*/
	fmt.Println(x2, ok)

	x3, ok := <-ch1
	/*
		0 false
	*/
	fmt.Println(x3, ok)
}
