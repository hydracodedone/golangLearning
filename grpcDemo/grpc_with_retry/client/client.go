package main

import (
	"context"
	"fmt"
	"grpc_with_retry/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//retry
	// 更多配置信息查看官方文档： https://github.com/grpc/grpc/blob/master/doc/service_config.md
	// service这里语法为<package>.<service> package就是proto文件中指定的package，service也是proto文件中指定的 Service Name。
	// method 可以不指定 即当前service下的所以方法都使用该配置。

	//客户端需要配置环境变量 export GRPC_GO_RETRY=on 也貌似不需要
	retryPolicy := `{
		"methodConfig": [{
		  "name": [{"service": "meta.RPCDemo","method":"DemoRequest"}],
		  "retryPolicy": {
			  "MaxAttempts": 5,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
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
