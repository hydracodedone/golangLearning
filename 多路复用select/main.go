package main

import (
	"fmt"
	"time"
)

/*
当多个需要从多个chan中读取或写入时，会先"轮询"一遍所有的case，然后在所有处于就绪（可读/可写）的chan中随机挑选一个进行读取或写入操作，并执行其语句块。
如果所有case都未就绪，则执行default语句，如未提供default语句，则当前协程被阻塞。
轮询的前提是各个chan是获取到的,如果chan都没有获取到,也就谈不上接收数据
结论:
请尽量在返回chan的函数中减少耗时操作；
select本身就是无序的 不要依赖执行顺序
*/
func producer(chanList []chan string) {
	value := 0
	for {
		for _, each := range chanList {
			product := fmt.Sprintf("The value is %d", value)
			value++
			select {
			case each <- product:
				{
					fmt.Printf("send success <%d>\n", value)
				}
			default: //未提供default语句，则当前协程被阻塞。
				{
					fmt.Println("PRODUCER PASS")
				}
			}
			time.Sleep(time.Second)
		}
	}
}
func consumer(chanList []chan string) {
	for {
		for _, each := range chanList {
			select {
			case product := <-each:
				{
					fmt.Printf("recv success <%s>\n", product)
				}
			default:
				{
					fmt.Println("CONSUMER PASS")
				}
			}
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var chanList []chan string
	for i := 0; i < 10; i++ {
		chanList = append(chanList, make(chan string, 2))
	}
	var sig chan int
	go consumer(chanList)
	go producer(chanList)
	<-sig
}
