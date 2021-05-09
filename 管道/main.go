package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myChan chan int
	myOnce sync.Once
	wg     sync.WaitGroup
)

func read() {
TAG:
	for {
		select {
		case number, status := <-myChan:
			if !status {
				break TAG
			} else {
				fmt.Println(number, status)
				time.Sleep(time.Second)
			}
		default:
			myOnce.Do(func() { close(myChan) })
			break TAG
		}
	}
	defer wg.Done()
}
func demo() {
	myChan = make(chan int, 5)
	for i := 0; i < 5; i++ {
		myChan <- i
	}
	wg.Add(2)
	go read()
	go read()
	wg.Wait()
}
func main() {
}
