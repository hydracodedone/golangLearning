package main

import "fmt"

/*
slice是无固定长度的数组，底层结构是一个结构体，包含如下3个属性
一个 slice 在 golang 中占用 24 个 bytes

	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

array : 包含了一个指向一个数组的指针，数据实际上存储在这个指针指向的数组上，占用 8 bytes
len: 当前 slice 使用到的长度，占用8 bytes
cap : 当前 slice 的容量，同时也是底层数组 array 的长度， 8 bytes
slice不支持并发读写，所以并不是线程安全的
多个切片如果共享同一个底层数组，这种情况下，对其中一个切片或者底层数组的更改，会影响到其他切片

如果我们只用到一个slice的一小部分，那么底层的整个数组也将继续保存在内存当中。当这个底层数组很大，或者这样的场景很多时，可能会造成内存急剧增加，造成崩溃。
空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素。
最后需要说明的一点是。不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的。
*/
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
	a = append(a, []int{1, 2, 3}...)
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
	//需要注意数组切片的cap
	var d []int = c[1:2]
	fmt.Printf("The array c is %v\n", c)
	fmt.Printf("The address of the d is %p, d is %v,cap is %d,len is %d\n", d, d, cap(d), len(d))
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
	h = append(h, 1111, 1111, 1111, 1111, 1111, 1111, 1111, 1111, 1111)
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
	first := b[:1]
	second := b[3:]
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", first, first, cap(first), len(first))
	fmt.Printf("The address of the  is %p,  %v,cap is %d,len is %d\n", second, second, cap(second), len(second))
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
		slice无法比较,slice只能和nil比较
	*/
}
