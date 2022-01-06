package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发map

//var m = make(map[string]int)         // 并发不安全
//
//func get(key string) int {
//	return m[key]
//}
//
//func set(key string, value int) {
//	m[key] = value
//}
//
//func main() {
//	wg := sync.WaitGroup{}
//		for i := 0; i < 21; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			// fatal error: concurrent map writes
//			set(key, n)
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

// 针对上面出现的map并发写错误，我们将代码改造
var m = sync.Map{} // 并发安全map

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 写入操作必须使用Store接口完成
			value, _ := m.Load(key) // 取值通过Load接口实现
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
