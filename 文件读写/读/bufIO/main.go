package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/home/hydra/Project/GolangLearning/src/文件读写/读/bufIO/main.go")
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("close file fail")
		}
	}()
	if err != nil {
		fmt.Println(err)
	} else {
		var infoAll []byte
		buf := bufio.NewReader(file)
		for {
			infoSlice, err := buf.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return
			} else if err == io.EOF {
				infoAll = append(infoAll, infoSlice...)
				break
			} else {
				infoAll = append(infoAll, infoSlice...)
			}
		}
		fmt.Println("---start---")
		fmt.Println(string(infoAll))
		fmt.Println("---end---")
	}
}
