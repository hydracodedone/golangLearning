package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func tracerInitialize(service_name string) (opentracing.Tracer, io.Closer) {
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
		ServiceName: service_name, // 服务名字，也可以在下面NewTracer的时候传入，不过弃用了
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
func singleSpanDemo() {
	tracer, closer := tracerInitialize("single-span-test")
	defer closer.Close()
	span := tracer.StartSpan("span_root")
	defer span.Finish()
	func() {
		span.SetTag("now", time.Now().String())
		span.LogFields(
			log.String("event", "print now"),
		)
	}()
	time.Sleep(2 * time.Second)
}
func causality() {
	x := 4
	y := 5
	tracer, closer := tracerInitialize("causality")
	defer closer.Close()
	span := tracer.StartSpan("begin causality")
	span.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	defer span.Finish()
	res := calculate1(span, x, y)
	span.LogFields(
		log.Int("output", res),
	)

}
func calculate1(span opentracing.Span, x int, y int) int {
	cspan := span.Tracer().StartSpan("calculate1", opentracing.ChildOf(span.Context()))
	defer cspan.Finish()
	cspan.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	z := x + y
	cspan.LogFields(
		log.Int("middle", z),
	)
	time.Sleep(1 * time.Second)
	res := calculate2(cspan, z, x*y)
	cspan.LogFields(
		log.Int("output", res),
	)
	return res
}
func calculate2(span opentracing.Span, x int, y int) int {
	cspan := span.Tracer().StartSpan("calculate2", opentracing.ChildOf(span.Context()))
	defer cspan.Finish()
	cspan.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	time.Sleep(2 * time.Second)
	res := x + y
	cspan.LogFields(
		log.Int("output", res),
	)
	return res
}

func causality2(ctx context.Context) {
	x := 4
	y := 5
	tracer, closer := tracerInitialize("causality")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	span := tracer.StartSpan("begin causality")
	span.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	defer span.Finish()
	spanCtx := opentracing.ContextWithSpan(ctx, span)
	res := calculate1_1(spanCtx, x, y)
	span.LogFields(
		log.Int("output", res),
	)

}
func calculate1_1(ctx context.Context, x int, y int) int {

	cspan, spanCtx := opentracing.StartSpanFromContext(ctx, "calculate1_1")
	defer cspan.Finish()
	cspan.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	z := x + y
	cspan.LogFields(
		log.Int("middle", z),
	)
	time.Sleep(1 * time.Second)
	res := calculate2_1(spanCtx, z, x*y)
	cspan.LogFields(
		log.Int("output", res),
	)
	return res
}
func calculate2_1(ctx context.Context, x int, y int) int {
	cspan, _ := opentracing.StartSpanFromContext(ctx, "calculate2_1")
	defer cspan.Finish()
	cspan.LogFields(
		log.Int("input", x),
		log.Int("input", y),
	)
	time.Sleep(2 * time.Second)
	res := x + y
	cspan.LogFields(
		log.Int("output", res),
	)
	return res
}

func main() {
	// singleSpanDemo()
	ctx := context.Context(context.Background())
	causality2(ctx)
}
