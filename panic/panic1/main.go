package main

import (
	"fmt"
)

// panic不会被defer内层的defer处理
func demo() {
	defer func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("the err is %s\n", err)
			} else {
				fmt.Println("nothing")
			}
		}()
	}()
	fmt.Println("begin")
	panic("demo")
}

func demo1() {
	defer func() {
		fmt.Printf("defer1\n")
		//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。
		//任何未捕获的错误都会沿调用堆栈向外传递。
		err := recover()
		if err != nil {
			fmt.Printf("recover error:[%v]\n", err)
		}
	}()
	fmt.Println("begin")
	defer func() {
		//触发了panic后,在panic处之前的defer会开始执行,如果都没有捕获到异常,则会抛出异常
		fmt.Printf("defer2\n")
	}()
	func() {
		panic("panic")
	}()
	defer func() {
		fmt.Printf("defer3\n")
		err := recover()
		if err != nil {
			fmt.Printf("recover error:[%v]\n", err)
		}
	}()
	fmt.Println("middle")
	fmt.Println("end")
}
func demo2() {
	//延迟调用中引发的错误，可被后续延迟调用捕获，但仅最后一个错误可被捕获。
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recover error:[%v]\n", err)
		}
	}()
	defer func() {
		fmt.Println("begin panic1")
		panic("panic1")
	}()
	defer func() {
		fmt.Println("begin panic2")
		panic("panic2")
	}()
	func() {
		panic("main panic")
	}()
}
func demo3() {
	//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()
	panic("test panic")
}

func main() {
	demo1()
	demo2()
	demo3()
}
