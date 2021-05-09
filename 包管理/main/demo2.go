package main

import "fmt"

var a int = 1

func init() {
	fmt.Println("init1")
	fmt.Printf("The value of the a is %d\n", a)
	a++
}

func init() {
	fmt.Println("init2")
	fmt.Printf("The value of the a is %d\n", a)
	a++

}
func main() {
	fmt.Println("main")
	fmt.Printf("The value of the a is %d\n", a)
}
