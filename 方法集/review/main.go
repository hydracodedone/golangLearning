package main

import "fmt"

// 接口调用不会做receiver的自动转换。
// 只有当接口存储的类型和对象都为nil时，接口才等于nil。
// 对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。!!

type sayner interface {
	say()
}
type people struct {
}
type people2 struct {
}

func (p people) say() {
}
func (p *people2) say() {
}
func demo() {
	p11 := people{}
	p12 := new(people)
	p11.say()
	p12.say()
	var saynerInstance sayner
	saynerInstance = p11
	saynerInstance.say()
	saynerInstance = p12
	saynerInstance.say()
	p21 := people2{}
	p22 := new(people2)
	p21.say()
	p22.say()
	var saynerInstance2 sayner
	// saynerInstance2 = p21
	// saynerInstance2.say()
	saynerInstance2 = p22
	saynerInstance2.say()
}

type changer interface {
	change(string)
}
type changerPeople1 struct {
	name string
}

func (c changerPeople1) change(newName string) {
	c.name = newName
}

type changerPeople2 struct {
	name string
}

func (c *changerPeople2) change(newName string) {
	c.name = newName
}
func demo2() {
	c11 := changerPeople1{}
	c12 := new(changerPeople1)
	c11.change("newName")
	c12.change("newName")
	fmt.Printf("the c11 newName is %s\n", c11.name)
	fmt.Printf("the c12 newName is %s\n", c12.name)
	var changerInsstance1 changer
	changerInsstance1 = c11
	changerInsstance1.change("newNames")
	changerInsstance1 = c12
	changerInsstance1.change("newNames")
	c21 := changerPeople2{}
	c22 := new(changerPeople2)
	c21.change("newName") //自动取指针
	c22.change("newName")
	var changerInsstance2 changer
	// changerInsstance2 = c21 //不会自动取指针
	// changerInsstance2.change("newNames")
	changerInsstance2 = c22
	changerInsstance2.change("newNames")
	fmt.Printf("the c21 newName is %s\n", c21.name)
	fmt.Printf("the c22 newName is %s\n", c22.name)
}

type inner struct {
}
type outer1 struct {
	inner
}
type outer2 struct {
	*inner
}
type innerSayner interface {
	say()
}
type innerPointSayner interface {
	pointSay()
}

func (i inner) say() {

}
func (i *inner) pointSay() {

}
func demo3() {
	outer1Instance := outer1{}
	outer2Instance := outer2{&inner{}}
	outer1Instance.say()
	outer2Instance.say()
	var sayner innerSayner
	sayner = outer1Instance
	sayner.say()
	sayner = outer2Instance
	sayner.say()
	outer1Instance.pointSay()
	outer2Instance.pointSay()
	var pointSayner innerPointSayner
	// pointSayner = outer1Instance
	// pointSayner.pointSay()
	pointSayner = outer2Instance
	pointSayner.pointSay()
	pointSayner = &outer1Instance
	pointSayner.pointSay()
}

type cat struct {
	name string
}
type runner interface {
	run()
}

func (c cat) run() {

}
func (c cat) eat() {

}

func demo4() {
	catInstance := cat{}
	var runnerInstance runner
	runnerInstance = catInstance
	runnerInstance.run()
	// runnerInstance.eat()
	// fmt.Println(runnerInstance.name)
}
func demo5() {
	recvAnyParameterFunc := func(any interface{}) {
		switch v := any.(type) {
		case string:
			new_v := v + "!!!"
			fmt.Printf("the any is %v\n", new_v)
		case int:
			new_v := v + 1
			fmt.Printf("the any is %v\n", new_v)
		default:
			fmt.Printf("not supported\n")
		}
	}
	recvAnyParameterFunc(23)
}
func main() {
	demo5()
}
