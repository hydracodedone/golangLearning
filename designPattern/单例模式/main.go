package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct{}

var singletonInstance *singleton
var mutex sync.Mutex
var singletonFlag uint32
var singletonInitOnce sync.Once

func getSingleton1() *singleton {
	mutex.Lock()
	defer mutex.Unlock()
	if singletonInstance == nil {
		singletonInstance = &singleton{}
		return singletonInstance
	}
	return singletonInstance
}

// check-lock-check
/*
第二个check是解决锁竞争情况下的问题，假设现在两个线程来请求
A、B线程同时发现​​singleton​​​实例对象为空，因为我们在第一次check方法上没有加锁，然后A线程率先获得锁，
进入同步代码块，new了一个​​singleton​​​实例对象，之后释放锁，接着B线程获得了这个锁，
发现​​singleton​​​实例对象已经被创建了，就直接释放锁，退出同步代码块。所以这就是​​Check-Lock-Check​​​
*/
func getSingleton2() *singleton {
	if atomic.LoadUint32(&singletonFlag) == 1 {
		return singletonInstance
	}
	mutex.Lock()
	defer mutex.Unlock()
	//解决
	if singletonFlag == 0 {
		singletonInstance = &singleton{}
		atomic.StoreUint32(&singletonFlag, 1)
	}
	return singletonInstance
}
// sync.Once 就是check-lock-check
func getSingleton3() *singleton {
	singletonInitOnce.Do(func() { singletonInstance = &singleton{} })
	return singletonInstance
}

func main() {

}
