package main

import "fmt"

func demo() {
	var mychan chan func(string) string = make(chan func(string) string, 3)
	chan1 := func(in string) string {
		temp := []rune(in)
		temp[0] = 'A'
		return (string(temp))
	}
	chan2 := func(in string) string {
		temp := []rune(in)
		temp[1] = 'B'
		return (string(temp))
	}
	mychan <- chan1
	mychan <- chan2
	close(mychan)
	temp := "EDF"
	for eachFunction := range mychan {
		temp = eachFunction(temp)
	}
	fmt.Println(temp)
}

func main() {
	demo()
}
