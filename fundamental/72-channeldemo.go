package main

import (
	"fmt"
	"sync"
)

// channel示例

// 声明连个channel

var wg2 sync.WaitGroup

func generateData(ch1 chan int)  {
	defer wg2.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func receiveData(ch1, ch2 chan int)  {
	defer wg2.Done()
	for {
		x, ok := <- ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main()  {
	// 定义两个channel
	a := make(chan int, 100)
	b := make(chan int, 100)

	wg2.Add(2)
	go generateData(a)
	go receiveData(a, b)
	wg2.Wait()
	for ret := range b {
		fmt.Println("ret========>", ret)
	}
}