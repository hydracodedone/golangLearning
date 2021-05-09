package main

import "fmt"

type student struct {
	name      string
	age       int
	gradelist []int
}
type studentSimple struct {
	name string
	age  int
}

/*
当且仅当结构体的所有字段都是可比较的，两个该结构体类型的实例才能用==比较，此时会对其每个字段进行比较
1 匿名结构体可以之间相互比较
2 非匿名结构体之间只有同类型的可以相互比较
3 非匿名结构体可以和匿名结构体(字段是可比较类型的)可以相互比较
*/
func demo1() {
	simpleStudentA := studentSimple{
		"Hydra",
		23,
	}
	simpleStudentB := studentSimple{
		"Hydra",
		23,
	}
	fmt.Printf("simpleStudentA and simpleStudentB is equal: %t\n", simpleStudentA == simpleStudentB)
	studentA := student{
		"Hydra",
		23,
		[]int{99, 98, 97},
	}
	studentB := student{
		"Hydra",
		23,
		[]int{99, 98, 97},
	}
	fmt.Printf("studentA %#v\n", studentA)
	fmt.Printf("studentB %#v\n", studentB)
	// fmt.Printf("studentA and studentB is equal: %t\n", studentA == studentB)
}

func demo2() {
	type people1 struct {
		name string
		age  int
	}
	type people2 struct {
		name string
		age  int
	}
	p1 := people1{"", 0}
	p2 := people2{"", 0}
	p3 := struct {
		name string
		age  int
	}{}
	p4 := struct {
		name string
		age  int
	}{}
	// fmt.Printf("p1==p2 is %T\n", p1 == p2)
	fmt.Printf("p1 is %#v\n", p1)
	fmt.Printf("p2 is %#v\n", p2)
	fmt.Printf("p3 is %#v\n", p3)
	fmt.Printf("p4 is %#v\n", p4)
	fmt.Printf("p1==p3 is %t\n", p1 == p3)
	fmt.Printf("p1==p4 is %t\n", p2 == p3)
	fmt.Printf("p3==p4 is %t\n", p3 == p4)
}
func demo3() {
	type people struct {
		name  string
		grade [3]int
	}
	var a people = people{"", [3]int{}}
	b := struct {
		name  string
		grade [3]int
	}{}
	fmt.Printf("a==b is %t\n", a == b)
}
func main() {
	demo1()
	demo2()
	demo3()
}
