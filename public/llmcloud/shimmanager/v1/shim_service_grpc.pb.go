// Copyright 2023 Groq Inc.
// All Rights Reserved.
//
// Definition of the Shim Service
// This is the inverse of the Inference API Service.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: public/llmcloud/shimmanager/v1/shim_service.proto

package shimmanagerv1

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
	ShimService_GetQuery_FullMethodName = "/public.llmcloud.shimmanager.v1.ShimService/GetQuery"
)

// ShimServiceClient is the client API for ShimService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShimServiceClient interface {
	// returns the query at the the top of the priority queue
	GetQuery(ctx context.Context, in *ShimServiceGetQueryRequest, opts ...grpc.CallOption) (*ShimServiceGetQueryResponse, error)
}

type shimServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShimServiceClient(cc grpc.ClientConnInterface) ShimServiceClient {
	return &shimServiceClient{cc}
}

func (c *shimServiceClient) GetQuery(ctx context.Context, in *ShimServiceGetQueryRequest, opts ...grpc.CallOption) (*ShimServiceGetQueryResponse, error) {
	out := new(ShimServiceGetQueryResponse)
	err := c.cc.Invoke(ctx, ShimService_GetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShimServiceServer is the server API for ShimService service.
// All implementations must embed UnimplementedShimServiceServer
// for forward compatibility
type ShimServiceServer interface {
	// returns the query at the the top of the priority queue
	GetQuery(context.Context, *ShimServiceGetQueryRequest) (*ShimServiceGetQueryResponse, error)
	mustEmbedUnimplementedShimServiceServer()
}

// UnimplementedShimServiceServer must be embedded to have forward compatible implementations.
type UnimplementedShimServiceServer struct {
}

func (UnimplementedShimServiceServer) GetQuery(context.Context, *ShimServiceGetQueryRequest) (*ShimServiceGetQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuery not implemented")
}
func (UnimplementedShimServiceServer) mustEmbedUnimplementedShimServiceServer() {}

// UnsafeShimServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShimServiceServer will
// result in compilation errors.
type UnsafeShimServiceServer interface {
	mustEmbedUnimplementedShimServiceServer()
}

func RegisterShimServiceServer(s grpc.ServiceRegistrar, srv ShimServiceServer) {
	s.RegisterService(&ShimService_ServiceDesc, srv)
}

func _ShimService_GetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShimServiceGetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).GetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShimService_GetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).GetQuery(ctx, req.(*ShimServiceGetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShimService_ServiceDesc is the grpc.ServiceDesc for ShimService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShimService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.llmcloud.shimmanager.v1.ShimService",
	HandlerType: (*ShimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuery",
			Handler:    _ShimService_GetQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/llmcloud/shimmanager/v1/shim_service.proto",
}