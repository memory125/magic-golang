package main

import "fmt"

// 面事 - 有缓冲的channel

var strBuf string
var bufferedCh = make(chan int, 10)

func bufferedChanTest() {
	strBuf = "hello, buffered channel"
	bufferedCh <- 0
}

func main() {
	go bufferedChanTest()
	temp := <-bufferedCh
	/*
		hello, buffered channel
		0
	*/
	fmt.Println(strBuf)
	fmt.Println(temp)
	defer close(bufferedCh)
}
