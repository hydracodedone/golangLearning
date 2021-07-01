package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
before runtime.Goexit() executed, all register defer function will be called
*/
func demo() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				fmt.Printf("The Defer Value Is %d\n", i)
				wg.Done()
			}()
			if i == 3 {
				runtime.Goexit()
			} else {
				fmt.Printf("The Value Is %d\n", i)
			}
		}(i)
	}
	wg.Wait()
}

func main() {
	demo()
}
