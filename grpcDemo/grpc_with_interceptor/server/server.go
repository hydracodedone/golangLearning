package main

import (
	"context"
	"fmt"
	"grpc_with_interceptor/types"
	"net"
	"time"

	"google.golang.org/grpc"
)

// embed
type Server struct {
	types.UnimplementedRPCDemoServer
}

// implementation interaface
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	reqID := req.GetId()
	return &types.Response{
		Id: reqID,
	}, nil
}

// unaryInterceptor 一个简单的 unary interceptor 示例。
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	m, err := handler(ctx, req)
	end := time.Now()
	// 记录请求参数 耗时 错误信息等数据
	fmt.Printf("RPC: %s,req:%v start time: %s, end time: %s, err: %v", info.FullMethod, req, start.Format(time.RFC3339), end.Format(time.RFC3339), err)
	return m, err
}

// create server
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
