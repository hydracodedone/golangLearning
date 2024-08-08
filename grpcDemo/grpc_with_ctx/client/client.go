package main

import (
	"context"
	"fmt"
	"grpc_with_ctx/types"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	timeoutCtx,cacelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cacelFunc()
	gResp, err := gClient.DemoRequest(timeoutCtx, &types.Request{
		Id: int32(1),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
