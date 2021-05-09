package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwLock sync.RWMutex

/*
可以同时申请多个读锁,多个读锁之间不会发生阻塞
有读锁时申请写锁将阻塞，有写锁时申请读锁将阻塞
只要有写锁，后续申请读锁和写锁都将阻塞
*/
/*
Lock()和Unlock()用于申请和释放写锁
RLock()和RUnlock()用于申请和释放读锁
一次RUnlock()操作只是对读锁数量减1，即减少一次读锁的引用计数
如果不存在写锁，则Unlock()引发panic，如果不存在读锁，则RUnlock()引发panic
*/
/*
无论是Mutex还是RWMutex都不会和goroutine进行关联，
这意味着它们的锁申请行为可以在一个goroutine中操作，
释放锁行为可以在另一个goroutine中操作。
*/
/*
由于RLock()和Lock()都能保证数据不被其它goroutine修改，
所以在RLock()与RUnlock()之间的，以及Lock()与Unlock()之间的代码区都是critical section。
*/
/*
写锁权限高于读锁，有写锁时优先进行写锁定。
*/
/*

一个读写锁的RLocker方法的结果值的Lock方法或Unlock方法进行调用的时候实际上是在调用该读写锁的RLock方法或RUnlock方法。
通过读写锁的RLocker方法获得这样一个结果值的实际意义在于，我们可以在之后以相同的方式对该读写锁中的“写锁”和“读锁”进行操作。
这为相关操作的灵活适配和替换提供了方便。
*/
var value int

func reader(name string) {
	for {
		rand.Seed(time.Now().UnixNano())
		rwLock.RLock()
		fmt.Printf("The value from reader <%s> is %d\n", name, value)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		rwLock.RUnlock()
	}
}
func reader2(name string) {
	for {
		rand.Seed(time.Now().UnixNano())
		rwLock.RLocker().Lock()
		fmt.Printf("The value from reader <%s> is %d\n", name, value)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		rwLock.RLocker().Unlock()
	}
}
func writer() {
	for {
		rwLock.Lock()
		value++
		fmt.Printf("The value set by writer is %d\n", value)
		time.Sleep(time.Second)
		rwLock.Unlock()
	}
}
func demo() {
	var sig = make(chan int)
	for i := 0; i < 5; i++ {
		go reader(fmt.Sprintf("READER %d", i))
	}
	go writer()
	<-sig
}
func demo2() {
	var sig = make(chan int)
	for i := 0; i < 5; i++ {
		go reader(fmt.Sprintf("READER %d", i))
	}
	go writer()
	<-sig
}
func main() {
	demo2()
}
