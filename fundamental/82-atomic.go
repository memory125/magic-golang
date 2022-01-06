package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作

var x int64
var wg6 sync.WaitGroup

func addOp1() {
	defer wg6.Done()
	//x++
	// 原子操作
	/*
		  类似于下述的操作：
			lock.Lock()
			x++
			lock.Unlock()
	*/
	atomic.AddInt64(&x, 1)
}

func main() {
	for i := 0; i < 1000000; i++ {
		wg6.Add(1)
		go addOp1()
	}
	wg6.Wait()

	fmt.Printf("Automic Operation ========== > finally x is %v.\n", x)
}
