# OPENTRACING
Opentracing 是分布式链路追踪的一种规范标准，是 CNCF（云原生计算基金会）下的项目之一。

Opentracing 不是传输协议，消息格式层面上的规范标准，而是一种语言层面上的 API 标准。以 Go 语言为例，只要某链路追踪系统实现了 Opentracing 规定的接口（interface），符合Opentracing 定义的表现行为，那么就可以说该应用符合 Opentracing 标准。

这意味着开发者只需修改少量的配置代码，就可以在符合 Opentracing 标准的链路追踪系统之间自由切换。
## Data Model
### Span
Span 是一条追踪链路中的基本组成要素，一个 Span 表示一个独立的工作单元，比如可以表示一次函数调用，一次 HTTP 请求等等。Span 会记录如下基本要素:

    服务名称(operation name)
    服务的开始时间和结束时间
    K/V形式的Tags
    K/V形式的Logs
    SpanContext
    References：该span对一个或多个span的引用（通过引用SpanContext）

#### Tag

    Tags以K/V键值对的形式保存用户自定义标签，主要用于链路追踪结果的查询过滤。   
        例如： 
            http.method="GET",http.status_code=200。
            其中key值必须为字符串，value必须是字符串，布尔型或者数值型。Span 中的 tag 仅自己可见，不会随着 SpanContext 传递给后续 Span。
#### Logs

    Logs 与 tags 类似，也是 K/V 键值对形式。
    与 tags 不同的是，logs 还会记录写入 logs 的时间，因此 logs 主要用于记录某些事件发生的时间。
    logs 的 key 值同样必须为字符串，但对 value 类型则没有限制

#### SpanContext

    SpanContext携带着一些用于跨服务通信的（跨进程）数据，主要包含：

    足够在系统中标识该span的信息，比如：span_id,trace_id。

#### Baggage Items，为整条追踪连保存跨服务（跨进程）的K/V格式的用户自定义数据。

    Baggage Items 与 tags 类似，也是 K/V键值对。与 tags 不同的是：

    其 key 跟 value 都只能是字符串格式
    Baggage items 不仅当前 span 可见，其会随着 SpanContext 传递给后续所有的子 span。
    要小心谨慎的使用baggage items——因为在所有的span中传递这些K,V会带来不小的网络和CPU开销。

#### References
    
    Opentracing 定义了两种引用关系:ChildOf和FollowFrom。

        ChildOf: 父span的执行依赖子span的执行结果时，此时子span对父span的引用关系是ChildOf。比如对于一次RPC调用，服务端的span（子span）与客户端调用的span（父span）是ChildOf关系。

        FollowFrom：父span的执不依赖子span执行结果时，此时子span对父span的引用关系是FollowFrom。FollowFrom常用于异步调用的表示，例如消息队列中consumerspan与producerspan之间的关系。
### Trace
Trace表示一次完整的追踪链路，trace由一个或多个span组成。下图示例表示了一个由8个span组成的trace
```
        [Span A]  ←←←(the root span)
            |
     +------+------+
     |             |
 [Span B]      [Span C] ←←←(Span C is a `ChildOf` Span A)
     |             |
 [Span D]      +---+-------+
               |           |
           [Span E]    [Span F] >>> [Span G] >>> [Span H]
                                       ↑
                                       ↑
                                       ↑
                         (Span G `FollowsFrom` Span F)
```


### Inject/Extract
为了实现分布式系统中的链路追踪，Opentracing 提供了 Inject/Extract 用于在请求中注入 SpanContext 或者从请求中提取出 SpanContext。

    客户端通过 Inject 将 SpanContext 注入到载体中，随着请求一起发送到服务端。
    服务端则通过 Extract 将 SpanContext 提取出来,进行后续处理。

## 产品选型
分布式链路追踪有大量相关产品，具体如下：

    Twitter：Zipkin
    Uber：Jaeger
    Elastic Stack：Elastic APM
    Apache：SkyWalking
    Naver：Pinpoint
    阿里：鹰眼
    大众点评：Cat
    京东：Hydra

### Jaeger
#### 组成
Jaeger Client 
    
    为不同语言实现了符合 OpenTracing 标准的 SDK。应用程序通过 API 写入数据，client library 把 trace 信息按照应用程序指定的采样策略传递给 jaeger-agent。

Agent 
    
    它是一个监听在 UDP 端口上接收 span 数据的网络守护进程，它会将数据批量发送给 collector。它被设计成一个基础组件，部署到所有的宿主机上。Agent 将 client library 和 collector 解耦，为 client library 屏蔽了路由和发现 collector 的细节。

Collector 

    接收 jaeger-agent 发送来的数据，然后将数据写入后端存储。Collector 被设计成无状态的组件，因此您可以同时运行任意数量的 jaeger-collector。

Data Store 
    
    后端存储被设计成一个可插拔的组件，支持将数据写入 cassandra、elastic search。

Query 
    
    接收查询请求，然后从后端存储系统中检索 trace 并通过 UI 进行展示。Query 是无状态的，您可以启动多个实例，把它们部署在 nginx 这样的负载均衡器后面。

分布式追踪系统发展很快，种类繁多，但核心步骤一般有三个：代码埋点，数据存储、查询展示
#### docker
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 14250:14250 \
-p 9411:9411 \
jaegertracing/all-in-one:latest

docker run -di --rm --name jaeger -p6831:6831/udp -p16686:16686 jaegertracing/all-in-one:latest
#### 因果关系
span 是链路追踪里的最小组成单元，为了保留各个功能之间的因果关系，必须在各个方法之间传递 span 并且新建span时指定opentracing.ChildOf(rootSpan.Context()),否则新建的span会是独立的，无法构成一个完整的 trace。
#### 使用context传递
如果再各个方法中传递span,会污染整个程序 因此,Go 语言中的 context.Context对象来进行传递。

实例代码如下：
```GO
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
func main(){
    tracer, closer := tracerInitialize("main")
    defer closer.Close()
    span := tracer.StartSpan("main")
    defer span.Finish()
    opentracing.SetGlobalTracer(tracer)
    ctx := context.Background()
    ctx = opentracing.ContextWithSpan(ctx, span)
    callFunc()
}
func callFunc(cxt context.Context){
    span, spanCtx := opentracing.StartSpanFromContext(ctx, "callFunc")
    defer span.Finish()
    callFunc2(spanCtx)
}
func callFunc2(cxt context.Context){
    span, _spanCtx_ := opentracing.StartSpanFromContext(ctx, "callFunc")
    defer span.Finish()
}

```
需要注意:

    opentracing.StartSpanFromContext()返回的第二个参数是子ctx,如果需要的话可以将该子ctx继续往下传递，而不是传递父ctx。

    需要注意的是opentracing.StartSpanFromContext()默认使用GlobalTracer来开始一个新的 span，所以使用之前需要设置 GlobalTracer。
    opentracing.SetGlobalTracer(tracer)