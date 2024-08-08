package main

import (
	"fmt"
	"sync"
	"time"
)

var cond sync.Cond = *sync.NewCond(&sync.Mutex{})
var data int

func groutineFather() {
	fmt.Println("father groutine begin")
	cond.L.Lock()
	defer cond.L.Unlock()
	defer cond.Broadcast()
	data += 1
}
func groutineChild(gName string) {
	fmt.Printf("child groutine %s begin\n", gName)
	cond.L.Lock()
	cond.Wait()
	defer cond.L.Unlock()
	data += 1
}
func main() {
	for i := 0; i < 10; i++ {
		go groutineChild(fmt.Sprintf("groutine %d", i))
	}
	go groutineFather()
	time.Sleep(time.Second * 1)
}
