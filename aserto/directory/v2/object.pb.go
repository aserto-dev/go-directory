// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/directory/v2/object.proto

package directory

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Flag int32

const (
	Flag_FLAG_UNKNOWN  Flag = 0 // default, no special object behavior
	Flag_FLAG_HIDDEN   Flag = 1 // hidden object
	Flag_FLAG_READONLY Flag = 2 // read-only object
	Flag_FLAG_SYSTEM   Flag = 4 // system object
	Flag_FLAG_SHADOW   Flag = 8 // shadow object by type+key assosciated to parent object
)

// Enum value maps for Flag.
var (
	Flag_name = map[int32]string{
		0: "FLAG_UNKNOWN",
		1: "FLAG_HIDDEN",
		2: "FLAG_READONLY",
		4: "FLAG_SYSTEM",
		8: "FLAG_SHADOW",
	}
	Flag_value = map[string]int32{
		"FLAG_UNKNOWN":  0,
		"FLAG_HIDDEN":   1,
		"FLAG_READONLY": 2,
		"FLAG_SYSTEM":   4,
		"FLAG_SHADOW":   8,
	}
)

func (x Flag) Enum() *Flag {
	p := new(Flag)
	*p = x
	return p
}

func (x Flag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Flag) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_directory_v2_object_proto_enumTypes[0].Descriptor()
}

func (Flag) Type() protoreflect.EnumType {
	return &file_aserto_directory_v2_object_proto_enumTypes[0]
}

func (x Flag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Flag.Descriptor instead.
func (Flag) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_v2_object_proto_rawDescGZIP(), []int{0}
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // ID (uuid)
	Type        int32                  `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`                                 // object type ID
	Key         string                 `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`                                    // external key/ID (case-sensistive)
	DisplayName string                 `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // display name object
	Properties  *structpb.Struct       `protobuf:"bytes,5,opt,name=properties,proto3" json:"properties,omitempty"`                      // property bag
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // created at timestamp (UTC)
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // last updated timestamp (UTC)
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`      // deleted timestamp (UTC)
	Hash        string                 `protobuf:"bytes,23,opt,name=hash,proto3" json:"hash,omitempty"`                                 // object hash
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Object) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Object) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Object) GetProperties() *structpb.Struct {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *Object) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Object) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Object) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Object) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type ObjectType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                     // object type ID
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                  // object type name (unique)
	DisplayName string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // object type display name
	IsSubject   bool                   `protobuf:"varint,4,opt,name=is_subject,json=isSubject,proto3" json:"is_subject,omitempty"`      // object type is a subject (user|group)
	Ordinal     int32                  `protobuf:"varint,5,opt,name=ordinal,proto3" json:"ordinal,omitempty"`                           // sort ordinal
	Status      uint32                 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`                             // status flag bitmap
	Schema      *structpb.Struct       `protobuf:"bytes,20,opt,name=schema,proto3" json:"schema,omitempty"`                             // object type schema definition (JSON)
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // created at timestamp (UTC)
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // last updated timestamp (UTC)
}

func (x *ObjectType) Reset() {
	*x = ObjectType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectType) ProtoMessage() {}

func (x *ObjectType) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectType.ProtoReflect.Descriptor instead.
func (*ObjectType) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_object_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ObjectType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ObjectType) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ObjectType) GetIsSubject() bool {
	if x != nil {
		return x.IsSubject
	}
	return false
}

func (x *ObjectType) GetOrdinal() int32 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *ObjectType) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ObjectType) GetSchema() *structpb.Struct {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *ObjectType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ObjectType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ObjectDependency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceType     int32  `protobuf:"varint,1,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`              // object type id of source object
	SourceTypeName string `protobuf:"bytes,2,opt,name=source_type_name,json=sourceTypeName,proto3" json:"source_type_name,omitempty"` // string represenantion of source_type (READONLY)
	SourceId       string `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`                     // object id (uuid) of source object
	SourceKey      string `protobuf:"bytes,4,opt,name=source_key,json=sourceKey,proto3" json:"source_key,omitempty"`                  // object search key of source object
	Relation       int32  `protobuf:"varint,5,opt,name=relation,proto3" json:"relation,omitempty"`                                    // relation identifier
	RelationName   string `protobuf:"bytes,6,opt,name=relation_name,json=relationName,proto3" json:"relation_name,omitempty"`         // string representation of relation (READONLY)
	TargetType     int32  `protobuf:"varint,7,opt,name=target_type,json=targetType,proto3" json:"target_type,omitempty"`              // object type id of target object
	TargetTypeName string `protobuf:"bytes,8,opt,name=target_type_name,json=targetTypeName,proto3" json:"target_type_name,omitempty"` // string represenantion of target_type (READONLY)
	TargetId       string `protobuf:"bytes,9,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`                     // object id (uuid) of target object
	TargetKey      string `protobuf:"bytes,10,opt,name=target_key,json=targetKey,proto3" json:"target_key,omitempty"`                 // object search key of target object
	Depth          int32  `protobuf:"varint,11,opt,name=depth,proto3" json:"depth,omitempty"`                                         // dependency depth
	IsCycle        bool   `protobuf:"varint,12,opt,name=is_cycle,json=isCycle,proto3" json:"is_cycle,omitempty"`                      // dependency cycle
	Path           string `protobuf:"bytes,13,opt,name=path,proto3" json:"path,omitempty"`                                            // dependency path
}

