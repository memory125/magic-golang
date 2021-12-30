package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wait group

var wg sync.WaitGroup

func frand()  {
	rand.Seed(time.Now().UnixNano())     // 随机数种子
	for i := 0; i < 5; i++{
		r1 := rand.Int()                 // int64
		r2 := rand.Intn(10)   // 返回0 ~ n的整型数值
		fmt.Printf("r1 = %v, r2 = %v \n", r1, r2)
	}
}

func trand(i int)  {
	defer wg.Done()           // 判断函数执行结束
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func main()  {
	//frand()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go trand(i)
	}
	wg.Wait()       // 等待wg的计数器减为0
}
