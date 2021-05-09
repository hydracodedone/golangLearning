package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"hydracode.com/rpcDemo/Demo2_1/service"
)

func gateWithMutualCredentialRpcClient() {
	cert, err := tls.LoadX509KeyPair("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/client.pem", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/client.key")
	if err != nil {
		log.Fatalf("lood server.pem and server.key fail:<%s>", err.Error())
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/ca.pem")
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
			ProductId: 10,
		},
	)
	if err != nil {
		log.Fatalf("grpc GetProductStock fail:<%s>", err.Error())
	}
	fmt.Printf("The productResponse is %s", productResponseInstance)
}

func gateWithMutualCredentialHttpClient() {
	//cert, err := tls.LoadX509KeyPair("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/client.pem", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/client.key")
	//if err != nil {
	//	log.Fatalf("load server.pem and server.key fail:<%s>", err.Error())
	//}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/ca.pem")
	if err != nil {
		log.Fatalf("lood ca.pem fail:<%s>", err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	tr := &http.Transport{
		////把从服务器传过来的非叶子证书，添加到中间证书的池中，使用设置的根证书和中间证书对叶子证书进行验证。
		TLSClientConfig: &tls.Config{RootCAs: certPool},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://localhost:8001/v1/product/1")
	if err != nil {
		log.Fatalf("request fail:<%s>", err.Error())
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read request fail:<%s>", err.Error())
	}
	fmt.Printf("the http response is %s\n", data)
}
func main() {
	gateWithMutualCredentialHttpClient()
}
