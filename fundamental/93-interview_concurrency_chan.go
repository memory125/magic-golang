package main

import "fmt"

// 面试 - 并发
/*
  创建一个goroutine, 与主线程按顺序相互发送信息若干次并打印
*/

// 声明用来沟通的channel
var comCh chan string

// routine函数
func communicationToMain() {
	i := 1
	for {
		fmt.Println(<-comCh)
		comCh <- fmt.Sprintf("Msg from communicationToMain.... #%v.", i)
		i++
	}
}

/*
	Msg from main.... #1.
	Msg from communicationToMain.... #1.
	Msg from main.... #2.
	Msg from communicationToMain.... #2.
	Msg from main.... #3.
	Msg from communicationToMain.... #3.
	Msg from main.... #4.
	Msg from communicationToMain.... #4.
	Msg from main.... #5.
	Msg from communicationToMain.... #5.
	Msg from main.... #6.
	Msg from communicationToMain.... #6.
	Msg from main.... #7.
	Msg from communicationToMain.... #7.
	Msg from main.... #8.
	Msg from communicationToMain.... #8.
	Msg from main.... #9.
	Msg from communicationToMain.... #9.
	Msg from main.... #10.
	Msg from communicationToMain.... #10.
*/

func main() {
	// 无缓存通道
	comCh = make(chan string)
	go communicationToMain()
	for i := 1; i <= 10; i++ {
		comCh <- fmt.Sprintf("Msg from main.... #%v.", i)
		fmt.Println(<-comCh)
	}
}
