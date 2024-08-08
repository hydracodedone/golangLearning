package main

import (
	"context"
	"fmt"
	"grpc_name_resolver/name_resolver"
	"grpc_name_resolver/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func init() {
	resolver.Register(&name_resolver.CustomResolverBuilder{})
}

func main() {
	//Dial
	target := fmt.Sprintf("%s:///%s", name_resolver.GrpcServiceScheme, name_resolver.GrpcServiceName)
	gCon, err := grpc.DialContext(context.Background(), target, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
