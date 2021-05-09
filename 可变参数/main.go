package main

import "fmt"

func test(value ...int) {
	fmt.Printf("The type of the value is %T,the value is %#v\n", value, value)
}
func test2(value []int) {
	fmt.Printf("The type of the value is %T,the value is %#v\n", value, value)
}
func main() {
	//可变参数可以接受0或者多个变量
	//切片作为参数时候,必须要申明,最起码是nil
	test(1, 2, 3)
	temp := []int{1, 2, 3}
	test2(temp)
	test()
	var temp2 []int
	test2(temp2)
	//test2()
}
