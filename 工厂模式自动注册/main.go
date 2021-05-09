package main

import (
	"fmt"

	"./factory"

	_ "./package1"

	_ "./package2"
)

func main() {
	res := factory.GolabalFactory.Create("people")
	fmt.Printf("The res is %#v\n", res)
	res2 := res()("hydra", 23)
	fmt.Printf("The res2 is %#v\n", res2)

}
