package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	path := "/home/hydra/Project/GolangLearning/src/文件读写/读/ioUtil/main.go"
	info, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---start---")
		fmt.Printf("%s", string(info))
		fmt.Println("---end---")
	}
}
