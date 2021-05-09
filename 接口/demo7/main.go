package main

import (
	"fmt"
	"reflect"
)

type writer interface {
	write() string
}
type people struct {
}

func (p people) write() string {
	return "hello world"
}

/*
nil在Go语言中只能被赋值给指针和接口。
接口在底层的实现有两部分：type和data
在源码中，显示地将nil赋值给接口时，接口的type和data都将为nil。此时，接口与nil值判断是相等的。
但如果将一个带有类型的nil赋值给接口时，只有data为nil，而type不为nil，此时，接口与nil判断将不相等。
*/
func demo() {
	var writerInterface writer
	fmt.Printf("the writerInterface is %v,type is %v, equal nil is %v \n", writerInterface, reflect.TypeOf(writerInterface), writerInterface == nil)
	var peopleInstance people
	writerInterface = peopleInstance
	fmt.Printf("the writerInterface is %v,type is %v, equal nil is %v \n", writerInterface, reflect.TypeOf(writerInterface), writerInterface == nil)
}
func main() {
	demo()
}
