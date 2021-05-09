package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var endSignal chan int = make(chan int)

func handlerFunc(dataChan chan int) {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		data, ok := <-dataChan
		if !ok {
			break
		} else {
			fmt.Printf("handle time:[%v],data:[%v]\n", time.Now(), data)
		}
	}
	ticker.Stop()
	endSignal <- 1
}
func main() {
	var dataChan chan int = make(chan int, 5)
	for i := 1; i < 6; i++ {
		dataChan <- i
	}
	close(dataChan)
	go handlerFunc(dataChan)
	<-endSignal
}
