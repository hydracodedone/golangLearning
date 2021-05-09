package main

import (
	"fmt"

	"factoryAutoRegisterDemo/factory"
)

func main() {
	res := factory.GlobalFactory.Create("people")
	res2 := res()("hydra", 23)
	fmt.Printf("The res2 is %#v\n", res2)
}
