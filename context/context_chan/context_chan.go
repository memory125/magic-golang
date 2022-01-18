package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine之间通信 - chan

// 全局变量
var wgChan sync.WaitGroup
var exitChan chan bool = make(chan bool, 1)

// go routine测试函数
func testChan() {
	defer wgChan.Done()
	var i int = 1

LOOP:
	for {
		fmt.Println("=========testChan======== Time: ", i)
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
		i++
	}
}

/*
	=========testChan======== Time:  1
	=========testChan======== Time:  2
	=========testChan======== Time:  3
	=========testChan======== Time:  4
	=========testChan======== Time:  5
	=========testChan======== Time:  6
	=========testChan======== Time:  7
	=========testChan======== Time:  8
	=========testChan======== Time:  9
	=========testChan======== Time:  10
*/
func main() {
	wgChan.Add(1)
	go testChan()
	time.Sleep(time.Second * 5)
	// 通知goroutine子程序退出
	exitChan <- true
	wgChan.Wait()
}
