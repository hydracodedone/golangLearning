package main

import (
	"log"
	"net"
	"rpc_demo/Demo2_2/service"

	"google.golang.org/grpc"
)

func tcpServer() {
	rpcServer := grpc.NewServer()
	service.RegisterQueryServiceServer(rpcServer, new(service.QueryService))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("rpc tcp listen fail:<%s>", err.Error())
	}
}

func main() {
	tcpServer()
}
