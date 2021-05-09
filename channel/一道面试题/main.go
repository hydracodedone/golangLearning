package main

import (
	"fmt"
	"sync"
)

//10个goroutine顺序输出1-10

func worker(wg *sync.WaitGroup, once *sync.Once, queue chan int, index int) {
	once.Do(func() {
		queue <- 1
	})
	defer wg.Done()
	i := <-queue
	fmt.Printf("worker :[%d] print [%d]\n", index, i)
	queue <- i + 1
}

func main() {
	once := sync.Once{}
	wg := &sync.WaitGroup{}
	queue := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(wg, &once, queue, i)
	}
	wg.Wait()
}
