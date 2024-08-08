package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"grpc_gateway/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	types.UnimplementedRPCDemoServer
}

func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	reqID := req.GetId()
	fmt.Println(reqID)
	return &types.Response{
		Id: reqID,
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer()
	types.RegisterRPCDemoServer(gServer, &Server{})

	go func() {
		err := gServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()

	gwmux := runtime.NewServeMux()
	err = types.RegisterRPCDemoHandler(context.Background(), gwmux, gCon)
	if err != nil {
		panic(err)
	}
	
	gwServer := &http.Server{
		Addr:    "localhost:9001",
		Handler: gwmux,
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
