package main

import (
	demo "demo_for_struct_factory/private"
	"fmt"
)

func main() {
	res := demo.FactoryForStruct("hydra", 23)
	fmt.Printf("The res is %T\n", res)
	res2 := demo.PublicatStruct{Name: "Hydra"}
	res2.Age = 1
	fmt.Printf("The res is %v\n", res)

}
