切换到根目录的上级目录
执行:

protoc -I . --go_out=. --go-grpc_out=. grpc_with_ctx/proto/meta.proto
