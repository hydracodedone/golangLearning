package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func readByIOutil() {
	bytesSlice, err := ioutil.ReadFile("./readFile.txt")
	if err != nil {
		panic(err)
	}
	data := string(bytesSlice)
	fmt.Printf("the read data is: \n%s\n", data)
}
func readByOS() {
	file, err := os.Open("./readFile.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var bytesSlice []byte = make([]byte, 1024)
	_, err = file.Read(bytesSlice)
	if err != nil {
		panic(err)
	}
	data := string(bytesSlice)
	fmt.Printf("the read data is: \n%s\n", data)
}
func commonOsOpen() (file *os.File, err error) {
	file, err = os.Open("./readFile.txt")
	return
}
func fileOpenAndReadByByteSlice() {
	file, err := commonOsOpen()
	if err != nil {
		panic(err)
	}
	var allReadBuf []byte = make([]byte, 1024)
	var eachReadBuf []byte = make([]byte, 2)
	for {
		readSize, err := file.Read(eachReadBuf)
		if readSize == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		allReadBuf = append(allReadBuf, eachReadBuf[0:readSize]...)
	}
	data := string(allReadBuf)
	fmt.Printf("the read data is: \n%s\n", data)
}
func fileOpenAndReadByReader() {
	file, err := commonOsOpen()
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	var allReadBuf []byte = make([]byte, 1024)
	var eachReadBuf []byte = make([]byte, 3)
	for {
		readSize, err := reader.Read(eachReadBuf)
		fmt.Println(readSize)
		if readSize == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		allReadBuf = append(allReadBuf, eachReadBuf[0:readSize]...)
	}
	data := string(allReadBuf)
	fmt.Printf("the read data is: \n%s\n", data)
}

func main() {
	// readByIOutil()
	// readByOS()
	// fileOpenAndReadByReader()
}
