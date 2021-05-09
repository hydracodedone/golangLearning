package main

import "fmt"

type cat struct {
	Name string
	Age  int
}

type blackCat struct {
	*cat
	otherInfo string
}

func newCat(Name string, Age int) *cat {
	temp := new(cat)
	temp.Name = Name
	temp.Age = Age
	return temp
}

func newBlackCat(Name string, Age int) *blackCat {
	temp := new(blackCat)
	temp.cat = newCat(Name, Age)
	temp.otherInfo = "THIS IS BlackCat"
	return temp
}

func newBlackCat2() *blackCat {
	temp := new(blackCat)
	temp.otherInfo = "THIS IS BlackCat"
	return temp
}

func main() {
	catInstance := newBlackCat("Hydra", 23)
	catInstance2 := newBlackCat2()
	fmt.Printf("The Cat is %#v\n", catInstance.cat)
	fmt.Printf("The Cat is %#v\n", catInstance2.cat)
}
