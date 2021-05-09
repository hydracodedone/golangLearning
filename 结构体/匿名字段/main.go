package main

import "fmt"

func anonymouseElement() {
	type student struct {
		string
		int
	}
	var stu student
	stu.string = "Hydra"
	stu.int = 23
	var stu2 student = student{
		"Hydra",
		23,
	}
	var stu3 student = student{
		string: "Hydra",
		int:    23,
	}
	fmt.Printf("The stu is %v\n", stu)
	fmt.Printf("The stu is %v\n", stu2)
	fmt.Printf("The stu is %v\n", stu3)
}

func main() {
	anonymouseElement()
}
