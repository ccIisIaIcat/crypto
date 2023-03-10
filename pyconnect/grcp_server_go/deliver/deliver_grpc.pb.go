// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0
// source: deliver.proto

package deliver

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

// BarDataRevicerClient is the client API for BarDataRevicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BarDataRevicerClient interface {
	BarDataRevicer(ctx context.Context, in *BarData, opts ...grpc.CallOption) (*Response, error)
}

type barDataRevicerClient struct {
	cc grpc.ClientConnInterface
}

func NewBarDataRevicerClient(cc grpc.ClientConnInterface) BarDataRevicerClient {
	return &barDataRevicerClient{cc}
}

func (c *barDataRevicerClient) BarDataRevicer(ctx context.Context, in *BarData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/BarDataRevicer/BarDataRevicer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BarDataRevicerServer is the server API for BarDataRevicer service.
// All implementations must embed UnimplementedBarDataRevicerServer
// for forward compatibility
type BarDataRevicerServer interface {
	BarDataRevicer(context.Context, *BarData) (*Response, error)
	mustEmbedUnimplementedBarDataRevicerServer()
}

// UnimplementedBarDataRevicerServer must be embedded to have forward compatible implementations.
type UnimplementedBarDataRevicerServer struct {
}

func (UnimplementedBarDataRevicerServer) BarDataRevicer(context.Context, *BarData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BarDataRevicer not implemented")
}
func (UnimplementedBarDataRevicerServer) mustEmbedUnimplementedBarDataRevicerServer() {}

// UnsafeBarDataRevicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BarDataRevicerServer will
// result in compilation errors.
type UnsafeBarDataRevicerServer interface {
	mustEmbedUnimplementedBarDataRevicerServer()
}

func RegisterBarDataRevicerServer(s grpc.ServiceRegistrar, srv BarDataRevicerServer) {
	s.RegisterService(&BarDataRevicer_ServiceDesc, srv)
}

func _BarDataRevicer_BarDataRevicer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BarDataRevicerServer).BarDataRevicer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BarDataRevicer/BarDataRevicer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BarDataRevicerServer).BarDataRevicer(ctx, req.(*BarData))
	}
	return interceptor(ctx, in, info, handler)
}

// BarDataRevicer_ServiceDesc is the grpc.ServiceDesc for BarDataRevicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BarDataRevicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BarDataRevicer",
	HandlerType: (*BarDataRevicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BarDataRevicer",
			Handler:    _BarDataRevicer_BarDataRevicer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deliver.proto",
}

// TickDataRevicerClient is the client API for TickDataRevicer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TickDataRevicerClient interface {
	TickDataRevicer(ctx context.Context, in *TickData, opts ...grpc.CallOption) (*Response, error)
}

type tickDataRevicerClient struct {
	cc grpc.ClientConnInterface
}

func NewTickDataRevicerClient(cc grpc.ClientConnInterface) TickDataRevicerClient {
	return &tickDataRevicerClient{cc}
}

func (c *tickDataRevicerClient) TickDataRevicer(ctx context.Context, in *TickData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/TickDataRevicer/TickDataRevicer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TickDataRevicerServer is the server API for TickDataRevicer service.
// All implementations must embed UnimplementedTickDataRevicerServer
// for forward compatibility
type TickDataRevicerServer interface {
	TickDataRevicer(context.Context, *TickData) (*Response, error)
	mustEmbedUnimplementedTickDataRevicerServer()
}

// UnimplementedTickDataRevicerServer must be embedded to have forward compatible implementations.
type UnimplementedTickDataRevicerServer struct {
}

func (UnimplementedTickDataRevicerServer) TickDataRevicer(context.Context, *TickData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TickDataRevicer not implemented")
}
func (UnimplementedTickDataRevicerServer) mustEmbedUnimplementedTickDataRevicerServer() {}

// UnsafeTickDataRevicerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TickDataRevicerServer will
// result in compilation errors.
type UnsafeTickDataRevicerServer interface {
	mustEmbedUnimplementedTickDataRevicerServer()
}

func RegisterTickDataRevicerServer(s grpc.ServiceRegistrar, srv TickDataRevicerServer) {
	s.RegisterService(&TickDataRevicer_ServiceDesc, srv)
}

func _TickDataRevicer_TickDataRevicer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TickData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TickDataRevicerServer).TickDataRevicer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TickDataRevicer/TickDataRevicer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TickDataRevicerServer).TickDataRevicer(ctx, req.(*TickData))
	}
	return interceptor(ctx, in, info, handler)
}

// TickDataRevicer_ServiceDesc is the grpc.ServiceDesc for TickDataRevicer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TickDataRevicer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TickDataRevicer",
	HandlerType: (*TickDataRevicerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TickDataRevicer",
			Handler:    _TickDataRevicer_TickDataRevicer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deliver.proto",
}
