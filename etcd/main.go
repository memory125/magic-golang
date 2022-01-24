package main

import (
	"fmt"
	"wing.com/magic-golang/etcd/conf"
	"wing.com/magic-golang/etcd/get"
	"wing.com/magic-golang/etcd/put"
	"wing.com/magic-golang/etcd/watch"
)

// etcd 示例

func main() {
	// 1. 初始化etcd
	err := conf.Init([]string{"127.0.0.1:2379"})
	if err != nil {
		fmt.Printf("Init etcd connection failed, error is %v.\n", err)
		return
	}

	// 2. put操作
	err = put.PutData(conf.GetCli(), "test", "hello")
	if err != nil {
		fmt.Printf("etcd put data failed, error is %v.\n", err)
		return
	}

	// 3. get操作
	err = get.GetInfo(conf.GetCli(), "test")
	if err != nil {
		fmt.Printf("etcd put data failed, error is %v.\n", err)
		return
	}

	// 4. watch操作
	watch.WatchKey(conf.GetCli(), "test")
}
