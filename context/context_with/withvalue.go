package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context - WithValue示例

type TraceCode string

var wgWithValue sync.WaitGroup

func testWithValue(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("Invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("testWithValue, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("testWithValue done!")
	wgWithValue.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "232242323124")
	wgWithValue.Add(1)
	go testWithValue(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wgWithValue.Wait()
	fmt.Println("Over")
}
