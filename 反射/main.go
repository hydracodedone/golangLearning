package main

import (
	"fmt"
	"reflect"
)

type Myint int
type Subject struct {
	Math    string
	Chinese string
	English string
}
type Student struct {
	Name string
	Age  Myint `this is a custom type`
	Info any
	Subject
}
type getSetName interface {
	getName() string
	setName(string)
}

func (s Student) PublicGetName() string {
	return s.Name
}
func (s Student) PublicSetName(otherName string) {
	s.Name = otherName
}
func (s Student) getName() string {
	return s.Name
}
func (s *Student) setName(newName string) {
	s.Name = newName
}
func demo1() {
	s := Student{
		Name: "Hydra",
		Age:  23,
		Subject: Subject{
			Math:    "Good",
			Chinese: "Bad",
			English: "Excellent",
		},
	}
	typeOfStudent := reflect.TypeOf(s)
	fmt.Printf("The reflect.TypeOf is:<%v>\n", typeOfStudent)
	fmt.Printf("The reflect.TypeOf Size is:<%v>\n", typeOfStudent.Size())
	fmt.Printf("The reflect.TypeOf PkgPath is:<%v>\n", typeOfStudent.PkgPath())
	fmt.Printf("The reflect.TypeOf Name is:<%v>\n", typeOfStudent.Name())
	fmt.Printf("The reflect.TypeOf Kind is:<%v>\n", typeOfStudent.Kind())                                                      //Kind是原始类型
	fmt.Printf("The reflect.TypeOf Implements is:<%v>\n", typeOfStudent.Implements(reflect.TypeOf((*getSetName)(nil)).Elem())) //判断类型是否实现了某接口

	fmt.Printf("The reflect.TypeOf(s.PublicSetName) is:<%v>\n", reflect.TypeOf(s.PublicSetName))                 //直接获取某个方法的信息
	fmt.Printf("The reflect.TypeOf(s.PublicSetName).NumIn() is:<%v>\n", reflect.TypeOf(s.PublicSetName).NumIn()) //获取入参个数
	fmt.Printf("The reflect.TypeOf(s.PublicSetName).In(0) is:<%v>\n", reflect.TypeOf(s.PublicSetName).In(0))     //获取入参

	fmt.Printf("The reflect.TypeOf(s.PublicSetName).NumOut() is:<%v>\n", reflect.TypeOf(s.PublicSetName).NumOut()) //获取入参个数

	//指针类型变量可以获取到值类型变量对应的方法,反之不可以
	fmt.Printf("The reflect.TypeOf NumMethod is:<%v>\n", typeOfStudent.NumMethod()) //For a non-interface type, it returns the number of exported methods.
	fmt.Printf("The reflect.TypeOf Method(0) is:<%v>\n", typeOfStudent.Method(0))
	fmt.Printf("The reflect.TypeOf Method(0).IsExported() is:<%v>\n", typeOfStudent.Method(0).IsExported()) //判断方法是否是可导出的
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(s)
	fmt.Printf("The reflect.TypeOf Method(0) Call is:<%v>\n", typeOfStudent.Method(0).Func.Call(args)) //通过反射调用方法,需要注意第一个参数一定是接收者本身
	fmt.Printf("The reflect.TypeOf Method(1) is:<%v>\n", typeOfStudent.Method(1))
	args2 := make([]reflect.Value, 2)
	args2[0] = reflect.ValueOf(s)
	args2[1] = reflect.ValueOf("newName")
	fmt.Printf("The reflect.TypeOf Method(1) Call is:<%v>\n", typeOfStudent.Method(1).Func.Call(args2)) //调用方法的时候第一个参数是自身的reflect.ValueOf结果
	fmt.Printf("The reflect.TypeOf NumField() is:<%v>\n", typeOfStudent.NumField())
	fmt.Printf("The reflect.TypeOf Field(1) is:<%v>\n", typeOfStudent.Field(1))
	fmt.Printf("The reflect.TypeOf Field(1).Anonymous is:<%v>\n", typeOfStudent.Field(1).Anonymous) //是否为匿名变量
	fmt.Printf("The reflect.TypeOf Field(1).Offset is:<%v>\n", typeOfStudent.Field(1).Offset)       //字段的内存偏移
	fmt.Printf("The reflect.TypeOf Field(1).Tag is:<%v>\n", typeOfStudent.Field(1).Tag)             //获取标签
	fmt.Printf("The reflect.TypeOf Field(1).IsEx IsExported():<%v>\n", typeOfStudent.Field(1).IsExported())

	fmt.Printf("The reflect.TypeOf Comparable() is:<%v>\n", typeOfStudent.Comparable()) //类型是否可比较

	valueOfStudent := reflect.ValueOf(&s)
	fmt.Printf("The reflect.ValueOf is:<%v>\n", valueOfStudent)
	fmt.Printf("The reflect.ValueOf.Type() is:<%v>\n", valueOfStudent.Elem().Type()) //Value转Type

	fmt.Printf("The reflect.ValueOf Elem().Kind() is:<%v>\n", valueOfStudent.Elem().Kind()) //指针类型的reflect.ValueOf需要通过Elem来获取指向的数据
	fmt.Printf("The reflect.ValueOf Elem().NumField() is:<%v>\n", valueOfStudent.Elem().NumField())
	fmt.Printf("The reflect.ValueOf Elem().Interface().(*Student) is:<%v>\n", valueOfStudent.Interface().(*Student)) //reflect.Value转换为某种类型

	valueOfStudent.Elem().Field(1).Set(reflect.ValueOf(Myint(99)))

	valueOfStudent.Elem().Field(3).Addr().Elem().Field(0).SetString("Terrible") //对于内嵌结构,可以先取Addr,获取到内嵌结构体的指针,再通过Elem获取到对应的元素完成修改
	fmt.Println(s)
}

