package main

import (
	"context"
	"net"
	"rpc_demo/Demo01/types"

	"google.golang.org/grpc"
)

// embed
type Server struct {
	types.UnimplementedRPCDemoServer
}

// rewrite
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	reqID := req.GetId()
	return &types.Response{
		Id: reqID,
	}, nil
}

//create server

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
