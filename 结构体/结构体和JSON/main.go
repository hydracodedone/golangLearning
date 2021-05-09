package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type info struct {
	InfoContent string
}
type AnotherInfo struct {
	AnotherInfoContent string
}
type people struct {
	name        string
	Age         int `json:"people_age"`
	Information info
	AnotherInfo
}

func main() {
	var p people = people{
		name: "Hydra",
		Age:  23,
		Information: info{
			InfoContent: "This is a test",
		},
		AnotherInfo: AnotherInfo{
			AnotherInfoContent: "This also is a test",
		},
	}
	res, err := json.Marshal(&p)
	if err != nil {
		log.Fatal("Marshal Fail")
	} else {
		fmt.Printf("The Result is %s\n", res)
	}
	unmarshalValue := []byte("{\"people_age\":25,\"name\":\"Hydra\",\"Information\":{\"InforContent\":\"This is a test\"}}")
	unmarshalStruct := new(people)
	err = json.Unmarshal(unmarshalValue, unmarshalStruct)
	if err != nil {
		log.Fatalf("Json Unmarshal Fail:<%s>", err.Error())
	}
	//反序列化的时候知会反序列化struct中指定的字段
	fmt.Printf("The unmarshal struct is %+v\n", unmarshalStruct)
}
