# RPC
rpc，全称 remote process call（远程过程调用），是微服务架构下的一种通信模式. 这种通信模式下，一台服务器在调用远程机器的接口时，能够获得像调用本地方法一样的良好体验.

    rpc 调用基于 sdk 方式，调用方法和出入参协议固定，stub 文件本身还能起到接口文档的作用，很大程度上优化了通信双方约定协议达成共识的成本.
    rpc 在传输层协议 tcp 基础之上，可以由实现框架自定义填充应用层协议细节，理论上存在着更高的上限
## GRPC
gRPC  是一个高性能、开源和通用的 RPC 框架，面向移动和 HTTP/2 设计。目前提供 C、Java 和 Go 语言版本，分别是：grpc, grpc-java, grpc-go.

gRPC 基于 HTTP/2 标准设计，带来诸如双向流、流控、头部压缩、单 TCP 连接上的多复用请求等特。这些特性使得其在移动设备上表现更好，更省电和节省空间占用。

## proto buffers

protocol buffers，是一套结构数据序列化机制（当然也可以使用其他数据格式如 JSON）用 proto files 创建 gRPC 服务，用 protocol buffers 消息类型来定义方法参数和返回类型

    代码生成
    序列化与反序列化

## 环境安装

### 安装 grpc

    go get google.golang.org/grpc@latest
### 安装 protocol buffer

根据操作系统型号，下载安装好对应版本的 protobuf 应用：

https://github.com/google/protobuf/releases

需要将 protobuf 执行文件所在的目录添加到环境变量 $PATH 当中.

安装完成后，可以通过查看 protobuf 版本指令，校验安装是否成功

    protoc --version
### 安装插件protoc-gen-go

不要使用 github.com/golang/protobuf/protoc-gen-go 这个版本
安装 protobuf -> pb.go 插件

    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

该插件的作用是，能够基于 .proto 文件一键生成 _pb.go 文件，对应内容为通信请求/响应参数的对象模型.

go install 指令默认会将插件安装到 GOPATH/bin 目录下. 需要确保 GOPATH/bin 路径有被添加到环境路径 $PATH 当中.

安装完成后，可以通过查看插件版本指令，校验安装是否成功

    protoc-gen-go --version

### 安装插件protoc-gen-go-grpc
安装 protobuf -> grpc.pb.go 插件

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

该插件的作用是，能够基于 .proto 文件生成 _grpc.pb.go，对应内容为通信服务框架代码.

安装完成后，可以通过查看插件版本指令，校验安装是否成功

    protoc-gen-go-grpc --version


## 使用
### proto文件编写
正如其他 RPC 系统，gRPC 基于如下思想：定义一个服务， 指定其可以被远程调用的方法及其参数和返回类型。gRPC 默认使用 protocol buffers 作为接口定义语言，来描述服务接口和有效载荷消息结构

```
service HelloService {
  rpc SayHello (HelloRequest) returns (HelloResponse);
}

message HelloRequest {
  required string greeting = 1;
}

message HelloResponse {
  required string reply = 1;
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
### proto3 注意事项
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

### 代码生成

    protoc --go_out=. xxx.proto
    protoc --go-grpc_out=. xxx.proto


### 注意事项

    package 关键字用于proto,在引用的时起作用
        package 声明的包名是proto的包名,用于proto语法解析import时候调用

    option go_package用于生成的.pb.go文件,在引用和生成go包名时起作用
        option go_package = "{out_path};out_go_package"
        前一个参数用于指定生成文件的位置，后一个参数指定生成的 .go 文件的 package
        这里指定的 out_path 并不是绝对路径，只是相对路径或者说只是路径的一部分，和 protoc 的 --go_out 拼接后才是完整的路径。
        第一个参数也会用于生成的go代码的import


### 实践
基于下列结构生成rpc项目
    
    1.整个项目mod为grpc_demo
    2.有2个模块,分别为data和user
    3.data模块有两个proto,分别为data.proto和data2.proto,认为这两个proto是一个package,记为 proto package data
    4.user模块有一个proto,记为proto package user
    5.user的proto依赖于data的两个proto 
```
.
├── data
│   ├── proto
│   │   ├── data2.proto
│   │   └── data.proto
│   └── data_types
│       ├── data2.pb.go
│       └── data.pb.go
├── go.mod
├── go.sum
└── user
    ├── proto
    │   └── user.proto
    └── user_types
        └── user.pb.go

```
步骤1:

    生成data.proto和data2.proto,需要注意的是
        data.proto和data2.proto的package 均为 data,表示这两个proto文件是一个包
        data.proto和data2.proto生成的pb.go文件对应包都是data,这个参数由 option go_packge="data/data_types;data_types"中的第二项控制
        option go_packge="xxx;data"中的第一项 data/data_types有两个用途 
            1 配合protoc 中的 --go_out="yyy" 路径来生成绝对路径为"yyy/data/data_types"的pb.go文件 
            2 在其他引入了该proto的proto生成pb.go代码时候,go代码的倒入原proto pb.go中的路径
        !!!建议go_packge 中包名和路径名一致,如data/data_types;data_types"中包名data_types!!!
步骤2:

    在go.mod的路径下执行:
        protoc --go_out="." /data/proto/data.proto
        protoc --go_out="." /data/proto/data2.proto

步骤3:

    生成user.proto,需要注意的是
        import "data/proto/data2.proto";
        import "data/proto/data.proto";
        倒入不能使用相对路径,不然会报" Backslashes, consecutive slashes, ".", or ".." are not allowed in the virtual path"因此import路径要从根路径出发,只有这样才能导入data的绝对路径
        导入后的使用则是按照导入包的package name进行使用

步骤4:

    在go.mod的路径下执行:
        protoc --go_out="." /user/proto/user.proto

生成的user.pb.go中会有一处错误
```Go
import (
	types "data_types/types" //could not import "data/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)
```
原因是初始化了go mod,需要手动修改,添加go mod 定义的mod 名
```Go
import (
	types "grpc_demo/data_types/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)
```

步骤5:

    sesrver端的编写

步骤6: 

    client端的编写


## 安全性

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


### 密钥生成

    生成私钥
        openssl genrsa -out server.key 2048
    生成证书
        openssl req -new -x509 -key server.key -out server.crt -days 36500
    生成csr(证书签名请求)
        openssl req -new -key server.key -out server.csr