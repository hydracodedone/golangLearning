package main

import (
	"error_demo/middle"
	"error_demo/middle/inner"
	"errors"
	"fmt"
)

func Demo1() {
	err1 := middle.Middle()
	if err1 != nil {
		fmt.Println(err1)
		fmt.Println(err1 == inner.InnerError)
	}
	err2 := middle.Middle()
	if err1 != nil {
		fmt.Println(err2)
		fmt.Println(err2 == inner.InnerError)
	}
	fmt.Println(err2 == err1)
}

func Demo2() {
	err1 := middle.MiddleWithWrap()
	if err1 != nil {
		fmt.Printf("%q\n", err1)
		fmt.Println(errors.Is(err1, inner.InnerError))
	}
}
func Demo3() {
	err1 := middle.MiddleWithStack()
	if err1 != nil {
		fmt.Printf("%+v\n", err1)
		fmt.Println(errors.Is(err1, inner.InnerError))
	}
}

func main() {
	Demo3()
}
