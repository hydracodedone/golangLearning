package main

import (
	"fmt"
	"reflect"
	"runtime"
)

type anyList = []interface{}
type registerCenter map[string]func(anyList) anyList

// var registerCenterInstance registerCenter = make(registerCenter, 10)

func register(registerFunc func(anyList) anyList) {
	name := runtime.FuncForPC(reflect.ValueOf(registerFunc).Pointer()).Name()
	// reflect.FuncOf()
	fmt.Println(name)
}

func testDemo(anyList) anyList {
	return nil
}
func main() {
	register(testDemo)
}
