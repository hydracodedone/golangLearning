package main

import (
	"fmt"
)

type CalAdd interface {
	add(x int, y int) (res int)
}
type People struct {
	Name string
	Age  int
}

func (p *People) add(a int, b int) int {
	return a + b
}
func main() {
	p := new(People)
	p.Name = "Test"
	p.add(1, 2)
	fmt.Printf("The type of the p is %T\n", p)
	var calAdd CalAdd = p
	fmt.Printf("The type of the calAdd is %T\n", calAdd)
	calAdd.add(2, 3)
	calAdd2 := CalAdd(p) //另一种接口变量的赋值模式
	calAdd2.add(3, 4)
}
