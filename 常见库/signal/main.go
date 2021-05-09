package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 程序无法捕获信号 SIGKILL 和 SIGSTOP （终止和暂停进程），因此 os/signal 包对这两个信号无效。
func demo() {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan) //监听所有信号
	exitChan := make(chan int)
	go func() {
		for eachSignal := range signalChan {
			fmt.Printf("signal received: %v\n", eachSignal)
			tranSignal, ok := eachSignal.(syscall.Signal)
			if !ok {
				exitChan <- -99
			}
			exitChan <- int(tranSignal)
		}
	}()
	go func() {
		time.Sleep(time.Second * 10)
		exitChan <- 0
	}()
	res := <-exitChan
	fmt.Printf("received signal: %v\n", res)
	os.Exit(res)
}

func main() {
	demo()
}
