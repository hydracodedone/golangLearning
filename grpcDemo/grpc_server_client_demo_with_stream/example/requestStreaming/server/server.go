package main

import (
	"fmt"
	"grpc_server_client_demo_with_stream/types"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

/*
1. for循环中通过stream.Recv()不断接收client传来的数据
2. err == io.EOF表示客户端已经发送完毕关闭连接了,此时在等待服务端处理完并返回消息
3. stream.SendAndClose() 发送消息并关闭连接(虽然在客户端流里服务器这边并不需要关闭 但是方法还是叫的这个名字，内部也只会调用Send())
*/
type Server struct {
	types.UnimplementedRPCDemoServer
}

func (s *Server) ResquestStreaming(stream types.RPCDemo_ResquestStreamingServer) error {
	// 1.for循环接收客户端发送的消息
	var builder strings.Builder
	for {
		// 2. 通过 Recv() 不断获取客户端 send()推送的消息
		req, err := stream.Recv() // Recv内部也是调用RecvMsg
		builder.WriteString(req.GetRequestMessage())
		builder.WriteString(" ")
		// 3. err == io.EOF表示已经获取全部数据
		if err == io.EOF {
			log.Println("client send finished")
			// 4.SendAndClose 返回并关闭连接
			// 在客户端发送完毕后服务端即可返回响应
			return stream.SendAndClose(&types.Response{ResponseMessage: fmt.Sprintf("Hello world %s", builder.String())})
		}
		if err != nil {
			return err
		}
	}
}
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
