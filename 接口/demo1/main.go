package main

import "fmt"

/*
接口是一种特殊的类型,规定了变量应该实现哪些方法
*/
type getValueInterface interface {
	getValue() int
}

type myInt struct {
	value int
}
type myFloat struct {
	value float32
}

func (t myInt) getValue() int {
	return t.value
}
func (t myFloat) getValue() int {
	return int(t.value)
}

func sum(t []getValueInterface) {
	var result int
	for _, value := range t {
		result += value.getValue()
		fmt.Printf("The Type is %T\n", value)
	}
	fmt.Printf("The Result is %d\n", result)
}
func main() {
	var a *myFloat
	var b *myInt
	a = new(myFloat)
	b = new(myInt)
	a.value = 100.
	b.value = 200
	var slice []getValueInterface //通过接口变量实现了解耦
	slice = append(slice, a)
	slice = append(slice, b)
	sum(slice)
}
