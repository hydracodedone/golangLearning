# RPC
rpc，全称 remote process call（远程过程调用），是微服务架构下的一种通信模式. 这种通信模式下，一台服务器在调用远程机器的接口时，能够获得像调用本地方法一样的良好体验.

    rpc 调用基于 sdk 方式，调用方法和出入参协议固定，stub 文件本身还能起到接口文档的作用，很大程度上优化了通信双方约定协议达成共识的成本.
    rpc 在传输层协议 tcp 基础之上，可以由实现框架自定义填充应用层协议细节，理论上存在着更高的上限
# GRPC
gRPC  是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计。目前提供 C、Java 和 Go 语言版本，分别是：grpc, grpc-java, grpc-go.

gRPC 基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。

# proto buffers

protocol buffers，是一套结构数据序列化机制（当然也可以使用其他数据格式如 JSON）用 proto files 创建 gRPC 服务，用 protocol buffers 消息类型来定义方法参数和返回类型

    代码生成
    序列化与反序列化
    支持跨平台
# 环境安装

## 安装 grpc

    go get google.golang.org/grpc@latest
## 安装 protocol buffer

根据操作系统型号，下载安装好对应版本的 protobuf 应用：

https://github.com/google/protobuf/releases

需要将 protobuf 执行文件所在的目录添加到环境变量 $PATH 当中.

安装完成后，可以通过查看 protobuf 版本指令，校验安装是否成功

    protoc --version
## 安装插件protoc-gen-go

不要使用 github.com/golang/protobuf/protoc-gen-go 这个版本
安装 protobuf -> pb.go 插件

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

该插件的作用是，能够基于 .proto 文件一键生成 _pb.go 文件，对应内容为通信请求/响应参数的对象模型.

go install 指令默认会将插件安装到 GOPATH/bin 目录下. 需要确保 GOPATH/bin 路经有被添加到环境路经 $PATH 当中.

安装完成后，可以通过查看插件版本指令，校验安装是否成功

    protoc-gen-go --version

## 安装插件protoc-gen-go-grpc
安装 protobuf -> grpc.pb.go 插件

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

该插件的作用是，能够基于 .proto 文件生成 _grpc.pb.go，对应内容为通信服务框架代码.

安装完成后，可以通过查看插件版本指令，校验安装是否成功

    protoc-gen-go-grpc --version



# proto语法
正如其他 RPC 系统，gRPC 基于如下思想：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构

```
message HelloRequest {
  required string greeting = 1;
}

message HelloResponse {
  required string reply = 1;
}

service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}
```
gRPC 允许你定义四类服务方法：

1. 单项 RPC，即客户端发送一个请求给服务端，从服务端获取一个应答，就像一次普通的函数调用。
```
    rpc SayHello(HelloRequest) returns (HelloResponse){
    }
```
2. 服务端流式 RPC，即客户端发送一个请求给服务端，可获取一个数据流用来读取一系列消息。客户端从返回的数据流里一直读取直到没有更多消息为止
```
    rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
    }
```


3. 客户端流式 RPC，即客户端用提供的一个数据流写入并发送一系列消息给服务端。一旦客户端完成消息写入，就等待服务端读取这些消息并返回应答
```
    rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
    }
```
4. 双向流式 RPC，即两边都可以分别通过一个读写数据流来发送一系列消息。这两个数据流操作是相互独立的，所以客户端和服务端能按其希望的任意顺序读写，例如：服务端可以在写应答前等待所有的客户端消息，或者它可以先读一个消息再写一个消息，或者是读写相结合的其他方式。每个数据流里消息的顺序会被保持。
```
    rpc BidiHello(stream HelloRequest) returns (stream HelloResponse){
    }
```
## proto3语法注意事项
    字段默认为optional,无需额外指定
    不能使用required
    不能使用extension
    在proto3中，repeated修饰的字段默认使用packed编码
    使用reserved来保留要删除的字段号或者字段名称
    你可以使用max关键字来声明你保留的条目值的范围一直到最大值。
    string：默认值是空字符串
    bytes：默认值是空byte数组
    bool：默认是false
    数字：默认是0
    enum：默认值是定义的第一个枚举值，且第一个枚举值一定要是0
    message：它的默认值取决于语言的实现。
    repeated字段的默认值为空（通常是一个空数组）

