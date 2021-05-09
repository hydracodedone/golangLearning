package main

import "fmt"

func demo1() {
	name := "hydra"
	age := 18
	finalString := fmt.Sprintf("name is %s,age is %d\n", name, age)
	fmt.Printf("The actual string is %s\n", finalString)
}

func demo2() {
	defer func() {
		recoverErr := recover()
		if recoverErr != nil {
			fmt.Printf("the err is %#v\n", recoverErr)
		}
	}()
	err := fmt.Errorf("this is a error demo")
	panic(err)
}

func main() {
	demo2()
}
