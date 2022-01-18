package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine之间通信 - bool

// 全局变量
var wgBool sync.WaitGroup
var notify bool

// go routine测试函数
func testBool() {
	defer wgBool.Done()
	var i int = 1

	for {
		fmt.Println("=========TestBool======== Time: ", i)
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
		i++
	}
}

/*
	=========TestBool======== Time:  1
	=========TestBool======== Time:  2
	=========TestBool======== Time:  3
	=========TestBool======== Time:  4
	=========TestBool======== Time:  5
	=========TestBool======== Time:  6
	=========TestBool======== Time:  7
	=========TestBool======== Time:  8
	=========TestBool======== Time:  9
	=========TestBool======== Time:  10
*/
func main() {
	wgBool.Add(1)
	go testBool()
	time.Sleep(time.Second * 5)
	// 通知goroutine子程序退出
	notify = true
	wgBool.Wait()
}
