package main

import (
	"fmt"
	"sync"
)

// 锁

var x = 0
var wg4 sync.WaitGroup
var lock1 sync.Mutex     // 互斥锁

func addData()  {
	defer wg4.Done()
	for i := 0; i < 5000000; i++ {
		// 加锁
		lock1.Lock()
		x++
		// 解锁
		lock1.Unlock()
	}
}

func main()  {
	wg4.Add(1)
	go addData()
	wg4.Add(1)
	go addData()
	wg4.Wait()

	/*
		x is 10000000.
	 */
	fmt.Printf("x is %v.\n", x)
}
