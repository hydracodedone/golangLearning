package main

import "fmt"

//fatal error: all goroutines are asleep - deadlock!
/*
原因:对于同步函数,在执行的过程中,由于申请的chan是无缓冲区的,因此当执行了myChan<-22后,由于没有其他的
groutine在此刻运行,因此,代码就阻塞在main函数
此刻就只有main函数这一个groutine,下文的groutine还没有执行到,因此所有的groutine(即main函数对应的groutine)
都处于阻塞状态,因此程序异常
*/
func main() {
	var myChan chan int = make(chan int)
	myChan <- 22
	go func() {
		fmt.Printf("The value from chan is %d\n", <-myChan)
	}()

}
