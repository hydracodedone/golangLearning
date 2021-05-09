package main

import "fmt"

// error是一个接口,需要实现Error()string方法
type myError interface {
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
func main() {
	var myErrorStructInstance myErrorStruct = myErrorStruct{}
	myErrorStructInstance.setErrorInfo(404, "NotFound")
	var err1 error
	err1 = &myErrorStructInstance
	fmt.Println(err1)

}
