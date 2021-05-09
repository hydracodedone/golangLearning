package main

import "fmt"

func func1(value ...int) {
	fmt.Printf("The type of the value is %T,the value is %#v\n", value, value)
	if len(value) == 3 {
		fmt.Println("begin handle")
		value[0] = 99
	}
}
func func2(value []int) {
	fmt.Printf("The type of the value is %T,the value is %#v\n", value, value)
	if len(value) == 3 {
		value[0] = 99
	}
}
func demo1() {
	//可变参数可以接受0或者多个变量
	//切片作为参数时候,必须要申明,最起码是nil,但是要注意nil切片的调用会导致panic
	a := 1
	b := 2
	c := 3
	func1(a, b, c)
	fmt.Printf("after func1 the result is %d,%d,%d\n", a, b, c)
	d := 4
	e := 5
	f := 6
	temp := []int{d, e, f}
	func1(temp...)
	fmt.Printf("after func1 the result is %v\n", temp)
	fmt.Printf("after func1 the result is %d,%d,%d\n", d, e, f)

	temp2 := []int{1, 2, 3}
	func1(temp2...) //会影响slice
	fmt.Printf("after func1 the result is %v\n", temp2)
}
func demo2() {
	temp := []int{1, 2, 3}
	func2(temp)
	d := 4
	e := 5
	f := 6
	temp2 := []int{d, e, f}
	func2(temp2)
	fmt.Printf("after func1 the result is %v\n", temp2)
	fmt.Println(d, e, f)
}
func main() {
	demo1()
}