func demo2() {
	var a float64 = 3.14
	res, ok := reflect.ValueOf(a).Interface().(float64) //这里的类型转换指的是接口类型到其他类型的转换,不是强制类型转换
	if ok {
		fmt.Println(res)
	}
	reflect.ValueOf(&a).Elem().Set(reflect.ValueOf(3.1415926))
	fmt.Println(a)
}

// 如何获取接口的相关信息
func demo3() {
	fmt.Println(reflect.TypeOf((*getSetName)(nil))) //通过对nil的接口的强制转换
}
func demo4() {
	var a []int
	var b any
	fmt.Println(b == nil)                 //true
	valid := reflect.ValueOf(b).IsValid() //IsValid判断是否给接口变量赋值
	fmt.Printf("before assign Isvalid is %v\n", valid)
	if valid {
		isNil := reflect.ValueOf(b).IsNil() //判断接口变量指向的变量是否为nil(以Slice为例)
		fmt.Printf("before assign IsNil is %v\n", isNil)

	}
	b = a
	valid = reflect.ValueOf(b).IsValid()
	fmt.Printf("after assign Isvalid is %v\n", valid)
	if valid {
		isNil := reflect.ValueOf(b).IsNil()
		fmt.Printf("after assign IsNil is %v\n", isNil)

	}
	fmt.Println(b == nil) //false
	a = make([]int, 3)
	b = a
	valid = reflect.ValueOf(b).IsValid()
	fmt.Printf("after allocate  Isvalid is %v\n", valid)
	if valid {
		isNil := reflect.ValueOf(b).IsNil()
		fmt.Printf("after allocate IsNil is %v\n", isNil)
	}
}

// 通过反射创建slice
func demo5() {
	sliceType := reflect.TypeOf(([]int)(nil))
	someSlice := reflect.MakeSlice(sliceType, 3, 5)
	fmt.Println(someSlice)
	someSlice.Index(0).Set(reflect.ValueOf(12))
	fmt.Println(someSlice)
}
func demo6() {
	newVar := reflect.New(reflect.TypeOf((string)("nil")))
	fmt.Println(newVar)
}

func main() {
	demo5()
}
