package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"hydracode.com/rpcDemo/Demo2_2/service"
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
