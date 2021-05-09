package main

import "fmt"

type myError interface {
	error
}
type myError2 interface {
	Error() string
}
type myError3 interface {
	error
	setErrorInfo(int, string)
}

type myErrorStruct struct {
	errorCode int
	errorInfo string
}

func (self *myErrorStruct) Error() string {
	return fmt.Sprintf("Error Code:<%d>,Error Info:<%s>\n", self.errorCode, self.errorInfo)
}
func (self *myErrorStruct) setErrorInfo(errorCode int, errorInfo string) {
	self.errorCode = errorCode
	self.errorInfo = errorInfo
}
func demo() {
	var myErrorStructInstance myErrorStruct = myErrorStruct{}
	myErrorStructInstance.setErrorInfo(404, "NotFound")
	var err1 error
	var err2 myError2
	var err3 myError3
	err1 = &myErrorStructInstance
	err2 = &myErrorStructInstance
	err3 = &myErrorStructInstance
	fmt.Println(err1.Error())
	fmt.Println(err2.Error())
	fmt.Println(err3.Error())
}

func main() {
	demo()
}
