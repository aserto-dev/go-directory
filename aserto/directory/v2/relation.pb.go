// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aserto/directory/v2/relation.proto

package directory

import (
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

type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectType  string                 `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`    // source object type ID
	ObjectId    string                 `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`          // source object ID (UUID)
	Relation    string                 `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`                          // relation type ID
	SubjectType string                 `protobuf:"bytes,4,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"` // target object type ID
	SubjectId   string                 `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`       // target object ID (UUID)
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // created at timestamp (UTC)
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // last updated timestamp (UTC)
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`      // deleted timestamp (UTC)
	Hash        string                 `protobuf:"bytes,23,opt,name=hash,proto3" json:"hash,omitempty"`                                 // object hash
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_relation_proto_rawDescGZIP(), []int{0}
}

func (x *Relation) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *Relation) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *Relation) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *Relation) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *Relation) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Relation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Relation) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Relation) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Relation) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type RelationType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                     // relation type ID
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                  // relation name (unique object_type+name)
	DisplayName string                 `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"` // relation display name
	ObjectType  string                 `protobuf:"bytes,4,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`    // object type referenced by relation (on the left hand side)
	Ordinal     int32                  `protobuf:"varint,5,opt,name=ordinal,proto3" json:"ordinal,omitempty"`                           // sort ordinal
	Status      uint32                 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`                             // status bitmap
	Unions      []string               `protobuf:"bytes,7,rep,name=unions,proto3" json:"unions,omitempty"`                              // relations union-ed with relation type instance
	Permissions []string               `protobuf:"bytes,8,rep,name=permissions,proto3" json:"permissions,omitempty"`                    // permissions associated to relation type instance
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // created at timestamp (UTC)
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // last updated timestamp (UTC)
	DeletedAt   *timestamppb.Timestamp `protobuf:"bytes,22,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`      // deleted timestamp (UTC)
	Hash        string                 `protobuf:"bytes,23,opt,name=hash,proto3" json:"hash,omitempty"`                                 // object hash
}

func (x *RelationType) Reset() {
	*x = RelationType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationType) ProtoMessage() {}

func (x *RelationType) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationType.ProtoReflect.Descriptor instead.
func (*RelationType) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_relation_proto_rawDescGZIP(), []int{1}
}

func (x *RelationType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RelationType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RelationType) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *RelationType) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *RelationType) GetOrdinal() int32 {
	if x != nil {
		return x.Ordinal
	}
	return 0
}

func (x *RelationType) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RelationType) GetUnions() []string {
	if x != nil {
		return x.Unions
	}
	return nil
}

func (x *RelationType) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *RelationType) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RelationType) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *RelationType) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *RelationType) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

type RelationParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectType  string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`    // source object type ID (OPTIONAL)
	ObjectId    string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`          // source object ID (UUID)
	Relation    string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`                          // relation type ID
	SubjectType string `protobuf:"bytes,4,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"` // target object type ID (OPTIONAL)
	SubjectId   string `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`       // target object ID (UUID)
}

func (x *RelationParam) Reset() {
	*x = RelationParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationParam) ProtoMessage() {}

func (x *RelationParam) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationParam.ProtoReflect.Descriptor instead.
func (*RelationParam) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_relation_proto_rawDescGZIP(), []int{2}
}

func (x *RelationParam) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *RelationParam) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *RelationParam) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationParam) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *RelationParam) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type RelationTypeKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                               // relation type name
	ObjectType string `protobuf:"bytes,2,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"` // relation type associated object type
}

func (x *RelationTypeKey) Reset() {
	*x = RelationTypeKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTypeKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTypeKey) ProtoMessage() {}

func (x *RelationTypeKey) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTypeKey.ProtoReflect.Descriptor instead.
func (*RelationTypeKey) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationTypeKey) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RelationTypeKey) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

type RelationTypeParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Opt:
	//	*RelationTypeParam_Id
	//	*RelationTypeParam_Key
	Opt isRelationTypeParam_Opt `protobuf_oneof:"opt"`
}

func (x *RelationTypeParam) Reset() {
	*x = RelationTypeParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aserto_directory_v2_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTypeParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTypeParam) ProtoMessage() {}

func (x *RelationTypeParam) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_v2_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTypeParam.ProtoReflect.Descriptor instead.
func (*RelationTypeParam) Descriptor() ([]byte, []int) {
	return file_aserto_directory_v2_relation_proto_rawDescGZIP(), []int{4}
}

func (m *RelationTypeParam) GetOpt() isRelationTypeParam_Opt {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (x *RelationTypeParam) GetId() int32 {
	if x, ok := x.GetOpt().(*RelationTypeParam_Id); ok {
		return x.Id
	}
	return 0
}

func (x *RelationTypeParam) GetKey() *RelationTypeKey {
	if x, ok := x.GetOpt().(*RelationTypeParam_Key); ok {
		return x.Key
	}
	return nil
}

type isRelationTypeParam_Opt interface {
	isRelationTypeParam_Opt()
}

type RelationTypeParam_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"` // relation type by id
}

