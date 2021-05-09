package main

import "fmt"

type myFunc func(int, int) int

type myFuncWithVarName func(x, y int) (z int) //使用type申明函数类型时,对函数的入参名称无要求,但是对命名返回值的名称是有要求的

type myFuncReturnWithFunc func(x int, y int) func(a, b int) (z int) //返回值是一个函数

type myFuncEnterWithFunc func(enterFunc func(c, d int) (f int)) (result int) //形参是一个函数

func test(x, y int) (z int) {
	z = x + y
	return
}
func demoForTypeOfFunction() {
	demo := test
	fmt.Printf("The demo type is %T\n", demo)
}
func demoForUseFuncType() {
	var myfunc myFunc
	myfunc = func(x, y int) (z int) {
		z = x + y
		return
	}
	var myfunc2 myFunc
	myfunc2 = func(x, y int) (z int) {
		z = x + y
		return
	}
	fmt.Printf("The myfunc is %#v\n", myfunc)
	fmt.Printf("The myfunc2 is %#v\n", myfunc2)
}
func demoForUseFuncTypeWithVarName() {
	var myfunc myFuncWithVarName
	myfunc = func(a int, b int) (c int) {
		c = a + b
		return
	}
	var myfunc2 myFuncWithVarName
	myfunc2 = func(a, b int) (df int) {
		df = a + b
		return df
	}
	fmt.Printf("The myfunc is %#v\n", myfunc)
	fmt.Printf("The myfunc2 is %#v\n", myfunc2)
}

func main() {
	demoForUseFuncTypeWithVarName()
}
