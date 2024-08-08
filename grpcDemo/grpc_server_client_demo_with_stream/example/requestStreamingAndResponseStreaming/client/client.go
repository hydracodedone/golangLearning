package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_stream/types"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
1. 建立连接 获取client
2. 通过client获取stream
3. 开两个goroutine 分别用于Recv()和Send()
	3.1 一直Recv()到err==io.EOF(即服务端关闭stream)
	3.2 Send()则由自己控制
4. 发送完毕调用 stream.CloseSend()关闭stream 必须调用关闭 否则Server会一直尝试接收数据 一直报错...
*/

// wrappedStream  用于包装 grpc.ClientStream 结构体并拦截其对应的方法。
type wrappedClientStream struct {
	grpc.ClientStream
}

func newWrappedClientStream(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedClientStream{s}
}

func (w *wrappedClientStream) RecvMsg(m interface{}) error {
	fmt.Printf("C RecvMsg0: %v\n", m)
	err := w.ClientStream.RecvMsg(m)
	fmt.Printf("C RecvMsg1: %v\n", m)
	return err
}

func (w *wrappedClientStream) SendMsg(m interface{}) error {
	fmt.Println("C SendMsg:", m)
	return w.ClientStream.SendMsg(m)
}

func ClientStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	// 返回的是自定义的封装过的 stream
	return newWrappedClientStream(s), nil
}

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainStreamInterceptor(ClientStreamInterceptor))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	requestStreamAndResponseStreamServer, err := gClient.ResquestStreamingAndResponseStreaming(context.Background())
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	// 3.开两个goroutine 分别用于Recv()和Send()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			req, err := requestStreamAndResponseStreamServer.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Println("Recv Finished")
					break
				} else {
					panic(err)
				}
			} else {
				fmt.Printf("Recv Data:%v \n", req.GetResponseMessage())
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		requestStreamAndResponseStreamServer.Send(&types.Request{RequestMessage: "Hydra1"})
		time.Sleep(time.Second)
		requestStreamAndResponseStreamServer.Send(&types.Request{RequestMessage: "Hydra2"})
		err := requestStreamAndResponseStreamServer.CloseSend()
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
}
