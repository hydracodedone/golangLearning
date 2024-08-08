package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func TraceMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
			ServiceName: "gin-service",
		}
		tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
		}
		defer closer.Close()
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(ctx.Request.Header))
		if err == nil {
			ctx.Set("spanCtx", spanCtx)
		}
		ctx.Set("tracer", tracer)
		ctx.Next()
	}
}
