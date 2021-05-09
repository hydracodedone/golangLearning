package main

type sayner interface {
	say()
}

type people struct{}

// 若以值语义 T 作为receiver实现接口，不管是T类型的值，还是T类型的指针，都实现了该接口
func (p people) say() {}

type people2 struct{}

// 若以指针语义 *T 作为receiver实现接口，只有T类型的指针实现了该接口
func (p *people2) say() {}

func demo2() {
	var s sayner
	p := people{}
	s = p
	s.say()
	s = &p
	s.say()
}
func demo3() {
	var s sayner
	p := people2{}
	s = &p
	s.say()
}

func main() {
	demo2()
	demo3()
}
