package main

import (
	"log"
	"net"
	"rpc_demo/Demo3/service"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterStudentInfoQueryServiceServer(rpcServer, new(service.StudentInfoQueryService))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
}
