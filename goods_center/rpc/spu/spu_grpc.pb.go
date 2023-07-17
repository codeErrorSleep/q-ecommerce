// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: spu.proto

package spu

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

const (
	Spu_GetSpuInfo_FullMethodName = "/spu.Spu/GetSpuInfo"
)

// SpuClient is the client API for Spu service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpuClient interface {
	GetSpuInfo(ctx context.Context, in *GetSpuInfoRequest, opts ...grpc.CallOption) (*GetSpuInfoResponse, error)
}

type spuClient struct {
	cc grpc.ClientConnInterface
}

func NewSpuClient(cc grpc.ClientConnInterface) SpuClient {
	return &spuClient{cc}
}

func (c *spuClient) GetSpuInfo(ctx context.Context, in *GetSpuInfoRequest, opts ...grpc.CallOption) (*GetSpuInfoResponse, error) {
	out := new(GetSpuInfoResponse)
	err := c.cc.Invoke(ctx, Spu_GetSpuInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpuServer is the server API for Spu service.
// All implementations must embed UnimplementedSpuServer
// for forward compatibility
type SpuServer interface {
	GetSpuInfo(context.Context, *GetSpuInfoRequest) (*GetSpuInfoResponse, error)
	mustEmbedUnimplementedSpuServer()
}

// UnimplementedSpuServer must be embedded to have forward compatible implementations.
type UnimplementedSpuServer struct {
}

func (UnimplementedSpuServer) GetSpuInfo(context.Context, *GetSpuInfoRequest) (*GetSpuInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpuInfo not implemented")
}
func (UnimplementedSpuServer) mustEmbedUnimplementedSpuServer() {}

// UnsafeSpuServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpuServer will
// result in compilation errors.
type UnsafeSpuServer interface {
	mustEmbedUnimplementedSpuServer()
}

func RegisterSpuServer(s grpc.ServiceRegistrar, srv SpuServer) {
	s.RegisterService(&Spu_ServiceDesc, srv)
}

func _Spu_GetSpuInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpuInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpuServer).GetSpuInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Spu_GetSpuInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpuServer).GetSpuInfo(ctx, req.(*GetSpuInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Spu_ServiceDesc is the grpc.ServiceDesc for Spu service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Spu_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spu.Spu",
	HandlerType: (*SpuServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSpuInfo",
			Handler:    _Spu_GetSpuInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spu.proto",
}
