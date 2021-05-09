package main

import (
	"fmt"
	"os"
)

func demo() {
	a := 24
	b := "hello,world"
	res := fmt.Sprintf("%s Hydra,your age is %#v\n", b, a)
	fmt.Println(res)
}

func demo2() {
	file, err := os.Create("./Test.txt")
	if err != nil {
		return
	}
	fmt.Fprintf(file, "Hello,World %s", "hydra")
	file.Close()
}
func main() {
	demo()
	demo2()
}
