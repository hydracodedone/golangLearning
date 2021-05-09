package main

//如果是下列的包导入,则要求GOlang学习处于$GOPATH/src下面
import (
	pk1_demo "Golang学习/包管理/PK1"
	pk2_demo "Golang学习/包管理/PK2"

	pk2_1demo "Golang学习/包管理/PK2/PK2_1"

	pk2_1_1demo "Golang学习/包管理/PK2/PK2_1/PK2_1_1"
)

func main() {
	pk1_demo.Pk1DemoFunction()
	pk2_demo.Pk2DemoFunction()
	pk2_1demo.Pk2_1DemoFunction()
	pk2_1_1demo.Pk2_1_1DemoFunction()
}
