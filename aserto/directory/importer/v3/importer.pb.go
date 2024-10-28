// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: aserto/directory/importer/v3/importer.proto

package importer

import (
	v3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
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

type Opcode int32

const (
	Opcode_OPCODE_UNKNOWN               Opcode = 0
	Opcode_OPCODE_SET                   Opcode = 1
	Opcode_OPCODE_DELETE                Opcode = 2
	Opcode_OPCODE_DELETE_WITH_RELATIONS Opcode = 3
)

// Enum value maps for Opcode.
var (
	Opcode_name = map[int32]string{
		0: "OPCODE_UNKNOWN",
		1: "OPCODE_SET",
		2: "OPCODE_DELETE",
		3: "OPCODE_DELETE_WITH_RELATIONS",
	}
	Opcode_value = map[string]int32{
		"OPCODE_UNKNOWN":               0,
		"OPCODE_SET":                   1,
		"OPCODE_DELETE":                2,
		"OPCODE_DELETE_WITH_RELATIONS": 3,
	}
)

func (x Opcode) Enum() *Opcode {
	p := new(Opcode)
	*p = x
	return p
}

func (x Opcode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Opcode) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_directory_importer_v3_importer_proto_enumTypes[0].Descriptor()
}

func (Opcode) Type() protoreflect.EnumType {
	return &file_aserto_directory_importer_v3_importer_proto_enumTypes[0]
}

func (x Opcode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Opcode.Descriptor instead.
func (Opcode) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v3_importer_proto_rawDescGZIP(), []int{0}
}

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// operation Opcode enum value
	OpCode Opcode `protobuf:"varint,1,opt,name=op_code,json=opCode,proto3,enum=aserto.directory.importer.v3.Opcode" json:"op_code,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*ImportRequest_Object
	//	*ImportRequest_Relation
	Msg isImportRequest_Msg `protobuf_oneof:"msg"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v3_importer_proto_rawDescGZIP(), []int{0}
}

func (x *ImportRequest) GetOpCode() Opcode {
	if x != nil {
		return x.OpCode
	}
	return Opcode_OPCODE_UNKNOWN
}

func (m *ImportRequest) GetMsg() isImportRequest_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ImportRequest) GetObject() *v3.Object {
	if x, ok := x.GetMsg().(*ImportRequest_Object); ok {
		return x.Object
	}
	return nil
}

func (x *ImportRequest) GetRelation() *v3.Relation {
	if x, ok := x.GetMsg().(*ImportRequest_Relation); ok {
		return x.Relation
	}
	return nil
}

type isImportRequest_Msg interface {
	isImportRequest_Msg()
}

type ImportRequest_Object struct {
	// object import message
	Object *v3.Object `protobuf:"bytes,5,opt,name=object,proto3,oneof"`
}

type ImportRequest_Relation struct {
	// relation import message
	Relation *v3.Relation `protobuf:"bytes,6,opt,name=relation,proto3,oneof"`
}

func (*ImportRequest_Object) isImportRequest_Msg() {}