## 注意事项
    package 该关键字用于申明proto文件对应的包,主要作用是其他proto文件在import该package时候,访问该package中的message的时侯通过package.message进行调用

    option go_package用于生成的.pb.go文件,在引用和生成go包名时起作用
        option go_package = "{out_path};out_go_package"
        前一个参数用于指定生成文件的位置，以及生成go文件被引用时候的导入路经
        后一个参数指定生成的 .go 文件的 package
        这里指定的 out_path 并不是绝对路经，只是相对路经或者说只是路经的一部分，和 protoc 的 --go_out 拼接后才是完整的路经

        实践经验:
            由于go module的作用,建议out_path使用绝对路经,配置--go_out=.来生成

            go-zero中不建议使用out_go_package字段,在不使用out_go_package的时候,生成的go的package包名会使用out_path最后的路经作为package名
                如指定的out_go_package为"github.com/zero-micro/utils",则生成的go文件的package为utils
        
# 安全性

    SSL/TLS认证方式
    基于Token的认证方式
    自定义身份认证
    不进行身份认证

TLS主要是解决三方面的问题

    1. 信息保密(信息加密传输)
    2. 通过MAC校验机制,防止篡改
    3. 认证,双方都可以配备证书,防止身份被冒充

SSL(Secure Sockets Layer 安全套阶层)标准化后称为TLS(传输层Transport Layer Security安全协议)
HTTP + (SSL/TLS) = HTTPS
TCP协议是HTTP协议的基石,HTTP协议依赖TCP协议传输数据,TCP为传输层协议,HTTP为应用层协议
对称加密: 加密和解密使用相同的密钥
不对称加密: 加密和解密使用不同的密钥

gRPC 内置了以下 encryption 机制：

    SSL / TLS：通过证书进行数据加密；
    ALTS：Google开发的一种双向身份验证和传输加密系统。只有运行在 Google Cloud Platform 才可用，一般不用考虑。
gRPC 中的连接类型一共有以下3种：

    insecure connection：不使用TLS加密
    server-side TLS：仅服务端TLS加密
    mutual TLS：客户端、服务端都使用TLS加密
### 密钥生成
#### 单向证书
    生成私钥(有密码)
    openssl genrsa -des3 -out server.key 2048
    生成私钥(无密码)
    openssl genrsa -out server.key 2048

    生成csr(证书签名请求)
    openssl req -new -key server.key -out server.csr(注意COMMON NAMES 使用域名 *.xxx.com)

    自签名生成证书(用csr)
    openssl x509 -req -days 36500 -in server.csr -signkey server.key  -out server.crt
    openssl x509 -req -days 36500 -in server.csr -key server.key  -out server.crt
    

##### 双向证书
生成带 SAN 的证书

SAN（Subject Alternative Name）是 SSL 标准 x509 中定义的一个扩展。使用了 SAN 字段的 SSL 证书，可以扩展此证书支持的域名，使得一个证书可以支持多个不同域名的解析

    拷贝/etc/ssl/openssl.conf 到当前目录
    修改
    打开 copy_extensions= copy
    打开 req_extensions= v3_req
    找到[v3_req],添加 subjectAltName = @alt_names
    添加新标签
    [ alt_names ]
    DNS.1=*.xxx.com
    
    生成CA证书
    
        openssl genpkey -algorithm RSA -out=server.key
        openssl req -x509 -new -nodes -key server.key  -days 5000 -out server.pem
    
    生成服务端证书
    openssl req -new -nodes -config ./openssl.cnf -keyout server.key   -out server.csr
    生成客户端证书


#### 解析

    key 指的是服务器上的私钥 加密要发送给客户端的数据,解密收到客户端的数据
    csr 证书签名的请求文件,用于提交给ca对证书进行签名
    crt ca颁发的证书或者开发者的自签名证书,包含了证书持有人的信息,持有人的公钥,签署者的签名信息
    pem 基于base64编码的证书扩展名包括 pem,crt,cer

# grpc 拦截器
推荐一下这个 go-grpc-middleware
https://github.com/grpc-ecosystem/go-grpc-middleware
## 拦截器分类与定义 
    
    类似于middlewarem,本质上就是一个特定类型的函数，所以实现拦截器只需要实现对应类型方法即可。
### 根据数据类型分类
    一元拦截器
    流拦截器
