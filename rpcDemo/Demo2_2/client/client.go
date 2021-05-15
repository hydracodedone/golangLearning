package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"hydracode.com/rpcDemo/Demo2_2/service"
)

func tcpClientWithServerStream() {
	grpcClient, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Grpc Dial Fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("Grpc Close Fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewQueryServiceClient(grpcClient)
	streamQueryServiceQueryClientInstance, err := productClient.QueryWithServerStream(
		context.Background(),
		&service.RequestIdList{
			RequestIdList: []*service.RequestId{
				{
					RequestId: 1,
				},
				{
					RequestId: 2,
				},
				{
					RequestId: 3,
				},
				{
					RequestId: 4,
				},
				{
					RequestId: 5,
				},
				{
					RequestId: 6,
				},
				{
					RequestId: 7,
				},
				{
					RequestId: 8,
				},
				{
					RequestId: 9,
				},
				{
					RequestId: 10,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("Grpc Query Fail:<%s>", err.Error())
	}
	for {
		responseInstance, err := streamQueryServiceQueryClientInstance.Recv()
		// We Can Use Goroutine To Handle The Receive Information
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Stream Response Fail:<%s>", err.Error())
			}
		} else {
			fmt.Printf("The Stream Response is %v\n", responseInstance.ResponseIdList)
		}
	}
}

func tcpClientWithClientStream() {
	grpcClient, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Grpc Dial Fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("Grpc Close Fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewQueryServiceClient(grpcClient)
	streamQueryServiceQueryClientInstance, err := productClient.QueryWithClientStream(
		context.TODO(),
	)
	if err != nil {
		log.Fatalf("Get Stream Client Fail:<%s>", err.Error())
	}
	requestInstance := service.RequestIdList{
		RequestIdList: []*service.RequestId{},
	}
	for i := 0; i < 100; i++ {
		requestInstance.RequestIdList = append(requestInstance.RequestIdList, &service.RequestId{
			RequestId: int32(i),
		})
		if i > 0 && i%5 == 0 {
			err := streamQueryServiceQueryClientInstance.Send(&requestInstance)
			if err != nil {
				log.Fatalf("Stream Client Request Fail:<%s>", err.Error())
			}
			requestInstance.RequestIdList = []*service.RequestId{}
		}
		//final send
		if len(requestInstance.RequestIdList) > 0 {
			err := streamQueryServiceQueryClientInstance.Send(&requestInstance)
			if err != nil {
				log.Fatalf("Stream Client Request Fail:<%s>", err.Error())
			}
		}
	}
	responseInstance, err := streamQueryServiceQueryClientInstance.CloseAndRecv()
	if err != nil {
		log.Fatalf("Stream Close And Receive Fail:<%s>", err.Error())
	}
	fmt.Printf("The Response is %v\n", responseInstance.ResponseIdList)
}

func main() {
	tcpClientWithClientStream()
}
