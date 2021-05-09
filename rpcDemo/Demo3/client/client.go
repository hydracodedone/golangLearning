package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"hydracode.com/rpcDemo/Demo3/service"
)

func main() {
	grpcClient, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("grpc close fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewStudentInfoQueryServiceClient(grpcClient)
	productResponseInstance, err := productClient.QueryStudentInfo(
		context.Background(),
		&service.StudentQueryId{
			Id: 10,
		},
	)
	if err != nil {
		log.Fatalf("grpc GetProductStock fail:<%s>", err.Error())
	}
	fmt.Printf("The productResponse is %+v\n", productResponseInstance)
}
