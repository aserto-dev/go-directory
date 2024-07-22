// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aserto/directory/exporter/v2/exporter.proto

package exporter

import (
	v2 "github.com/aserto-dev/go-directory/aserto/directory/common/v2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Deprecated: Marked as deprecated in aserto/directory/exporter/v2/exporter.proto.
type Option int32

const (
	// nothing selected (default initialization value)
	Option_OPTION_UNKNOWN Option = 0
	// object type metadata
	Option_OPTION_METADATA_OBJECT_TYPES Option = 1
	// relation type metadata
	Option_OPTION_METADATA_RELATION_TYPES Option = 2
	// permission metadata
	Option_OPTION_METADATA_PERMISSIONS Option = 4
	// all metadata = OPTION_METADATA_OBJECT_TYPES | OPTION_METADATA_RELATION_TYPES | OPTION_METADATA_PERMISSIONS
	Option_OPTION_METADATA Option = 7
	// object instances
	Option_OPTION_DATA_OBJECTS Option = 8
	// relation instances
	Option_OPTION_DATA_RELATIONS Option = 16
	// relation instances with key values
	Option_OPTION_DATA_RELATIONS_WITH_KEYS Option = 32
	// all data = OPTION_DATA_OBJECTS | OPTION_DATA_RELATIONS
	Option_OPTION_DATA Option = 24
	// all data with keys = OPTION_DATA_OBJECTS | OPTION_DATA_RELATIONS_WITH_KEYS
	Option_OPTION_DATA_WITH_KEYS Option = 40
	// all metadata and data = OPTION_METADATA | OPTION_DATA
	Option_OPTION_ALL Option = 31
	// all metadata and data with keys = OPTION_METADATA | OPTION_DATA_WITH_KEYS
	Option_OPTION_ALL_WITH_KEYS Option = 47
)

// Enum value maps for Option.
var (
	Option_name = map[int32]string{
		0:  "OPTION_UNKNOWN",
		1:  "OPTION_METADATA_OBJECT_TYPES",
		2:  "OPTION_METADATA_RELATION_TYPES",
		4:  "OPTION_METADATA_PERMISSIONS",
		7:  "OPTION_METADATA",
		8:  "OPTION_DATA_OBJECTS",
		16: "OPTION_DATA_RELATIONS",
		32: "OPTION_DATA_RELATIONS_WITH_KEYS",
		24: "OPTION_DATA",
		40: "OPTION_DATA_WITH_KEYS",
		31: "OPTION_ALL",
		47: "OPTION_ALL_WITH_KEYS",
	}
	Option_value = map[string]int32{
		"OPTION_UNKNOWN":                  0,
		"OPTION_METADATA_OBJECT_TYPES":    1,
		"OPTION_METADATA_RELATION_TYPES":  2,
		"OPTION_METADATA_PERMISSIONS":     4,
		"OPTION_METADATA":                 7,
		"OPTION_DATA_OBJECTS":             8,
		"OPTION_DATA_RELATIONS":           16,
		"OPTION_DATA_RELATIONS_WITH_KEYS": 32,
		"OPTION_DATA":                     24,
		"OPTION_DATA_WITH_KEYS":           40,
		"OPTION_ALL":                      31,
		"OPTION_ALL_WITH_KEYS":            47,
	}
)

func (x Option) Enum() *Option {
	p := new(Option)
	*p = x
	return p
}

func (x Option) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Option) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_directory_exporter_v2_exporter_proto_enumTypes[0].Descriptor()
}

func (Option) Type() protoreflect.EnumType {
	return &file_aserto_directory_exporter_v2_exporter_proto_enumTypes[0]
}

func (x Option) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Option.Descriptor instead.
func (Option) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_exporter_v2_exporter_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in aserto/directory/exporter/v2/exporter.proto.
type ExportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data export options mask
	Options uint32 `protobuf:"varint,1,opt,name=options,proto3" json:"options,omitempty"`
	// start export from timestamp (UTC)
	StartFrom *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_exporter_v2_exporter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_exporter_v2_exporter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportRequest.ProtoReflect.Descriptor instead.
func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_exporter_v2_exporter_proto_rawDescGZIP(), []int{0}
}

