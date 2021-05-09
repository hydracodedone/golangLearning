package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"rpc_demo/Demo2/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func client() {
	grpcClient, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("grpc close fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewProductServiceClient(grpcClient)
	productResponseInstance, err := productClient.GetProductStock(
		context.Background(),
		&service.ProductRequest{
			ProductId: []int32{1, 2},
		},
	)
	if err != nil {
		log.Fatalf("grpc GetProductStock fail:<%s>", err.Error())
	}
	fmt.Printf("The productResponse is %s", productResponseInstance)
}
func clientWithCredential() {
	cred, err := credentials.NewClientTLSFromFile("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/certificate/server.cert", "hydra.com")
	if err != nil {
		log.Fatalf("generate credentials fail:<%s>", err.Error())
	}
	grpcClient, err := grpc.Dial(":8000", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("grpc dial fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("grpc close fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewProductServiceClient(grpcClient)
	productResponseInstance, err := productClient.GetProductStock(
		context.Background(),
		&service.ProductRequest{
			ProductId: []int32{1, 2},
		},
	)
	if err != nil {
		log.Fatalf("grpc GetProductStock fail:<%s>", err.Error())
	}
	fmt.Printf("The productResponse is %s", productResponseInstance)
}
func tcpClientWithMutualCredential() {
	cert, err := tls.LoadX509KeyPair("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/mutualCertification/client.pem", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/mutualCertification/client.key")
	if err != nil {
		log.Fatalf("lood server.pem and server.key fail:<%s>", err.Error())
	}
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/mutualCertification/ca.pem")
	if err != nil {
		log.Fatalf("lood ca.pem fail:<%s>", err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	grpcClient, err := grpc.Dial(":8000", grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatalf("grpc dial fail:<%s>", err.Error())
	}
	defer func() {
		err = grpcClient.Close()
		if err != nil {
			log.Fatalf("grpc close fail:<%s>", err.Error())
		}
	}()
	productClient := service.NewProductServiceClient(grpcClient)
	productResponseInstance, err := productClient.GetProductStock(
		context.Background(),
		&service.ProductRequest{
			ProductId: []int32{1, 2},
		},
	)
	if err != nil {
		log.Fatalf("grpc GetProductStock fail:<%s>", err.Error())
	}
	fmt.Printf("The productResponse is %s", productResponseInstance)
}
func main() {
	tcpClientWithMutualCredential()
}
