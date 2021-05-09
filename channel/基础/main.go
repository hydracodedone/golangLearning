package main

// chan
import (
	"fmt"
	"time"
)

func demo1() {
	var myChan chan int
	myChan = make(chan int, 3)
	myChan <- 1
	close(myChan)
	for each := range myChan {
		//range 会接受关闭之前的chan的数据
		fmt.Printf("the each is [%d]\n", each)
		time.Sleep(1 * time.Second)
	}
}
func demo2() {
	var myChan chan int
	myChan = make(chan int, 3)
	myChan <- 1
	close(myChan)
	// close(myChan)//关闭一个已经关闭的channel会报错panic: close of closed channel
	// myChan <- 1 //向一个关闭的channel发送会报错panic: send on closed channel
	for {
		each, ok := <-myChan //读取一个关闭的channnel会获取定义的零值,因此正确的读取时range或者data,ok:=<-xxx 这种形式
		if ok {
			fmt.Printf("the each is [%v],the ok is [%v]\n", each, ok)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Printf("the each is [%v],the ok is [%v],break\n", each, ok)
			break
		}
	}
	fmt.Printf("the recv is [%v]\n", <-myChan)
}
func demo3() {
	//定义只读channel
	var myChan <-chan int
	myChan = make(<-chan int, 3)
	fmt.Printf("the myChan is [%#v]\n", myChan)
	var myChan2 chan<- int
	myChan2 = make(chan<- int, 3)
	fmt.Printf("the myChan is [%#v]\n", myChan2)
}
func demo4() {
	var myChan chan int
	myChan = make(chan int, 1)
	myChan <- 1
	close(myChan)
	recv, ok := <-myChan
	fmt.Printf("recv is [%#v],ok is[%v]\n", recv, ok)
	recv, ok = <-myChan
	fmt.Printf("recv is [%#v],ok is[%v]\n", recv, ok)
}
func main() {
	demo3()
}
