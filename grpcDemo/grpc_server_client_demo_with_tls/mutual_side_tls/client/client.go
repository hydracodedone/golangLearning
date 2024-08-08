package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"grpc_server_client_demo_with_tls/types"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

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
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	gResp, err := gClient.DemoRequest(context.Background(), &types.Request{
		Id: int32(1),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
