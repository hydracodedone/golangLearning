切换到根目录的上级目录
执行:

protoc -I ./grpc_gateway/proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. grpc_gateway/proto/meta.proto

annnotations.proto 文件下载路经
https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto