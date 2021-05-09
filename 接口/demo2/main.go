package main

import "fmt"

type sum interface {
	getSum(a int, b int) (c int)
}
type people struct {
}

//实现接口与接口申明时候制定的变量名无关
func (people) getSum(first int, second int) (result int) {
	result = first + second
	return
}
func main() {
	//接口不可实例化
	var sumTemp sum
	var peopleTemp people
	sumTemp = peopleTemp
	result := sumTemp.getSum(1, 2)
	fmt.Printf("The result is %d\n", result)
}