func (x *ObjectDependency) Reset() {
	*x = ObjectDependency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectDependency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectDependency) ProtoMessage() {}

func (x *ObjectDependency) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectDependency.ProtoReflect.Descriptor instead.
func (*ObjectDependency) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_object_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectDependency) GetSourceType() int32 {
	if x != nil {
		return x.SourceType
	}
	return 0
}

func (x *ObjectDependency) GetSourceTypeName() string {
	if x != nil {
		return x.SourceTypeName
	}
	return ""
}

func (x *ObjectDependency) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *ObjectDependency) GetSourceKey() string {
	if x != nil {
		return x.SourceKey
	}
	return ""
}

func (x *ObjectDependency) GetRelation() int32 {
	if x != nil {
		return x.Relation
	}
	return 0
}

func (x *ObjectDependency) GetRelationName() string {
	if x != nil {
		return x.RelationName
	}
	return ""
}

func (x *ObjectDependency) GetTargetType() int32 {
	if x != nil {
		return x.TargetType
	}
	return 0
}

func (x *ObjectDependency) GetTargetTypeName() string {
	if x != nil {
		return x.TargetTypeName
	}
	return ""
}

func (x *ObjectDependency) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *ObjectDependency) GetTargetKey() string {
	if x != nil {
		return x.TargetKey
	}
	return ""
}

func (x *ObjectDependency) GetDepth() int32 {
	if x != nil {
		return x.Depth
	}
	return 0
}

func (x *ObjectDependency) GetIsCycle() bool {
	if x != nil {
		return x.IsCycle
	}
	return false
}

func (x *ObjectDependency) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ObjectKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"` // object type ID
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`    // external key/ID (case-sensistive)
}

func (x *ObjectKey) Reset() {
	*x = ObjectKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_object_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectKey) ProtoMessage() {}

func (x *ObjectKey) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_object_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectKey.ProtoReflect.Descriptor instead.
func (*ObjectKey) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_object_proto_rawDescGZIP(), []int{3}
}

func (x *ObjectKey) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ObjectKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_aserto_directory_v2_object_proto protoreflect.FileDescriptor

var file_aserto_directory_v2_object_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xcb, 0x02, 0x0a, 0x0a, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2f, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa6, 0x03, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x70, 0x74, 0x68,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x43, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x31, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x2a, 0x5e, 0x0a, 0x04, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x0c, 0x46, 0x4c,
	0x41, 0x47, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x46, 0x4c, 0x41, 0x47, 0x5f, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x10,
	0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x53, 0x48, 0x41, 0x44, 0x4f, 0x57,
	0x10, 0x08, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_v2_object_proto_rawDescOnce sync.Once
	file_aserto_directory_v2_object_proto_rawDescData = file_aserto_directory_v2_object_proto_rawDesc
)

func file_aserto_directory_v2_object_proto_rawDescGZIP() []byte {
	file_aserto_directory_v2_object_proto_rawDescOnce.Do(func() {
		file_aserto_directory_v2_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_v2_object_proto_rawDescData)
	})
	return file_aserto_directory_v2_object_proto_rawDescData
}

var file_aserto_directory_v2_object_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_v2_object_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_aserto_directory_v2_object_proto_goTypes = []interface{}{
	(Flag)(0),                     // 0: aserto.directory.v2.Flag
	(*Object)(nil),                // 1: aserto.directory.v2.Object
	(*ObjectType)(nil),            // 2: aserto.directory.v2.ObjectType
	(*ObjectDependency)(nil),      // 3: aserto.directory.v2.ObjectDependency
	(*ObjectKey)(nil),             // 4: aserto.directory.v2.ObjectKey
	(*structpb.Struct)(nil),       // 5: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_aserto_directory_v2_object_proto_depIdxs = []int32{
	5, // 0: aserto.directory.v2.Object.properties:type_name -> google.protobuf.Struct
	6, // 1: aserto.directory.v2.Object.created_at:type_name -> google.protobuf.Timestamp
	6, // 2: aserto.directory.v2.Object.updated_at:type_name -> google.protobuf.Timestamp
	6, // 3: aserto.directory.v2.Object.deleted_at:type_name -> google.protobuf.Timestamp
	5, // 4: aserto.directory.v2.ObjectType.schema:type_name -> google.protobuf.Struct
	6, // 5: aserto.directory.v2.ObjectType.created_at:type_name -> google.protobuf.Timestamp
	6, // 6: aserto.directory.v2.ObjectType.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_aserto_directory_v2_object_proto_init() }
func file_aserto_directory_v2_object_proto_init() {
	if File_aserto_directory_v2_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_directory_v2_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_aserto_directory_v2_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectType); i {
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
		file_aserto_directory_v2_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectDependency); i {
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
		file_aserto_directory_v2_object_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectKey); i {
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
			RawDescriptor: file_aserto_directory_v2_object_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_directory_v2_object_proto_goTypes,
		DependencyIndexes: file_aserto_directory_v2_object_proto_depIdxs,
		EnumInfos:         file_aserto_directory_v2_object_proto_enumTypes,
		MessageInfos:      file_aserto_directory_v2_object_proto_msgTypes,
	}.Build()
	File_aserto_directory_v2_object_proto = out.File
	file_aserto_directory_v2_object_proto_rawDesc = nil
	file_aserto_directory_v2_object_proto_goTypes = nil
	file_aserto_directory_v2_object_proto_depIdxs = nil
}
