


切换到根目录的上级目录
执行:

protoc -I . --go_out=. --go-grpc_out=. grpc_with_interceptor/proto/meta.proto

流式拦截器见grpc_server_client_demo_with_tls