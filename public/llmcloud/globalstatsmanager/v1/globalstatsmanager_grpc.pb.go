// Copyright 2023 Groq Inc.
// All Rights Reserved.
//
// Definition of the Inference API Service.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: public/llmcloud/globalstatsmanager/v1/globalstatsmanager.proto

package globalstatsmanagerv1

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
	GlobalStatsManagerService_GetGlobalStats_FullMethodName = "/public.llmcloud.globalstatsmanager.v1.GlobalStatsManagerService/GetGlobalStats"
)

// GlobalStatsManagerServiceClient is the client API for GlobalStatsManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GlobalStatsManagerServiceClient interface {
	GetGlobalStats(ctx context.Context, in *GetGlobalStatsRequest, opts ...grpc.CallOption) (*GetGlobalStatsResponse, error)
}

type globalStatsManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGlobalStatsManagerServiceClient(cc grpc.ClientConnInterface) GlobalStatsManagerServiceClient {
	return &globalStatsManagerServiceClient{cc}
}

func (c *globalStatsManagerServiceClient) GetGlobalStats(ctx context.Context, in *GetGlobalStatsRequest, opts ...grpc.CallOption) (*GetGlobalStatsResponse, error) {
	out := new(GetGlobalStatsResponse)
	err := c.cc.Invoke(ctx, GlobalStatsManagerService_GetGlobalStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GlobalStatsManagerServiceServer is the server API for GlobalStatsManagerService service.
// All implementations must embed UnimplementedGlobalStatsManagerServiceServer
// for forward compatibility
type GlobalStatsManagerServiceServer interface {
	GetGlobalStats(context.Context, *GetGlobalStatsRequest) (*GetGlobalStatsResponse, error)
	mustEmbedUnimplementedGlobalStatsManagerServiceServer()
}

// UnimplementedGlobalStatsManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGlobalStatsManagerServiceServer struct {
}

func (UnimplementedGlobalStatsManagerServiceServer) GetGlobalStats(context.Context, *GetGlobalStatsRequest) (*GetGlobalStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalStats not implemented")
}
func (UnimplementedGlobalStatsManagerServiceServer) mustEmbedUnimplementedGlobalStatsManagerServiceServer() {
}

// UnsafeGlobalStatsManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GlobalStatsManagerServiceServer will
// result in compilation errors.
type UnsafeGlobalStatsManagerServiceServer interface {
	mustEmbedUnimplementedGlobalStatsManagerServiceServer()
}

func RegisterGlobalStatsManagerServiceServer(s grpc.ServiceRegistrar, srv GlobalStatsManagerServiceServer) {
	s.RegisterService(&GlobalStatsManagerService_ServiceDesc, srv)
}

func _GlobalStatsManagerService_GetGlobalStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GlobalStatsManagerServiceServer).GetGlobalStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GlobalStatsManagerService_GetGlobalStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GlobalStatsManagerServiceServer).GetGlobalStats(ctx, req.(*GetGlobalStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GlobalStatsManagerService_ServiceDesc is the grpc.ServiceDesc for GlobalStatsManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GlobalStatsManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "public.llmcloud.globalstatsmanager.v1.GlobalStatsManagerService",
	HandlerType: (*GlobalStatsManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGlobalStats",
			Handler:    _GlobalStatsManagerService_GetGlobalStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public/llmcloud/globalstatsmanager/v1/globalstatsmanager.proto",
}