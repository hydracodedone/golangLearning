package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup
var myChan = make(chan int)

func closeChan() {
	func() {
		close(myChan)
	}()
}
func worker(name string) {
	defer wg.Done()
	res := once.Do(closeChan)
	if res {
		fmt.Printf("worker:<%s> close the chan", name)
	}
}

func demo1() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(fmt.Sprintf("WORKER%d", i))
	}
	wg.Wait()
}
func demo2() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		once.Do(func() { fmt.Println("sss") })
	}
}

func main() {
	demo1()
	demo2()
}
