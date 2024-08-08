// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: meta.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RPCDemoClient is the client API for RPCDemo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCDemoClient interface {
	DemoRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type rPCDemoClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCDemoClient(cc grpc.ClientConnInterface) RPCDemoClient {
	return &rPCDemoClient{cc}
}

func (c *rPCDemoClient) DemoRequest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/meta.RPCDemo/DemoRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCDemoServer is the server API for RPCDemo service.
// All implementations must embed UnimplementedRPCDemoServer
// for forward compatibility
type RPCDemoServer interface {
	DemoRequest(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedRPCDemoServer()
}

// UnimplementedRPCDemoServer must be embedded to have forward compatible implementations.
type UnimplementedRPCDemoServer struct {
}

func (UnimplementedRPCDemoServer) DemoRequest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoRequest not implemented")
}
func (UnimplementedRPCDemoServer) mustEmbedUnimplementedRPCDemoServer() {}

// UnsafeRPCDemoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCDemoServer will
// result in compilation errors.
type UnsafeRPCDemoServer interface {
	mustEmbedUnimplementedRPCDemoServer()
}

func RegisterRPCDemoServer(s grpc.ServiceRegistrar, srv RPCDemoServer) {
	s.RegisterService(&RPCDemo_ServiceDesc, srv)
}

func _RPCDemo_DemoRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCDemoServer).DemoRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta.RPCDemo/DemoRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCDemoServer).DemoRequest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// RPCDemo_ServiceDesc is the grpc.ServiceDesc for RPCDemo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPCDemo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta.RPCDemo",
	HandlerType: (*RPCDemoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DemoRequest",
			Handler:    _RPCDemo_DemoRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta.proto",
}
