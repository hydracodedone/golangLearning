package main

import "fmt"

//接口
type Call interface {
	call()
}

//自定义函数类型A
type MyFunc func()

//为自定义函数类型A的实例实现该接口
func (myFuncInstance *MyFunc) call() {
	fmt.Println("Hello,World")
}

func main() {
	var caller Call
	var myFuncInstance MyFunc = func() {}
	//因为myFuncInstance为自定义函数类型的实例,且该类型实现了接口,因此可以将myFuncInstance赋值给接口类型的变量caller
	caller = &myFuncInstance
	caller.call()
	(&myFuncInstance).call()
}
