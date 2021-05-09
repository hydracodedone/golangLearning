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
				&service.RequestId{
					RequestId: 1,
				},
				&service.RequestId{
					RequestId: 2,
				},
				&service.RequestId{
					RequestId: 3,
				},
				&service.RequestId{
					RequestId: 4,
				},
				&service.RequestId{
					RequestId: 5,
				},
				&service.RequestId{
					RequestId: 6,
				},
				&service.RequestId{
					RequestId: 7,
				},
				&service.RequestId{
					RequestId: 8,
				},
				&service.RequestId{
					RequestId: 9,
				},
				&service.RequestId{
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
		if err != nil && err != io.EOF {
			log.Fatalf("Stream Response Fail:<%s>", err.Error())
		} else if err == nil {
			fmt.Printf("The Stream Response is %v\n", responseInstance.ResponseIdList)
		} else {
			break
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
		context.Background(),
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
			err := streamQueryServiceQueryClientInstance.SendMsg(&requestInstance)
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
