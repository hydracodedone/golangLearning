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
func notificationStruct1() {
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
func notificationStruct2() {
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
func notificationStruct3() {
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
func notificationStruct4() {
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
func review() {
	type innerInner struct {
		innerInnerParam1 int
		innerParam1      int
	}
	type inner struct {
		innerParam1 int
		innerInner
	}
	type outer struct {
		outerParam1 int
		int
		inner
	}
	var outerInstance = outer{
		outerParam1: 1,
		int:         2,
		inner: inner{
			innerParam1: 3,
			innerInner: innerInner{
				innerInnerParam1: 4,
				innerParam1:      5,
			},
		},
	}
	var outerInstance2 = outer{
		1,
		2,
		inner{
			3,
			innerInner{
				4,
				5,
			},
		},
	}
	var outerInstance3 = outer{
		1,
		2,
		inner{
			innerParam1: 3,
			innerInner: innerInner{
				4,
				5,
			},
		},
	}

	fmt.Printf("The outerInstance  is %#v\n", outerInstance)
	fmt.Printf("The outerInstance.innerParam1 is %d\n", outerInstance.innerParam1)
	fmt.Printf("The outerInstance.inner.innerParam1 is %v\n", outerInstance.inner.innerInnerParam1)
	fmt.Printf("The outerInstance.inner.innerInner.innerParam1 is %v\n", outerInstance.inner.innerInner.innerParam1)
	fmt.Printf("The outerInstance2 is %#v\n", outerInstance2)
	fmt.Printf("The outerInstance3 is %#v\n", outerInstance3)

}

func main() {
	notificationStruct1()
	notificationStruct2()
	notificationStruct3()
	notificationStruct4()
	review()
}
