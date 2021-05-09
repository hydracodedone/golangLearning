package main

import (
	"fmt"
	"sync"
)

func main() {

	var a = 0b111101
	var b = 0b000010
	fmt.Printf("%08b,%08b\n", a, b)
	var m sync.Mutex
	m.Lock()
}
