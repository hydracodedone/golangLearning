package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_stream/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
ctx：Go语言中的上下文，一般和 Goroutine 配合使用，起到超时控制的效果
method：当前调用的 RPC 方法名
req：本次请求的参数，只有在处理前阶段修改才有效
reply：本次请求响应，需要在处理后阶段才能获取到
cc：gRPC 连接信息
invoker：可以看做是当前 RPC 方法，一般在拦截器中调用 invoker 能达到调用 RPC 方法的效果，当然底层也是 gRPC 在处理。
opts：本次调用指定的 options 信息
*/
func CustomUnaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre-processing
	fmt.Printf("Before RPC Interceptor RPC Method %v, Req %v Reply %v\n", method, req, reply)
	err := invoker(ctx, method, req, reply, cc, opts...) // invoking RPC method
	// post-processing
	fmt.Printf("After RPC Interceptor RPC Method %v, Req %v Reply %v\n", method, req, reply)
	return err
}

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(CustomUnaryInterceptor))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	gResp, err := gClient.Unary(context.Background(), &types.Request{
		RequestMessage: "Hydra",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(gResp.GetResponseMessage())
}
