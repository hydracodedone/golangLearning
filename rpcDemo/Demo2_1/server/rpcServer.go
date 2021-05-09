package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"rpc_demo/Demo2_1/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GateWithMutualCredentialRpcServer() {
	cert, err := tls.LoadX509KeyPair("/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/server.pem", "/home/hydra/Project/GolangLearning/src/rpcDemo/Demo2_1/mutualCertification/server.key")
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
