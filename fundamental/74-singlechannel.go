package main

import "fmt"

// 单向通道

/*
	chan<- int是一个只写单向通道（只能对其写入int类型值），可以对其执行发送操作但是不能执行接收操作；
	<-chan int是一个只读单向通道（只能从其读取int类型值），可以对其执行接收操作但是不能执行发送操作。
*/

// out是一个只写的通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

// out是一个只写的通道，in是一个只读的通道
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

// in是一个只读的通道
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