func (x *ExportRequest) GetOptions() uint32 {
	if x != nil {
		return x.Options
	}
	return 0
}

func (x *ExportRequest) GetStartFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.StartFrom
	}
	return nil
}

// Deprecated: Marked as deprecated in aserto/directory/exporter/v2/exporter.proto.
type ExportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*ExportResponse_Object
	//	*ExportResponse_ObjectType
	//	*ExportResponse_Relation
	//	*ExportResponse_RelationType
	//	*ExportResponse_Permission
	Msg isExportResponse_Msg `protobuf_oneof:"msg"`
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_exporter_v2_exporter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_exporter_v2_exporter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportResponse.ProtoReflect.Descriptor instead.
func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_exporter_v2_exporter_proto_rawDescGZIP(), []int{1}
}

func (m *ExportResponse) GetMsg() isExportResponse_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ExportResponse) GetObject() *v2.Object {
	if x, ok := x.GetMsg().(*ExportResponse_Object); ok {
		return x.Object
	}
	return nil
}

func (x *ExportResponse) GetObjectType() *v2.ObjectType {
	if x, ok := x.GetMsg().(*ExportResponse_ObjectType); ok {
		return x.ObjectType
	}
	return nil
}

func (x *ExportResponse) GetRelation() *v2.Relation {
	if x, ok := x.GetMsg().(*ExportResponse_Relation); ok {
		return x.Relation
	}
	return nil
}

func (x *ExportResponse) GetRelationType() *v2.RelationType {
	if x, ok := x.GetMsg().(*ExportResponse_RelationType); ok {
		return x.RelationType
	}
	return nil
}

func (x *ExportResponse) GetPermission() *v2.Permission {
	if x, ok := x.GetMsg().(*ExportResponse_Permission); ok {
		return x.Permission
	}
	return nil
}

type isExportResponse_Msg interface {
	isExportResponse_Msg()
}

type ExportResponse_Object struct {
	// object instance (data)
	Object *v2.Object `protobuf:"bytes,2,opt,name=object,proto3,oneof"`
}

type ExportResponse_ObjectType struct {
	// object type instance (metadata)
	ObjectType *v2.ObjectType `protobuf:"bytes,3,opt,name=object_type,json=objectType,proto3,oneof"`
}

type ExportResponse_Relation struct {
	// relation instance (data)
	Relation *v2.Relation `protobuf:"bytes,4,opt,name=relation,proto3,oneof"`
}

type ExportResponse_RelationType struct {
	// relation type instance (metadata)
	RelationType *v2.RelationType `protobuf:"bytes,5,opt,name=relation_type,json=relationType,proto3,oneof"`
}

type ExportResponse_Permission struct {
	// permission instance (metadata)
	Permission *v2.Permission `protobuf:"bytes,6,opt,name=permission,proto3,oneof"`
}

func (*ExportResponse_Object) isExportResponse_Msg() {}

func (*ExportResponse_ObjectType) isExportResponse_Msg() {}

func (*ExportResponse_Relation) isExportResponse_Msg() {}

func (*ExportResponse_RelationType) isExportResponse_Msg() {}

func (*ExportResponse_Permission) isExportResponse_Msg() {}

var File_aserto_directory_exporter_v2_exporter_proto protoreflect.FileDescriptor

var file_aserto_directory_exporter_v2_exporter_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a, 0x27, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x3a, 0x02, 0x18, 0x01, 0x22,
	0x83, 0x03, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x49, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52,
	0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x48, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0xcb, 0x02, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x53, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x53, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x4f, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4f,
	0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x07,
	0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f,
	0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x53, 0x10, 0x08, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x53, 0x10, 0x10, 0x12, 0x23, 0x0a, 0x1f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x10, 0x20, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x18, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x50,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x4b,
	0x45, 0x59, 0x53, 0x10, 0x28, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x4c, 0x4c, 0x10, 0x1f, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x4c, 0x4c, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x4b, 0x45, 0x59, 0x53, 0x10, 0x2f, 0x1a,
	0x02, 0x18, 0x01, 0x32, 0x76, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12,
	0x6a, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72,
	0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x30, 0x01, 0x42, 0x4a, 0x5a, 0x48, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x3b, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_exporter_v2_exporter_proto_rawDescOnce sync.Once
	file_aserto_directory_exporter_v2_exporter_proto_rawDescData = file_aserto_directory_exporter_v2_exporter_proto_rawDesc
)

