// Copyright 2023 Groq Inc.
// All Rights Reserved.
//
// Definition of the Inference API Service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: public/llmcloud/requestmanager/v1/requestmanager.proto

package requestmanagerv1

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

type GetTextCompletionStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// request_id will only be included in the first response
	RequestId *string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3,oneof" json:"request_id,omitempty"`
	// the newly generated token data
	Content *string           `protobuf:"bytes,2,opt,name=content,proto3,oneof" json:"content,omitempty"`
	Stats   *PerformanceStats `protobuf:"bytes,3,opt,name=stats,proto3,oneof" json:"stats,omitempty"`
}

func (x *GetTextCompletionStreamResponse) Reset() {
	*x = GetTextCompletionStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextCompletionStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextCompletionStreamResponse) ProtoMessage() {}

func (x *GetTextCompletionStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextCompletionStreamResponse.ProtoReflect.Descriptor instead.
func (*GetTextCompletionStreamResponse) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP(), []int{0}
}

func (x *GetTextCompletionStreamResponse) GetRequestId() string {
	if x != nil && x.RequestId != nil {
		return *x.RequestId
	}
	return ""
}

func (x *GetTextCompletionStreamResponse) GetContent() string {
	if x != nil && x.Content != nil {
		return *x.Content
	}
	return ""
}

func (x *GetTextCompletionStreamResponse) GetStats() *PerformanceStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type GetTextCompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the UID of the model to run this request against
	ModelId string `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	// the prompt that the system will use to define its response behavior
	SystemPrompt *string `protobuf:"bytes,2,opt,name=system_prompt,json=systemPrompt,proto3,oneof" json:"system_prompt,omitempty"`
	// All of the messages in this chat so far
	History []*GetTextCompletionRequest_HistoryMessage `protobuf:"bytes,3,rep,name=history,proto3" json:"history,omitempty"`
	// the prompt provided by the user
	UserPrompt string `protobuf:"bytes,4,opt,name=user_prompt,json=userPrompt,proto3" json:"user_prompt,omitempty"`
	// the seed to use when generating tokens
	// Randomization seed for the RNG when doing distribution sampling
	Seed *uint32 `protobuf:"varint,5,opt,name=seed,proto3,oneof" json:"seed,omitempty"`
	// the maximum amount of tokens to generate before terminating
	MaxTokens *uint32 `protobuf:"varint,6,opt,name=max_tokens,json=maxTokens,proto3,oneof" json:"max_tokens,omitempty"`
	// controls the randomness of the response
	// Lower temperature is lower randomness
	// range(0,2]
	Temperature *float32 `protobuf:"fixed32,7,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// controls the top percentile of tokens that are eligible to be sampled
	// range[0,1]
	TopP *float32 `protobuf:"fixed32,8,opt,name=top_p,json=topP,proto3,oneof" json:"top_p,omitempty"`
	// controls the top k most likely tokens that are eligible to be sampled
	TopK *uint32 `protobuf:"varint,9,opt,name=top_k,json=topK,proto3,oneof" json:"top_k,omitempty"`
}

func (x *GetTextCompletionRequest) Reset() {
	*x = GetTextCompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextCompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextCompletionRequest) ProtoMessage() {}

func (x *GetTextCompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextCompletionRequest.ProtoReflect.Descriptor instead.
func (*GetTextCompletionRequest) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP(), []int{1}
}

func (x *GetTextCompletionRequest) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *GetTextCompletionRequest) GetSystemPrompt() string {
	if x != nil && x.SystemPrompt != nil {
		return *x.SystemPrompt
	}
	return ""
}

func (x *GetTextCompletionRequest) GetHistory() []*GetTextCompletionRequest_HistoryMessage {
	if x != nil {
		return x.History
	}
	return nil
}

func (x *GetTextCompletionRequest) GetUserPrompt() string {
	if x != nil {
		return x.UserPrompt
	}
	return ""
}

func (x *GetTextCompletionRequest) GetSeed() uint32 {
	if x != nil && x.Seed != nil {
		return *x.Seed
	}
	return 0
}

func (x *GetTextCompletionRequest) GetMaxTokens() uint32 {
	if x != nil && x.MaxTokens != nil {
		return *x.MaxTokens
	}
	return 0
}

func (x *GetTextCompletionRequest) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *GetTextCompletionRequest) GetTopP() float32 {
	if x != nil && x.TopP != nil {
		return *x.TopP
	}
	return 0
}

func (x *GetTextCompletionRequest) GetTopK() uint32 {
	if x != nil && x.TopK != nil {
		return *x.TopK
	}
	return 0
}

type GetTextCompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// the full response
	Content string            `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Stats   *PerformanceStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *GetTextCompletionResponse) Reset() {
	*x = GetTextCompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextCompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextCompletionResponse) ProtoMessage() {}

func (x *GetTextCompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextCompletionResponse.ProtoReflect.Descriptor instead.
func (*GetTextCompletionResponse) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP(), []int{2}
}

func (x *GetTextCompletionResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *GetTextCompletionResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *GetTextCompletionResponse) GetStats() *PerformanceStats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type PerformanceStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the total generation time to this point in the stream. In seconds
	TimeGenerated float64 `protobuf:"fixed64,1,opt,name=time_generated,json=timeGenerated,proto3" json:"time_generated,omitempty"`
	// total number of tokens generated to this point in the stream
	TokensGenerated uint32 `protobuf:"varint,2,opt,name=tokens_generated,json=tokensGenerated,proto3" json:"tokens_generated,omitempty"`
	// the processing time to generate the first token. In seconds
	TimeProcessed float64 `protobuf:"fixed64,3,opt,name=time_processed,json=timeProcessed,proto3" json:"time_processed,omitempty"`
	// total number of tokens processed as input for this stream
	TokensProcessed uint32 `protobuf:"varint,4,opt,name=tokens_processed,json=tokensProcessed,proto3" json:"tokens_processed,omitempty"`
}

func (x *PerformanceStats) Reset() {
	*x = PerformanceStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerformanceStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerformanceStats) ProtoMessage() {}

func (x *PerformanceStats) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerformanceStats.ProtoReflect.Descriptor instead.
func (*PerformanceStats) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP(), []int{3}
}

func (x *PerformanceStats) GetTimeGenerated() float64 {
	if x != nil {
		return x.TimeGenerated
	}
	return 0
}

func (x *PerformanceStats) GetTokensGenerated() uint32 {
	if x != nil {
		return x.TokensGenerated
	}
	return 0
}

func (x *PerformanceStats) GetTimeProcessed() float64 {
	if x != nil {
		return x.TimeProcessed
	}
	return 0
}

func (x *PerformanceStats) GetTokensProcessed() uint32 {
	if x != nil {
		return x.TokensProcessed
	}
	return 0
}

type GetTextCompletionRequest_HistoryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the user message: stringified tokens
	UserPrompt string `protobuf:"bytes,1,opt,name=user_prompt,json=userPrompt,proto3" json:"user_prompt,omitempty"`
	// the assistant message: stringified tokens
	AssistantResponse string `protobuf:"bytes,2,opt,name=assistant_response,json=assistantResponse,proto3" json:"assistant_response,omitempty"`
}

func (x *GetTextCompletionRequest_HistoryMessage) Reset() {
	*x = GetTextCompletionRequest_HistoryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTextCompletionRequest_HistoryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTextCompletionRequest_HistoryMessage) ProtoMessage() {}

func (x *GetTextCompletionRequest_HistoryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTextCompletionRequest_HistoryMessage.ProtoReflect.Descriptor instead.
func (*GetTextCompletionRequest_HistoryMessage) Descriptor() ([]byte, []int) {
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetTextCompletionRequest_HistoryMessage) GetUserPrompt() string {
	if x != nil {
		return x.UserPrompt
	}
	return ""
}

func (x *GetTextCompletionRequest_HistoryMessage) GetAssistantResponse() string {
	if x != nil {
		return x.AssistantResponse
	}
	return ""
}

var File_public_llmcloud_requestmanager_v1_requestmanager_proto protoreflect.FileDescriptor

var file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDesc = []byte{
	0x0a, 0x36, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x48, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x22, 0xa2, 0x06, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x28, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x88, 0x01, 0x01, 0x12, 0x64, 0x0a, 0x07,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x27, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x73,
	0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x04, 0x73, 0x65, 0x65,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0f, 0xba,
	0x48, 0x0c, 0x0a, 0x0a, 0x1d, 0x00, 0x00, 0x00, 0x40, 0x25, 0x00, 0x00, 0x00, 0x00, 0x48, 0x03,
	0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x42,
	0x0f, 0xba, 0x48, 0x0c, 0x0a, 0x0a, 0x1d, 0x00, 0x00, 0x80, 0x3f, 0x2d, 0x00, 0x00, 0x00, 0x00,
	0x48, 0x04, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0b, 0xba, 0x48, 0x08, 0x2a,
	0x06, 0x18, 0x80, 0xfa, 0x01, 0x28, 0x00, 0x48, 0x05, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x88,
	0x01, 0x01, 0x1a, 0x60, 0x0a, 0x0e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x3a, 0xb2, 0x01, 0xba, 0x48, 0xae, 0x01, 0x1a, 0xab, 0x01, 0x0a, 0x28,
	0x67, 0x65, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x44, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x20, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x70, 0x6c, 0x75,
	0x73, 0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x20,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x62, 0x65, 0x20, 0x6c,
	0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x61, 0x6e, 0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x1a, 0x39,
	0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x29, 0x20, 0x2b, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68,
	0x69, 0x73, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x29, 0x20, 0x3c, 0x20, 0x36, 0x35, 0x35, 0x33, 0x36, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x73, 0x65, 0x65, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x22, 0x9f, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54,
	0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x49,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x10, 0x50, 0x65,
	0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x5f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x32, 0xac, 0x03, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcc, 0x01, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c,
	0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0xc3, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c,
	0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x75, 0x6c,
	0x6c, 0x42, 0xb4, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x6f, 0x71, 0x2e, 0x69, 0x6f, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x6c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x50, 0x4c, 0x52, 0xaa, 0x02, 0x21, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x2e, 0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x22, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x4c, 0x6c, 0x6d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x5c, 0x4c, 0x6c, 0x6d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x24, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x3a, 0x3a, 0x4c, 0x6c, 0x6d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescOnce sync.Once
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescData = file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDesc
)

func file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescGZIP() []byte {
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescOnce.Do(func() {
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescData)
	})
	return file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDescData
}

var file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_public_llmcloud_requestmanager_v1_requestmanager_proto_goTypes = []interface{}{
	(*GetTextCompletionStreamResponse)(nil),         // 0: public.llmcloud.requestmanager.v1.GetTextCompletionStreamResponse
	(*GetTextCompletionRequest)(nil),                // 1: public.llmcloud.requestmanager.v1.GetTextCompletionRequest
	(*GetTextCompletionResponse)(nil),               // 2: public.llmcloud.requestmanager.v1.GetTextCompletionResponse
	(*PerformanceStats)(nil),                        // 3: public.llmcloud.requestmanager.v1.PerformanceStats
	(*GetTextCompletionRequest_HistoryMessage)(nil), // 4: public.llmcloud.requestmanager.v1.GetTextCompletionRequest.HistoryMessage
}
var file_public_llmcloud_requestmanager_v1_requestmanager_proto_depIdxs = []int32{
	3, // 0: public.llmcloud.requestmanager.v1.GetTextCompletionStreamResponse.stats:type_name -> public.llmcloud.requestmanager.v1.PerformanceStats
	4, // 1: public.llmcloud.requestmanager.v1.GetTextCompletionRequest.history:type_name -> public.llmcloud.requestmanager.v1.GetTextCompletionRequest.HistoryMessage
	3, // 2: public.llmcloud.requestmanager.v1.GetTextCompletionResponse.stats:type_name -> public.llmcloud.requestmanager.v1.PerformanceStats
	1, // 3: public.llmcloud.requestmanager.v1.RequestManagerService.GetTextCompletionStream:input_type -> public.llmcloud.requestmanager.v1.GetTextCompletionRequest
	1, // 4: public.llmcloud.requestmanager.v1.RequestManagerService.GetTextCompletion:input_type -> public.llmcloud.requestmanager.v1.GetTextCompletionRequest
	0, // 5: public.llmcloud.requestmanager.v1.RequestManagerService.GetTextCompletionStream:output_type -> public.llmcloud.requestmanager.v1.GetTextCompletionStreamResponse
	2, // 6: public.llmcloud.requestmanager.v1.RequestManagerService.GetTextCompletion:output_type -> public.llmcloud.requestmanager.v1.GetTextCompletionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_public_llmcloud_requestmanager_v1_requestmanager_proto_init() }
func file_public_llmcloud_requestmanager_v1_requestmanager_proto_init() {
	if File_public_llmcloud_requestmanager_v1_requestmanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextCompletionStreamResponse); i {
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
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextCompletionRequest); i {
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
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextCompletionResponse); i {
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
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerformanceStats); i {
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
		file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTextCompletionRequest_HistoryMessage); i {
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
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_llmcloud_requestmanager_v1_requestmanager_proto_goTypes,
		DependencyIndexes: file_public_llmcloud_requestmanager_v1_requestmanager_proto_depIdxs,
		MessageInfos:      file_public_llmcloud_requestmanager_v1_requestmanager_proto_msgTypes,
	}.Build()
	File_public_llmcloud_requestmanager_v1_requestmanager_proto = out.File
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_rawDesc = nil
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_goTypes = nil
	file_public_llmcloud_requestmanager_v1_requestmanager_proto_depIdxs = nil
}