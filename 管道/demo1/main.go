package main

import (
	"fmt"
	"sync"
	"time"
)

/*
fatal error: all goroutines are asleep - deadlock!
在使用管道的时候很容易出现上述错误
运行时发现所有的 goroutine（包括main）都处于等待 goroutine。'
也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码
*/

func demo() {
	// 当 channel是nil的时候，无论是传入数据还是取出数据，都会永久的block。
	var temp chan int
	temp <- 23
}

func demo2() {
	var channel chan int = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for {
			fmt.Printf("The Result Consume  is %d\n", <-channel)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		temp := 0
		for {
			channel <- temp
			fmt.Printf("The Result Produce is %d\n", temp)
			temp++
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

}
func demo3() {
	var chan1 chan int = make(chan int)
	var chan2 chan int = make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		temp := 0
		for {
			chan1 <- temp
			fmt.Printf("The Result Producd by chan1 is %d\n", temp)
			temp++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		temp := 0
		for {
			chan2 <- temp
			fmt.Printf("The Result Producd by chan2 is %d\n", temp)
			temp++
			time.Sleep(time.Second)
		}
	}()
	go func() {
		temp := 0
		for {
			select {
			case temp = <-chan1:
				fmt.Printf("The Result Consumed by chan1 is %d\n", temp)
			case temp = <-chan2:
				fmt.Printf("The Result Consumed by chan2 is %d\n", temp)
			default:
				fmt.Println("wait...")
				time.Sleep(time.Second / 2)
			}
		}
	}()
	wg.Wait()
}
func demo4() {
	var chan1 chan int = make(chan int)
	var chan2 chan int = make(chan int)
	var wg sync.WaitGroup
	close(chan1)
	wg.Add(2)
	go func() {
		time.Sleep(time.Second)
		chan2 <- 999
		defer wg.Done()
	}()
	go func() {
		var temp int
	Tag:
		for {
			select {
			case temp = <-chan1:
				fmt.Printf("Recv from the chan2 is %3d\n", temp)
			case temp = <-chan2:
				fmt.Printf("Recv from the chan1 is ---------->%3d\n", temp)
				break Tag
			}
			time.Sleep(time.Millisecond * 10)
		}
		defer wg.Done()
	}()
	wg.Wait()
}
func main() {
	demo4()
}