func (*ImportRequest_Relation) isImportRequest_Msg() {}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// object import counter
	//
	// Deprecated: Marked as deprecated in aserto/directory/importer/v3/importer.proto.
	Object *ImportCounter `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// relation import counter
	//
	// Deprecated: Marked as deprecated in aserto/directory/importer/v3/importer.proto.
	Relation *ImportCounter `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*ImportResponse_Status
	//	*ImportResponse_Counter
	Msg isImportResponse_Msg `protobuf_oneof:"msg"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v3_importer_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Marked as deprecated in aserto/directory/importer/v3/importer.proto.
func (x *ImportResponse) GetObject() *ImportCounter {
	if x != nil {
		return x.Object
	}
	return nil
}

// Deprecated: Marked as deprecated in aserto/directory/importer/v3/importer.proto.
func (x *ImportResponse) GetRelation() *ImportCounter {
	if x != nil {
		return x.Relation
	}
	return nil
}

func (m *ImportResponse) GetMsg() isImportResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ImportResponse) GetStatus() *ImportStatus {
	if x, ok := x.GetMsg().(*ImportResponse_Status); ok {
		return x.Status
	}
	return nil
}

func (x *ImportResponse) GetCounter() *ImportCounter {
	if x, ok := x.GetMsg().(*ImportResponse_Counter); ok {
		return x.Counter
	}
	return nil
}

type isImportResponse_Msg interface {
	isImportResponse_Msg()
}

type ImportResponse_Status struct {
	// import status message
	Status *ImportStatus `protobuf:"bytes,4,opt,name=status,proto3,oneof"`
}

type ImportResponse_Counter struct {
	// import counter per type
	Counter *ImportCounter `protobuf:"bytes,5,opt,name=counter,proto3,oneof"`
}

func (*ImportResponse_Status) isImportResponse_Msg() {}

func (*ImportResponse_Counter) isImportResponse_Msg() {}

type ImportCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// number of messages received
	Recv uint64 `protobuf:"varint,1,opt,name=recv,proto3" json:"recv,omitempty"`
	// number of messages with OPCODE_SET
	Set uint64 `protobuf:"varint,2,opt,name=set,proto3" json:"set,omitempty"`
	// number of messages with OPCODE_DELETE
	Delete uint64 `protobuf:"varint,3,opt,name=delete,proto3" json:"delete,omitempty"`
	// number of messages resulting in error
	Error uint64 `protobuf:"varint,4,opt,name=error,proto3" json:"error,omitempty"`
	// counter of type (object|relation)
	Type string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ImportCounter) Reset() {
	*x = ImportCounter{}
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportCounter) ProtoMessage() {}

func (x *ImportCounter) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportCounter.ProtoReflect.Descriptor instead.
func (*ImportCounter) Descriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v3_importer_proto_rawDescGZIP(), []int{2}
}

func (x *ImportCounter) GetRecv() uint64 {
	if x != nil {
		return x.Recv
	}
	return 0
}

func (x *ImportCounter) GetSet() uint64 {
	if x != nil {
		return x.Set
	}
	return 0
}

func (x *ImportCounter) GetDelete() uint64 {
	if x != nil {
		return x.Delete
	}
	return 0
}

func (x *ImportCounter) GetError() uint64 {
	if x != nil {
		return x.Error
	}
	return 0
}

func (x *ImportCounter) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ImportStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// gRPC status code (google.golang.org/grpc/codes)
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// gRPC status message (google.golang.org/grpc/status)
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// req contains the original import request message
	Req *ImportRequest `protobuf:"bytes,3,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *ImportStatus) Reset() {
	*x = ImportStatus{}
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportStatus) ProtoMessage() {}

func (x *ImportStatus) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v3_importer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportStatus.ProtoReflect.Descriptor instead.
func (*ImportStatus) Descriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v3_importer_proto_rawDescGZIP(), []int{3}
}

func (x *ImportStatus) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ImportStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ImportStatus) GetReq() *ImportRequest {
	if x != nil {
		return x.Req
	}
	return nil
}

var File_aserto_directory_importer_v3_importer_proto protoreflect.FileDescriptor

var file_aserto_directory_importer_v3_importer_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x27, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xbc,
	0x02, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4b, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x02, 0x18, 0x01, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x47, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x77, 0x0a,
	0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x65, 0x63, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65,
	0x63, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x03, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x73, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3d, 0x0a, 0x03,
	0x72, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x03, 0x72, 0x65, 0x71, 0x2a, 0x61, 0x0a, 0x06, 0x4f,
	0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c,
	0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x03, 0x32, 0x75,
	0x0a, 0x08, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x69, 0x0a, 0x06, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67,
	0x6f, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x3b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_importer_v3_importer_proto_rawDescOnce sync.Once
	file_aserto_directory_importer_v3_importer_proto_rawDescData = file_aserto_directory_importer_v3_importer_proto_rawDesc
)

