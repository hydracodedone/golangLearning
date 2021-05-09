package main

import "fmt"

func demo1() int {
	var x int = 5
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	//return 对于匿名返回值相当于对x进行了一次拷贝
	return x
}
func demo2() (x int) {
	x = 5
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	//return 对于命名返回值返回不是拷贝
	return
}
func demo3() (x int) {
	x = 5
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	//return 对于命名返回值返回不是拷贝
	return x
}
func demo4() (x int) {
	x = 5
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("The value of the x is %d\n", x)
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	//return 对于命名返回值返回不是拷贝
	return 22
}
func demo5() (x int) {
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("The value of the x is %d\n", x)
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	return 22
}
func demo6() (y int) {
	x := 5
	fmt.Printf("the address of the x is %p\n", &x)
	defer func() {
		//闭包,可以访问外部变量
		fmt.Printf("The value of the x is %d\n", x)
		fmt.Printf("the address of the x is %p\n", &x)
		x++
	}()
	return x
}
func main() {
	value2 := demo2()
	fmt.Printf("The value of the value2 is %d\n", value2)
	value3 := demo3()
	fmt.Printf("The value of the value3 is %d\n", value3)
	value4 := demo4()
	fmt.Printf("The value of the value4 is %d\n", value4)
	value5 := demo5()
	fmt.Printf("The value of the value5 is %d\n", value5)
	value6 := demo6()
	fmt.Printf("The value of the value6 is %d\n", value6)
}
