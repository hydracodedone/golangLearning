切换到根目录的上级目录
执行:

protoc -I . --go_out=. --go-grpc_out=. grpc_server_client_demo_with_stream/proto/meta.proto
