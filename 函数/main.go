package main

import "fmt"

type myFunc func(int, int) int

type myFuncWithVarName func(x, y int) (z int) //使用type申明函数类型时,对函数的入参名称无要求,但是对命名返回值的名称是有要求的

type myFuncReturnWithFunc func(int, int) func(int, int) int //返回值是一个函数

// type myFuncEnterWithFunc func(func(int, int) int) int //形参是一个函数

func test(x, y int) (z int) {
	z = x + y
	return
}
func demoForTypeOfFunction() {
	demo := test
	fmt.Printf("The demo type is %T\n", demo)
	var demo2 myFunc = demo
	fmt.Printf("The demo2 type is %T\n", demo2)

}
func demoForUseFuncType() {
	var myfunc myFunc = func(x, y int) (z int) {
		z = x + y
		return
	}
	myfunc2 := func(x, y int) (z int) {
		z = x + y
		return
	}
	fmt.Printf("The myfunc type is %T\n", myfunc)
	fmt.Printf("The myfunc2 type is %T\n", myfunc2)
}
func demoForUseFuncTypeWithVarName() {
	var myfunc myFuncWithVarName = func(a int, b int) (c int) {
		c = a + b
		return
	}
	var myfunc2 myFuncWithVarName = func(a, b int) (df int) {
		df = a + b
		return df
	}
	fmt.Printf("The myfunc type is %T\n", myfunc)
	fmt.Printf("The myfunc2 type is %T\n", myfunc2)
}
func demoForComplexFuncType() {
	temp := func(x int, y int) func(int, int) int {
		temp2 := func(a, b int) int {
			return a + b
		}
		return temp2
	}
	var temp2 myFuncReturnWithFunc = temp
	fmt.Printf("The temp2 type is %T\n", temp2)

}

func demoForVariableLengthArgument() {
	demoFunction := func(demo ...int) {
		fmt.Printf("The demo is %v\n", demo)
		fmt.Printf("The demo type is %T\n", demo)

	}
	demoFunction(1, 2, 3)
	demoFunction([]int{1, 2, 3}...)
}
func demoForVariableSlice() {
	demoFunction := func(demo []int) {
		fmt.Printf("The demo is %v\n", demo)
		fmt.Printf("The demo type is %T\n", demo)
	}
	demoFunction([]int{1, 2, 3})
}

func demoForInterface() {
	temp := func(anyVar ...interface{}) {
		fmt.Printf("The anyVar is %v\n", anyVar)
		fmt.Printf("The anyVar type is %T\n", anyVar)
	}
	temp(1, "2", nil)
}
func demoForNamedReturnValue() {
	temp := func(x int, y int) (z int) {
		fmt.Printf("the z is %v\n", z)
		var z2 int = x + y
		return z2
	}
	res := temp(1, 2)
	fmt.Printf("the res is %d\n", res)
}

func main() {
	demoForTypeOfFunction()
	demoForUseFuncType()
	demoForUseFuncTypeWithVarName()
	demoForComplexFuncType()
	demoForVariableLengthArgument()
	demoForVariableSlice()
	demoForInterface()
	demoForNamedReturnValue()
}
