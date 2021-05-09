package main

import (
	"fmt"
	"unicode/utf8"
)

/*
golang中有一个byte数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：
byte 等同于int8，常用来处理ascii字符
rune 等同于int32,常用来处理unicode或utf-8字符
*/
func demo() {
	var stringDemo string
	stringDemo = "Hello 中国"
	fmt.Printf("The length of the stringdemo1 is %d\n", len(stringDemo))
	fmt.Printf("The length of the stringdemo1 is %d\n", len([]rune(stringDemo)))
	fmt.Printf("The length of the stringdemo1 is %d\n", utf8.RuneCountInString(stringDemo))
	var temp rune
	temp = '中'
	fmt.Printf("The type of the temp is %T\n", temp)
}

func main() {
	demo()
}
