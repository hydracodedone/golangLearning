package main

import "fmt"

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
