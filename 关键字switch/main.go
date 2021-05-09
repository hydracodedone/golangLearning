package main

import "fmt"

func demo1() {
	var a interface{}
	//type switch
	switch i := a.(type) {
	case nil:
		fmt.Printf("i is %v,type is nil\n", i)
	case string:
		fmt.Printf("i is %v,type is string\n", i)
	case int:
		fmt.Printf("i is %v,type is int\n", i)
	default:
		fmt.Printf("i don't know")
	}
}
func demo2(a int) {
	b := 10
	switch a {
	default: //可以写在前面,但是一般写在最后
		fmt.Println("default")
	case 1:
		fmt.Println("a is 1")
	case 2:
		fmt.Println("a is 2")
	case 3, 5, 7:
		fmt.Println("a is in [3 5 7]")
		fallthrough //不能用作switch的最后一个case
	case b + 2:
		fmt.Println("a is b+2")
	}
}
func main() {
	demo1()
	demo2(3)
	demo2(33)

}
