package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_tls/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//Cred
	cred, err := credentials.NewClientTLSFromFile("ca.crt", "www.hydracode.com")
	if err != nil {
		panic(err)
	}
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(cred))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	gResp, err := gClient.DemoRequest(context.Background(), &types.Request{
		Id: int32(1),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
