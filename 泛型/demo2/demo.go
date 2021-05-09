package main

import "fmt"

func myAdd[T int | string](a, b T) int {
	var myTypeA interface{} = a
	var myTypeB interface{} = b
	result := 0
	switch myTypeA.(type) {
	case string:
		result = len(string(a)) + len(string(b))
	case int:
		result = myTypeA.(int) + myTypeB.(int)
	}
	return result
}
func main() {
	a1 := 1
	b1 := 2
	a2 := "1"
	b2 := "2"
	fmt.Println(myAdd(a1, b1))
	fmt.Println(myAdd(a2, b2))
}
