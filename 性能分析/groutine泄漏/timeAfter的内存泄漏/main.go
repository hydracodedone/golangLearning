package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func demo1() {
	sig := make(chan int)
	go func() {
		for {
			time.Sleep(time.Millisecond * 10)
			sig <- 1
		}
	}()
	go func() {
		if err := http.ListenAndServe(":9000", nil); err != nil {
			fmt.Println("fail")
		}
	}()
Loop:
	for {
		select {
		case <-time.After(time.Second * 100):
			fmt.Println("wait already 2 seconds")
			break Loop
		case <-sig:
			fmt.Println("recv data,continue")
		}
	}
}
func main() {
	demo1()
}