### 根据拦截器所处位置
    服务端拦截器
    客户端拦截器。

一共有以下4种类型:

    grpc.UnaryServerInterceptor
```Go
/*

ctx context.Context：请求上下文
req interface{}：RPC 方法的请求参数
info *UnaryServerInfo：RPC 方法的所有信息
handler UnaryHandler：RPC 方法真正执行的逻辑
*/
type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error) 
```
    grpc.UnaryClientInterceptor
```Go
/*
ctx：Go语言中的上下文，一般和 Goroutine 配合使用，起到超时控制的效果
method：当前调用的 RPC 方法名
req：本次请求的参数，只有在处理前阶段修改才有效
reply：本次请求响应，需要在处理后阶段才能获取到
cc：gRPC 连接信息
invoker：可以看做是当前 RPC 方法，一般在拦截器中调用 invoker 能达到调用 RPC 方法的效果，当然底层也是 gRPC 在处理。
opts：本次调用指定的 options 信息
*/
type UnaryClientInterceptor func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
```
    grpc.StreamServerInterceptor
```Go
type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
```
    grpc.StreamClientInterceptor
```Go
type StreamClientInterceptor func(ctx context.Context, desc *StreamDesc, cc *ClientConn, method string, streamer Streamer, opts ...CallOption) (ClientStream, error)
```
## 拦截器执行过程

一元拦截器

        预处理
        调用RPC方法
        后处理
流拦截器

        预处理
        用RPC方法 获取 Streamer
        后处理
            调用 SendMsg 、RecvMsg 之前
            调用 SendMsg 、RecvMsg
            调用 SendMsg 、RecvMsg 之后

## 多个拦截器执行顺序

    配置多个拦截器时，会按照参数传入顺序依次执行

    所以，如果想配置一个 Recovery 拦截器则必须放在第一个，放在最后则无法捕获前面执行的拦截器中触发的 panic。

    同时也可以将 一元和流拦截器一起配置，gRPC 会根据不同方法选择对应类型的拦截器执行。


# grpc gateway

    将gRPC 转为 RESTful HTTP API
    简单来说就是生成了一个 HTTP 服务，在具体处理逻辑中去请求 gRPC 服务

    当 HTTP 请求到达 gRPC-Gateway 时，它将 JSON 数据解析为 Protobuf 消息。然后，它使用解析的 Protobuf 消息发出正常的 Go gRPC 客户端请求。Go gRPC 客户端将 Protobuf 结构编码为 Protobuf 二进制格式，然后将其发送到 gRPC 服务器。gRPC 服务器处理请求并以 Protobuf 二进制格式返回响应。Go gRPC 客户端将其解析为 Protobuf 消息，并将其返回到 gRPC-Gateway，后者将 Protobuf 消息编码为 JSON 并将其返回给原始客户端
## grpc-gateway插件安装
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
## 生成对应的pb.gw.go文件
只需要在 protoc 编译时传递不同参数即可
如protoc --go_out . --go-grpc_out . --grpc-gateway_out . xxx.proto

## 生成步骤
1.引入文件

    引入annotations.proto

    import "google/api/annotations.proto";
    引入annotations.proto文件，因为添加的注解依赖该文件。

    该文件需要手动从 grpc-gateway/third_party/googleapis 目录复制到自己的项目中。

    该文件需要手动从 grpc-gateway/third_party/googleapis 目录复制到自己的项目中。
2.修改proto文件

    为每个方法都必须添加 google.api.http 注解后 gRPC-Gateway 才能生成对应 http 方法。

```Go
rpc SayHello (HelloRequest) returns (HelloReply) {
    option (google.api.http) = {
      get: "/v1/greeter/sayhello"
      body: "*"
    };
  }
```
3.编译生成grpc代码
3.编写server端

    编写逻辑为
    在server编写一个用于转发rest请求到grpc的grpc客户端服务
    开server端编写一个http服务,用于接收rest请求