type RelationTypeParam_Key struct {
	Key *RelationTypeKey `protobuf:"bytes,2,opt,name=key,proto3,oneof"` // relation type by object_type + name
}

func (*RelationTypeParam_Id) isRelationTypeParam_Opt() {}

func (*RelationTypeParam_Key) isRelationTypeParam_Opt() {}

var File_aserto_directory_v2_relation_proto protoreflect.FileDescriptor

var file_aserto_directory_v2_relation_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb, 0x02, 0x0a, 0x08, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
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
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xa7, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x22, 0xab, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x66, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x38, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x76, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4b,
	0x65, 0x79, 0x48, 0x00, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x05, 0x0a, 0x03, 0x6f, 0x70, 0x74,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_v2_relation_proto_rawDescOnce sync.Once
	file_aserto_directory_v2_relation_proto_rawDescData = file_aserto_directory_v2_relation_proto_rawDesc
)

func file_aserto_directory_v2_relation_proto_rawDescGZIP() []byte {
	file_aserto_directory_v2_relation_proto_rawDescOnce.Do(func() {
		file_aserto_directory_v2_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_v2_relation_proto_rawDescData)
	})
	return file_aserto_directory_v2_relation_proto_rawDescData
}

var file_aserto_directory_v2_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_aserto_directory_v2_relation_proto_goTypes = []interface{}{
	(*Relation)(nil),              // 0: aserto.directory.v2.Relation
	(*RelationType)(nil),          // 1: aserto.directory.v2.RelationType
	(*RelationParam)(nil),         // 2: aserto.directory.v2.RelationParam
	(*RelationTypeKey)(nil),       // 3: aserto.directory.v2.RelationTypeKey
	(*RelationTypeParam)(nil),     // 4: aserto.directory.v2.RelationTypeParam
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_aserto_directory_v2_relation_proto_depIdxs = []int32{
	5, // 0: aserto.directory.v2.Relation.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: aserto.directory.v2.Relation.updated_at:type_name -> google.protobuf.Timestamp
	5, // 2: aserto.directory.v2.Relation.deleted_at:type_name -> google.protobuf.Timestamp
	5, // 3: aserto.directory.v2.RelationType.created_at:type_name -> google.protobuf.Timestamp
	5, // 4: aserto.directory.v2.RelationType.updated_at:type_name -> google.protobuf.Timestamp
	5, // 5: aserto.directory.v2.RelationType.deleted_at:type_name -> google.protobuf.Timestamp
	3, // 6: aserto.directory.v2.RelationTypeParam.key:type_name -> aserto.directory.v2.RelationTypeKey
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_aserto_directory_v2_relation_proto_init() }
func file_aserto_directory_v2_relation_proto_init() {
	if File_aserto_directory_v2_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aserto_directory_v2_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_aserto_directory_v2_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationType); i {
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
		file_aserto_directory_v2_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationParam); i {
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
		file_aserto_directory_v2_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTypeKey); i {
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
		file_aserto_directory_v2_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTypeParam); i {
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
	file_aserto_directory_v2_relation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*RelationTypeParam_Id)(nil),
		(*RelationTypeParam_Key)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_v2_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_directory_v2_relation_proto_goTypes,
		DependencyIndexes: file_aserto_directory_v2_relation_proto_depIdxs,
		MessageInfos:      file_aserto_directory_v2_relation_proto_msgTypes,
	}.Build()
	File_aserto_directory_v2_relation_proto = out.File
	file_aserto_directory_v2_relation_proto_rawDesc = nil
	file_aserto_directory_v2_relation_proto_goTypes = nil
	file_aserto_directory_v2_relation_proto_depIdxs = nil
}
