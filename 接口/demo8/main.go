package main

import "fmt"

func demo0() {
	var a []interface{} = make([]interface{}, 2)
	a[0] = 1
	a[1] = "string"
	fmt.Println(a)
}
func demo1() {
	myFunc := func(anySlice []interface{}) {
		fmt.Println(anySlice)
	}
	var a []int = make([]int, 2)
	/*
		每个 interface{} 占两个字节
		一个字节用于存放 interface{} 真正传入的数据的类型；
		另一个字节用于指向实际的数据。
		而 一个slice申请的内存空间为 length*(sizeof(elemnent)),不同类型的元素,同length对应的空间不同
	*/
	var b []interface{} = make([]interface{}, len(a))
	for k, v := range a {
		b[k] = v
	}
	myFunc(b)
}
func demo2() {
	var a int = 1
	var b interface{}
	b = a
	var c int
	// c = b
	/*
		interface{} 不能直接赋值给具体的变量,需要进行类型断言
	*/
	c = b.(int)
	fmt.Println(c)
	v, ok := b.(string)
	if ok {
		fmt.Println(v)
	}
}
func main() {
	demo2()
}
