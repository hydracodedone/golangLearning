package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"grpc_with_credentials/types"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

// 自定义拦截器
func CustomInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 如果返回err不为nil则说明token验证未通过
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	} else {
		fmt.Println(md)
	}
	return handler(ctx, req)
}

// create server
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
	// 创建CertPool，后续就用池里的证书来校验客户端证书有效性
	// 所以如果有多个客户端 可以给每个客户端使用不同的 CA 证书，来实现分别校验的目的
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("ca.crt")
	if err != nil {
		panic(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		panic("failed to append certs")
	}

	// 构建基于 TLS 的 TransportCredentials
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})

	gServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(CustomInterceptor))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
