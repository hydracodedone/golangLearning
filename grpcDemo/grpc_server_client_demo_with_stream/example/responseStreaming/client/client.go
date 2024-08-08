package main

import (
	"context"
	"fmt"
	"grpc_server_client_demo_with_stream/types"
	"io"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	request := types.Request{RequestMessage: "Hydra"}
	//send
	responseStreamingServer, err := gClient.ResponseStreaming(context.Background(), &request)
	if err != nil {
		panic(err)
	}
	builder := strings.Builder{}
	//recieve
	for {
		response, err := responseStreamingServer.Recv()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		} else {
			builder.WriteString(response.GetResponseMessage())
			builder.WriteString(" ")
		}
	}
	fmt.Println(builder.String())
}
