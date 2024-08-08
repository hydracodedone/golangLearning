package main

import (
	"fmt"
	"grpc_server_client_demo_with_stream/types"
	"io"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
}

func newWrappedServerStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedServerStream{s}
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	fmt.Println("S RecvMsg0:", m)
	err := w.ServerStream.RecvMsg(m)
	fmt.Println("S RecvMsg1:", m)
	return err
}

func (w *wrappedServerStream) SendMsg(m interface{}) error {
	fmt.Println("S SendMsg:", m)
	return w.ServerStream.SendMsg(m)
}

func ServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// 包装 grpc.ServerStream 以替换 RecvMsg SendMsg这两个方法。
	err := handler(srv, newWrappedServerStream(ss))
	if err != nil {
		panic(err)
	}
	return err
}

/*
// 1. 开两个goroutine（使用 chan 传递数据） 分别用于Recv()和Send()
// 1.1 一直Recv()到err==io.EOF(即客户端关闭stream)
// 1.2 Send()则自己控制什么时候Close 服务端stream没有close 只要跳出循环就算close了。 具体见https://github.com/grpc/grpc-go/issues/444
*/
type Server struct {
	types.UnimplementedRPCDemoServer
}

func (s *Server) ResquestStreamingAndResponseStreaming(stream types.RPCDemo_ResquestStreamingAndResponseStreamingServer) error {
	var (
		waitGroup sync.WaitGroup
		msgCh     = make(chan string)
	)
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		number := 0
		for v := range msgCh {
			err := stream.Send(&types.Response{ResponseMessage: fmt.Sprintf("%s-%d", v, number)})
			number++
			if err != nil {
				panic(err)
			}
		}
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			msgCh <- req.GetRequestMessage()
		}
		close(msgCh)
		fmt.Println("Recv Finished")
	}()
	waitGroup.Wait()
	// 返回nil表示已经完成响应
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer(grpc.ChainStreamInterceptor(ServerStreamInterceptor))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
