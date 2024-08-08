package main

import (
	"context"
	"grpc_server_client_demo_with_tls/types"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	cred, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer(grpc.Creds(cred))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
