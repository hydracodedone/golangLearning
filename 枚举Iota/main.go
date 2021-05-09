package main

import (
	"fmt"
)

/*
IOTA在const出现的时候初始化为0
const每新增一行,则IOTA增加1,即使IOTA没有被使用,也就是说,如果IOTA在第一行没有出现,但是在第二行出现,那么IOTA值是1,而不是2
如果CONST某一个变量没有被赋值,则该变量和上一个变量(常量保持一直,包含IOTA的表达式则表达式一致)保持一致
IOTA在每一行出现多次,每次的值都是固定值
*/
const (
	first  = iota
	second // iota
	third  = 100
	fourth
	fifth = iota
)
const (
	sixth   = iota
	seventh = 1000
	eighth
	ninth = 100 + iota
	tenth
)
const (
	a1, a2 = iota + 1, iota + 2
	_      = iota
	a3, a4 = 100, iota
)

func main() {
	fmt.Printf("The first is %d\n", first)
	fmt.Printf("The second is %d\n", second)
	fmt.Printf("The third is %d\n", third)
	fmt.Printf("The fourth is %d\n", fourth)
	fmt.Printf("The fifth is %d\n", fifth)
	fmt.Printf("The sixth is %d\n", sixth)
	fmt.Printf("The seventh is %d\n", seventh)
	fmt.Printf("The eigth is %d\n", eighth)
	fmt.Printf("The ninth is %d\n", ninth)
	fmt.Printf("The tenth is %d\n", tenth)
	fmt.Printf("The a1 is %d\n", a1)
	fmt.Printf("The a2 is %d\n", a2)
	fmt.Printf("The a3 is %d\n", a3)
	fmt.Printf("The a4 is %d\n", a4)
}