```Go
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"grpc_gateway/types"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	types.UnimplementedRPCDemoServer
}
//实现grpc的业务逻辑
func (s *Server) DemoRequest(ctx context.Context, req *types.Request) (resp *types.Response, err error) {
	reqID := req.GetId()
	fmt.Println(reqID)
	return &types.Response{
		Id: reqID,
	}, nil
}
func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
    //创建grpc服务
	gServer := grpc.NewServer()
    //绑定业务
	types.RegisterRPCDemoServer(gServer, &Server{})
    //启动grpc服务
	go func() {
		err := gServer.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	//创建用于rest请求到grpc接口的客户端服务
	gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer gCon.Close()

	gwmux := runtime.NewServeMux()
	err = types.RegisterRPCDemoHandler(context.Background(), gwmux, gCon)
	if err != nil {
		panic(err)
	}
	//创建rest接口,并启动服务
	gwServer := &http.Server{
		Addr:    "localhost:9001",
		Handler: gwmux,
	}
	err = gwServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```
# grpc 自定义验证
在 gRPC 中，身份验证被抽象为了credentials.PerRPCCredentials接口
```Go
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}
所以需要实现自定义验证只需要实现上述接口即可
```
各方法作用如下：

GetRequestMetadata：以 map 的形式返回本次调用的授权信息，ctx 是用来控制超时的，并不是从这个 ctx 中获取。

RequireTransportSecurity：指该 Credentials 的传输是否需要需要 TLS 加密，如果返回 true 则说明该 Credentials 需要在一个有 TLS 认证的安全连接上传输，如果当前连接并没有使用 TLS 则会报错：

## 具体流程

    客户端请求时带上 Credentials；
    服务端取出 Credentials，并验证有效性，一般配合拦截器使用。
## 客户端
client发出请求之前，gRPC 会将 Credentials 存放在 metadata 中进行传递,在真正发起调用之前，gRPC 会通过 GetRequestMetadata函数，将用户定义的 Credentials 提取出来，并添加到 metadata 中，随着请求一起传递到服务端

客户端添加 Credentials 有两种方式：

    1.建立连接时指定,授权信息保存在 conn 对象上，通过该连接发起的每个调用都会附带上该授权信息
    2.发起调用时指定,可以为每个调用指定不同的授权信息
## 服务端
服务端从 metadata 中取出 Credentials 进行有效性校验
type MD map[string][]string
可以看到 MD 是一个 map ，授权信息在这个map中具体怎么存的由 PerRPCCredentials接口的GetRequestMetadata函数实现
# grpc retry
## 环境变量的配置
有可能需要

export GRPC_GO_RETRY=on

## 主要是客户端的配置
```GO
retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "xx.xxx","method":"xx"}],
		  "retryPolicy": {
			  "MaxAttempts": 4,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`
gCon, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
	if err != nil {
		panic(err)
	}
