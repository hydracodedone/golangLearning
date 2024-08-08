package main

import (
	"context"
	"distribute_tracing_grpc/common"
	"distribute_tracing_grpc/types"

	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// pre-processing
	cfg := &config.Configuration{
		// 采样率暂配置，设置为1，全部采样
		// 如果每个请求都保存到jeager中，压力会大，所以可以设置采集速率
		// 如：rateLimiting:每秒spans数
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,             // 是否打印日志
			LocalAgentHostPort: "localhost:6831", // jeager默认端口是6831
		},
		ServiceName: "grpc-client",
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	span := tracer.StartSpan("grpc_client")
	md := metadata.New(nil)
	tracer.Inject(span.Context(), opentracing.TextMap, common.MDReaderWriter{md})
	mdCtx := metadata.NewOutgoingContext(ctx, md)
	err = invoker(mdCtx, method, req, reply, cc, opts...) // invoking RPC method
	// post-processing
	return err
}

func main() {
	//Dial
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(unaryInterceptor))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()
	//Create Client
	gClient := types.NewRPCDemoClient(gCon)
	//Rpc call
	timeoutCtx, cacelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cacelFunc()
	gResp, err := gClient.DemoRequest(timeoutCtx, &types.Request{
		Id: int32(1),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("the ID is %v\n", gResp.GetId())
}
