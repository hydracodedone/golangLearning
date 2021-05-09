package main

import "fmt"

type errorType string

const (
	notFound  errorType = "not_found"
	notHandle errorType = "not_handle"
	unknown   errorType = "unknown"
)

type CustomError struct {
	ErrorType errorType
	ErrorMsg  string
}

func NewCustomError(eT errorType, eM string) *CustomError {
	return &CustomError{
		ErrorType: eT,
		ErrorMsg:  eM,
	}
}
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error Type:[%s],Error Message:[%s]", e.ErrorType, e.ErrorMsg)
}
func errorHandler(err error) {
	//检查错误类型
	if err != nil {
		switch err.(type) {
		case *CustomError:
			errorInfo := err.Error()
			fmt.Printf("handle Error: <%s>\n", errorInfo)
		default:
			fmt.Printf("handle Error: <%s>\n", err)
		}
	}

	//处理错误
}
func topFunc() {
	err := innerFunc()
	if err != nil {
		errorHandler(err)
		return
	}
	fmt.Println("continue topFunc")

}
func innerFunc() error {
	return NewCustomError(unknown, "unknown error")
}
func demo1() {
	topFunc()
}

func main() {
	demo1()
}
