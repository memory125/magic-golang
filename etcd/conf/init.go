package conf

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd - conf

var (
	etcdCli *clientv3.Client
)

func Init(address []string) (err error) {
	etcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")

	return
}

func GetCli() *clientv3.Client {
	return etcdCli
}
