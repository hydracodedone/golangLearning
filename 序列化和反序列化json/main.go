package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name    string `json:"name"`
	age     int    //私有变量无法打tag
	address struct {
		home string
	}
}

func structLoadToJSON() {
	var stu1 student = student{
		"Hydra",
		23,
		struct{ home string }{
			"CHINA",
		},
	}
	var stu2 student
	fmt.Printf("The stu1 is %#v\n", stu1)
	var jsonStr []byte
	jsonStr, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("Marshal Fail")
	} else {
		fmt.Printf("stu1 after marshaler is %s\n", jsonStr)
	}
	err = json.Unmarshal(jsonStr, &stu2)
	if err != nil {
		fmt.Println("Unmarshal Fail")
	} else {
		fmt.Printf("jsonstr after unmarshaler is %#v\n", stu2)

	}
}
func main() {
	structLoadToJSON()
}
