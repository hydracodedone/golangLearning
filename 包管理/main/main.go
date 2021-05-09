package main

import (
	"fmt"

	"demo_for_package_management/PK1"

	pk2demo "demo_for_package_management/PK2"
	pk21demo "demo_for_package_management/PK2/PK2_1"
	anotherDemo "demo_for_package_management/PK2/PK2_1/PK2_1_1"
	pk3 "demo_for_package_management/PK3"
)

var a int = 1

func init() {
	fmt.Println("init1")
	fmt.Printf("The value of the a is %d\n", a)
	a++
}

func init() {
	fmt.Println("init2")
	fmt.Printf("The value of the a is %d\n", a)
	a++

}

func main() {
	PK1.Pk1DemoFunction()
	pk2demo.Pk2DemoFunction()
	pk21demo.Pk21DemoFunction()
	anotherDemo.Pk211DemoFunction()
	pk3.Pk3DemoFunction()
}
