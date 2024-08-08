package main

import (
	"fmt"
	"unsafe"
)

func demo1() {
	a := [3]int{1, 2, 3}
	firstElementPtr := unsafe.Pointer(&a[0])
	thirdElementPtr1 := unsafe.Pointer(&a[2])
	thirdElementPtr2 := unsafe.Pointer(uintptr(firstElementPtr) + 2*unsafe.Sizeof(a[0]))
	fmt.Println(thirdElementPtr1, thirdElementPtr2)
}
func demo2() {
	var a int64 = 32
	var b int8
	b = *(*int8)(unsafe.Pointer(&a))
	fmt.Println(b)
	var c int64
	c = *(*int64)(unsafe.Pointer(&b))
	fmt.Println(c) //value is not correct
}

func main() {
	demo2()
}
