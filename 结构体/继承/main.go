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

type father struct {
	name string
	age  int
}
type son struct {
	grade int
	father
}

func (f *father) work() {
	fmt.Printf("name : %s is working\n", f.name)
}
func (s *son) learning() {
	fmt.Printf("son get grade :%d\n", s.grade)
}
func inherit() {
	var temp son = son{
		99,
		father{
			"hydra",
			23,
		},
	}
	temp.work()
	temp.learning()
}

func main() {
	demo()
	inherit()
}

// type Engine interface {
// 	start()
// 	end()
// }

// type Car struct {
// 	Engine
// }
// type ActualEngine struct {
// }

// func (actualEngineInstance ActualEngine) start() {}

// func (actualEngineInstance *ActualEngine) start() {}
// func (actualEngineInstance *ActualEngine) end()   {}

// func demo() {
// 	var actualEngine ActualEngine = ActualEngine{}
// 	var engine Engine = actualEngine
// 	fmt.Printf("The engine is %v\n", engine)
// 	var a Engine = ActualEngine{}
// 	actualCar := Car{
// 		Engine: a,
// 	}
// 	fmt.Printf("The actualCar is: %#v\n", actualCar)
// 	fmt.Printf("The actualCar.Engine is: %#v\n", actualCar.Engine)

// }

// func main() {
// 	demo()
// }
