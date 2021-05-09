package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var Wg sync.WaitGroup
var Lock sync.Mutex

func demo1() {
	now := time.Now()
	Wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 150; j++ {
				count++
				fmt.Printf("The Count now is %d\n", count)
				time.Sleep(time.Microsecond)
			}
			Wg.Done()
		}()
	}
	Wg.Wait()
	end := time.Now()
	fmt.Printf("The Durations is %v,%v\n", end, now)
}

func demo2() {
	now := time.Now()
	Wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			for j := 0; j < 15; j++ {
				Lock.Lock()
				count++
				fmt.Printf("The Count now is %d\n", count)
				Lock.Unlock()
				time.Sleep(time.Second)
			}
			Wg.Done()
		}()
	}
	Wg.Wait()
	end := time.Now()
	fmt.Printf("The Durations is %v,%v\n", end, now)
}

func main() {
	demo1()
	demo2()
}
