package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func demo1() {
	path := "/home/hydra/Project/GolangLearning/src/GolangLearn/文件读写/main.go"
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Close File Fail:<%s>", err.Error())
		}
	}(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.Write([]byte("package main\n"))
	if err != nil {
		log.Fatalf("Write File Fail:<%s>", err.Error())
	}
	_, err = file.WriteString(
		"import \"fmt\"\nfunc main(){\n\tfmt.Println(\"Hello,World\")\n}",
	)
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
Write方法首先会判断写入的数据长度是否大于设置的缓冲长度，
如果小于，则会将数据copy到缓冲中；但不会执行flush
当数据长度大于缓冲长度时，如果数据特别大，则会跳过copy环节，直接写入文件。

*/
func demo2() {
	path := "/home/hydra/Project/GolangLearning/src/GolangLearn/文件读写/main.go"
	file, err := os.OpenFile(path, os.O_WRONLY, 0755)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Close File Fail:<%s>", err.Error())
		}
	}(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bufio.NewWriter(file)
	info := "package main\nimport \"fmt\"\nfunc main(){\n\tfmt.Println(\"Hello,World\")\n}"
	_, err = buf.WriteString(info)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = buf.Flush(); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	demo2()
}
