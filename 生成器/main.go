package main

import "fmt"

func generator() func() int {
	myChan := make(chan int)
	count := 0
	go func() {
		for {
			myChan <- count
			count++
		}
	}()
	return func() int {
		res := <-myChan
		return res
	}
}

func demo() {
	myGenerator := generator()
	for i := 0; i < 10; i++ {
		res := myGenerator()
		fmt.Println(res)
	}
}

func main() {
	demo()
}
