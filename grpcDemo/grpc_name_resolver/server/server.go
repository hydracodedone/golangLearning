package main

import (
	"context"
	"grpc_name_resolver/types"
	"net"

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

// create server
func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
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
