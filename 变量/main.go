package main

import (
	"fmt"
	"unsafe"
)

func demo1() {
	/*
		int   类型⼤⼩为 8 字节
		int8  类型⼤⼩为 1 字节
		int16 类型⼤⼩为 2 字节
		int32 类型⼤⼩为 4 字节
		int64 类型⼤⼩为 8 字节
		go语⾔中的int的⼤⼩是和操作系统位数相关的
		如果是32位操作系统 int类型的⼤⼩就是4字节
		如果是64位操作系统 int类型的⼤⼩就是8个字节
	*/
	var a uint8
	var b byte
	var c int32
	var d rune
	var e int
	var f int8
	var g int64
	a = 1
	b = 1
	c = 1
	d = 1
	e = 1
	f = 1
	g = 1
	fmt.Printf("The size of the uint8 is %d\n", unsafe.Sizeof(a))
	fmt.Printf("The size of the byte is %d\n", unsafe.Sizeof(b))
	fmt.Printf("The size of the int32 is %d\n", unsafe.Sizeof(c))
	fmt.Printf("The size of the rune is %d\n", unsafe.Sizeof(d))
	fmt.Printf("The size of the int is %d\n", unsafe.Sizeof(e))
	fmt.Printf("The size of the int8 is %d\n", unsafe.Sizeof(f))
	fmt.Printf("The size of the int64 is %d\n", unsafe.Sizeof(g))
}

func demo2() {
	/*
		Go 语言和C语言一样，编译器会尽量提高精确度，以避免计算中的精度损失。
		所以说,除非手动指定精度,否则浮点类型会自动推导为float64
	*/
	var a = 1
	b := 2.0
	fmt.Printf("The a is %d\n", a)
	fmt.Printf("The b is %f\n", b)
	fmt.Printf("The size of the a is %d\n", unsafe.Sizeof(a))
}

var a = 1

func demo3() {
	fmt.Printf("The a is %d\n", a)
	a += 10
	fmt.Printf("The a is %d\n", a)
	a := 2.0
	fmt.Printf("The a is %f\n", a)
}

func demo4() {
	var a float64 = 1.111111111111111111111111111111111
	var b float64 = 1.111111111111111111111111111111111111
	fmt.Printf("The a is %.30f\n", a)
	fmt.Printf("The b is %.30f\n", b)
	fmt.Printf("The a is equal to b is %t\n", a == b)
}

func demo5() {
	var a = 13.14159265358979323846
	fmt.Printf("The a is %2.3g\n", a)
	fmt.Printf("The a is %e\n", a)
}

func demo6() {
	var a uint16 = 256
	var b uint8 = uint8(a)
	fmt.Printf("the b is %d\n", b)
}

func demo7() {
	var a uint = 3
	var b = ^a
	fmt.Printf("The binary a %b\n", a)
	fmt.Printf("The binary b %b\n", b)
}

type myInt int
type sameInt = int

func (value myInt) getValue() {
	fmt.Printf("The value is %d\n", value)
}

func demo8() {
	var a int = 1
	var b sameInt = 2
	var c myInt = 3
	b = a
	fmt.Printf("The b value is %d\n", b)
	c = myInt(a)
	c.getValue()
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
	demo7()
	demo8()
}
