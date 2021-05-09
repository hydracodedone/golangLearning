//golang 实现循环队列
//golang 实现LRU算法

package main

import (
	"errors"
	"fmt"
)

type LoopQueue struct {
	lenght int
	data   []interface{}
	flag   int
	head   int
	rear   int
}

func Init(length int) *LoopQueue {
	return &LoopQueue{
		lenght: length + 1,
		data:   make([]interface{}, length+1),
		flag:   0,
		head:   0,
		rear:   0,
	}
}

func (l *LoopQueue) Push(data interface{}) error {
	if l.Full() {
		return errors.New("full")
	} else {
		l.rear = (l.rear + 1) % l.lenght
		l.data[l.rear] = data
		return nil
	}
}

func (l *LoopQueue) Pop() (interface{}, error) {
	if l.Empty() {
		return nil, errors.New("empty")
	} else {
		l.head = (l.head + 1) % l.lenght
		data := l.data[l.head]
		return data, nil
	}
}
func (l *LoopQueue) Full() bool {
	if (l.rear+1)%l.lenght == l.head {
		return true
	} else {
		return false
	}
}
func (l *LoopQueue) Empty() bool {
	if l.rear == l.head {
		return true
	} else {
		return false
	}
}
func demo() {
	a := Init(3)
	fmt.Println(a.Push(1))
	fmt.Println(a.Push(2))
	fmt.Println(a.Push(3))
	fmt.Println(a.Push(4))
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
}
func main() {
	var a [3]int = [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0] = 111
			a[1] = 222
		}
		fmt.Println(k, a[k], v)
		a[k] += v
	}
	fmt.Println(a)

	var b []int = []int{1, 2, 3}
	for k, v := range b {
		if k == 0 {
			b[0] = 111
			b[1] = 222
		}
		fmt.Println(k, b[k], v)
		b[k] += v
	}
	fmt.Println(b)
}
