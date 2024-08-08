package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_stream/types"
	"net"

	"google.golang.org/grpc"
)

/*
ctx context.Context：请求上下文
req interface{}：RPC 方法的请求参数
info *UnaryServerInfo：RPC 方法的所有信息
handler UnaryHandler：RPC 方法真正执行的逻辑
*/
func CustomUnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// pre-processing
	fmt.Printf("Before RPC Interceptor RPC Req %v, Info %v Handler %v\n", req, info, handler)
	resp, err := handler(ctx, req)
	// post-processing
	fmt.Printf("Before RPC Interceptor RPC Req %v, Info %v Handler %v\n", req, info, handler)
	return resp, err
}

type Server struct {
	types.UnimplementedRPCDemoServer
}

func (s *Server) Unary(ctx context.Context, request *types.Request) (*types.Response, error) {
	requestMessage := request.GetRequestMessage()
	return &types.Response{ResponseMessage: fmt.Sprintf("hello,world %s\n", requestMessage)}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer(grpc.ChainUnaryInterceptor(CustomUnaryInterceptor))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
