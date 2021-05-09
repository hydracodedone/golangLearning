package main

import "fmt"

func demo1() {
	/*
		对slice进行append,元素个数超过原有的slice的cap后,会重新生成一个slice
	*/
	var a []int
	fmt.Printf("The type of the a is %T\n", a)
	a = make([]int, 3, 5)
	fmt.Printf("The address of the a is %p, a is %v,cap is %d,len is %d\n", a, a, cap(a), len(a))
	a = append(a, 2, 3)
	fmt.Printf("The address of the a is %p, a is %v,cap is %d,len is %d\n", a, a, cap(a), len(a))
	a = append(a, 2, 3)
	fmt.Printf("The address of the a is %p, a is %v,cap is %d,len is %d\n", a, a, cap(a), len(a))
	/*
		切片的另一种申明方式
	*/
	var b []int = []int{0: 1, 2: 3}
	fmt.Printf("The address of the b is %p, b is %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	/*
		切片和数组之间的关系
		切片可以直接申明,则对应的底层数组不可见
		切片也可以由数组切片实现,是对数组的引用,但此时的切片一旦进行append操作后,则有可能失去对原有的数组的引用
		1 如果对数组切片得到的切片进行append操作时候,其len小于cap,则append的元素相当于对原数组的元素进行修改
		2 如果对数组切片得到的切片进行append操作时候,其len不小于cap,则append的元素不会对元素组的元素进行修改
		3 通过数组slice得到的切片的容量是切片第一个元素对应数组中的位置到数组最后一个元素的之间的元素的个数
	*/
	var c [3]int = [3]int{1, 2, 3}
	var d []int = c[1:2]
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
	fmt.Printf("The array c is %v\n", c)
	d[0] = 100
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
	fmt.Printf("The array c is %v\n", c)
	d = append(d, 1111)
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
	fmt.Printf("The array c is %v\n", c)
	d = append(d, 2222)
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
	fmt.Printf("The array c is %v\n", c)
	d = append(d, 3333)
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
	fmt.Printf("The array c is %v\n", c)
	/*
		切片的切片
	*/
	var f [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("The array f is %v\n", f)
	var g []int = f[1:4]
	fmt.Printf("The address of the g is %p, g is %v,cap is %d,len is %d\n", g, g, cap(g), len(g))
	var h []int = f[2:3]
	fmt.Printf("The address of the h is %p, h is %v,cap is %d,len is %d\n", h, h, cap(h), len(h))
	h = append(h, 100)
	fmt.Printf("The array f is %v\n", f)
	fmt.Printf("The address of the g is %p, g is %v,cap is %d,len is %d\n", g, g, cap(g), len(g))
	fmt.Printf("The address of the h is %p, h is %v,cap is %d,len is %d\n", h, h, cap(h), len(h))
	/*
		nil 与 长度为0的切片之间的关系
		空切片与nil切片
	*/
	var l []int
	fmt.Printf("l is nil is %v\n", l == nil)
	fmt.Printf("The address of the l is %p, l is %v,cap is %d,len is %d\n", l, l, cap(l), len(l))
	l = make([]int, 0)
	fmt.Printf("l is nil is %v\n", l == nil)
	fmt.Printf("The address of the l is %p, l is %v,cap is %d,len is %d\n", l, l, cap(l), len(l))
}

func demo2() {
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("The a is %v\n", a)
	var b []int = a[0:3]
	fmt.Printf("The b is %v\n", b)
	b = append(b, 111)
	fmt.Printf("The b is %v\n", b)
	fmt.Printf("The a is %v\n", a)
}
func demo3() {
	/*
		难点
		append 对底层数组的影响
	*/
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("The a is %v\n", a)
	var b []int = a[0:5]
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	b = append(b[:1], b[3:]...)
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	fmt.Printf("The a is %v\n", a)

}
func demo4() {
	/*
		难点
		append 对底层数组的影响
	*/
	var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("The a is %v\n", a)
	var b []int = a[0:5]
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	b = append(b[:3], 111, 111, 111)
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	fmt.Printf("The a is %v\n", a)
	b = append(b[:3], 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111, 111)
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", b, b, cap(b), len(b))
	fmt.Printf("The a is %v\n", a)
}

func demo5() {
	/*
		map无法比较
	*/
}

func main() {
	demo4()
}
