package main

import "fmt"

func demo() {
	//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
	defer func() {
		//effective
		err := recover()
		if err != nil {
			fmt.Printf("the panic is capture in this moment0 %v\n", err)
		}
	}()
	//not effective
	defer recover()
	//not effective
	defer fmt.Printf("the panic is capture in this moment1 %v\n", recover())
	defer func() {
		//not effective
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("the panic is capture in this moment2 %v\n", err)
			}
		}()
	}()
	defer func() {
		//effective
		fmt.Printf("the panic is capture in this moment3 %v\n", recover())
	}()
	capture := func() {
		//effective
		err := recover()
		if err != nil {
			fmt.Printf("the panic is capture in this moment4 %v\n", err)
		}
	}
	defer capture()
	panic("panic in main")
}
func main() {
	demo()
}
