package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"rpc_demo/Demo2_1/service"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GateWithMutualCredentialHttpServer() {
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
	gateWayMux := runtime.NewServeMux()
	opt := []grpc.DialOption{
		grpc.WithTransportCredentials(cred),
	}
	err = service.RegisterProductServiceHandlerFromEndpoint(
		context.Background(),
		gateWayMux, ":8000", //RPC
		opt,
	)
	if err != nil {
		log.Fatalf("rpc register fail :<%s>", err.Error())
	}
	httpServer := &http.Server{
		Addr:    "localhost:8001",
		Handler: gateWayMux,
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("https listen and server fail:<%s>", err.Error())
	}
}
