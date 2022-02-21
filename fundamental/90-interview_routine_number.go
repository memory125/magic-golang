package main

import (
	"fmt"
	"sync"
)

// 面试
/*
创建协程交替输出数据。一个协程输出偶数，一个协程奇数
*/

/*
 通过chan实现：
	evenCh：存放偶数
    oddCh：存放奇数
*/

// 输出偶数
/*
  此时：奇数oddCh是只写的，偶数evenCh是只读的
*/
func EvenNumber(routinename string, wg *sync.WaitGroup, oddCh chan<- int, evenCh <-chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// 给奇数通道oddCh赋值
		oddCh <- 2*i - 1
		// 输出偶数
		fmt.Printf("routine ==> %v, %v\n", routinename, <-evenCh)
	}
}

// 输出奇数
/*
  此时：奇数evenCh是只写的，偶数oddCh是只读的
*/
func OddNumber(routinename string, wg *sync.WaitGroup, oddCh <-chan int, evenCh chan<- int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		// 输出奇数
		fmt.Printf("routine ==> %v, %v\n", routinename, <-oddCh)
		// 将偶数通道ch1赋值
		evenCh <- 2 * i
	}
}

func main() {
	// waitGroup
	wg := &sync.WaitGroup{}
	// 偶数chan
	evenCh := make(chan int)
	// 奇数chan
	oddCh := make(chan int)
	wg.Add(2)
	/*
		routine ==> Odd, 1
		routine ==> Even, 2
		routine ==> Odd, 3
		routine ==> Even, 4
		routine ==> Odd, 5
		routine ==> Even, 6
		routine ==> Odd, 7
		routine ==> Even, 8
		routine ==> Odd, 9
		routine ==> Even, 10
		routine ==> Odd, 11
		routine ==> Even, 12
		routine ==> Odd, 13
		routine ==> Even, 14
		routine ==> Odd, 15
		routine ==> Even, 16
		routine ==> Odd, 17
		routine ==> Even, 18
		routine ==> Odd, 19
		routine ==> Even, 20
	*/
	go OddNumber("Odd", wg, oddCh, evenCh)
	go EvenNumber("Even", wg, oddCh, evenCh)
	wg.Wait()
}
