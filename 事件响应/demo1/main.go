package main

import "fmt"

var event map[string][]func([]interface{}) []interface{}

func init() {
	event = make(map[string][]func([]interface{}) []interface{}, 10)
}
func Register(name string, callback func([]interface{}) []interface{}) {
	temp := event[name]
	temp = append(temp, callback)
	event[name] = temp
}
func CallEvent(name string, param []interface{}) {
	temp := event[name]
	for _, callback := range temp {
		callback(param)
	}
}

func test(param []interface{}) []interface{} {
	params := param[0]
	fmt.Println(params)
	return nil
}

func main() {
	Register("serialization", test)
	temp := make([]interface{}, 1)
	temp[0] = "hello,world"
	CallEvent("serialization", temp)
}
