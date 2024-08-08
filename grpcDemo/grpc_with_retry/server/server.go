package main

import (
	"context"
	"grpc_with_retry/types"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// embed
type Server struct {
	types.UnimplementedRPCDemoServer
	mu         sync.Mutex
	reqCounter uint
}

func (s *Server) failRequest() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reqCounter++
	time.Sleep(time.Second)
	if s.reqCounter%4 == 0 {
		return nil
	}
	return status.Errorf(codes.Unavailable, "FAIL")
}

// implementation interaface
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	err = s.failRequest()
	if err != nil {
		return nil, err
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
