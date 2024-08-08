package main

import (
	"grpc_server_client_demo_with_stream/types"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	types.UnimplementedRPCDemoServer
}

func (s *Server) ResponseStreaming(request *types.Request, responseStreamServer types.RPCDemo_ResponseStreamingServer) error {
	// 1.for循环接收客户端发送的消息
	err := responseStreamServer.Send(&types.Response{ResponseMessage: "hello"})
	if err != nil {
		panic(err)
	}
	err = responseStreamServer.Send(&types.Response{ResponseMessage: "world"})
	if err != nil {
		panic(err)
	}
	err = responseStreamServer.Send(&types.Response{ResponseMessage: request.GetRequestMessage()})
	if err != nil {
		panic(err)
	}
	return nil
}
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer()
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
