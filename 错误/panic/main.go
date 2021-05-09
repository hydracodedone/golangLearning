package main

import (
	"fmt"
)

func syntaxErrorHandler() {
	err := recover()
	fmt.Println("THE ERROR IS:", err)
	if err != nil {
		fmt.Println("THE ERROR IS:", err)
	} else {
	}
}

/*
一般而言，当宕机发生时，程序会中断运行，并立即执行在该 goroutine（可以先理解成线程）中被延迟的函数（defer 机制），随后，程序崩溃并输出日志信息，日志信息包括 panic value 和函数调用的堆栈跟踪信息，panic value 通常是某种错误信息。
Go语言的类型系统会在编译时捕获很多错误，但有些错误只能在运行时检查，如数组访问越界、空指针引用等，这些运行时错误会引起宕机。
panic() 的声明如下：
func panic(v interface{})    //panic() 的参数可以是任意类型的。
当 panic() 触发的宕机发生时，panic() 后面的代码将不会被运行，但是在 panic() 函数前面已经运行过的 defer 语句依然会在宕机发生时发生作
Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。
*/

func demo(x int) {
	a := [3]int{1, 3, 5}

	defer syntaxErrorHandler() //访问了demo函数的panic的上下文，因此原理是一个闭包
	a[x] = 100
	a[x+2] = 100

}

func main() {
	demo(5)
}
