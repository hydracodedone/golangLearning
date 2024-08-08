package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Student struct {
	Name string
	Age  int
}
type People struct {
	Names string
	Ages  int32
}

func (p *People) Name(newName string) {
	p.Names = newName
}
func (s *Student) Ages() int {
	return s.Age
}
func Demo1() {
	fromValue := &Student{
		Name: "Hydra",
		Age:  23,
	}
	toValue := &People{}
	copier.CopyWithOption(toValue, fromValue, copier.Option{})
	fmt.Println(toValue)
}
func Demo2() {
	var fromValue []int = []int{1, 2, 3}
	var toValue []int = make([]int, 0)
	err := copier.CopyWithOption(&toValue, fromValue, copier.Option{})
	if err != nil {
		panic(err)
	}
	fromValue[0] = 100
	fmt.Println(fromValue, toValue)
}
func Demo3() {
	first := 1
	second := 2
	third := 3
	forth := 4
	var fromValue []int = []int{first, second, third, forth}
	var toValue []int = make([]int, 3)
	err := copier.CopyWithOption(&toValue, &fromValue, copier.Option{DeepCopy: false})
	if err != nil {
		panic(err)
	}
	fromValue[0] = 100
	fmt.Println(fromValue, toValue)
}

type FromType struct {
	SliceData []int
	MapData   map[int]string
}

type ToType struct {
	SliceData2 []int          `copier:"SliceData"`
	MapData2   map[int]string `copier:"MapData"`
}

func Demo4() {
	sliceData := []int{1, 2, 3}
	mapData := map[int]string{1: "first", 2: "second"}
	fromData := FromType{
		SliceData: sliceData,
		MapData:   mapData,
	}
	toData := ToType{}
	err := copier.CopyWithOption(&toData, fromData, copier.Option{DeepCopy: true})
	if err != nil {
		panic(err)
	}
	sliceData[0] = 100
	fmt.Println(toData)
}
func Demo5() {
	fromValue := &Student{
		Name: "",
		Age:  23,
	}
	toValue := &People{
		Names: "Hydra2",
	}
	copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: true})
	fmt.Println(toValue)
}
func main() {
	Demo5()
}
