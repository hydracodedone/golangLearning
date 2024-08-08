切换到protobuf_demo0的目录的上一级,执行:

protoc -I .  --go_out=.  --go-grpc_out=.  protobuf_demo0/user/proto/user.proto
-I 表示搜索路经为当前路经 ../protobuf_demo
而import path 为
protobuf_demo/data/proto/data2.proto
那么搜索路经为

../protobuf_demo/protobuf_demo/data/proto/data2.proto


由于引入了go.mod而当前的package名字为protobuf_demo,如果该go package 下面的package在被引用的时候,导入的写法为
import "protobuf_demo/xxx"
而我们定义
go_package="protobuf_demo/user/user_types;user_types";
是符合规范的

需要注意的是go_package也是生成的pb.go文件的目录,我们指定了最开始的为go module的名称,因此,生成protobuf应该从go_module的上级目录开始生成!


根目录的上层
protoc -I .  --go_out=.  --go-grpc_out=.  protobuf_demo/user/proto/user.proto
protoc -I .  --go_out=.  --go-grpc_out=.  protobuf_demo/data/proto/data.proto
protoc -I .  --go_out=.  --go-grpc_out=.  protobuf_demo/data/proto/data2.proto
