package main

import (
	"context"
	"fmt"
	"grpc_with_metadata/types"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// embed
type Server struct {
	types.UnimplementedRPCDemoServer
}

// implementation interaface
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	md, exists := metadata.FromIncomingContext(ctx)
	if exists {
		token, tokenExists := md["token"]
		if tokenExists {
			fmt.Printf("token is %s\n", token)
		}
	}
	reqID := req.GetId()
	return &types.Response{
		Id: reqID,
	}, nil
}

// create server
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
