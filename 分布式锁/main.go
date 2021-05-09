package main

import (
	_ "github.com/coreos/etcd/clientv3"
	_ "github.com/xiaoxuxiansheng/redis_lock"
	_ "go.etcd.io/etcd/clientv3/concurrency"
)

func main() {
}
