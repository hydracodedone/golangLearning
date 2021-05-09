package main

import "fmt"

// MyString is string
type MyString string

//MyInt is int
type MyInt int

// MyPrint is a method for type MyString
func (s MyString) MyPrint() {
	fmt.Printf("The address of the string is %p\n", &s)
}

//MyPrint is a method for type MyInt
func (s MyInt) MyPrint() {
	fmt.Printf("The address of the string is %p\n", &s)
}
func _() {
	var s MyString = "hello,world"
	s.MyPrint()
	var s2 MyInt = 2
	s2.MyPrint()
}

type student struct {
	Name string
	Age  int
}

func (s student) changeName(newName string) {
	s.Name = newName
}
func (s *student) changeNameForPointer(newName string) {
	s.Name = newName
}
func methodInfluenceStruct() {
	stu := student{
		"hydra",
		23,
	}
	stu.changeName("Hydra")
	fmt.Printf("The name change is %s\n", stu.Name)
	stu.changeNameForPointer("Hydra2") //优先级
	fmt.Printf("The name change is %s\n", stu.Name)
	(&stu).changeNameForPointer("Hydra3") //优先级
	fmt.Printf("The name change is %s\n", stu.Name)
	stu2 := new(student)
	stu2.changeName("123")
	fmt.Printf("The stu2 name is %s\n", stu2.Name)
}

func main() {
	methodInfluenceStruct()
}
