package main

import "fmt"

type student struct {
	Name string
	Age  int
	grade
}
type student2 struct {
	Name         string
	Age          int
	StudentGrade grade
}
type student3 struct {
	Name         string
	Age          int
	StudentGrade grade
	pointer      float64
}
type student4 struct {
	Name string
	Age  int
	grade
	pointer float64
}
type grade struct {
	MathGrade    float64
	EnglishGrade float64
	pointer      float64
}

/*
嵌套结构的struct初始化的时候必须要遵循结构体定义，否则会编译报错
*/
func nestficationStruct() {
	stu1 := student{
		"Hydra",
		23,
		grade{
			92,
			90,
			4.0,
		},
	}
	fmt.Printf("The stu1 mathgrade is %f\n", stu1.grade.MathGrade)
	fmt.Printf("The stu1 mathgrade is %f\n", stu1.MathGrade)
	fmt.Printf("The stu1 pointer is %f\n", stu1.grade.pointer)
	fmt.Printf("The stu1 pointer is %f\n", stu1.pointer)
}
func nestficationStruct2() {
	//不是匿名嵌套结构体,不能使用语法塘
	stu1 := student2{
		"Hydra",
		23,
		grade{
			92,
			90,
			4.0,
		},
	}
	fmt.Printf("The stu1 mathgrade is %f\n", stu1.StudentGrade.MathGrade)
	fmt.Printf("The stu1 pointer is %f\n", stu1.StudentGrade.pointer)
}
func nestficationStruct3() {
	stu1 := student3{
		"Hydra",
		23,
		grade{
			92,
			90,
			4.0,
		},
		5.0,
	}
	fmt.Printf("The stu1 pointer inner is %f\n", stu1.StudentGrade.pointer)
	fmt.Printf("The stu1 pointer is %f\n", stu1.pointer)
}
func nestficationStruct4() {
	stu1 := student4{
		"Hydra",
		23,
		grade{
			92,
			90,
			4.0,
		},
		5.0,
	}
	fmt.Printf("The stu1 pointer is %f\n", stu1.pointer)
	fmt.Printf("The stu1 inner pointer is %f\n", stu1.grade.pointer)
}
func main() {
	nestficationStruct4()
}
