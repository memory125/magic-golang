package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// pprof - 示例

// 一段有问题的代码
func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c: // chan c未初始化，为nil，这里会阻塞
			fmt.Printf("recv from chan, value:%v\n", v)
		default: // 代码执行default，陷入死循环

		}
	}
}

func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	/*
			步骤：
			1. go build pprof.go
		    2. pprof.exe -cpu=true
		    3. go tool pprof cpu.pprof -> top 3 / top 4
			(pprof) top 3
			Showing nodes accounting for 76.83s, 99.74% of 77.03s total
			Dropped 23 nodes (cum <= 0.39s)
		      flat  flat%   sum%        cum   cum%
		    42.15s 54.72% 54.72%     42.15s 54.72%  runtime.chanrecv
		    24.15s 31.35% 86.07%     66.31s 86.08%  runtime.selectnbrecv
		    10.53s 13.67% 99.74%     76.91s 99.84%  main.logicCode
		-------------------------------------------------------------------------
			flat：当前函数占用CPU的耗时
			flat：:当前函数占用CPU的耗时百分比
			sun%：函数占用CPU的耗时累计百分比
			cum：当前函数加上调用当前函数的函数占用CPU的总耗时
			cum%：当前函数加上调用当前函数的函数占用CPU的总耗时百分比
			最后一列：函数名称
	*/
	if isCPUPprof {
		f1, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(f1)

		// 使用匿名函数
		defer func() {
			pprof.StopCPUProfile()
			f1.Close()
		}()
	}
	for i := 0; i < 6; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		f2, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(f2)
		f2.Close()
	}
}
