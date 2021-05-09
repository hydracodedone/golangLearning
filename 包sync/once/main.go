package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var wg sync.WaitGroup
var myChan = make(chan int)

func demo1() {
	var wg sync.WaitGroup
	var configMap map[string]string
	initConfifFunc := func() {
		fmt.Printf("the configuration is beign")
		configMap = make(map[string]string, 5)
		configMap["configRoot"] = "root"
		configMap["configAdmin"] = "admin"
	}
	actualFunc := func() {
		defer wg.Done()
		if configMap == nil {
			initConfifFunc()
		}
		fmt.Printf("actualFunc is workding\n")
	}
	for i := -0; i < 10; i++ {
		go actualFunc()
		wg.Add(1)
	}
	wg.Wait()
}
func demo2() {
	var wg sync.WaitGroup
	var once sync.Once
	var configMap map[string]string
	initConfifFunc := func() {
		fmt.Printf("the configuration is beign")
		configMap = make(map[string]string, 5)
		configMap["configRoot"] = "root"
		configMap["configAdmin"] = "admin"
	}
	actualFunc := func() {
		defer wg.Done()
		once.Do(initConfifFunc)
		fmt.Printf("actualFunc is workding\n")
	}
	for i := -0; i < 10; i++ {
		go actualFunc()
		wg.Add(1)
	}
	wg.Wait()
}
func demo3() {
	var wg sync.WaitGroup
	var myChan chan int = make(chan int)
	runFunc := func() {
		defer func() {
			once.Do(func() {
				fmt.Printf("close myChan")
				close(myChan)
			})
			wg.Done()
		}()
		fmt.Println("worker worked")
	}
	for i := 0; i < 10; i++ {
		go runFunc()
		wg.Add(1)
	}
	wg.Wait()

}
func main() {
	demo3()
}
