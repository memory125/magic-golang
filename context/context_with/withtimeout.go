package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context - WithTimeout示例

var wgWithTimeOut sync.WaitGroup

func testWithTimeOut(ctx context.Context) {
LOOP:
	for {
		fmt.Println("Database connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("testWithTimeOut done!")
	wgWithTimeOut.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wgWithTimeOut.Add(1)
	go testWithTimeOut(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wgWithTimeOut.Wait()
	fmt.Println("Over")
}
