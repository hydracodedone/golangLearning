package main

import "fmt"

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
	inherit()
}
