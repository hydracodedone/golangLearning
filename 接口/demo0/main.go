package main

import "fmt"

type myInterface interface {
	saySomething(info string) (syaInfo string) //接口声明中,函数的参数名与实现的参数名称不需要保持一致,实际上申明不需要说明形参的名称
}

type myStruct struct{}

func (myStructInstance myStruct) saySomething(greeting string) (greetingInfo string) {
	greetingInfo = greeting
	return
}
func main() {
	var myInter myInterface
	myInter = myStruct{}
	info := myInter.saySomething("Hello,World")
	fmt.Println(info)
}
