// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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

// manifest type
type Manifest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Manifest payload
	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Model representation of manifest
	Model *structpb.Struct `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	// last updated timestamp (UTC)
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// manifest instance etag
	Etag          string `protobuf:"bytes,23,opt,name=etag,proto3" json:"etag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Manifest) Reset() {
	*x = Manifest{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Manifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manifest) ProtoMessage() {}

func (x *Manifest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Manifest.ProtoReflect.Descriptor instead.
func (*Manifest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{0}
}

func (x *Manifest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Manifest) GetModel() *structpb.Struct {
	if x != nil {
		return x.Model
	}
	return nil
}

func (x *Manifest) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Manifest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

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
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{1}
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
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{2}
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

// object identifier
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
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObjectIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectIdentifier) ProtoMessage() {}

func (x *ObjectIdentifier) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ObjectIdentifier.ProtoReflect.Descriptor instead.
func (*ObjectIdentifier) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{3}
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

// relation identifier
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
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationIdentifier) ProtoMessage() {}

func (x *RelationIdentifier) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use RelationIdentifier.ProtoReflect.Descriptor instead.
func (*RelationIdentifier) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{4}
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

// pagination request
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
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationRequest) ProtoMessage() {}

func (x *PaginationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PaginationRequest.ProtoReflect.Descriptor instead.
func (*PaginationRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{5}
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

// pagination response
type PaginationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// next page token, when empty there are no more pages to fetch
	NextToken     string `protobuf:"bytes,1,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationResponse) Reset() {
	*x = PaginationResponse{}
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationResponse) ProtoMessage() {}

func (x *PaginationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_common_v3_common_proto_msgTypes[6]
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
	return file_aserto_directory_common_v3_common_proto_rawDescGZIP(), []int{6}
}

func (x *PaginationResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

var File_aserto_directory_common_v3_common_proto protoreflect.FileDescriptor

const file_aserto_directory_common_v3_common_proto_rawDesc = "" +
	"\n" +
	"'aserto/directory/common/v3/common.proto\x12\x1aaserto.directory.common.v3\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xab\x01\n" +
	"\bManifest\x12\x12\n" +
	"\x04body\x18\x01 \x01(\fR\x04body\x122\n" +
	"\x05model\x18\x03 \x01(\v2\x17.google.protobuf.StructB\x03\xe0A\x03R\x05model\x12>\n" +
	"\n" +
	"updated_at\x18\x15 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tupdatedAt\x12\x17\n" +
	"\x04etag\x18\x17 \x01(\tB\x03\xe0A\x01R\x04etag\"\xb5\x02\n" +
	"\x06Object\x12\x17\n" +
	"\x04type\x18\x01 \x01(\tB\x03\xe0A\x02R\x04type\x12\x13\n" +
	"\x02id\x18\x02 \x01(\tB\x03\xe0A\x02R\x02id\x12&\n" +
	"\fdisplay_name\x18\x03 \x01(\tB\x03\xe0A\x01R\vdisplayName\x12<\n" +
	"\n" +
	"properties\x18\x04 \x01(\v2\x17.google.protobuf.StructB\x03\xe0A\x01R\n" +
	"properties\x12>\n" +
	"\n" +
	"created_at\x18\x14 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tcreatedAt\x12>\n" +
	"\n" +
	"updated_at\x18\x15 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tupdatedAt\x12\x17\n" +
	"\x04etag\x18\x17 \x01(\tB\x03\xe0A\x01R\x04etag\"\x88\x03\n" +
	"\bRelation\x12$\n" +
	"\vobject_type\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"objectType\x12 \n" +
	"\tobject_id\x18\x02 \x01(\tB\x03\xe0A\x02R\bobjectId\x12\x1f\n" +
	"\brelation\x18\x03 \x01(\tB\x03\xe0A\x02R\brelation\x12&\n" +
	"\fsubject_type\x18\x04 \x01(\tB\x03\xe0A\x02R\vsubjectType\x12\"\n" +
	"\n" +
	"subject_id\x18\x05 \x01(\tB\x03\xe0A\x02R\tsubjectId\x12.\n" +
	"\x10subject_relation\x18\x06 \x01(\tB\x03\xe0A\x01R\x0fsubjectRelation\x12>\n" +
	"\n" +
	"created_at\x18\x14 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tcreatedAt\x12>\n" +
	"\n" +
	"updated_at\x18\x15 \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\tupdatedAt\x12\x17\n" +
	"\x04etag\x18\x17 \x01(\tB\x03\xe0A\x01R\x04etag\"Z\n" +
	"\x10ObjectIdentifier\x12$\n" +
	"\vobject_type\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"objectType\x12 \n" +
	"\tobject_id\x18\x02 \x01(\tB\x03\xe0A\x02R\bobjectId\"\xf9\x01\n" +
	"\x12RelationIdentifier\x12$\n" +
	"\vobject_type\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"objectType\x12 \n" +
	"\tobject_id\x18\x02 \x01(\tB\x03\xe0A\x02R\bobjectId\x12\x1f\n" +
	"\brelation\x18\x03 \x01(\tB\x03\xe0A\x02R\brelation\x12&\n" +
	"\fsubject_type\x18\x04 \x01(\tB\x03\xe0A\x02R\vsubjectType\x12\"\n" +
	"\n" +
	"subject_id\x18\x05 \x01(\tB\x03\xe0A\x02R\tsubjectId\x12.\n" +
	"\x10subject_relation\x18\x06 \x01(\tB\x03\xe0A\x01R\x0fsubjectRelation\"G\n" +
	"\x11PaginationRequest\x12\x17\n" +
	"\x04size\x18\x01 \x01(\x05B\x03\xe0A\x01R\x04size\x12\x19\n" +
	"\x05token\x18\x02 \x01(\tB\x03\xe0A\x01R\x05token\"8\n" +
	"\x12PaginationResponse\x12\"\n" +
	"\n" +
	"next_token\x18\x01 \x01(\tB\x03\xe0A\x03R\tnextTokenBFZDgithub.com/aserto-dev/go-directory/aserto/directory/common/v3;commonb\x06proto3"

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

var file_aserto_directory_common_v3_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_aserto_directory_common_v3_common_proto_goTypes = []any{
	(*Manifest)(nil),              // 0: aserto.directory.common.v3.Manifest
	(*Object)(nil),                // 1: aserto.directory.common.v3.Object
	(*Relation)(nil),              // 2: aserto.directory.common.v3.Relation
	(*ObjectIdentifier)(nil),      // 3: aserto.directory.common.v3.ObjectIdentifier
	(*RelationIdentifier)(nil),    // 4: aserto.directory.common.v3.RelationIdentifier
	(*PaginationRequest)(nil),     // 5: aserto.directory.common.v3.PaginationRequest
	(*PaginationResponse)(nil),    // 6: aserto.directory.common.v3.PaginationResponse
	(*structpb.Struct)(nil),       // 7: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_aserto_directory_common_v3_common_proto_depIdxs = []int32{
	7, // 0: aserto.directory.common.v3.Manifest.model:type_name -> google.protobuf.Struct
	8, // 1: aserto.directory.common.v3.Manifest.updated_at:type_name -> google.protobuf.Timestamp
	7, // 2: aserto.directory.common.v3.Object.properties:type_name -> google.protobuf.Struct
	8, // 3: aserto.directory.common.v3.Object.created_at:type_name -> google.protobuf.Timestamp
	8, // 4: aserto.directory.common.v3.Object.updated_at:type_name -> google.protobuf.Timestamp
	8, // 5: aserto.directory.common.v3.Relation.created_at:type_name -> google.protobuf.Timestamp
	8, // 6: aserto.directory.common.v3.Relation.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
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
			NumMessages:   7,
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
