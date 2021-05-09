package main

import "fmt"

func demo1() {
	var a [3]int
	var b [3]int
	a = [3]int{0: 0, 2: 2}
	b = a
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func demo2() {
	a := [3]int{1, 2, 3}
	b := [3]int{1, 2, 3}
	fmt.Printf("a == b is %t\n", a == b)
}
func main() {
	demo2()
}
