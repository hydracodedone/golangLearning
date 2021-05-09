package main

import "fmt"

const CONST = 1.1111111111111111111111111111111111111
const CONST1 float32 = 1.1111111111111111111111111111111111111
const CONST2 float64 = 1.1111111111111111111111111111111111111

func main() {
	var c float64 = CONST
	var d float32 = CONST
	var e float32 = CONST1
	var f float64 = CONST2
	fmt.Printf("The CONST is %v\n", CONST)
	fmt.Printf("The CONST is %.150f\n", CONST)
	fmt.Printf("The c is %.150f\n", c)
	fmt.Printf("The d is %.150f\n", d)
	fmt.Printf("The e is %.150f\n", e)
	fmt.Printf("The f is %.150f\n", f)
}
