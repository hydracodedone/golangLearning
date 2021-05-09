package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func tryExcept(actualFunction func(int, int) int, errorHandler func(interface{}) int) (rerturnResult int) {
	rerturnResult = 0
	defer func() {
		if err := recover(); err != nil {
			rerturnResult = errorHandler(err)
		}
	}()
	rerturnResult = actualFunction(2, 0)
	return
}
func main() {
	divisonFunc := func(x, y int) int {
		if y == 0 {
			panic(ErrDivByZero)
		} else {
			return x / y
		}
	}
	errorHandlerFunc := func(errors interface{}) int {
		fmt.Printf("in devision,meet error [%v]\n", errors)
		return 0
	}
	res := tryExcept(divisonFunc, errorHandlerFunc)
	fmt.Printf("the result is %v\n", res)
}
