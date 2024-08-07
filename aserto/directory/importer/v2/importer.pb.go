// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/directory/importer/v2/importer.proto

package importer

import (
	v2 "github.com/aserto-dev/go-directory/aserto/directory/common/v2"
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

// Deprecated: Marked as deprecated in aserto/directory/importer/v2/importer.proto.
type Opcode int32

const (
	Opcode_OPCODE_UNKNOWN Opcode = 0
	Opcode_OPCODE_SET     Opcode = 1
	Opcode_OPCODE_DELETE  Opcode = 2
)

// Enum value maps for Opcode.
var (
	Opcode_name = map[int32]string{
		0: "OPCODE_UNKNOWN",
		1: "OPCODE_SET",
		2: "OPCODE_DELETE",
	}
	Opcode_value = map[string]int32{
		"OPCODE_UNKNOWN": 0,
		"OPCODE_SET":     1,
		"OPCODE_DELETE":  2,
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
	return file_aserto_directory_importer_v2_importer_proto_enumTypes[0].Descriptor()
}

func (Opcode) Type() protoreflect.EnumType {
	return &file_aserto_directory_importer_v2_importer_proto_enumTypes[0]
}

func (x Opcode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Opcode.Descriptor instead.
func (Opcode) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_importer_v2_importer_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in aserto/directory/importer/v2/importer.proto.
type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// operation Opcode enum value
	OpCode Opcode `protobuf:"varint,1,opt,name=op_code,json=opCode,proto3,enum=aserto.directory.importer.v2.Opcode" json:"op_code,omitempty"`
	// Types that are assignable to Msg:
	//
	//	*ImportRequest_ObjectType
	//	*ImportRequest_Permission
	//	*ImportRequest_RelationType
	//	*ImportRequest_Object
	//	*ImportRequest_Relation
	Msg isImportRequest_Msg `protobuf_oneof:"msg"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_aserto_directory_importer_v2_importer_proto_rawDescGZIP(), []int{0}
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

func (x *ImportRequest) GetObjectType() *v2.ObjectType {
	if x, ok := x.GetMsg().(*ImportRequest_ObjectType); ok {
		return x.ObjectType
	}
	return nil
}

func (x *ImportRequest) GetPermission() *v2.Permission {
	if x, ok := x.GetMsg().(*ImportRequest_Permission); ok {
		return x.Permission
	}
	return nil
}

func (x *ImportRequest) GetRelationType() *v2.RelationType {
	if x, ok := x.GetMsg().(*ImportRequest_RelationType); ok {
		return x.RelationType
	}
	return nil
}

func (x *ImportRequest) GetObject() *v2.Object {
	if x, ok := x.GetMsg().(*ImportRequest_Object); ok {
		return x.Object
	}
	return nil
}

func (x *ImportRequest) GetRelation() *v2.Relation {
	if x, ok := x.GetMsg().(*ImportRequest_Relation); ok {
		return x.Relation
	}
	return nil
}

type isImportRequest_Msg interface {
	isImportRequest_Msg()
}

type ImportRequest_ObjectType struct {
	// object_type import message
	ObjectType *v2.ObjectType `protobuf:"bytes,2,opt,name=object_type,json=objectType,proto3,oneof"`
}

type ImportRequest_Permission struct {
	// permission import message
	Permission *v2.Permission `protobuf:"bytes,3,opt,name=permission,proto3,oneof"`
}

type ImportRequest_RelationType struct {
	// relation_type import message
	RelationType *v2.RelationType `protobuf:"bytes,4,opt,name=relation_type,json=relationType,proto3,oneof"`
}

type ImportRequest_Object struct {
	// object import message
	Object *v2.Object `protobuf:"bytes,5,opt,name=object,proto3,oneof"`
}

type ImportRequest_Relation struct {
	// relation import message
	Relation *v2.Relation `protobuf:"bytes,6,opt,name=relation,proto3,oneof"`
}

func (*ImportRequest_ObjectType) isImportRequest_Msg() {}

func (*ImportRequest_Permission) isImportRequest_Msg() {}

func (*ImportRequest_RelationType) isImportRequest_Msg() {}

func (*ImportRequest_Object) isImportRequest_Msg() {}

func (*ImportRequest_Relation) isImportRequest_Msg() {}

// Deprecated: Marked as deprecated in aserto/directory/importer/v2/importer.proto.
type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// object_type import counter
	ObjectType *ImportCounter `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// object_type import counter
	Permission *ImportCounter `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// object_type import counter
	RelationType *ImportCounter `protobuf:"bytes,3,opt,name=relation_type,json=relationType,proto3" json:"relation_type,omitempty"`
	// object import counter
	Object *ImportCounter `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	// object_type import counter
	Relation *ImportCounter `protobuf:"bytes,5,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_aserto_directory_importer_v2_importer_proto_rawDescGZIP(), []int{1}
}

func (x *ImportResponse) GetObjectType() *ImportCounter {
	if x != nil {
		return x.ObjectType
	}
	return nil
}

func (x *ImportResponse) GetPermission() *ImportCounter {
	if x != nil {
		return x.Permission
	}
	return nil
}

func (x *ImportResponse) GetRelationType() *ImportCounter {
	if x != nil {
		return x.RelationType
	}
	return nil
}

func (x *ImportResponse) GetObject() *ImportCounter {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ImportResponse) GetRelation() *ImportCounter {
	if x != nil {
		return x.Relation
	}
	return nil
}

// Deprecated: Marked as deprecated in aserto/directory/importer/v2/importer.proto.
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
}

func (x *ImportCounter) Reset() {
	*x = ImportCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportCounter) ProtoMessage() {}

func (x *ImportCounter) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_importer_v2_importer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_aserto_directory_importer_v2_importer_proto_rawDescGZIP(), []int{2}
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

var File_aserto_directory_importer_v2_importer_proto protoreflect.FileDescriptor

var file_aserto_directory_importer_v2_importer_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a, 0x27, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x03, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x07, 0x6f, 0x70, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x48, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0d, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48,
	0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x02, 0x18,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x8f, 0x03, 0x0a, 0x0e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4b, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x47, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x67, 0x0a, 0x0d, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x63, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x65, 0x63, 0x76, 0x12,
	0x10, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x73, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x3a,
	0x02, 0x18, 0x01, 0x2a, 0x43, 0x0a, 0x06, 0x4f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x50, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x02, 0x1a, 0x02, 0x18, 0x01, 0x32, 0x78, 0x0a, 0x08, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x12, 0x6c, 0x0a, 0x06, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2b,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73,
	0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2f, 0x76, 0x32, 0x3b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_importer_v2_importer_proto_rawDescOnce sync.Once
	file_aserto_directory_importer_v2_importer_proto_rawDescData = file_aserto_directory_importer_v2_importer_proto_rawDesc
)

func file_aserto_directory_importer_v2_importer_proto_rawDescGZIP() []byte {
	file_aserto_directory_importer_v2_importer_proto_rawDescOnce.Do(func() {
		file_aserto_directory_importer_v2_importer_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_importer_v2_importer_proto_rawDescData)
	})
	return file_aserto_directory_importer_v2_importer_proto_rawDescData
}

var file_aserto_directory_importer_v2_importer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_importer_v2_importer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aserto_directory_importer_v2_importer_proto_goTypes = []any{
	(Opcode)(0),             // 0: aserto.directory.importer.v2.Opcode
	(*ImportRequest)(nil),   // 1: aserto.directory.importer.v2.ImportRequest
	(*ImportResponse)(nil),  // 2: aserto.directory.importer.v2.ImportResponse
	(*ImportCounter)(nil),   // 3: aserto.directory.importer.v2.ImportCounter
	(*v2.ObjectType)(nil),   // 4: aserto.directory.common.v2.ObjectType
	(*v2.Permission)(nil),   // 5: aserto.directory.common.v2.Permission
	(*v2.RelationType)(nil), // 6: aserto.directory.common.v2.RelationType
	(*v2.Object)(nil),       // 7: aserto.directory.common.v2.Object
	(*v2.Relation)(nil),     // 8: aserto.directory.common.v2.Relation
}
var file_aserto_directory_importer_v2_importer_proto_depIdxs = []int32{
	0,  // 0: aserto.directory.importer.v2.ImportRequest.op_code:type_name -> aserto.directory.importer.v2.Opcode
	4,  // 1: aserto.directory.importer.v2.ImportRequest.object_type:type_name -> aserto.directory.common.v2.ObjectType
	5,  // 2: aserto.directory.importer.v2.ImportRequest.permission:type_name -> aserto.directory.common.v2.Permission
	6,  // 3: aserto.directory.importer.v2.ImportRequest.relation_type:type_name -> aserto.directory.common.v2.RelationType
	7,  // 4: aserto.directory.importer.v2.ImportRequest.object:type_name -> aserto.directory.common.v2.Object
	8,  // 5: aserto.directory.importer.v2.ImportRequest.relation:type_name -> aserto.directory.common.v2.Relation
	3,  // 6: aserto.directory.importer.v2.ImportResponse.object_type:type_name -> aserto.directory.importer.v2.ImportCounter
	3,  // 7: aserto.directory.importer.v2.ImportResponse.permission:type_name -> aserto.directory.importer.v2.ImportCounter
	3,  // 8: aserto.directory.importer.v2.ImportResponse.relation_type:type_name -> aserto.directory.importer.v2.ImportCounter
	3,  // 9: aserto.directory.importer.v2.ImportResponse.object:type_name -> aserto.directory.importer.v2.ImportCounter
	3,  // 10: aserto.directory.importer.v2.ImportResponse.relation:type_name -> aserto.directory.importer.v2.ImportCounter
	1,  // 11: aserto.directory.importer.v2.Importer.Import:input_type -> aserto.directory.importer.v2.ImportRequest
	2,  // 12: aserto.directory.importer.v2.Importer.Import:output_type -> aserto.directory.importer.v2.ImportResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_aserto_directory_importer_v2_importer_proto_init() }
func file_aserto_directory_importer_v2_importer_proto_init() {
	if File_aserto_directory_importer_v2_importer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_directory_importer_v2_importer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ImportRequest); i {
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
		file_aserto_directory_importer_v2_importer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ImportResponse); i {
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
		file_aserto_directory_importer_v2_importer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ImportCounter); i {
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
	file_aserto_directory_importer_v2_importer_proto_msgTypes[0].OneofWrappers = []any{
		(*ImportRequest_ObjectType)(nil),
		(*ImportRequest_Permission)(nil),
		(*ImportRequest_RelationType)(nil),
		(*ImportRequest_Object)(nil),
		(*ImportRequest_Relation)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_importer_v2_importer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_importer_v2_importer_proto_goTypes,
		DependencyIndexes: file_aserto_directory_importer_v2_importer_proto_depIdxs,
		EnumInfos:         file_aserto_directory_importer_v2_importer_proto_enumTypes,
		MessageInfos:      file_aserto_directory_importer_v2_importer_proto_msgTypes,
	}.Build()
	File_aserto_directory_importer_v2_importer_proto = out.File
	file_aserto_directory_importer_v2_importer_proto_rawDesc = nil
	file_aserto_directory_importer_v2_importer_proto_goTypes = nil
	file_aserto_directory_importer_v2_importer_proto_depIdxs = nil
}
