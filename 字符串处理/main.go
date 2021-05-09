package main

import (
	"fmt"
	"log"
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
		log.Fatalf("Create File Fail:<%s>", err.Error())
		return
	}
	_, err = fmt.Fprintf(file, "Hello,World %s", "hydra")
	if err != nil {
		log.Fatalf("Fprintf Fail:<%s>", err.Error())
		return
	}
	err = file.Close()
	if err != nil {
		log.Fatalf("Close File Fail:<%s>", err.Error())
		return
	}
}
func main() {
	demo()
	demo2()
}
