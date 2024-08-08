package main
import "fmt"

type sayner1 interface {
	say()
}
type sayner2 interface {
	say()
}
type singer interface {
	sing()
}
type people1 struct{}

func (p *people1) say()  {}
func (p *people1) sing() {}

type people2 struct{}

func (p *people2) say() {}
func demo1() {
	var s sayner1
	fmt.Println(s == nil)
	var p *people1
	fmt.Println(p == nil)
	s = p
	fmt.Println(s == nil)
}
func demo2() {
	var s1 sayner1
	var s2 sayner2
	var p *people1
	s1 = p
	s2 = p
	fmt.Println(s1 == s2)
}
func demo3() {
	var s11 sayner1
	var s12 sayner1
	var p11 *people1 = &people1{}
	var p12 *people1 = &people1{}
	fmt.Println(p11 == p12)
	s11 = p11
	s12 = p12
	fmt.Println(s11 == s12)
}
func main() {
	demo3()
}
