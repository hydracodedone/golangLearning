package main

import "fmt"

func testForStruct() {
	//使用new来分配内存空间给一个struct,返回一个指向该struct类型的指针
	type people struct {
		Name string
		age  int
	}
	var peopleDemo people = *(new(people))
	fmt.Printf("The peopleDemo is %#v\n", peopleDemo)
	fmt.Printf("The name is %s,The age is %d\n", peopleDemo.Name, peopleDemo.age)

}

func testForPointerOfStruct() {
	//结构体中的语法糖,使用结构体指针来访问结构体中的成员变量
	temp := &struct {
		Name string
		Age  int
	}{
		Name: "Hydra",
		Age:  23,
	}
	fmt.Printf("The Name is %s\n", temp.Name)
	fmt.Printf("The Name is %s\n", (*temp).Name)
	fmt.Printf("The temp is %#v\n", temp)
}

func testForAnonymousStructEqual() {
	//IMPORTANT
	//匿名结构体的申请
	//结构体变量之间的比较
	temp1 := struct {
		Name string
		Age  int
	}{}
	temp2 := struct {
		Name string
		Age  int
	}{}
	fmt.Printf("The address of the temp1 is %p\n", &temp1)
	fmt.Printf("The address of the temp1 is %p\n", &temp2)
	fmt.Println(temp1 == temp2)
	temp1.Name = "ss"
	fmt.Println(temp1 == temp2)
}
func testForAnonymousStruct() {
	//匿名结构体变量的两种定义方式
	var s struct {
		name string
		age  int
	}
	s2 := struct {
		name string
		age  int
	}{}
	fmt.Printf("The s is %#v\n", s)
	fmt.Printf("The s2 is %#v\n", s2)
}
func structInit() {
	type student struct {
		name string
		age  int
	}
	var stu1 student
	stu1.name = "Hydra"
	stu1.age = 23
	var stu2 = student{
		"Hydra",
		23, //逗号不可少
	}
	var stu3 = student{
		name: "Hydra",
		age:  23,
	}
	var stu4 = student{
		age:  23,
		name: "Hydra",
	}
	var stu5 = student{
		age: 23,
	}
	stu6 := struct {
		name string
		age  int
	}{
		name: "hydra",
		age:  23,
	}
	stu7 := struct {
		name string
		age  int
	}{
		"hydra",
		23,
	}
	fmt.Printf("%v\n", stu1)
	fmt.Printf("%v\n", stu2)
	fmt.Printf("%v\n", stu3)
	fmt.Printf("%v\n", stu4)
	fmt.Printf("%v\n", stu5)
	fmt.Printf("%v\n", stu6)
	fmt.Printf("%v\n", stu7)
}
func typeAlias() {
	type number struct {
	}
	type nu number
	temp := number{}
	fmt.Printf("The type of the temp is %T\n", temp)
	var temp2 nu = nu(temp)
	fmt.Printf("The type of the temp2 is %T\n", temp2)
	var temp3 nu = nu{}
	fmt.Printf("The type of the temp3 is %T\n", temp3)
	var temp4 *nu = new(nu)
	fmt.Printf("The type of the temp3 is %T\n", temp4)
}
func main() {
	testForStruct()
	testForAnonymousStruct()
	testForAnonymousStructEqual()
	testForPointerOfStruct()
	structInit()
	typeAlias()
}
