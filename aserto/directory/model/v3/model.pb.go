// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/directory/model/v3/model.proto

package model

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetManifestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// empty request
	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *GetManifestRequest) Reset() {
	*x = GetManifestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManifestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManifestRequest) ProtoMessage() {}

func (x *GetManifestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManifestRequest.ProtoReflect.Descriptor instead.
func (*GetManifestRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{0}
}

func (x *GetManifestRequest) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type GetManifestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*GetManifestResponse_Metadata
	//	*GetManifestResponse_Body
	//	*GetManifestResponse_Model
	Msg isGetManifestResponse_Msg `protobuf_oneof:"msg"`
}

func (x *GetManifestResponse) Reset() {
	*x = GetManifestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManifestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManifestResponse) ProtoMessage() {}

func (x *GetManifestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManifestResponse.ProtoReflect.Descriptor instead.
func (*GetManifestResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{1}
}

func (m *GetManifestResponse) GetMsg() isGetManifestResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *GetManifestResponse) GetMetadata() *Metadata {
	if x, ok := x.GetMsg().(*GetManifestResponse_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *GetManifestResponse) GetBody() *Body {
	if x, ok := x.GetMsg().(*GetManifestResponse_Body); ok {
		return x.Body
	}
	return nil
}

func (x *GetManifestResponse) GetModel() *structpb.Struct {
	if x, ok := x.GetMsg().(*GetManifestResponse_Model); ok {
		return x.Model
	}
	return nil
}

type isGetManifestResponse_Msg interface {
	isGetManifestResponse_Msg()
}

type GetManifestResponse_Metadata struct {
	// Manifest metadata
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type GetManifestResponse_Body struct {
	// Manifest content
	Body *Body `protobuf:"bytes,2,opt,name=body,proto3,oneof"`
}

type GetManifestResponse_Model struct {
	// Model
	Model *structpb.Struct `protobuf:"bytes,3,opt,name=model,proto3,oneof"`
}

func (*GetManifestResponse_Metadata) isGetManifestResponse_Msg() {}

func (*GetManifestResponse_Body) isGetManifestResponse_Msg() {}

func (*GetManifestResponse_Model) isGetManifestResponse_Msg() {}

type SetManifestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// manifest instance
	//
	// Types that are assignable to Msg:
	//
	//	*SetManifestRequest_Body
	Msg isSetManifestRequest_Msg `protobuf_oneof:"msg"`
}

func (x *SetManifestRequest) Reset() {
	*x = SetManifestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManifestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManifestRequest) ProtoMessage() {}

func (x *SetManifestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManifestRequest.ProtoReflect.Descriptor instead.
func (*SetManifestRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{2}
}

func (m *SetManifestRequest) GetMsg() isSetManifestRequest_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *SetManifestRequest) GetBody() *Body {
	if x, ok := x.GetMsg().(*SetManifestRequest_Body); ok {
		return x.Body
	}
	return nil
}

type isSetManifestRequest_Msg interface {
	isSetManifestRequest_Msg()
}

type SetManifestRequest_Body struct {
	// Manifest content
	Body *Body `protobuf:"bytes,1,opt,name=body,proto3,oneof"`
}

func (*SetManifestRequest_Body) isSetManifestRequest_Msg() {}

type SetManifestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// empty result
	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SetManifestResponse) Reset() {
	*x = SetManifestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetManifestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetManifestResponse) ProtoMessage() {}

func (x *SetManifestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetManifestResponse.ProtoReflect.Descriptor instead.
func (*SetManifestResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{3}
}

func (x *SetManifestResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteManifestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// empty request
	Empty *emptypb.Empty `protobuf:"bytes,1,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *DeleteManifestRequest) Reset() {
	*x = DeleteManifestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManifestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManifestRequest) ProtoMessage() {}

func (x *DeleteManifestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManifestRequest.ProtoReflect.Descriptor instead.
func (*DeleteManifestRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteManifestRequest) GetEmpty() *emptypb.Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

type DeleteManifestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// empty result
	Result *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteManifestResponse) Reset() {
	*x = DeleteManifestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteManifestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteManifestResponse) ProtoMessage() {}

func (x *DeleteManifestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteManifestResponse.ProtoReflect.Descriptor instead.
func (*DeleteManifestResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteManifestResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// last updated timestamp (UTC)
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// object instance etag
	Etag string `protobuf:"bytes,23,opt,name=etag,proto3" json:"etag,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{6}
}

func (x *Metadata) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Metadata) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// manifest content
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_model_v3_model_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_model_v3_model_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_aserto_directory_model_v3_model_proto_rawDescGZIP(), []int{7}
}

func (x *Body) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_aserto_directory_model_v3_model_proto protoreflect.FileDescriptor

var file_aserto_directory_model_v3_model_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x76, 0x33, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc7,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x2f, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x52, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x48, 0x00, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x45, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x45, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x16, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x63, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x25, 0x0a, 0x04, 0x42, 0x6f, 0x64,
	0x79, 0x12, 0x1d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x09, 0xba, 0x48, 0x06, 0x7a, 0x04, 0x18, 0x80, 0x80, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x88, 0x04, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x76, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x00,
	0x30, 0x01, 0x12, 0x76, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x06, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x00, 0x28, 0x01, 0x12, 0x8e, 0x02, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x30, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x31, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x96, 0x01, 0x92, 0x41, 0x71, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x12, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6d, 0x61, 0x6e, 0x69,
	0x66, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20, 0x6d, 0x61, 0x6e,
	0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x2a, 0x1c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x62, 0x23, 0x0a, 0x13, 0x0a, 0x0f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x41, 0x50, 0x49, 0x4b, 0x65, 0x79, 0x12, 0x00, 0x0a, 0x0c, 0x0a, 0x08, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a,
	0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_model_v3_model_proto_rawDescOnce sync.Once
	file_aserto_directory_model_v3_model_proto_rawDescData = file_aserto_directory_model_v3_model_proto_rawDesc
)

