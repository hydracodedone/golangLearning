package main

import (
	"context"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
)

func main() {
	//初始化etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: time.Second,
	})
	if err != nil {
		panic(err)
	}
	//创建一个session，并根据业务情况设置锁的ttl
	s, err := concurrency.NewSession(cli, concurrency.WithTTL(300))
	if err != nil {
		panic(err)
	}
	defer s.Close()
	//初始化一个锁的实例，并进行加锁解锁操作。
	mu := concurrency.NewMutex(s, "mutex-linugo")
	if err := mu.Lock(context.TODO()); err != nil {
		log.Fatal("m lock err: ", err)
	}
	//do something
	if err := mu.Unlock(context.TODO()); err != nil {
		log.Fatal("m unlock err: ", err)
	}
}
