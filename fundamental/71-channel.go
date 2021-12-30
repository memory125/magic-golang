package main

import (
	"fmt"
	"sync"
)

// 通道 - channel

// 声明channel
var ch chan int

var wg sync.WaitGroup

func noCacheChan()  {
	fmt.Println("Ch===1==>", ch)
	ch = make(chan int)       // 不带缓冲区的初始化，通过make函数初始化，slice，map，chan
	// ch = make(chan int, 10)     // 带缓冲区的初始化
	/*
		Ch===2==> 0xc00004c120
	*/
	fmt.Println("Ch===2==>", ch)
	wg.Add(1)
	go func() {        // 从通道中接收数据
		defer wg.Done()
		// 接收值：匿名routine
		va := <- ch
		fmt.Println("va=========>", va)
	}()
	// 发送值: main routine，发送数据到通道
	ch <- 10
	fmt.Println("send value to channel =========>", ch)
	wg.Wait()
	close(ch)
}

func cacheChan()  {
	fmt.Println("Ch===1==>", ch)
	ch = make(chan int, 10)     // 带缓冲区的初始化，初始化多少，channel中就缓存多少数据
	/*
		Ch===2==> 0xc00004c120
	*/
	fmt.Println("Ch===2==>", ch)
	// 发送值: main routine，发送数据到通道
	ch <- 10
	fmt.Println("send value to channel =========>", ch)
	va := <- ch
	fmt.Println("va=========>", va)
	close(ch)
}

func main()  {
	// noCacheChan()
	cacheChan()
}
