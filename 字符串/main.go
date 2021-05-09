package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

/*
golang中有一个byte数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：
byte 等同于int8，常用来处理ascii字符
rune 等同于int32,常用来处理unicode或utf-8字符
*/
func demo() {
	stringDemo := "Hello 中国"
	fmt.Printf("The length of the stringdemo1 is %d\n", len(stringDemo))
	fmt.Printf("The length of the stringdemo1 is %d\n", len([]rune(stringDemo)))
	fmt.Printf("The length of the stringdemo1 is %d\n", utf8.RuneCountInString(stringDemo))
	var temp rune = '中'
	fmt.Printf("The type of the temp is %T\n", temp)
}

func demo1() {
	var a = "abc"
	fmt.Printf("the len of the a is %d\n", len(a))
	fmt.Printf("The utf8.RuneCountInString is %d\n", utf8.RuneCountInString(a))
}

func demo2() {
	var a = "中"
	fmt.Printf("the len of the a is %d\n", len(a))
	fmt.Printf("The utf8.RuneCountInString is %d\n", utf8.RuneCountInString(a))
	fmt.Printf("The rune len of the a is %d\n", len([]rune(a)))
}

func demo3() {
	// 单引号 不能用来表示字符串
	// 双引号 支持转义
	// 反引号 不支持转义
	var a = "中"
	var b = '中'
	fmt.Printf("The type of the a is %s\n", reflect.TypeOf(a))
	fmt.Printf("The type of the b is %s\n", reflect.TypeOf(b))
	fmt.Printf("The char b is %c\n", b)
	fmt.Printf("The []rune() of the a is %v\n", []rune(a))
	fmt.Printf("The utf-8 code of the a is %U\n", []rune(a)[0])
	fmt.Printf("The \u4e2d value is %d\n", '\u4e2d')
	fmt.Printf("The hex of the 20013 is %x\n", 20013)
	var c = '\u4e2d'
	var d = '\U00004e2d'
	var e int = '\U00004e2d'
	fmt.Printf("The type of the c is %s\n", reflect.TypeOf(c))
	fmt.Printf("The type of the d is %s\n", reflect.TypeOf(d))
	fmt.Printf("The type of the e is %s\n", reflect.TypeOf(e))
	fmt.Printf("The \u4e2d represent char is %c\n", c)
	fmt.Printf("The 'U' is %U\n", d)
}

func demo4() {
	var a = "Hello World"
	res := strings.Fields(a)
	fmt.Printf("The res is %v\n", res)
	fmt.Printf("The res2 is %v\n", strings.Split(a, " "))
}

func demo5() {
	var a = "hello 中国"
	reader := strings.NewReader(a)
	rune_data, size, err := reader.ReadRune()
	if err != nil {
		fmt.Printf("The err is %v\n", err)
	} else {
		fmt.Printf("the rune readed is %c,the size is %d\n", rune_data, size)
	}
}

func demo6() {
	var a = "hello 中国"
	for index, value := range a {
		fmt.Printf("the index is %d,the vlaue is %c\n", index, value)
	}
}

func main() {
	demo()
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
}
