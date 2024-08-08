package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func doSth1(done chan bool) {
	// 子协裎执行1秒
	time.Sleep(time.Second)

	// Notice 由于done是不带缓存的channel 如果done没有接收方，子协裎会一直hang在这里
	done <- true
}

func timeOut1(f func(chan bool)) error {
	// 不带缓存的channel ———— 发送/接收 得配对，否则都会夯住
	done := make(chan bool)

	go f(done)

	select {
	case <-done: // 接收done
		fmt.Println("done!")
		return nil
	case <-time.After(time.Millisecond): // 主协裎 1微秒 就超时
		//fmt.Println("timeOut!!!")
		return fmt.Errorf("timeout!")
	}
}

func main() {
	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second)
		timeOut1(doSth1)
	}
	// 主协裎执行
	fmt.Println("the numbr of the groutine is %d\n", runtime.NumGoroutine())
}
