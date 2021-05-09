package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	buf := new(bytes.Buffer)
	//把指针丢进去
	enc := gob.NewEncoder(buf)
	g := []string{
		"Hello",
		"World",
	}

	//调用Encode进行序列化
	if err := enc.Encode(g); err != nil {
		fmt.Println(err)
	}
	fmt.Println(buf)
}
