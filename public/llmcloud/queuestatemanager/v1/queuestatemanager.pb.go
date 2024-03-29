// Copyright 2023 Groq Inc.
// All Rights Reserved.
//
// Definition of the QueueManager API Service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: public/llmcloud/queuestatemanager/v1/queuestatemanager.proto

package queuestatemanagerv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequestStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the ID of the request
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequestStateRequest) Reset() {
	*x = GetRequestStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestStateRequest) ProtoMessage() {}

func (x *GetRequestStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestStateRequest.ProtoReflect.Descriptor instead.
func (*GetRequestStateRequest) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequestStateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequestStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the ID of the request
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// How many requests are in front of this request in the queue
	PlaceInQueue int32 `protobuf:"varint,2,opt,name=place_in_queue,json=placeInQueue,proto3" json:"place_in_queue,omitempty"`
}

func (x *GetRequestStateResponse) Reset() {
	*x = GetRequestStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequestStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestStateResponse) ProtoMessage() {}

func (x *GetRequestStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestStateResponse.ProtoReflect.Descriptor instead.
func (*GetRequestStateResponse) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequestStateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetRequestStateResponse) GetPlaceInQueue() int32 {
	if x != nil {
		return x.PlaceInQueue
	}
	return 0
}

var File_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto protoreflect.FileDescriptor

var file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x30, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x4f, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x32, 0xdf, 0x01, 0x0a, 0x18, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xc2, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x3c, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x12, 0x2a, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0xcc, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x16, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74,
	0x2e, 0x67, 0x72, 0x6f, 0x71, 0x2e, 0x69, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c,
	0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x50, 0x4c, 0x51, 0xaa, 0x02, 0x24, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x25,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x31, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c,
	0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x27, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x3a, 0x3a, 0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescOnce sync.Once
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescData = file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDesc
)

func file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescGZIP() []byte {
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescOnce.Do(func() {
		file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescData)
	})
	return file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDescData
}

var file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_goTypes = []interface{}{
	(*GetRequestStateRequest)(nil),  // 0: public.llmcloud.queuestatemanager.v1.GetRequestStateRequest
	(*GetRequestStateResponse)(nil), // 1: public.llmcloud.queuestatemanager.v1.GetRequestStateResponse
}
var file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_depIdxs = []int32{
	0, // 0: public.llmcloud.queuestatemanager.v1.QueueStateManagerService.GetRequestState:input_type -> public.llmcloud.queuestatemanager.v1.GetRequestStateRequest
	1, // 1: public.llmcloud.queuestatemanager.v1.QueueStateManagerService.GetRequestState:output_type -> public.llmcloud.queuestatemanager.v1.GetRequestStateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_init() }
func file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_init() {
	if File_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestStateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequestStateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_goTypes,
		DependencyIndexes: file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_depIdxs,
		MessageInfos:      file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_msgTypes,
	}.Build()
	File_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto = out.File
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_rawDesc = nil
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_goTypes = nil
	file_public_llmcloud_queuestatemanager_v1_queuestatemanager_proto_depIdxs = nil
}