```
## 具体的配置
```GO
// Configuration for a method.
message MethodConfig {
  // The names of the methods to which this configuration applies.
  // - MethodConfig without names (empty list) will be skipped.
  // - Each name entry must be unique across the entire ServiceConfig.
  // - If the 'method' field is empty, this MethodConfig specifies the defaults
  //   for all methods for the specified service.
  // - If the 'service' field is empty, the 'method' field must be empty, and
  //   this MethodConfig specifies the default for all methods (it's the default
  //   config).
  //
  // When determining which MethodConfig to use for a given RPC, the most
  // specific match wins. For example, let's say that the service config
  // contains the following MethodConfig entries:
  //
  // method_config { name { } ... }
  // method_config { name { service: "MyService" } ... }
  // method_config { name { service: "MyService" method: "Foo" } ... }
  //
  // MyService/Foo will use the third entry, because it exactly matches the
  // service and method name. MyService/Bar will use the second entry, because
  // it provides the default for all methods of MyService. AnotherService/Baz
  // will use the first entry, because it doesn't match the other two.
  //
  // In JSON representation, value "", value `null`, and not present are the
  // same. The following are the same Name:
  // - { "service": "s" }
  // - { "service": "s", "method": null }
  // - { "service": "s", "method": "" }
  message Name {
    string service = 1;  // Required. Includes proto package name.
    string method = 2;
  }
  repeated Name name = 1;

  // Whether RPCs sent to this method should wait until the connection is
  // ready by default. If false, the RPC will abort immediately if there is
  // a transient failure connecting to the server. Otherwise, gRPC will
  // attempt to connect until the deadline is exceeded.
  //
  // The value specified via the gRPC client API will override the value
  // set here. However, note that setting the value in the client API will
  // also affect transient errors encountered during name resolution, which
  // cannot be caught by the value here, since the service config is
  // obtained by the gRPC client via name resolution.
  google.protobuf.BoolValue wait_for_ready = 2;

  // The default timeout in seconds for RPCs sent to this method. This can be
  // overridden in code. If no reply is received in the specified amount of
  // time, the request is aborted and a DEADLINE_EXCEEDED error status
  // is returned to the caller.
  //
  // The actual deadline used will be the minimum of the value specified here
  // and the value set by the application via the gRPC client API.  If either
  // one is not set, then the other will be used.  If neither is set, then the
  // request has no deadline.
  google.protobuf.Duration timeout = 3;

  // The maximum allowed payload size for an individual request or object in a
  // stream (client->server) in bytes. The size which is measured is the
  // serialized payload after per-message compression (but before stream
  // compression) in bytes. This applies both to streaming and non-streaming
  // requests.
  //
  // The actual value used is the minimum of the value specified here and the
  // value set by the application via the gRPC client API.  If either one is
  // not set, then the other will be used.  If neither is set, then the
  // built-in default is used.
  //
  // If a client attempts to send an object larger than this value, it will not
  // be sent and the client will see a ClientError.
  // Note that 0 is a valid value, meaning that the request message
  // must be empty.
  google.protobuf.UInt32Value max_request_message_bytes = 4;

  // The maximum allowed payload size for an individual response or object in a
  // stream (server->client) in bytes. The size which is measured is the
  // serialized payload after per-message compression (but before stream
  // compression) in bytes. This applies both to streaming and non-streaming
  // requests.
  //
  // The actual value used is the minimum of the value specified here and the
  // value set by the application via the gRPC client API.  If either one is
  // not set, then the other will be used.  If neither is set, then the
  // built-in default is used.
  //
  // If a server attempts to send an object larger than this value, it will not
  // be sent, and a ServerError will be sent to the client instead.
  // Note that 0 is a valid value, meaning that the response message
  // must be empty.
  google.protobuf.UInt32Value max_response_message_bytes = 5;

  // The retry policy for outgoing RPCs.
  message RetryPolicy {
    // The maximum number of RPC attempts, including the original attempt.
    //
    // This field is required and must be greater than 1.
    // Any value greater than 5 will be treated as if it were 5.
    uint32 max_attempts = 1;

    // Exponential backoff parameters. The initial retry attempt will occur at
    // random(0, initial_backoff). In general, the nth attempt will occur at
    // random(0,
    //   min(initial_backoff*backoff_multiplier**(n-1), max_backoff)).
    // Required. Must be greater than zero.
    google.protobuf.Duration initial_backoff = 2;
    // Required. Must be greater than zero.
    google.protobuf.Duration max_backoff = 3;
    float backoff_multiplier = 4;  // Required. Must be greater than zero.

    // The set of status codes which may be retried.
    //
    // This field is required and must be non-empty.
    repeated google.rpc.Code retryable_status_codes = 5;
  }
    // The hedging policy for outgoing RPCs. Hedged RPCs may execute more than
  // once on the server, so only idempotent methods should specify a hedging
  // policy.
  message HedgingPolicy {
    // The hedging policy will send up to max_requests RPCs.
    // This number represents the total number of all attempts, including
    // the original attempt.
    //
    // This field is required and must be greater than 1.
    // Any value greater than 5 will be treated as if it were 5.
    uint32 max_attempts = 1;

    // The first RPC will be sent immediately, but the max_requests-1 subsequent
    // hedged RPCs will be sent at intervals of every hedging_delay. Set this
    // to 0 to immediately send all max_requests RPCs.
    google.protobuf.Duration hedging_delay = 2;

    // The set of status codes which indicate other hedged RPCs may still
    // succeed. If a non-fatal status code is returned by the server, hedged
    // RPCs will continue. Otherwise, outstanding requests will be canceled and
    // the error returned to the client application layer.
    //
    // This field is optional.
    repeated google.rpc.Code non_fatal_status_codes = 3;
  }

  // Only one of retry_policy or hedging_policy may be set. If neither is set,
  // RPCs will not be retried or hedged.
  oneof retry_or_hedging_policy {
    RetryPolicy retry_policy = 6;
    HedgingPolicy hedging_policy = 7;
  }
}
```
转为JSON
```GO
{
		"methodConfig": [{
		  "name": [{"service": "xx.xxx","method":"xxx"}],
          "wait_for_ready": false,
          "timeout": 1000ms,
          "max_request_message_bytes": 1024,
          "max_response_message_bytes": 1024,
		  "retryPolicy": {
			  "maxAttempts": 4,
			  "initialBackoff": ".01s",
			  "maxBackoff": ".01s",
			  "backoffMultiplier": 1.0,
			  "retryableStatusCodes": [ "UNAVAILABLE" ]
		  },
		  "hedgingPolicy":{
              "maxAttempts":4,
              "hedgingDelay":"0.1s",
              "nonFatalStatusCodes": [ "" ]
          }}]
}
```
其中:
    
    name 指定下面的配置信息作用的 RPC 服务或方法
    service：通过服务名匹配，语法为<package>.<service>package就是proto文件中指定的package，service也是proto文件中指定的 Service Name。
    method：匹配具体某个方法，proto文件中定义的方法名。

    MaxAttempts：最大尝试次数
    InitialBackoff：默认退避时间
    MaxBackoff：最大退避时间
    BackoffMultiplier：退避时间增加倍率
    RetryableStatusCodes：服务端返回什么错误码才重试
    重试机制一般会搭配退避算法一起使用。

    即假设第一次请求失败后，等1秒（随便取的一个数）再次请求，又失败后就等2秒在请求，一直重试直达超过指定重试次数或者等待时间就不在重试。

    如果不使用退避算法，失败后就一直重试只会增加服务器的压力。如果是因为服务器压力大，导致的请求失败，那么根据退避算法等待一定时间后再次请求可能就能成功。反之直接请求可能会因为压力过大导致服务崩溃。

    第一次重试间隔是 random(0, initialBackoff)
    第 n 次的重试间隔为 random(0, min( initialBackoff*backoffMultiplier**(n-1) , maxBackoff))

## 重试策略
gRPC 的重试策略有两种分别是:
    
    重试(retryPolicy)
    对冲(hedging)

一个RPC方法只能配置一种重试策略。

    对冲是指在不等待响应的情况主动发送单次调用的多个请求，如果一个方法使用对冲策略，那么首先会像正常的 RPC 调用一样发送第一次请求，如果 hedgingDelay 时间内没有响应，那么直接发送第二次请求，以此类推，直到发送了 maxAttempts 次。

    对冲在超过指定时间没有响应就会直接发起请求，而重试则必须要服务端响应后才会发起请求。

    注意： 使用对冲的时候，请求可能会访问到不同的后端(如果设置了负载均衡)，那么就要求方法在多次执行下是安全，并且符合预期的
# grpc name_resolver
gRPC 中的默认 name-system 是 DNS，同时在客户端以插件形式提供了自定义 name-system 的机制

    例如：默认使用 DNS name-system，我们只需要提供服务器的域名即端口号，NameResolver 就会使用 DNS 解析出域名对应的 IP 列表并返回。

gRPC NameResolver 会根据 name-system 选择对应的解析器，用以解析用户提供的服务器名，最后返回具体地址列表（IP+端口号）。


1）客户端启动时，注册自定义的 resolver 。

    一般在 init() 方法，构造自定义的 resolveBuilder，并将其注册到 grpc 内部的 resolveBuilder 表中（其实是一个全局 map，key 为协议名，value 为构造的 resolveBuilder）。
2）客户端启动时通过自定义 Dail() 方法构造 grpc.ClientConn 单例

    grpc.DialContext() 方法内部解析 URI，分析协议类型，并从 resolveBuilder 表中查找协议对应的 resolverBuilder。
    找到指定的 resolveBuilder 后，调用 resolveBuilder 的 Build() 方法，构建自定义 resolver，同时开启协程，通过此 resolver 更新被调服务实例列表。
    Dial() 方法接收主调服务名和被调服务名，并根据自定义的协议名，基于这两个参数构造服务的 URI
    Dial() 方法内部使用构造的 URI，调用 grpc.DialContext() 方法对指定服务进行拨号
## name resolver

https://www.lixueduan.com/posts/grpc/11-name-resolver/

resolver 包括 ResolverBuilder 和 Resolver两个部分。
分别需要实现Builder和Resolver接口

Resolver 是整个功能最核心的代码，用于将服务名解析成对应实例。
Builder 则采用 Builder 模式在包初始化时创建并注册构造自定义 Resolver 实例。当客户端通过 Dial 方法对指定服务进行拨号时，grpc resolver 查找注册的 Builder 实例调用其 Build() 方法构建自定义 Resolver。
```GO
type Builder interface {
	Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)
	Scheme() string
}
type Resolver interface {
	ResolveNow(ResolveNowOptions)
	Close()
```
```Go
type CustomResolverBuilder struct{}
type customResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

const GrpcServiceScheme = "ns"
const GrpcServiceName = "resolver.scheme.hydracode.io"
const grpcServerAddress = "localhost:9000"

func (r *customResolver) ResolveNow(o resolver.ResolveNowOptions) {
	// 直接从map中取出对于的addrList
	addrStrs := r.addrsStore[r.target.Endpoint()]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (*customResolver) Close() {}

func (*CustomResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &customResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			GrpcServiceName: {grpcServerAddress},
		},
	}
	r.ResolveNow(resolver.ResolveNowOptions{})
	return r, nil
}
func (*CustomResolverBuilder) Scheme() string { return GrpcServiceScheme }
```

## 源码分析
1.客户端初始化的时候指定默认的解析为DNS
```GO
func defaultDialOptions() dialOptions {
	return dialOptions{
		copts: transport.ConnectOptions{
			ReadBufferSize:  defaultReadBufSize,
			WriteBufferSize: defaultWriteBufSize,
			UseProxy:        true,
			UserAgent:       grpcUA,
		},
		bs:              internalbackoff.DefaultExponential,
		healthCheckFunc: internal.HealthCheckFunc,
		idleTimeout:     30 * time.Minute,
		recvBufferPool:  nopBufferPool{},
		defaultScheme:   "dns",
	}
}
```
2.注册自定义resolver
```Go
func init() {
	resolver.Register(&name_resolver.CustomResolverBuilder{})
}
var (
	// m is a map from scheme to resolver builder.
	m = make(map[string]Builder)
	// defaultScheme is the default scheme to use.
	defaultScheme = "passthrough"
)