func file_aserto_directory_importer_v3_importer_proto_rawDescGZIP() []byte {
	file_aserto_directory_importer_v3_importer_proto_rawDescOnce.Do(func() {
		file_aserto_directory_importer_v3_importer_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_importer_v3_importer_proto_rawDescData)
	})
	return file_aserto_directory_importer_v3_importer_proto_rawDescData
}

var file_aserto_directory_importer_v3_importer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_importer_v3_importer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_directory_importer_v3_importer_proto_goTypes = []any{
	(Opcode)(0),            // 0: aserto.directory.importer.v3.Opcode
	(*ImportRequest)(nil),  // 1: aserto.directory.importer.v3.ImportRequest
	(*ImportResponse)(nil), // 2: aserto.directory.importer.v3.ImportResponse
	(*ImportCounter)(nil),  // 3: aserto.directory.importer.v3.ImportCounter
	(*ImportStatus)(nil),   // 4: aserto.directory.importer.v3.ImportStatus
	(*v3.Object)(nil),      // 5: aserto.directory.common.v3.Object
	(*v3.Relation)(nil),    // 6: aserto.directory.common.v3.Relation
}
var file_aserto_directory_importer_v3_importer_proto_depIdxs = []int32{
	0, // 0: aserto.directory.importer.v3.ImportRequest.op_code:type_name -> aserto.directory.importer.v3.Opcode
	5, // 1: aserto.directory.importer.v3.ImportRequest.object:type_name -> aserto.directory.common.v3.Object
	6, // 2: aserto.directory.importer.v3.ImportRequest.relation:type_name -> aserto.directory.common.v3.Relation
	3, // 3: aserto.directory.importer.v3.ImportResponse.object:type_name -> aserto.directory.importer.v3.ImportCounter
	3, // 4: aserto.directory.importer.v3.ImportResponse.relation:type_name -> aserto.directory.importer.v3.ImportCounter
	4, // 5: aserto.directory.importer.v3.ImportResponse.status:type_name -> aserto.directory.importer.v3.ImportStatus
	3, // 6: aserto.directory.importer.v3.ImportResponse.counter:type_name -> aserto.directory.importer.v3.ImportCounter
	1, // 7: aserto.directory.importer.v3.ImportStatus.req:type_name -> aserto.directory.importer.v3.ImportRequest
	1, // 8: aserto.directory.importer.v3.Importer.Import:input_type -> aserto.directory.importer.v3.ImportRequest
	2, // 9: aserto.directory.importer.v3.Importer.Import:output_type -> aserto.directory.importer.v3.ImportResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_aserto_directory_importer_v3_importer_proto_init() }
func file_aserto_directory_importer_v3_importer_proto_init() {
	if File_aserto_directory_importer_v3_importer_proto != nil {
		return
	}
	file_aserto_directory_importer_v3_importer_proto_msgTypes[0].OneofWrappers = []any{
		(*ImportRequest_Object)(nil),
		(*ImportRequest_Relation)(nil),
	}
	file_aserto_directory_importer_v3_importer_proto_msgTypes[1].OneofWrappers = []any{
		(*ImportResponse_Status)(nil),
		(*ImportResponse_Counter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_importer_v3_importer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_importer_v3_importer_proto_goTypes,
		DependencyIndexes: file_aserto_directory_importer_v3_importer_proto_depIdxs,
		EnumInfos:         file_aserto_directory_importer_v3_importer_proto_enumTypes,
		MessageInfos:      file_aserto_directory_importer_v3_importer_proto_msgTypes,
	}.Build()
	File_aserto_directory_importer_v3_importer_proto = out.File
	file_aserto_directory_importer_v3_importer_proto_rawDesc = nil
	file_aserto_directory_importer_v3_importer_proto_goTypes = nil
	file_aserto_directory_importer_v3_importer_proto_depIdxs = nil
}
