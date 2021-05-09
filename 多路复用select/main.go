package main

import (
	"fmt"
	"time"
)

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
