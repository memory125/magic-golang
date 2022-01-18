package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// goroutine之间通信 - context标准库

// 全局变量
var wgContext sync.WaitGroup

// go routine测试函数
func testContext(ctx context.Context) {
	defer wgContext.Done()
	var i int = 1

LOOP:
	for {
		fmt.Println("=========testContext======== Time: ", i)
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
		i++
	}
}

/*
	=========testContext======== Time:  1
	=========testContext======== Time:  2
	=========testContext======== Time:  3
	=========testContext======== Time:  4
	=========testContext======== Time:  5
	=========testContext======== Time:  6
	=========testContext======== Time:  7
	=========testContext======== Time:  8
	=========testContext======== Time:  9
	=========testContext======== Time:  10
*/
func main() {
	wgContext.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go testContext(ctx)
	time.Sleep(time.Second * 5)
	// 通知goroutine子程序退出
	cancel()
	wgContext.Wait()
}
