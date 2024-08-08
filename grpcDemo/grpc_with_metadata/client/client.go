package main

import (
	"context"
	"fmt"
	"grpc_with_metadata/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	// md := metadata.Pairs("token", "hydra")
	//key 中大写字母会被转化为小写
	// grpc- 开头的键为 grpc 内部使用，如果再 metadata 中设置这样的键可能会导致一些错误

	// md := metadata.New(map[string]string{"token": "hydra"})
	// ctx := metadata.NewOutgoingContext(context.Background(), md)
	// 有两种发送 metadata 到 server 的方法
	// 推荐的方法是使用 AppendToOutgoingContext, 如果 metadata 已存在则会合并，不存在则添加
	// 而 NewOutgoingContext 则会覆盖 context 中 已有的 metadata
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", "hydra")
	metadata.FromOutgoingContext(ctx)
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	gResp, err := gClient.DemoRequest(ctx, &types.Request{
		Id: int32(1),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