// TODO(bar) install dns resolver in init(){}.

// Register registers the resolver builder to the resolver map. b.Scheme will
// be used as the scheme registered with this builder. The registry is case
// sensitive, and schemes should not contain any uppercase characters.
//
// NOTE: this function must only be called during initialization time (i.e. in
// an init() function), and is not thread-safe. If multiple Resolvers are
// registered with the same name, the one registered last will take effect.
func Register(b Builder) {
	m[b.Scheme()] = b
}
func Get(scheme string) Builder {
	if b, ok := m[scheme]; ok {
		return b
	}
	return nil
}
```

3.解析传入的target,设置cc的resolverBuilder
```GO
if err := cc.parseTargetAndFindResolver(); err != nil {
		channelz.RemoveEntry(cc.channelz.ID)
		return nil, err
	}


func parseTarget(target string) (resolver.Target, error) {
	u, err := url.Parse(target)
	if err != nil {
		return resolver.Target{}, err
	}

	return resolver.Target{URL: *u}, nil
}
//	resolvers 类型是[]resolver.Builder
func (cc *ClientConn) getResolver(scheme string) resolver.Builder {
	for _, rb := range cc.dopts.resolvers {
		if scheme == rb.Scheme() {
			return rb
		}
	}
	return resolver.Get(scheme) //如果从cc中自带的的resolvers中找不到对应的scheme,就去	"google.golang.org/grpc/resolver" 这个包中的Get去找,那么就从上文中的m全局变量找
}
func (cc *ClientConn) parseTargetAndFindResolver() error {
	channelz.Infof(logger, cc.channelz, "original dial target is: %q", cc.target)

	var rb resolver.Builder
	parsedTarget, err := parseTarget(cc.target)
	if err != nil {
		channelz.Infof(logger, cc.channelz, "dial target %q parse failed: %v", cc.target, err)
	} else {
		channelz.Infof(logger, cc.channelz, "parsed dial target is: %#v", parsedTarget)
		rb = cc.getResolver(parsedTarget.URL.Scheme) 
		if rb != nil {
			cc.parsedTarget = parsedTarget
			cc.resolverBuilder = rb
			return nil
		}
	}

	// We are here because the user's dial target did not contain a scheme or
	// specified an unregistered scheme. We should fallback to the default
	// scheme, except when a custom dialer is specified in which case, we should
	// always use passthrough scheme. For either case, we need to respect any overridden
	// global defaults set by the user.
	defScheme := cc.dopts.defaultScheme
	if internal.UserSetDefaultScheme {
		defScheme = resolver.GetDefaultScheme()
	}

	channelz.Infof(logger, cc.channelz, "fallback to scheme %q", defScheme)
	canonicalTarget := defScheme + ":///" + cc.target

	parsedTarget, err = parseTarget(canonicalTarget)
	if err != nil {
		channelz.Infof(logger, cc.channelz, "dial target %q parse failed: %v", canonicalTarget, err)
		return err
	}
	channelz.Infof(logger, cc.channelz, "parsed dial target is: %+v", parsedTarget)
	rb = cc.getResolver(parsedTarget.URL.Scheme)
	if rb == nil {
		return fmt.Errorf("could not get resolver for default scheme: %q", parsedTarget.URL.Scheme)
	}
	cc.parsedTarget = parsedTarget
	cc.resolverBuilder = rb
	return nil
}
```
# grpc实践
```md
文件目录
rpc_demo
├── go.mod
├── go.sum
├── proto
│   ├── info.proto
│   └── meta
│       └── meta.proto

