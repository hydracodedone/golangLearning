package main

import (
	"fmt"

	"github.com/sourcegraph/conc/panics"
)

func task(input string) {
	panic(fmt.Sprintf("input:%s is not valid", input))
}

// 传统的panic处理
func demo1() {
	bussinessData1 := "some"
	func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("handle panic: [%s]\n", err)
			}
		}()
		task(bussinessData1)
	}()
	fmt.Println("other tasks....")
}

// conc panic处理

func demo2() {
	bussinessData1 := "some"
	var task1PanicCather panics.Catcher
	//只能捕获当前goroutine的panic,新启动的goroutine的panic是无法捕获的
	task1PanicCather.Try(
		func() {
			task(bussinessData1)
		})
	if recv := task1PanicCather.Recovered(); recv != nil {
		fmt.Printf("handle panic: [%s]\n", recv.Value)
	}
	fmt.Println("other tasks....")

}
func main() {
	demo1()
	demo2()
}
