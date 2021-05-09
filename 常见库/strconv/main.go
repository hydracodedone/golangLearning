package main

import (
	"fmt"
	"strconv"
)

func demo() {
	a := "10"
	parsedInt, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("parsed failed: %v\n", err)
	}
	fmt.Printf("the parsed int is %d\n", parsedInt)

}
func main() {
	demo()
}
