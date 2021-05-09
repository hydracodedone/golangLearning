package main

import "fmt"

/*
   接口是一个或多个方法签名的集合。
   任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
   就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
   这称为Structural Typing。
   所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
   当然，该类型还可以有其他方法。

   接口只有方法声明，没有实现，没有数据字段。
   接口可以匿名嵌入其他接口，或嵌入到结构中。
   对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
   只有当接口存储的类型和对象都为nil时，接口才等于nil。
   接口调用不会做receiver的自动转换。
   接口同样支持匿名字段方法。
   接口也可实现类似OOP中的多态。
   空接口可以作为任何类型数据的容器。
   一个类型可实现多个接口。
   接口命名习惯以 er 结尾。
*/
type changer interface {
	changeName(newName string)
}

type people struct {
	name string
}

type pPeople struct {
	name string
}

func (p people) changeName(newName string) {
	p.name = newName
}

func (p *pPeople) changeName(newName string) {
	p.name = newName
}

func main() {
	var p people
	var c changer
	c = p
	c.changeName("people1")
	fmt.Println(p)
	var pp *people = &p
	c = pp
	c.changeName("people2")
	fmt.Println(p)

	var pP pPeople
	var c2 changer
	c2 = &pP
	c2.changeName("people1")
	fmt.Println(pP)
}
