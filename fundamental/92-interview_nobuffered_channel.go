package main

import "fmt"

// 面事 - 无缓冲的channel

var strNoBuf string
var noBufferedCh = make(chan int)

func noBufferedChanTest() {
	strNoBuf = "hello, no buffered channel"
	<-noBufferedCh
}

func main() {
	go noBufferedChanTest()
	noBufferedCh <- 0
	/*
		hello, no buffered channel
	*/
	fmt.Println(strNoBuf)
	defer close(noBufferedCh)
}
