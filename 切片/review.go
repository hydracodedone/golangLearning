package main

import (
	"fmt"
)

/*
1，扩容就生成新的底层数组，修改新的 slice 不会影响原来底层数组；
2，不扩容就不会生成新的数组，修改新的 slice 会修改原来的底层数组；
1.新增元素后slice的length不超出capacity，将对原底层数组的数据进行修改，如果该数组有多个切片，需要注意这种相互影响。
2.新增元素后slice的length超出capacity，将会申请新的底层数组并拷贝数据。
扩容策略如下：
1.如果新的size超出现有size的2倍，分配的大小就是新的size，如果新size是奇数还会加1（size为1是例外）。
2.否则；如果当前size小于1024，按每次2倍size分配空间，否则每次按当前size的四分之一扩容。
*/
func review1() {
	a := [3]int{1, 2, 3}
	fmt.Printf("The type of the a is %T,The content of the a is %v\n", a, a)
	b := a[:]
	c := a[:]
	fmt.Printf("The type of the b is %T,The content of the b is %v\n", b, b)
	b = append(b, 33)
	fmt.Printf("The type of the b is %T,The content of the b is %v\n", b, b)
	fmt.Printf("The type of the a is %T,The content of the a is %v\n", a, a)
	b[2] = 100
	fmt.Printf("The type of the b is %T,The content of the b is %v\n", b, b)
	fmt.Printf("The type of the a is %T,The content of the a is %v\n", a, a)
	fmt.Printf("The type of the c is %T,The content of the c is %v\n", c, c)
	c[2] = 99
	fmt.Printf("The type of the a is %T,The content of the a is %v\n", a, a)
	fmt.Printf("The type of the c is %T,The content of the c is %v\n", c, c)
}

/*
copy 函数值适用slice
copy(a,b)是把b放入a中
copy函数会返回实际拷贝的元素个数，这也取决于源和目标两者长度较短的那一个
第一个参数是要复制的目标 slice，第二个参数是源 slice，两个 slice 可以共享同一个底层数组，甚至有重叠也没有问题。
*/
func review2() {
	a := []int{1, 2, 3}
	b := []int{5, 6, 7, 8}
	some := copy(a, b)
	fmt.Printf("The a is %v,The b is %v\n", a, b)
	fmt.Println(some)
}

/*
append(x,int,int...)
*/
func review3() {
	a := [3]int{1, 2, 3}
	b := []int{}
	b = append(b, 1, 2, 3)
	b = append(b, a[:]...)
	fmt.Println(b)
}

/*
...语法糖只能用于对slice进行解包
或者表示为不定参数
在函数内部会将这些不定参数组合成一个slice
*/
func review4() {
	someFunc := func(args ...int) {
		fmt.Printf("The type of the args is %T,The content of the args is %v\n", args, args)
	}
	a := []int{1, 2, 3}
	someFunc(a...)
	someFunc(6, 7, 8, 9)
}

/*
append 返回一个slice,因此可以链式调用
*/
func review5() {
	a := []int{1, 2, 3, 4, 5}
	a = append(a[0:2], append(a[3:], 99)...)
	fmt.Println(a)
}

/*
内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作
*/
