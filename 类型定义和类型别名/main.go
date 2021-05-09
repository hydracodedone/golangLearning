package main

import "fmt"

type myInt int
type myAnotherInt = int

func main() {
	var myIntInstance myInt = 23
	var intInstance int = 23
	var myAnotherIntInstance myAnotherInt = 23
	fmt.Printf("myIntInstance type is %T\n", myIntInstance)
	fmt.Printf("myAnotherIntInstance type is %T\n", myAnotherIntInstance)
	fmt.Printf("myIntInstance is equal intInstance is:%t\n", int(myIntInstance) == intInstance)
	fmt.Printf("myAnotherIntInstance is equal intInstance is:%t\n", myAnotherIntInstance == intInstance)
}
