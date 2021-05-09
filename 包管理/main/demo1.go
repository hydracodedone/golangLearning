package main

import (
	"fmt"

	"packageManagement/PK1"
	"packageManagement/PK2"
	"packageManagement/PK2/PK2_1"
	"packageManagement/PK2/PK2_1/PK2_1_1"
)

func init() {
	fmt.Printf("pk2_1demo")
}
func main() {
	PK1.Pk1DemoFunction()
	pk2_demo.Pk2DemoFunction()
	pk2_1demo.Pk21DemoFunction()
	pk2_1_1demo.Pk211DemoFunction()
}
