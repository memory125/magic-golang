package watch

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

// watch demo

func WatchKey(cli *clientv3.Client, key string) {
	// watch key change
	rch := cli.Watch(context.Background(), key) // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