func file_aserto_directory_model_v3_model_proto_rawDescGZIP() []byte {
	file_aserto_directory_model_v3_model_proto_rawDescOnce.Do(func() {
		file_aserto_directory_model_v3_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_model_v3_model_proto_rawDescData)
	})
	return file_aserto_directory_model_v3_model_proto_rawDescData
}

var file_aserto_directory_model_v3_model_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aserto_directory_model_v3_model_proto_goTypes = []any{
	(*GetManifestRequest)(nil),     // 0: aserto.directory.model.v3.GetManifestRequest
	(*GetManifestResponse)(nil),    // 1: aserto.directory.model.v3.GetManifestResponse
	(*SetManifestRequest)(nil),     // 2: aserto.directory.model.v3.SetManifestRequest
	(*SetManifestResponse)(nil),    // 3: aserto.directory.model.v3.SetManifestResponse
	(*DeleteManifestRequest)(nil),  // 4: aserto.directory.model.v3.DeleteManifestRequest
	(*DeleteManifestResponse)(nil), // 5: aserto.directory.model.v3.DeleteManifestResponse
	(*Metadata)(nil),               // 6: aserto.directory.model.v3.Metadata
	(*Body)(nil),                   // 7: aserto.directory.model.v3.Body
	(*emptypb.Empty)(nil),          // 8: google.protobuf.Empty
	(*structpb.Struct)(nil),        // 9: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
}
var file_aserto_directory_model_v3_model_proto_depIdxs = []int32{
	8,  // 0: aserto.directory.model.v3.GetManifestRequest.empty:type_name -> google.protobuf.Empty
	6,  // 1: aserto.directory.model.v3.GetManifestResponse.metadata:type_name -> aserto.directory.model.v3.Metadata
	7,  // 2: aserto.directory.model.v3.GetManifestResponse.body:type_name -> aserto.directory.model.v3.Body
	9,  // 3: aserto.directory.model.v3.GetManifestResponse.model:type_name -> google.protobuf.Struct
	7,  // 4: aserto.directory.model.v3.SetManifestRequest.body:type_name -> aserto.directory.model.v3.Body
	8,  // 5: aserto.directory.model.v3.SetManifestResponse.result:type_name -> google.protobuf.Empty
	8,  // 6: aserto.directory.model.v3.DeleteManifestRequest.empty:type_name -> google.protobuf.Empty
	8,  // 7: aserto.directory.model.v3.DeleteManifestResponse.result:type_name -> google.protobuf.Empty
	10, // 8: aserto.directory.model.v3.Metadata.updated_at:type_name -> google.protobuf.Timestamp
	0,  // 9: aserto.directory.model.v3.Model.GetManifest:input_type -> aserto.directory.model.v3.GetManifestRequest
	2,  // 10: aserto.directory.model.v3.Model.SetManifest:input_type -> aserto.directory.model.v3.SetManifestRequest
	4,  // 11: aserto.directory.model.v3.Model.DeleteManifest:input_type -> aserto.directory.model.v3.DeleteManifestRequest
	1,  // 12: aserto.directory.model.v3.Model.GetManifest:output_type -> aserto.directory.model.v3.GetManifestResponse
	3,  // 13: aserto.directory.model.v3.Model.SetManifest:output_type -> aserto.directory.model.v3.SetManifestResponse
	5,  // 14: aserto.directory.model.v3.Model.DeleteManifest:output_type -> aserto.directory.model.v3.DeleteManifestResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_aserto_directory_model_v3_model_proto_init() }
func file_aserto_directory_model_v3_model_proto_init() {
	if File_aserto_directory_model_v3_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_directory_model_v3_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetManifestRequest); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetManifestResponse); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SetManifestRequest); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SetManifestResponse); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteManifestRequest); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteManifestResponse); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
		file_aserto_directory_model_v3_model_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Body); i {
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
	file_aserto_directory_model_v3_model_proto_msgTypes[1].OneofWrappers = []any{
		(*GetManifestResponse_Metadata)(nil),
		(*GetManifestResponse_Body)(nil),
		(*GetManifestResponse_Model)(nil),
	}
	file_aserto_directory_model_v3_model_proto_msgTypes[2].OneofWrappers = []any{
		(*SetManifestRequest_Body)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_model_v3_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_model_v3_model_proto_goTypes,
		DependencyIndexes: file_aserto_directory_model_v3_model_proto_depIdxs,
		MessageInfos:      file_aserto_directory_model_v3_model_proto_msgTypes,
	}.Build()
	File_aserto_directory_model_v3_model_proto = out.File
	file_aserto_directory_model_v3_model_proto_rawDesc = nil
	file_aserto_directory_model_v3_model_proto_goTypes = nil
	file_aserto_directory_model_v3_model_proto_depIdxs = nil
}
