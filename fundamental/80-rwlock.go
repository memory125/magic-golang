package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁 - RWLock

var (
	x1     int64
	x2     int64
	wg5    sync.WaitGroup
	wg6    sync.WaitGroup
	lock2  sync.Mutex
	rwlock sync.RWMutex
)

func readTest1() {
	defer wg5.Done()
	lock2.Lock()
	fmt.Printf("x1 = %v.\n", x1)
	lock2.Unlock()
}

func writeTest1() {
	defer wg5.Done()
	lock2.Lock()
	x1 = x1 + 1
	lock2.Unlock()
}

func readTest2() {
	defer wg6.Done()
	rwlock.RLock()
	fmt.Printf("x2 = %v.\n", x2)
	rwlock.RUnlock()
}

func writeTest2() {
	defer wg6.Done()
	rwlock.Lock()
	x2 = x2 + 1
	rwlock.Unlock()
}

func mutexTest() {
	fmt.Println(" ========= Mutex Testing ===========")
	// 互斥锁测试
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		wg5.Add(1)
		go writeTest1()
	}

	// 延时
	time.Sleep(time.Second)

	for i := 0; i < 10000; i++ {
		wg5.Add(1)
		go readTest1()
	}
	wg5.Wait()

	endTime := time.Now()
	fmt.Printf("Mutex Testing ======> finally time costs = %v.\n", endTime.Sub(startTime))
}

func rwmutexTest() {
	fmt.Println(" ========= RWMutex Testing ===========")
	// 读写互斥锁测试
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		wg6.Add(1)
		go writeTest2()
	}

	// 延时
	time.Sleep(time.Second)

	for i := 0; i < 10000; i++ {
		wg6.Add(1)
		go readTest2()
	}
	wg6.Wait()

	endTime := time.Now()
	fmt.Printf("RWMutex Testing ======> finally time costs = %v.\n", endTime.Sub(startTime))
}

func main() {
	mutexTest()
	rwmutexTest()
}
