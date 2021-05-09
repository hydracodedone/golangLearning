package main

import "fmt"

func demo() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)
	var b [3]int = [...]int{1, 2, 3}
	fmt.Println(b)
	var c = [3]int{0: 1, 2: 2}
	fmt.Println(c)
	d := [...][2]int{{1, 2}, {3, 4}}
	fmt.Println(d)
}

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
	demo()
	demo1()
	demo2()
}
