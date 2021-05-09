package main

import (
	"errors"
	"fmt"
)

type info struct {
	err error
}

func procedure1ErrorGenerator() error {
	return errors.New("error happened during inner procedure1")
}
func (i *info) procedure1() {
	//每个动作在执行前都要先通过err判断下是否能够执行
	if i.err != nil {
		return
	}
	err := procedure1ErrorGenerator()
	if err != nil {
		i.err = err
	}
	fmt.Println("process in procedure1")
}
func (i *info) procedure2() {
	if i.err != nil {
		return
	}
	fmt.Println("process in procedure2")
}
func (i *info) procedure3() {
	if i.err != nil {
		return
	}
	fmt.Println("process in procedure3")
}

func demo1() {
	i := &info{}
	i.procedure1()
	i.procedure2()
	i.procedure3()
}

type chainFunc func() error

func errorHandler(err error) error {
	fmt.Printf("handle error :[%s]\n", err)
	return nil
}

type i2 struct {
}

func (i *i2) procedure1() error {
	fmt.Println("procedure1")
	return nil
}
func (i *i2) procedure2() error {
	return errors.New("error happened during procedure2")

}
func (i *i2) procedure3() error {
	fmt.Println("procedure3")
	return nil
}

type chain struct {
	length       int
	index        int
	funcs        []chainFunc
	errorHandler func(error) error
}

func InitChain(handler func(error) error) *chain {
	return &chain{
		funcs:        make([]chainFunc, 0),
		errorHandler: handler,
	}
}
func (c *chain) Register(chainFunction chainFunc) {
	c.funcs = append(c.funcs, chainFunction)
	c.length += 1
}
func (c *chain) Process() error {
	if c.funcs == nil {
		return nil
	}
	for i, f := range c.funcs {
		err := f()
		if err != nil {
			if c.errorHandler == nil {
				return err
			}
			err := c.errorHandler(err)
			if err != nil {
				return err
			}
			break
		}
		c.index = i
	}
	return nil
}
func demo2() {
	i := &i2{}
	c := InitChain(errorHandler)
	c.Register(func() error {
		return i.procedure1()
	})
	c.Register(func() error {
		return i.procedure2()
	})
	c.Register(func() error {
		return i.procedure3()
	})
	if err := c.Process(); err != nil {
		panic(err)
	}
}

func main() {
	demo2()
}
