package main

import "fmt"

//接口
type Call interface {
	call()
}

//自定义函数类型A
type MyFunc func()

//为自定义函数类型A的实例实现该接口
func (myfunc *MyFunc) call() {
	fmt.Println("Hello,World")
}

func main() {
	var caller Call
	var myfunc MyFunc = func() {}
	//因为myfunc为自定义函数类型的实例,且该类型实现了接口,因此可以将myfunc赋值给接口类型的变量caller
	caller = &myfunc
	caller.call()
	(&myfunc).call()
}
