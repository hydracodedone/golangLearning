package main

import "fmt"

type People struct {
}
type Student struct {
	People
}
type Student2 struct {
	People People
}

func (people *People) walk() {
	fmt.Println("people is working")
}
func (people *People) say() {
	fmt.Println("people is saying")
}
func (student *Student) walk() {
	fmt.Println("student is working")
}
func (student *Student2) walk() {
	fmt.Println("student2 is working")
}
func demo() {
	stu := Student{}
	stu.walk()
	stu.People.walk()
	stu.say()
	stu2 := Student2{}
	stu2.walk()
	stu2.People.walk()
	stu2.People.say()
}
func main() {
	demo()
}
