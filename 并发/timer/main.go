package main

import (
	"fmt"
	"time"
)

// timer
func demo() {
	timer := time.NewTimer(time.Second)
	select {
	case res := <-timer.C:
		fmt.Printf("recv is %v\n", res)
	}
}

// Replace time.After
func demo1() {
	timer := time.NewTimer(time.Second)
	for {
		select {
		case res := <-timer.C:
			fmt.Printf("recv is %v\n", res)
			timer.Reset(time.Second)
		}
	}
}

// ticker
func demo2() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case recv := <-ticker.C:
			fmt.Printf("recv is %v\n", recv)
		}
	}
}

// the right way to use single ticker
func demo3() {
	ticker := time.NewTicker(time.Second)
	count := 0
	for recv := range ticker.C {
		fmt.Printf("recv is %v\n", recv)
		count += 1
		if count == 3 {
			ticker.Reset(time.Second * 2)
		}
		if count >= 5 {
			break
		}
	}
	//Stop does not close the channelStop does not close the channel
	ticker.Stop()
}

func demo3_2() {
	myChan1 := make(chan int)
	myChan2 := make(chan int)
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		select {
		case <-myChan1:
			fmt.Println("chan1 recv")
		case <-myChan2:
			fmt.Println("chan2 recv")
		default:
			fmt.Println("default")
		}
	}
}
func demo5() {
	timeAfter := time.After(time.Second)
stop1:
	for {
		select {
		case recv := <-timeAfter:
			fmt.Printf("the recv is %v", recv)
			break stop1
		}
	}
	fmt.Println("the main groutine is ending")
}
func demo6() {
	timeAfter := time.After(time.Second)
	for recv := range timeAfter {
		fmt.Printf("the recv is %v", recv) //timer.C不会自动关闭
	}
	fmt.Println("the main groutine is ending")
}
func main() {
}
