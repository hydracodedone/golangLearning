package main

import "fmt"

func anonymouseElement() {
	type grade struct {
		math    float32
		chinese float32
	}
	type student struct {
		string
		int
		grade
	}
	var stu student
	stu.string = "Hydra"
	stu.int = 23
	var stu2 student = student{
		"Hydra",
		23,
		grade{
			100.0,
			99.0,
		},
	}
	var stu3 student = student{
		string: "Hydra",
		int:    23,
		grade: grade{
			100.0,
			99.0,
		},
	}
	fmt.Printf("The stu is %v\n", stu)
	fmt.Printf("The stu2 is %v\n", stu2)
	fmt.Printf("The stu3 is %v\n", stu3)
}

func main() {
	anonymouseElement()
}
