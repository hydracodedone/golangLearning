package main

import "fmt"

/*
panic 会使得panic后续的defer语句实效
因此在处理一些释放资源的操作时候,为了避免遇到未知错误,应该将对应关闭资源的语句提前
*/
func main() {
	defer func() {
		fmt.Println("BEFORE")
	}()
	panic("This is a test panic")
	defer func() {
		fmt.Println("AFTER")
	}()
}
