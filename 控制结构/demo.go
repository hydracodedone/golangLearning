package main

import (
	"fmt"
)

func demo1() {
	a := 1
	if a := a; a == 1 {
		fmt.Println("The a is 1")
	} else {
		fmt.Printf("The a is not 1,is %d\n", a)
	}
}

func demo2() {
	a := 1
	switch b := a + 1; {
	case b > 0:
		fmt.Println("The b is over 0")
		fallthrough
	case b == 1:
		fmt.Println("The b is 1")
	case b == 2:
		fmt.Println("The b is 2")
	default:
		fmt.Println("The b is not 1 or 2")
	}
}
func demo3() {
	for i := 0; i < 5; i++ {
		for j := i; j < 10; j++ {
			println(i, j)
		}
	}
}
func demo4() {
	for i := 0; i < 5; i++ {
		for j := 5 - i; j > 0; j-- {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
}
func demo5() {
	var a = "ABCDE"
	for key, value := range a {
		fmt.Printf("The index is %d,The value is %c\n", key, value)
		value = 2 //val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质
	}
	fmt.Println(a)
}
func demo6() {
	var a = "ABCDE"
	for key, value := range a { //此处的a为外部变量a的拷贝
		fmt.Printf("The index is %d,The value is %c\n", key, value)
		a = "123" //val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质
		fmt.Printf("The a is %v\n", a)
	}
	fmt.Println(a)
}
func demo7() {
	var a = "ABCDE"
	for key, value := range a { //此处的a为外部变量a的拷贝
		fmt.Printf("The index is %d,The value is %c\n", key, value)
		fmt.Printf("The a is %v\n", a)
		a = "GGGG"
		fmt.Printf("The a is %v\n", a)
		value = 100
		fmt.Printf("The a is %v\n", a)
	}
	fmt.Println(a)
}
func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
	demo7()
}
