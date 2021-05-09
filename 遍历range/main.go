package main

import "fmt"

func forLoopAndAnonymousFunction() {
	a := []int{1, 2, 3, 4, 5}
	for key, value := range a {
		func() {
			fmt.Printf("The key is %d The value is %d\n", key, value)
			value++
		}()
	}
	for key, value := range a {
		func() {
			fmt.Printf("The key is %d The value is %d\n", key, value)
			value++
		}()
	}
}

func strangeRangeForArray() {
	a := [3]int{1, 2, 3}
	// 以下代码遍历数组和遍历slice结果不一样
	for i, v := range a { //对整个a进行了一次拷贝
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}
func strangeRangeForSlice() {
	a := []int{1, 2, 3}
	// 以下代码遍历数组和遍历slice结果不一样
	for i, v := range a { //拿到的是a的地址的拷贝
		if i == 0 {
			a[1], a[2] = 200, 300
			fmt.Println(a)
		}
		a[i] = v + 100
	}
	fmt.Println(a)
}
func main() {
	strangeRangeForArray()
	strangeRangeForSlice()
	forLoopAndAnonymousFunction()
}
