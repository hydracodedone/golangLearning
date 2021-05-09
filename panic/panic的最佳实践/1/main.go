package main

import (
	"fmt"
	"math"
)

func divison(x int, y int) (z int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recover error:[%v]\n", err)
		}
		z = math.MaxInt8
	}()
	fmt.Println("begin calculate")
	z = x / y
	fmt.Println("end calculate") //do not excute
	return
}
func demo4() {
	res := divison(2, 0)
	fmt.Printf("the result is %d\n", res)
}

// 一种处理panic的方式,相比于demo4,能够使得代码正常执行下去
// 尽量将recover函数和会出现问题的代码放在一个函数中处理
func newDivison(x int, y int) (z int) {
	fmt.Println("begin calculate")
	func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("recover error:[%v]\n", err)
			}
			z = math.MaxInt8
		}()
		z = x / y
	}()
	fmt.Println("end calculate") //excute
	return
}
func demo5() {
	res := newDivison(2, 0)
	fmt.Printf("the result is %d\n", res)
}

func main() {
	demo4()
	demo5()
}
