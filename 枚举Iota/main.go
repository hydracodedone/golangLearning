package main

import (
	"fmt"
)

/*
IOTA在const出现的时候初始化为0,简单地讲，每遇到一次 const 关键字，iota 就重置为 0
const每新增一行,则IOTA增加1,即使IOTA没有被使用,也就是说,如果IOTA在第一行没有出现,但是在第二行出现,那么IOTA值是1,而不是2
如果CONST某一个变量没有被赋值,则该变量和上一个变量(常量保持一直,包含IOTA的表达式则表达式一致)保持一致
IOTA在每一行出现多次,每次的值都是固定值
*/
const (
	first  = iota
	second // iota
	third  = 100
	fourth // fourth=4
	fifth  = iota
)
const (
	sixth   = iota
	seventh = 1000
	eighth
	ninth = 100 + iota
	tenth
)
const (
	a1, a2 = iota + 1, iota + 2 // 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
	_      = iota
	a3, a4 = 100, iota
)
const (
	temp1, temp2 = iota + 1, iota + 2
	temp3,
	temp4
)
const (
	_  = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota) // 1 << (10*1)
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
)

func demo() {
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
	fmt.Printf("The temp1 is %d\n", temp1)
	fmt.Printf("The temp2 is %d\n", temp2)
	fmt.Printf("The temp3 is %d\n", temp3)
	fmt.Printf("The temp4 is %d\n", temp4)
	fmt.Printf("The KB is %d\n", KB)
	fmt.Printf("The MB is %d\n", MB)
	fmt.Printf("The GB is %d\n", GB)
}

const (
	a = iota
	b
	c = 10
	d
	e, f = iota, iota + 1
)

func review() {
	fmt.Println("this is main")
	fmt.Printf("The a is %d\n", a)
	fmt.Printf("The b is %d\n", b)
	fmt.Printf("The c is %d\n", c)
	fmt.Printf("The d is %d\n", d)
	fmt.Printf("The e is %d\n", e)
	fmt.Printf("The f is %d\n", f)
}

func main() {
	review()
}
