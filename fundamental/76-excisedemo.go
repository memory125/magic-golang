package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
  使用goroutine和channel实现一个计算int64随机数各位数和的程序
	1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3. 主goroutine从resultChan取出结果并打印到终端输出
*/

// job结构体
type job struct {
	value int64
}

// result结构体
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg3 sync.WaitGroup

// 生成随机数
func generateJob(ch1 chan<- *job) {
	// 循环生成int64类型的随机数，发送到jobChan
	defer wg3.Done()
	for {
		n := rand.Int63() // 生成int64的随机数
		newJob := &job{
			value: n,
		}
		ch1 <- newJob
		time.Sleep(time.Second * 3)
	}
}

// 计算随机数各位的和
func computeSum(jobChan <-chan *job, resultChan chan<- *result) {
	defer wg3.Done()
	for {
		job := <-jobChan
		v := job.value
		sum := int64(0)
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}

		resultChan <- newResult
	}

}

func main() {
	wg3.Add(1)
	go generateJob(jobChan)
	// 开启24个goroutine
	wg3.Add(24)
	for i := 0; i < 24; i++ {
		go computeSum(jobChan, resultChan)
	}

	// 主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Printf("value: %d, sum: %d.\n", result.job.value, result.sum)
	}

	wg3.Wait()
}
