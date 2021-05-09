package main

import (
	"fmt"
)

func demo1() {
	var a map[int]string
	fmt.Printf("The a is %v\n", a)
	fmt.Printf("The a is nil now is %t\n", a == nil)
	a = make(map[int]string)
	fmt.Printf("The a is %v\n", a)
	fmt.Printf("The a is nil now is %t\n", a == nil)
}

func demo2() {
	var a map[int]string = make(map[int]string, 5) //分配容量
	a[1] = "one"
	a[2] = "two"
	a[3] = "three"
	a[4] = "four"
	a[5] = "five"
	a[6] = "six"
	delete(a, 1) //map删除一个不存在的参数不会报错
	delete(a, 8)
	fmt.Printf("The a[7] is %#v\n", a[7]) //map获取一个不存在的参数不会报错
	value, boolean := a[7]
	fmt.Printf("The a[7] is %#v,exist is %t\n ", value, boolean)
}

func demo3() {
	var selectFunc = make(map[int]func() string, 3)
	selectFunc[1] = func() string {
		return "Hello"
	}
	selectFunc[1] = func() string {
		return "World"
	}
	fmt.Println(selectFunc)
	fmt.Println(selectFunc[1]())
}

func demo4() {
	myTestFunc := func(in map[string]string) {
		for key := range in {
			if key == "name" {
				in[key] = "Hydra2"
			}
		}
	}
	myTestMap := make(map[string]string, 5)
	myTestMap["name"] = "Hydra"
	myTestMap["age"] = "23"
	fmt.Printf("The Begin is %v\n", myTestMap)
	myTestFunc(myTestMap)
	fmt.Printf("The End is %v\n", myTestMap)
}
func demo5() {
	a := map[int]string{
		1: "hello",
		2: "world",
	}
	value, ok := a[99]
	fmt.Printf("The value is %v,the ok is %v\n", value, ok)
	value, ok = a[1]
	if ok {
		fmt.Printf("The value is %v,the ok is %v\n", value, ok)
	} else {
		fmt.Printf("The ok is %v\n", ok)
	}
}
func demo6() {
	a := map[int]string{
		1: "hello",
		2: "world",
	}
	fmt.Printf("length of the a is %d\n", len(a))
	for k, v := range a {
		fmt.Printf("the key is %v,the value is %v\n", k, v)
	}
}

func demo7() {
	//删除键值对并不能保证背后的内存也被回收。
	//解决该问题的办法是重建map
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
	demo6()
	demo7()
}
