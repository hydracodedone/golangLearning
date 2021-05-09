package main

import (
	"encoding/json"
	"fmt"
)

func demo() {
	type people struct {
		Name  string
		age   int //私有变量无法打tag
		Sex   string `json:"-"`
		Grade int    `json:"grade string"`
	}
	a := people{
		Name:  "hydracode",
		age:   23,
		Sex:   "male",
		Grade: 80,
	}
	marshalledData, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("marshal data failed: %v\n", err)
	}
	stringData := string(marshalledData)
	fmt.Println(stringData)
	unMarshalData := "{\"Name\":\"hydracode\",\"age\":23,\"grade string\":80,\"Sex\":\"male\",\"demo\":\"demo\"}"
	byteSlice := []byte(unMarshalData)
	unMarshalledPeople := new(people)
	err = json.Unmarshal(byteSlice, unMarshalledPeople)
	if err != nil {
		fmt.Printf("unmarshal data failed: %v\n", err)
	}
	fmt.Printf("the unMarshalledPeople is %#v\n", unMarshalledPeople)
}
func main() {
	demo()
}
