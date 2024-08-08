package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 保证数据一致性的几个方法
func demo1() {
	var wg sync.WaitGroup
	var count int64 = 0
	countFunc := func(i int) {
		defer wg.Done()
		for {
			localCount := atomic.LoadInt64(&count)
			if localCount == 100 {
				break
			}
			atomic.AddInt64(&count, 1)
			fmt.Printf("groutine [%d] added ,now count is %d\n", i, count)
			runtime.Gosched()
		}
	}
	exitFunc := func() {
		defer wg.Done()
		for {
			localCount := atomic.LoadInt64(&count)
			if localCount == 100 {
				fmt.Printf("count is not reach 100,exit\n")
				break
			}
			runtime.Gosched()
		}

	}
	wg.Add(1)
	go exitFunc()
	for i := 0; i < 10; i++ {
		go countFunc(i)
		wg.Add(1)
	}
	wg.Wait()
}
func demo2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	var count int64 = 0
	var lock sync.Mutex
	countFunc := func(i int) {
		defer wg.Done()
		for {
			lock.Lock()
			runtime.Gosched() // invalid
			if count == 100 {
				lock.Unlock()
				break
			} else {
				count += 1
				fmt.Printf("groutine [%d] added ,now count is %d\n", i, count)
			}
			lock.Unlock()
			runtime.Gosched()
		}
	}
	exitFunc := func() {
		defer wg.Done()
		for {
			lock.Lock()
			runtime.Gosched() // invalid
			if count == 100 {
				fmt.Printf("count is not reach 100,exit\n")
				lock.Unlock()
				break
			}
			runtime.Gosched()
			lock.Unlock()
			runtime.Gosched()
		}
	}
	wg.Add(1)
	go exitFunc()
	for i := 0; i < 10; i++ {
		go countFunc(i)
		wg.Add(1)
	}
	wg.Wait()
}

func main() {
	demo1()
	demo2()
}
