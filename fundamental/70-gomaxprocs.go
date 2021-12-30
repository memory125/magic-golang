package main

// GOMAXPROCS

import (
	"fmt"
	"runtime"
	"sync"
)

var wg1 sync.WaitGroup

func a()  {
	defer wg1.Done()
	for i := 0; i < 10; i++  {
		fmt.Printf("A==========>: %d.\n", i)
	}
}

func b()  {
	defer wg1.Done()
	for i := 0; i < 10; i++  {
		fmt.Printf("B==========>: %d.\n", i)
	}
}

func main()  {
	runtime.GOMAXPROCS(4)       // 默认CPU的逻辑核心数，默认跑满整个CPU
	fmt.Println("CPU Cores ", runtime.NumCPU())     // 返回电脑的CPU核心数
	wg1.Add(2)
	go a()
	go b()
	wg1.Wait()
}
