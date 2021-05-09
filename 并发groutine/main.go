package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func consumer(name string, productLine <-chan string) {
	for res := range productLine {
		fmt.Printf("Name %s consume %s\n", name, res)
	}
	defer wg.Done()
}
func producer(productLine chan<- string) {
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("[product %d]", i)
		productLine <- value
	}
	defer wg.Done()
	defer close(productLine)
}
func main() {
	var myChan chan string = make(chan string)
	wg.Add(2)
	for i := 0; i < 3; i++ {
		go consumer(fmt.Sprintf("%d", i), myChan)
	}
	go producer(myChan)
	wg.Wait()
}
