package main

import "fmt"

func demo1() {
	var a map[*int]*int = make(map[*int]*int, 3)
	var b []int = []int{0, 1, 2}
	for key, value := range b {
		fmt.Println(&key, &value)
		a[&key] = &value
	}
	for k, v := range a {
		fmt.Println(*k, *v)
	}
}

func demo2() {
	a := []int{1, 2, 3}
	for _, v := range a {
		a = append(a, v)
	}
	fmt.Println(a)
}
func demo3() {
	a := map[int]int{1: 1, 2: 2}
	for k, v := range a {
		a[k+3] = v
	}
	fmt.Println(a)
	fmt.Println(len(a))
}

func main() {
	demo3()
}
