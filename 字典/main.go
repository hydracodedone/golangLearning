package main

import "fmt"

func demo1() {
	/*
		As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
		Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.
	*/
	var a map[int]string
	fmt.Printf("The a is %v\n", a)
	fmt.Printf("The a is nil now is %t\n", a == nil)
}

func demo2() {
	var a map[int]string
	a = make(map[int]string, 5) //分配容量
	a[1] = "one"
	a[2] = "two"
	a[3] = "three"
	a[4] = "four"
	a[5] = "five"
	a[6] = "six"
	fmt.Printf("The a[7] is %#v\n", a[7]) //map获取一个不存在的参数不会报错
	value, boolean := a[7]
	fmt.Printf("The a[7] is %#v,exist is %t\n ", value, boolean) //map获取一个不存在的参数不会报错
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

func main() {
	demo4()
}
