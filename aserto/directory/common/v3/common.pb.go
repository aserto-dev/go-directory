// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: aserto/directory/common/v3/common.proto

package common

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Object struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object type name
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// external object identifier (cs-string, no spaces or tabs)
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// display name object
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// property bag
	Properties *structpb.Struct `protobuf:"bytes,4,opt,name=properties,proto3" json:"properties,omitempty"`
	// created at timestamp (UTC)
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// last updated timestamp (UTC)
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// object instance etag
	Etag          string `protobuf:"bytes,23,opt,name=etag,proto3" json:"etag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Object) Reset() {
	*x = Object{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[0]
	if x != nil {
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
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
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

func (x *Object) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

type Relation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object type
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// object identifier
	ObjectId string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// object relation name
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// subject type
	SubjectType string `protobuf:"bytes,4,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	// subject identifier
	SubjectId string `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// optional subject relation name
	SubjectRelation string `protobuf:"bytes,6,opt,name=subject_relation,json=subjectRelation,proto3" json:"subject_relation,omitempty"`
	// created at timestamp (UTC)
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// last updated timestamp (UTC)
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// object instance etag
	Etag          string `protobuf:"bytes,23,opt,name=etag,proto3" json:"etag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Relation) Reset() {
	*x = Relation{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[1]
	if x != nil {
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
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{1}
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

func (x *Relation) GetSubjectRelation() string {
	if x != nil {
		return x.SubjectRelation
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

func (x *Relation) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

// Object identifier
type ObjectIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object type (lc-string)
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// object identifier (cs-string)
	ObjectId      string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ObjectIdentifier) Reset() {
	*x = ObjectIdentifier{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectIdentifier) ProtoMessage() {}

func (x *ObjectIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectIdentifier.ProtoReflect.Descriptor instead.
func (*ObjectIdentifier) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectIdentifier) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *ObjectIdentifier) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

// Relation identifier
type RelationIdentifier struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object type
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// object identifier
	ObjectId string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// object relation name
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// subject type
	SubjectType string `protobuf:"bytes,4,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
	// subject identifier
	SubjectId string `protobuf:"bytes,5,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// optional subject relation name
	SubjectRelation string `protobuf:"bytes,6,opt,name=subject_relation,json=subjectRelation,proto3" json:"subject_relation,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *RelationIdentifier) Reset() {
	*x = RelationIdentifier{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationIdentifier) ProtoMessage() {}

func (x *RelationIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationIdentifier.ProtoReflect.Descriptor instead.
func (*RelationIdentifier) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{3}
}

func (x *RelationIdentifier) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *RelationIdentifier) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *RelationIdentifier) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationIdentifier) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *RelationIdentifier) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *RelationIdentifier) GetSubjectRelation() string {
	if x != nil {
		return x.SubjectRelation
	}
	return ""
}

// Pagination request
type PaginationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// requested page size, valid value between 1-100 rows (default 100)
	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	// pagination start token, default ""
	Token         string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationRequest) Reset() {
	*x = PaginationRequest{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{4}
}

func (x *PaginationRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PaginationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Pagination response
type PaginationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// next page token, when empty there are no more pages to fetch
	NextToken     string `protobuf:"bytes,1,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationResponse.ProtoReflect.Descriptor instead.
func (*PaginationResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{5}
}

func (x *PaginationResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

var File_aserto_directory_common_v3_common_proto protoreflect.FileDescriptor

var file_aserto_directory_common_v3_common_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x02, 0x0a, 0x06, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x17, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x17, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x88, 0x03,
	0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0b, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x2e, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0f,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x3e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x22, 0x5a, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0b,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x22, 0xf9, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0b, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x20, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52,
	0x0f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x47, 0x0a, 0x11, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x12, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f,
	0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_aserto_directory_common_v3_common_proto_rawDescOnce sync.Once
	file_aserto_directory_common_v3_common_proto_rawDescData []byte
)

func file_aserto_directory_common_v3_common_proto_rawDescGZIP() []byte {
	file_aserto_directory_common_v3_common_proto_rawDescOnce.Do(func() {
		file_aserto_directory_common_v3_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_directory_common_v3_common_proto_rawDesc), len(file_aserto_directory_common_v3_common_proto_rawDesc)))
	})
	return file_aserto_directory_common_v3_common_proto_rawDescData
}

var file_aserto_directory_common_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aserto_directory_common_v3_common_proto_goTypes = []any{
	(*Object)(nil),                // 0: aserto.directory.common.v3.Object
	(*Relation)(nil),              // 1: aserto.directory.common.v3.Relation
	(*ObjectIdentifier)(nil),      // 2: aserto.directory.common.v3.ObjectIdentifier
	(*RelationIdentifier)(nil),    // 3: aserto.directory.common.v3.RelationIdentifier
	(*PaginationRequest)(nil),     // 4: aserto.directory.common.v3.PaginationRequest
	(*PaginationResponse)(nil),    // 5: aserto.directory.common.v3.PaginationResponse
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_aserto_directory_common_v3_common_proto_depIdxs = []int32{
	6, // 0: aserto.directory.common.v3.Object.properties:type_name -> google.protobuf.Struct
	7, // 1: aserto.directory.common.v3.Object.created_at:type_name -> google.protobuf.Timestamp
	7, // 2: aserto.directory.common.v3.Object.updated_at:type_name -> google.protobuf.Timestamp
	7, // 3: aserto.directory.common.v3.Relation.created_at:type_name -> google.protobuf.Timestamp
	7, // 4: aserto.directory.common.v3.Relation.updated_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_aserto_directory_common_v3_common_proto_init() }
func file_aserto_directory_common_v3_common_proto_init() {
	if File_aserto_directory_common_v3_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_directory_common_v3_common_proto_rawDesc), len(file_aserto_directory_common_v3_common_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_directory_common_v3_common_proto_goTypes,
		DependencyIndexes: file_aserto_directory_common_v3_common_proto_depIdxs,
		MessageInfos:      file_aserto_directory_common_v3_common_proto_msgTypes,
	}.Build()
	File_aserto_directory_common_v3_common_proto = out.File
	file_aserto_directory_common_v3_common_proto_goTypes = nil
	file_aserto_directory_common_v3_common_proto_depIdxs = nil
}
