package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"hydracode.com/rpcDemo/Demo2/service"
)

func main() {
	productRequestInstance := service.ProductRequest{
		ProductId: []int32{1, 2},
	}
	marshalResult, err := proto.Marshal(&productRequestInstance)
	if err != nil {
		log.Fatalf("Marshal Message Fail:<%>", err.Error())
	}
	fmt.Printf("The Marshal Message is:%v\n", marshalResult)
	unmarshalResult := service.ProductRequest{}
	err = proto.Unmarshal(marshalResult, &unmarshalResult)
	if err != nil {
		log.Fatalf("Marshal Message Fail:<%s>", err.Error())
	}
	fmt.Printf("The Unmarshal Message is:%v\n", unmarshalResult)
}
