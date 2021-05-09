package main

import "fmt"

func demo1() {
	defer func() {
		ok := recover()
		if ok != nil {
			fmt.Printf("recover,the panic info is [%v]\n", ok)
		}
	}()
	fmt.Println("begin")
	panic("err from panic")
}

func trace(functionName string) func() {
	fmt.Printf("%s begin\n", functionName)
	returnFunc := func() {
		fmt.Printf("%s end\n", functionName)
	}
	return returnFunc
}
func demo2() {
	defer trace("demo2")()
	return
}

func main() {
	demo2()
}
