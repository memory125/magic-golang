package get

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd - client get示例
// use etcd/clientv3

func GetInfo(cli *clientv3.Client, key string) (err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	return
}
