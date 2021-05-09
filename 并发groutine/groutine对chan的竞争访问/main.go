package main

import (
	"context"
	"fmt"
)

var resourceChan chan int = make(chan int, 10)

func produce() {
	for i := 0; i <= 100; i++ {
		resourceChan <- i
		fmt.Printf("The producer produce %d\n", i)
	}
}
func consumer(ctx context.Context, cancelFuncion func(), name string) {
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case product := <-resourceChan:
			fmt.Printf("The consumer:<%s> consume:<%d>\n", name, product)
			if product == 100 {
				cancelFuncion()
			}
		default:
		}
	}
}
func main() {
	bgCtx := context.Background()
	cancelCtx, cancelFuncion := context.WithCancel(bgCtx)
	go produce()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("->Consumer %d<-", i)
		go consumer(cancelCtx, cancelFuncion, name)
	}
	select {
	case <-cancelCtx.Done():
	}
}
