package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"raw_rpc_demo/common"
)

func rpcHttpClient() {
	parameterInstance := new(common.Parameter)
	parameterInstance.Length = 10
	parameterInstance.Width = 20
	rpcClient, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatalf("Dial Http Fail:<%s>", err.Error())
	}
	defer func() {
		err := rpcClient.Close()
		if err != nil {
			log.Fatalf("RpcClient Close Fail:<%s>", err.Error())
		}
	}()
	result := 0
	err = rpcClient.Call("Rect.Area", &parameterInstance, &result)
	if err != nil {
		log.Fatalf("Rpc Call Fail:<%s>", err.Error())
	}
	fmt.Printf("The Area is %d\n", result)
}
func rpcTcpClient() {
	parameterInstance := new(common.Parameter)
	parameterInstance.Length = 10
	parameterInstance.Width = 20
	rpcClient, err := rpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalf("Dial Http Fail:<%s>", err.Error())
	}
	defer func() {
		err := rpcClient.Close()
		if err != nil {
			log.Fatalf("RpcClient Close Fail:<%s>", err.Error())
		}
	}()
	result := 0
	err = rpcClient.Call("Rect.Area", &parameterInstance, &result)
	if err != nil {
		log.Fatalf("Rpc Call Fail:<%s>", err.Error())
	}
	fmt.Printf("The Area is %d\n", result)
}
func JsonRpcClient() {
	parameterInstance := new(common.Parameter)
	parameterInstance.Length = 10
	parameterInstance.Width = 20
	rpcClient, err := jsonrpc.Dial("tcp", ":8000")
	if err != nil {
		log.Fatalf("Dial Http Fail:<%s>", err.Error())
	}
	defer func() {
		err := rpcClient.Close()
		if err != nil {
			log.Fatalf("RpcClient Close Fail:<%s>", err.Error())
		}
	}()
	result := 0
	err = rpcClient.Call("Rect.Area", &parameterInstance, &result)
	if err != nil {
		log.Fatalf("Rpc Call Fail:<%s>", err.Error())
	}
	fmt.Printf("The Area is %d\n", result)
}
func main() {
	JsonRpcClient()
}
