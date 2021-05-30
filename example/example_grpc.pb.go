// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package expb

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

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	ExampleFiringRpc(ctx context.Context, in *ExampleRpcRequest, opts ...grpc.CallOption) (*ExampleRpcResponse, error)
	ExampleSilentRpc(ctx context.Context, in *ExampleRpcRequest, opts ...grpc.CallOption) (*ExampleRpcResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) ExampleFiringRpc(ctx context.Context, in *ExampleRpcRequest, opts ...grpc.CallOption) (*ExampleRpcResponse, error) {
	out := new(ExampleRpcResponse)
	err := c.cc.Invoke(ctx, "/example.ExampleService/ExampleFiringRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) ExampleSilentRpc(ctx context.Context, in *ExampleRpcRequest, opts ...grpc.CallOption) (*ExampleRpcResponse, error) {
	out := new(ExampleRpcResponse)
	err := c.cc.Invoke(ctx, "/example.ExampleService/ExampleSilentRpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	ExampleFiringRpc(context.Context, *ExampleRpcRequest) (*ExampleRpcResponse, error)
	ExampleSilentRpc(context.Context, *ExampleRpcRequest) (*ExampleRpcResponse, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) ExampleFiringRpc(context.Context, *ExampleRpcRequest) (*ExampleRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleFiringRpc not implemented")
}
func (UnimplementedExampleServiceServer) ExampleSilentRpc(context.Context, *ExampleRpcRequest) (*ExampleRpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExampleSilentRpc not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_ExampleFiringRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleFiringRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ExampleService/ExampleFiringRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleFiringRpc(ctx, req.(*ExampleRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_ExampleSilentRpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleRpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).ExampleSilentRpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.ExampleService/ExampleSilentRpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).ExampleSilentRpc(ctx, req.(*ExampleRpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExampleFiringRpc",
			Handler:    _ExampleService_ExampleFiringRpc_Handler,
		},
		{
			MethodName: "ExampleSilentRpc",
			Handler:    _ExampleService_ExampleSilentRpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/example.proto",
}
