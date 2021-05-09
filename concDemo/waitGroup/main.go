package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/sourcegraph/conc"
)

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("task finished")
}

// 传统模式的goroutine控制
func demo1() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	task(&wg)
	wg.Wait()
}
func task2() {
	time.Sleep(time.Second)
	panic("demo")
}

// conc模式的goroutine控制
func demo2() {
	wg := conc.WaitGroup{}
	wg.Go(task2)
	wg.WaitAndRecover()

}

func task3(i int) {
	fmt.Printf("task %d begin\n", i)
	time.Sleep(time.Second)
	fmt.Printf("task %d finished\n", i)
}

// conc避免groutine对循环变量的错误使用
func demo3() {
	wg := conc.WaitGroup{}
	for i := 0; i < 10; i++ {
		some := i
		wg.Go(func() {
			task3(some)
		})
	}
	wg.Wait()
}

// 传统模式的panic处理
// 如果waitGroup中由任意一个groutine产生的panic没有被处理,则会将整个的waitGroup中的groutine都异常退出
func task4(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		recover()
	}()
	fmt.Printf("task %d begin\n", i)
	time.Sleep(time.Second)
	if i > 3 {
		panic(fmt.Sprintf("panic %d", i))
	}
	fmt.Printf("task %d finishedd\n", i)
}

// 传统模式的异常
func demo4() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task4(i, &wg)
	}
	wg.Wait()
}

// conc模式的异常处理
// 相比之下可以在用户未手动完成goroutine的异常处理时候进行兜底
// 但是只会保存最后一次的panic的栈信息
func task5(i int) {
	fmt.Printf("task %d begin\n", i)
	time.Sleep(time.Second)
	if i > 3 {
		panic(fmt.Sprintf("panic %d", i))
	}
	fmt.Printf("task %d finishedd\n", i)
}
func demo5() {
	wg := conc.WaitGroup{}
	for i := 0; i < 10; i++ {
		some := i
		wg.Go(func() {
			task5(some)
		})
	}
	panicInfo := wg.WaitAndRecover()
	frames := runtime.CallersFrames(panicInfo.Callers)
	for i := 0; ; i++ {
		frame, more := frames.Next()
		fmt.Printf("file:[%s],line:[%d],function:[%s],address:[%v]\n", frame.File, frame.Line, frame.Function, frame.Entry)
		if !more {
			break
		}
	}

}

func main() {
	demo5()
}