其中go.mod 中定义的module为 module rpc_demo
info.proto中引入了meta.proto中的message
```
其中meta.proto文件内容:

```protobuf

syntax="proto3";

package meta;

option go_package="rpc_demo/rpc/types/meta;meta";

message req{
    int32 id=1;
}
message resp{
    int32 id=1;
    string name=2;
}
```

其中info.proto文件内容:

```protobuf
syntax="proto3";

package info;

option go_package="rpc_demo/rpc/types/info;info";

//import 的路经和protoc -I参数组合,建议从根目录写
import "rpc_dmeo/proto/meta/meta.proto";
service infoRpc {
    rpc infoQuey(meta.req) returns(meta.resp);
}
```
需要注意的是:
    
    1.pacakge为protobuf自己的package,而不是go的package,可以在一个文件目录下有多个不同的package申明

    2.go_package的写法为"{out_path};out_go_package",其中out_path表示的
        1.其他go中要引入生成的pb.go对应的package的文件的import路经
        2.实际生成的pb.go文件的存放路经


由于我们定义的go_module为rpc_demo,且根目录文件夹名也叫rpc_demo(根目录名称和module名称约定一致)
那么如果我们要在该目录下面定义个package,如在根目录下面创建一个package,路经为rpc_demo/utils/xxx.go 其中xxx.go的package为utils则
在根目录下的main.go下要引入该package的import语法是

    import "rpc_demo/utils"

因此我们可知,proto文件中go_package的写法,应该从定义的go module所在的根目录开始写,如下:
    
    option go_package="rpc_demo/utils;utils"


总结:
    1.go module应该和和根文件夹名称一致
    2.protobuf中的import推荐从根目录开始写起
    3.protoc 的执行路经为根目录的上一级目录,生成grpc的标准的写法为

        1.切换到根目录的上一级目录
        2.执行protoc -I . --go_out=. --go-grpc_out=. xxx/xxx/xxx.proto

执行命令参数解析

protoc -I xxx --go_out=xxx



命令解析：
-I 
    如果多个proto文件之间有互相依赖，生成某个proto文件时，需要import其他几个proto文件，这时候就要用-I来指定搜索目录,指定的路经和proto文件中import的路经拼接一起作为导入的proto文件的查询路经
--go_out 指
    定我们生成的目录
    
操作总结为:

    1.切换到根目录的上级目录
    2.执行:
        protoc -I . --go_out=. --go-grpc_out=. rpc_demo/proto/info.proto
        protoc -I . --go_out=. --go-grpc_out=. rpc_demo/proto/meta/meta.proto

最终生成的结果为
```
rpc_demo
├── go.mod
├── go.sum
├── proto
│   ├── info.proto
│   └── meta
│       └── meta.proto
└── rpc
    └── types
        ├── info
        │   ├── info_grpc.pb.go
        │   └── info.pb.go
        └── meta
            └── meta.pb.go
```
