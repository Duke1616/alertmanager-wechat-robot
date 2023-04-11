// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: apps/target/pb/target.proto

package target

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	// 创建消息通知群组信息
	CreateTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Target, error)
	// 查看群组信息
	DescribeTarget(ctx context.Context, in *DescribeTargetRequest, opts ...grpc.CallOption) (*Target, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) CreateTarget(ctx context.Context, in *Target, opts ...grpc.CallOption) (*Target, error) {
	out := new(Target)
	err := c.cc.Invoke(ctx, "/robot.target.RPC/CreateTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DescribeTarget(ctx context.Context, in *DescribeTargetRequest, opts ...grpc.CallOption) (*Target, error) {
	out := new(Target)
	err := c.cc.Invoke(ctx, "/robot.target.RPC/DescribeTarget", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	// 创建消息通知群组信息
	CreateTarget(context.Context, *Target) (*Target, error)
	// 查看群组信息
	DescribeTarget(context.Context, *DescribeTargetRequest) (*Target, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) CreateTarget(context.Context, *Target) (*Target, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTarget not implemented")
}
func (UnimplementedRPCServer) DescribeTarget(context.Context, *DescribeTargetRequest) (*Target, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeTarget not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_CreateTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Target)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).CreateTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.target.RPC/CreateTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).CreateTarget(ctx, req.(*Target))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_DescribeTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).DescribeTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/robot.target.RPC/DescribeTarget",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).DescribeTarget(ctx, req.(*DescribeTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "robot.target.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTarget",
			Handler:    _RPC_CreateTarget_Handler,
		},
		{
			MethodName: "DescribeTarget",
			Handler:    _RPC_DescribeTarget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/target/pb/target.proto",
}
