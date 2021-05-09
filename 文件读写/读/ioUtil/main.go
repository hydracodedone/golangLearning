package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/home/hydra/Project/GolangLearning/src/文件读写/读/ioUtil/main.go"
	//ioutil.ReadFile is deprecated: As of Go 1.16, this function simply calls os.ReadFile.deprecated(default)
	// info, err := ioutil.ReadFile(path)
	info, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("---start---")
		fmt.Printf("%s", string(info))
		fmt.Println("---end---")
	}
}
