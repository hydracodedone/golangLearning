package main

import (
	"fmt"
	"time"
)

//goroutine启动访问外部变量的情况
//因为groutine启动需要一定的时间,因此,启动groutine的时间远小于goroutine实际开始运行的时间
//如果是groutine访问外部变量,可以导致函数运行不符合预期
func goroutineWithOutVariable() {
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second)
			fmt.Printf("The value is %d\n", i)
		}()
	}
}

func goroutineSafetyUsage() {
	for i := 0; i < 100; i++ {
		go func(i int) {
			time.Sleep(time.Second)
			fmt.Printf("The value is %d\n", i)
		}(i)
	}
}

func main() {
	goroutineSafetyUsage()
	time.Sleep(time.Second * 2)
}
