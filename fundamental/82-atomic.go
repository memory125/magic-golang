package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作

var x3 int64
var wg7 sync.WaitGroup

func addOp1() {
	defer wg7.Done()
	//x++
	// 原子操作
	/*
		  类似于下述的操作：
			lock.Lock()
			x3++
			lock.Unlock()
	*/
	atomic.AddInt64(&x3, 1)
}

func main() {
	for i := 0; i < 1000000; i++ {
		wg7.Add(1)
		go addOp1()
	}
	wg7.Wait()

	fmt.Printf("Automic Operation ========== > finally x3 is %v.\n", x3)
}
