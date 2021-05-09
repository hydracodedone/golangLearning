package main

import "fmt"

func main() {
	//	go语言中return函数不是原子操作,而是给返回值赋值和执行RET执行两部分,而defer是在返回值赋值后,RET指令执行之前
	// return a 可以理解为  return_value =a  RET return_value
	// defer 为 return_value=a  def RET return_value
	excute2()
}

func deferTest() int {
	var a int = 1
	fmt.Printf("1.The address of the a is %p, the value is %d\n", &a, a)
	defer func() {
		a++
		fmt.Printf("3.The address of the a is %p, the value is %d\n", &a, a)
	}()
	fmt.Printf("2.The address of the a is %p, the value is %d\n", &a, a)
	return a
}

func excute1() {
	//需要理解的地方是,return是将待return的值赋值给return_value 相当于作了一次拷贝
	var a int = deferTest()
	fmt.Printf("4.The address of the a is %p, the value is %d\n", &a, a)
}

func deferTest2() *int {
	var a *int
	var b int = 1
	a = &b
	fmt.Printf("1.The address of the a is %p, the value is %d\n", a, *a)
	defer func() {
		(*a)++
		fmt.Printf("3.The address of the a is %p, the value is %d\n", a, *a)
	}()
	fmt.Printf("2.The address of the a is %p, the value is %d\n", a, *a)
	return a
}

func excute2() {
	//需要理解的地方是,return是将待return的值赋值给return_value 相当于作了一次拷贝
	var a *int = deferTest2()
	fmt.Printf("4.The address of the a is %p, the value is %d\n", a, *a)
}
