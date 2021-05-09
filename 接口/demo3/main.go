package main

import (
	"fmt"
)

type getValue interface {
	getValueFunction() int
}
type myInt struct {
	value int
}

type myFloat struct {
	value float32
}

func (temp myInt) getValueFunction() int {
	return temp.value
}
func (temp myFloat) getValueFunction() int {
	return int(temp.value)
}

func wrong1() {
	var intTemp myInt
	var floatTemp myFloat
	var temp []interface{}
	temp = append(temp, intTemp)
	temp = append(temp, floatTemp)

	for _, value := range temp {
		fmt.Printf("The Type of the value is %T\n", value)
		//value.getValue()
	}
}
func wrong2() {
	var intTemp myInt
	var floatTemp myFloat
	var temp []interface{}
	temp = append(temp, intTemp)
	temp = append(temp, floatTemp)

	for _, value := range temp {
		fmt.Printf("The Type of the value is %T\n", value)
		// var result getValue
		// result = value
	}
}

// func wrong3() {
// 	var intTemp myInt
// 	var floatTemp myFloat
// 	var temp []interface{}
// 	temp = append(temp, intTemp)
// 	temp = append(temp, floatTemp)

// 	for _, value := range temp {
// 		fmt.Printf("The Type of the value is %T\n", value)
// 		switch value.(type) {
// 		case myInt:
// 			fmt.Printf("The type is %T,The value is %d\n", value, myInt(value).getValue())
// 		case myFloat:
// 			fmt.Printf("The type is %T,The value is %d\n", value, myFloat(value).getValue())
// 		}
// 	}
// }

// func wrong4() {
// 	var intTemp myInt
// 	var floatTemp myFloat
// 	var temp []interface{}
// 	temp = append(temp, intTemp)
// 	temp = append(temp, floatTemp)

// 	for _, value := range temp {
// 		fmt.Printf("The Type of the value is %T\n", value)
// 		switch value.(type) {
// 		case myInt:
// 			fmt.Printf("The type is %T,The value is %d\n", value, reflect.ValueOf(value).Type().getValue())
// 		case myFloat:
// 			fmt.Printf("The type is %T,The value is %d\n", value, myFloat(value).getValue())
// 		}
// 	}
// }
func main() {
	// wrong3()
}
