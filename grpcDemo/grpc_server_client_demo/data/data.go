package data

import (
	"fmt"
	"grpc_server_client_demo/types"

	"google.golang.org/protobuf/proto"
)

func DataHandle() {
	req := &types.Request{Id: 1}
	marshalMsg, err := proto.Marshal(req)
	if err != nil {
		fmt.Println(marshalMsg)
	} else {
		fmt.Println(marshalMsg)
	}
	data := &types.Response{}
	err = proto.Unmarshal(marshalMsg, data)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(data)
	}
}
