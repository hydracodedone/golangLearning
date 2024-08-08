package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"raw_rpc_demo/common"
)

func rpcHttpServer() {
	var rpcService common.RpcService
	rectInstance := new(common.Rect)
	rpcService = rectInstance

	common.RpcServiceRegisterService(rpcService)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Http Listen And Server Fail:<%s>", err.Error())
	}
}
func rpcHttpServer2() {
	rectInstance := new(common.Rect)
	err := rpc.Register(rectInstance)
	if err != nil {
		log.Fatalf("RPC Register Fail:<%s>", err.Error())
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Listen Fail:<%s>", err.Error())
	}
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatalf("Http Serve Fail:<%s>", err.Error())
	}
}

func rpcTcpServer() {
	rectInstance := new(common.Rect)
	err := rpc.Register(rectInstance)
	if err != nil {
		log.Fatalf("RPC Register Fail:<%s>", err.Error())
	}
	ip := "127.0.0.1"
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatalf("Listen Fail:<%s>", err.Error())
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Listener Close Fail:%s", err.Error())
		}
	}()
	temp := make(chan int)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Accept Fail:<%s>", err.Error())
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				log.Fatalf("Connetion Close Fail:%s", err.Error())
			}
		}()
		rpc.ServeConn(conn)
	}()
	<-temp
}
func JsonRpcTcpServer() {
	rectInstance := new(common.Rect)
	err := rpc.Register(rectInstance)
	if err != nil {
		log.Fatalf("RPC Register Fail:<%s>", err.Error())
	}
	ip := "127.0.0.1"
	port := 8000
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		log.Fatalf("Listen Fail:<%s>", err.Error())
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Listener Close Fail:%s", err.Error())
		}
	}()
	temp := make(chan int)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Accept Fail:<%s>", err.Error())
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				log.Fatalf("Connetion Close Fail:%s", err.Error())
			}
		}()
		jsonrpc.ServeConn(conn)
	}()
	<-temp
}
func main() {
	rpcHttpServer()
}
