package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_stream/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
1. 建立连接并获取client
2. 获取 stream 并通过 Send 方法不断推送数据到服务端
3. 发送完成后通过stream.CloseAndRecv() 关闭steam并接收服务端返回结果
*/

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	requestStreamClient, err := gClient.ResquestStreaming(context.Background())
	if err != nil {
		panic(err)
	}
	err = requestStreamClient.Send(&types.Request{RequestMessage: "Hydra1"})
	if err != nil {
		panic(err)
	}
	err = requestStreamClient.Send(&types.Request{RequestMessage: "Hydra2"})
	if err != nil {
		panic(err)
	}
	resp, err := requestStreamClient.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.GetResponseMessage())
}
