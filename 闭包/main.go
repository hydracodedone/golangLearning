package main

import (
	"fmt"
	"log"
	"runtime"
)

func demoForClosure() {
	x := 5
	fmt.Printf("The address of x is %p,The value is %d\n", &x, x)
	demo := func() {
		fmt.Printf("The address of x is %p,The value is %d\n", &x, x)
	}
	demo()
}

//闭包在函数返回中,引用的变量并没有消亡
func demoForLeak() func() {
	x := 5
	fmt.Printf("The address of x is %p,The value is %d\n", &x, x)

	demo := func() {
		fmt.Printf("The address of x is %p,The value is %d\n", &x, x)
	}
	return demo
}

func demo2() {
	temp := demoForLeak()
	fmt.Println("test begin")
	temp()
	fmt.Println("test end")
}

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("  %s:%d", file, line)
}
var where2 = log.Print

func testForWhere() {
	func1 := func() {
		fmt.Printf("this is a func1\n")
		where()
	}
	func1()
	func2 := func() {
		fmt.Printf("this is a func2\n")
	}
	where()
	func2()
	func3 := func() {
		fmt.Printf("this is a func3\n")
	}
	where2()
	func3()
}

func main() {
	demoForClosure()
	demo2()
	testForWhere()
}
