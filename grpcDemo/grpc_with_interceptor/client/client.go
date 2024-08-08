package main

import (
	"context"
	"fmt"
	"grpc_with_interceptor/types"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// unaryInterceptor 一个简单的 unary interceptor 示例。
func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre-processing
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...) // invoking RPC method
	// post-processing
	end := time.Now()
	fmt.Printf("RPC: %s, req:%v start time: %s, end time: %s, err: %v", method, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return err
}

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(unaryInterceptor))
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
