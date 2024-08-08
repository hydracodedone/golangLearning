package main

import (
	"context"
	"distribute_tracing_grpc/common"
	"distribute_tracing_grpc/types"

	"fmt"
	"net"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// embed
type Server struct {
	types.UnimplementedRPCDemoServer
}

// metadata 读写
type MDReaderWriter struct {
	metadata.MD
}

// 为了 opentracing.TextMapReader ，参考 opentracing 代码
func (c MDReaderWriter) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range c.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

// 为了 opentracing.TextMapWriter，参考 opentracing 代码
func (c MDReaderWriter) Set(key, val string) {
	key = strings.ToLower(key)
	c.MD[key] = append(c.MD[key], val)
}
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var finalCtx context.Context
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
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
			ServiceName: "grpc-server",
		}
		tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}
		defer closer.Close()
		spanCtx, err := tracer.Extract(opentracing.TextMap, common.MDReaderWriter{md})
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot extract spanCtx: %v\n", err))
		}
		span := tracer.StartSpan("grpc_server", ext.RPCServerOption(spanCtx))
		finalCtx = opentracing.ContextWithSpan(ctx, span)
	} else {
		finalCtx = ctx
	}
	m, err := handler(finalCtx, req)
	return m, err
}

// implementation interaface
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	reqID := req.GetId()
	return &types.Response{
		Id: reqID,
	}, nil
}

// create server
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	gServer := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	types.RegisterRPCDemoServer(gServer, &Server{})
	err = gServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
