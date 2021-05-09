package main

import (
	"fmt"

	"factoryAutoRegisterDemo/factory"

	_ "factoryAutoRegisterDemo/package1"

	_ "factoryAutoRegisterDemo/package2"
)

func main() {
	res := factory.GlobalFactory.Create("people")
	res2 := res()("hydra", 23)
	fmt.Printf("The res2 is %#v\n", res2)
}
