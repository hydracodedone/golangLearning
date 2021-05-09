package main

import (
	"fmt"
)

func printInterfaceSlice(any []interface{}) {
	for index, value := range any {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
}
func printStringSlice(any []string) {
	for index, value := range any {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
}
func printInterfaceSliceWithTypeAssert(any interface{}) {
	for index, value := range any.([]string) {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
}
func printWithGenericInterface[T string | int](array []T) {
	for index, value := range array {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
}

/*
内置的泛型类型:
any 任意类型,等价于interfac{}
comparable 可比较类型 int uint float bool struct 指针
*/
func demo1() {
	someString := []string{"a", "b"}
	someInt := []int{1, 2}
	//don't work
	// printInterfaceSlice(some)
	//dose work,but not support []int
	printStringSlice(someString)
	//dose work,but not support []int
	printInterfaceSliceWithTypeAssert(someString)
	printWithGenericInterface(someString)
	printWithGenericInterface(someInt)
}

// 匿名函数不支持泛型
// 匿名结构体不支持泛型
func demo2Func[T int | float64](intOrFloat64 []T) {
	for index, value := range intOrFloat64 {
		fmt.Printf("index:%v, value:%v\n", index, value)
	}
}
func demo2() {
	type intSlice []int
	type floatSlice []float64
	type genericIntefaceSlice[T int | float64] []T
	type anyGenericIntefaceSlice[T any] []T
	//need to declare concrete type
	//泛型类型不能直接使用,必须实例化类型为具体的类型
	var intSliceInstance genericIntefaceSlice[int] = []int{1, 2, 3}
	fmt.Printf("intSliceInstance :%v\n", intSliceInstance)
	var floatSliceInstance genericIntefaceSlice[float64] = []float64{1.1, 2.2, 3.3}
	fmt.Printf("floatSliceInstance :%v\n", floatSliceInstance)
	demo2Func(intSliceInstance)
	demo2Func(floatSliceInstance)
	type genericIntefaceMap[KEY int | string, VALUE any] map[KEY]VALUE
	var stringStringMap genericIntefaceMap[string, string] = map[string]string{
		"name": "Hydra",
	}
	fmt.Printf("stringStringMap :%v\n", stringStringMap)
	type Mystruct[T string | int] struct {
		ID   T
		Name string
	}
	var intMystructInstance Mystruct[int] = Mystruct[int]{
		ID:   1,
		Name: "Hydra",
	}
	fmt.Println(intMystructInstance)
	var stringMystructInstance Mystruct[string] = Mystruct[string]{
		ID:   "1",
		Name: "Hydra",
	}
	fmt.Println(stringMystructInstance)
	//泛型嵌套注意类型范围取的是内层的子集,而且实际初始化是达不到内层的差集的
	//比如下面的myStruct 在定义的时候不能取 float32,因为float32不是内层的mySlice的取值的元素
	//在实例化过程中,虽然mySlice可以初始化的类型由float64,但是在myStruct初始化的过程中不能是float64
	type mySlice[T int | string | float64] []T //inner contains outer
	type myStruct[T int | string] struct {
		ID    T
		Name  string
		slice mySlice[T] //T of mySlice must contains int and string
	}
	//指针类型的泛型定义使用interface{}包裹约束类型或者加上逗号消除歧义(建议统一为interface{})
	type pointerSlice[T interface{ *int }] []T
	type pointerSlice2[T interface{ *int | *float64 }] []T
}

// 自定义泛型数据类型

type customTypeChoice interface {
	int | string
}
type myStruct[T customTypeChoice] struct {
	ID   T
	Name string
}

func (m myStruct[T]) getId() T {
	return m.ID
}
func demo3() {
	m := myStruct[string]{
		ID:   "123",
		Name: "Hydra",
	}
	fmt.Println(m.getId())
}

func main() {
	demo3()
}
