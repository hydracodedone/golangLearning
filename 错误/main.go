package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type myError struct {
	code      int
	errorInfo string
}

//标准的函数异常处理模式
//自定义异常类
/*
如何理解Error:
error类型是go语言的一种内置类型，使用的时候不用特定去import,他本质上是一个接口,
type error interface{
  	Error() string //Error()是每一个订制的error对象需要填充的错误消息,可以理解成是一个字段Error
}
*/
func (m *myError) Error() string {
	return fmt.Sprintf("The code is %d,The error info is %s\n", m.code, m.errorInfo)
}
func demo() (*os.File, error) {
	file, err := os.Open("./Demo.md")
	if err != nil {
		return nil, &myError{
			code:      404,
			errorInfo: err.Error(),
		}
	}
	return file, nil
}

func main() {
	file, err := demo()
	if err != nil {
		//error 类型断言
		if err, ok := err.(*os.PathError); ok {
			fmt.Println(err)
		} else {
			fmt.Printf("The Error is [%s]\n", err)
		}
	} else {
		reader := bufio.NewReader(file)
		readInfo, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Printf("The info from file is %s\n", readInfo)
		} else {
			fmt.Printf("Read from file fail ,fail info is %s\n", err)
		}
	}
}
