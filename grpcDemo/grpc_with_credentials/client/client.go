package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc_with_credentials/types"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type MyAuth struct {
	Username string
	Password string
}

const (
	Admin = "admin"
	Root  = "root"
)

func NewMyAuth() *MyAuth {
	return &MyAuth{
		Username: Admin,
		Password: Root,
	}
}

// GetRequestMetadata 定义授权信息的具体存放形式，最终会按这个格式存放到 metadata map 中。
func (a *MyAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"username": a.Username, "password": a.Password}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 加密连接进行安全传输
func (a *MyAuth) RequireTransportSecurity() bool {
	return false
}
func main() {
	//Cred
	certificate, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {
		log.Fatal(err)
	}
	// 构建CertPool以校验服务端证书有效性
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.hydracode.com", // NOTE: this is required!
		RootCAs:      certPool,
	})
	//make creadentials
	if err != nil {
		panic(err)
	}
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(NewMyAuth()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)

	//Rpc call
	gResp, err := gClient.DemoRequest(context.Background(), &types.Request{Id: int32(1)})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
