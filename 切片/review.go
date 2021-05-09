package main

import (
	"fmt"
	"sync"
	"time"
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
func review6() {
	var a [3]int = [3]int{1, 2, 3}
	b := a[0:0]
	fmt.Printf("the b is %v\n", b)
	fmt.Printf("the b is nil is %t\n", b == nil)
	var c []int = []int{}
	fmt.Printf("the c is %v\n", c)
	fmt.Printf("the c is nil is %t\n", c == nil)
	var d []int
	fmt.Printf("the d is %v\n", d)
	fmt.Printf("the d is nil is %t\n", d == nil)
	//仅声明切片的默认值是nil,打印输出却是[]
	//空切片已经分配了内存,不是nil
	//切片是动态结构，只能与 nil 判定相等，不能互相判定相等
}

func review7() {
	var a []int = []int{1, 2, 3}
	a = append([]int{5, 6, 7}, append(a, []int{1, 2, 3}...)...)
	fmt.Printf("the a is %v\n", a)
}

func reverse(target []int) []int {
	length := len(target)
	for i := 0; i <= length/2; i++ {
		target[i], target[length-i-1] = target[length-i-1], target[i]
	}
	return target
}

func review8() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("the reverse result is %v\n", reverse(a))
	fmt.Printf("the reverse result is %v\n", reverse(b))
}

func review9() {
	var a []int = []int{1, 2, 3, 4, 5}
	var b []int = []int{11, 22, 33}
	copy(a, b)
	fmt.Printf("the result is %v\n", a)
	var c []int = []int{1, 2, 3, 4, 5}
	var d []int = []int{11, 22, 33}
	copy(d, c)
	fmt.Printf("the result is %v\n", d)
}

func review10() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b []int = a[1:3]
	fmt.Printf("The b is %v,the cap of the b is %d,the len of the b is %d\n", b, cap(b), len(b))
	var c []int = a[5:]
	fmt.Printf("The c is %v,the cap of the c is %d,the len of the c is %d\n", c, cap(c), len(c))
	var d []int = b[1:]
	fmt.Printf("The d is %v,the cap of the d is %d,the len of the d is %d\n", d, cap(d), len(d))
	var e []int = []int{1, 2, 3}
	fmt.Printf("The e is %v,the cap of the e is %d,the len of the e is %d\n", e, cap(e), len(e))
}
func review11() {
	var e []int = []int{1, 2, 3}
	fmt.Printf("The e is %v,the cap of the e is %d,the len of the e is %d\n", e, cap(e), len(e))
	f := e
	f = append(f, 44, 55)
	fmt.Printf("The e is %v,the cap of the e is %d,the len of the e is %d\n", e, cap(e), len(e))
	fmt.Printf("The f is %v,the cap of the f is %d,the len of the f is %d\n", f, cap(f), len(f))
}
func review12() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b []int = a[1:3]
	b[1] = 100
	fmt.Printf("The a is %v\n", a)
	fmt.Printf("The b is %v\n", b)
	/*
		如果多个切片来源于同一个数组,则修改其中一个,在未进行扩容之前对其修改会影响到其他的切片,所以尽量减少slice对原有的array的依赖
	*/
}

func review13() {
	var a [10]int
	b := a[6:9]
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the a is %v\n", b, b, cap(b), a)
	b[0] = 1
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the a is %v\n", b, b, cap(b), a)
	b = append(b, 2)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the a is %v\n", b, b, cap(b), a)
	b = append(b, 3)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the a is %v\n", b, b, cap(b), a)
	b[0] = 100
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the a is %v\n", b, b, cap(b), a)
}

func review14() {
	var wg sync.WaitGroup

	var b []int = make([]int, 5)
	getSliceInfoFunc := func(b []int) {
		defer wg.Done()
		fmt.Printf("g The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
		time.Sleep(time.Second * 3)
		fmt.Printf("g The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
	}
	wg.Add(1)
	go getSliceInfoFunc(b)
	b = append(b, 1)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
	wg.Wait()
}
func review15() {
	var a [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var b []int = a[1:2]
	getSliceInfoFunc := func(b []int) {
		fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
		b = append(b, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1)
		fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
	}
	getSliceInfoFunc(b)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
	fmt.Printf("The a is %v\n", a)
}
func review16() {
	var b []int
	getSliceInfoFunc := func(b []int) {
		fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
		b = make([]int, 3, 4)
		fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
	}
	getSliceInfoFunc(b)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
}

func review17() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := append(a[0:3], a[5:]...)
	fmt.Printf("The b is %v,the address of the b is %p,the cap b is %d,the len b is %d\n", b, b, cap(b), len(b))
}
func review18() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for range a { //此处的a是拷贝
		//访问的是外部的a变量
		a = append(a, 1)
	}
	fmt.Println(a)
}
func review19() {
	a := [3]int{1, 2, 3} // 数组 (深拷贝)
	fmt.Printf("the address of the a is %p\n", &a)
	for k, v := range a { //k,v实际上是遍历的一个拷贝的数组,因此k,v不会发生变化
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Printf("the address of the a is %p\n", &a)
			fmt.Println(a)
		}
		a[k] = 100 + v
	}
	fmt.Printf("the address of the a is %p\n", &a)
	fmt.Println(a)
}

func review20() {
	var a [2][]int = [2][]int{{1, 2, 3}, {4, 5, 6}}
	b := a
	b[0] = []int{1, 1, 1}
	fmt.Println(a)
	fmt.Println(b)
	b[1][1] = 111
	fmt.Println(a)
	fmt.Println(b)
}
