package main

import (
	"fmt"
	"reflect"
)

type customIntType int
type intAlias = int

func demo1() {
	var customIntTypeInstance customIntType = 23
	var intAliasInstance intAlias = 23
	var intInstance int = 23

	intInstance = int(customIntTypeInstance)
	customIntTypeInstance = customIntType(intInstance)
	intAliasInstance = intInstance
	intInstance = intAliasInstance
	fmt.Printf("intInstance type is %T\n", intInstance)
	fmt.Printf("customIntTypeInstance type is %T\n", customIntTypeInstance)
	fmt.Printf("intAliasInstance type is %T\n", intAliasInstance)
}

type stack1 = []int

type stack2 []int

// 类型别名不能拥有自定义方法
// func (s1 stack1) push() {}

func (s2 stack2) push(data int) {
	if s2 == nil {
		s2 = make(stack2, 0)
		s2 = make([]int, 0)
		s2 = []int{}
	}
	s2 = append(s2, data) //stack2本质上是切片类型,因此append方 法可以对该自定义类型适用
}

func demo2() {
	var s []int
	var s2 stack2
	var s1 stack1
	s2 = s
	s2 = stack2(s)
	s2.push(1)
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(reflect.TypeOf(s2))
}

func main() {
	demo2()
}
