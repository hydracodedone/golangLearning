package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"rpc_demo/Demo2/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func tcpServer() {
	rpcServer := grpc.NewServer()
	service.RegisterProductServiceServer(rpcServer, new(service.ProductService))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("rpc tcp listen fail:<%s>", err.Error())
	}
}
func tcpServerWithCredential() {
	cred, err := credentials.NewClientTLSFromFile("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/certificate/server.cert", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/certificate/server_no_password.key")
	if err != nil {
		log.Fatalf("generate credentials fail:<%s>", err.Error())
	}
	rpcServer := grpc.NewServer(grpc.Creds(cred))
	service.RegisterProductServiceServer(rpcServer, new(service.ProductService))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("rpc tcp listen fail:<%s>", err.Error())
	}
}
func ServerWithMutualCredential() {
	cert, err := tls.LoadX509KeyPair("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/mutualCertification/server.pem", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/mutualCertification/server.key")
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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	rpcServer := grpc.NewServer(grpc.Creds(cred))
	service.RegisterProductServiceServer(rpcServer, new(service.ProductService))
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("tcp listen fail:<%s>", err.Error())
	}
}
func httpServer() {
	rpcServer := grpc.NewServer()
	service.RegisterProductServiceServer(rpcServer, new(service.ProductService))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("The proto is %v\n", request.Proto)
		fmt.Printf("The request url is %v\n", request.RequestURI)
		fmt.Printf("The head is %v\n", request.Header)
		fmt.Printf("The body is %+v\n", request.Body)

		rpcServer.ServeHTTP(writer, request)
	})
	httpServer := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	err := httpServer.ListenAndServeTLS("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/certificate/server.cert", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2/certificate/server_no_password.key")
	if err != nil {
		log.Fatalf("https listen and server fail:<%s>", err.Error())
	}
}

func main() {
	ServerWithMutualCredential()
}