func file_aserto_directory_exporter_v2_exporter_proto_rawDescGZIP() []byte {
	file_aserto_directory_exporter_v2_exporter_proto_rawDescOnce.Do(func() {
		file_aserto_directory_exporter_v2_exporter_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_exporter_v2_exporter_proto_rawDescData)
	})
	return file_aserto_directory_exporter_v2_exporter_proto_rawDescData
}

var file_aserto_directory_exporter_v2_exporter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_exporter_v2_exporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_directory_exporter_v2_exporter_proto_goTypes = []any{
	(Option)(0),                   // 0: aserto.directory.exporter.v2.Option
	(*ExportRequest)(nil),         // 1: aserto.directory.exporter.v2.ExportRequest
	(*ExportResponse)(nil),        // 2: aserto.directory.exporter.v2.ExportResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*v2.Object)(nil),             // 4: aserto.directory.common.v2.Object
	(*v2.ObjectType)(nil),         // 5: aserto.directory.common.v2.ObjectType
	(*v2.Relation)(nil),           // 6: aserto.directory.common.v2.Relation
	(*v2.RelationType)(nil),       // 7: aserto.directory.common.v2.RelationType
	(*v2.Permission)(nil),         // 8: aserto.directory.common.v2.Permission
}
var file_aserto_directory_exporter_v2_exporter_proto_depIdxs = []int32{
	3, // 0: aserto.directory.exporter.v2.ExportRequest.start_from:type_name -> google.protobuf.Timestamp
	4, // 1: aserto.directory.exporter.v2.ExportResponse.object:type_name -> aserto.directory.common.v2.Object
	5, // 2: aserto.directory.exporter.v2.ExportResponse.object_type:type_name -> aserto.directory.common.v2.ObjectType
	6, // 3: aserto.directory.exporter.v2.ExportResponse.relation:type_name -> aserto.directory.common.v2.Relation
	7, // 4: aserto.directory.exporter.v2.ExportResponse.relation_type:type_name -> aserto.directory.common.v2.RelationType
	8, // 5: aserto.directory.exporter.v2.ExportResponse.permission:type_name -> aserto.directory.common.v2.Permission
	1, // 6: aserto.directory.exporter.v2.Exporter.Export:input_type -> aserto.directory.exporter.v2.ExportRequest
	2, // 7: aserto.directory.exporter.v2.Exporter.Export:output_type -> aserto.directory.exporter.v2.ExportResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_aserto_directory_exporter_v2_exporter_proto_init() }
func file_aserto_directory_exporter_v2_exporter_proto_init() {
	if File_aserto_directory_exporter_v2_exporter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_directory_exporter_v2_exporter_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExportRequest); i {
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
		file_aserto_directory_exporter_v2_exporter_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExportResponse); i {
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
	file_aserto_directory_exporter_v2_exporter_proto_msgTypes[1].OneofWrappers = []any{
		(*ExportResponse_Object)(nil),
		(*ExportResponse_ObjectType)(nil),
		(*ExportResponse_Relation)(nil),
		(*ExportResponse_RelationType)(nil),
		(*ExportResponse_Permission)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_exporter_v2_exporter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_exporter_v2_exporter_proto_goTypes,
		DependencyIndexes: file_aserto_directory_exporter_v2_exporter_proto_depIdxs,
		EnumInfos:         file_aserto_directory_exporter_v2_exporter_proto_enumTypes,
		MessageInfos:      file_aserto_directory_exporter_v2_exporter_proto_msgTypes,
	}.Build()
	File_aserto_directory_exporter_v2_exporter_proto = out.File
	file_aserto_directory_exporter_v2_exporter_proto_rawDesc = nil
	file_aserto_directory_exporter_v2_exporter_proto_goTypes = nil
	file_aserto_directory_exporter_v2_exporter_proto_depIdxs = nil
}
