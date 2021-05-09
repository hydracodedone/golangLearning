package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/*
type  WaitGroup struct {
	noCopy noCopy
	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state, and the other 4 as storage
	// for the sema.
	state1 [3]uint32
}
waigGroup 是一个struct,因此在程序中不要作为形参传入
*/
func waitGroupDemo() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			wg.Add(1)
			fmt.Printf("The number is %d\n", i)
		}(i)
	}
	wg.Wait()
}

func main() {
	waitGroupDemo()
}
