package main

import (
	"fmt"
)

func demo() {
	var myChan chan func(string) string = make(chan func(string) string, 3)
	chan1 := func(in string) string {
		temp := []rune(in)
		temp[0] = 'A'
		return string(temp)
	}
	chan2 := func(in string) string {
		temp := []rune(in)
		temp[1] = 'B'
		return string(temp)
	}
	myChan <- chan1
	myChan <- chan2
	close(myChan)
	temp := "EDF"
	for eachFunction := range myChan {
		temp = eachFunction(temp)
	}
	fmt.Println(temp)
}

func main() {
	var myChan chan int
	go func() {
		for i := 0; i <= 30; i++ {
			fmt.Printf("num %v\n", i)
		}
	}()
	myChan <- 1
}
