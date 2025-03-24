// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/directory/writer/v3/writer.proto

package writer

import (
	v3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type SetObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object instance
	Object        *v3.Object `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetObjectRequest) Reset() {
	*x = SetObjectRequest{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetObjectRequest) ProtoMessage() {}

func (x *SetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetObjectRequest.ProtoReflect.Descriptor instead.
func (*SetObjectRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{0}
}

func (x *SetObjectRequest) GetObject() *v3.Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SetObjectResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object instance
	Result        *v3.Object `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetObjectResponse) Reset() {
	*x = SetObjectResponse{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetObjectResponse) ProtoMessage() {}

func (x *SetObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetObjectResponse.ProtoReflect.Descriptor instead.
func (*SetObjectResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{1}
}

func (x *SetObjectResponse) GetResult() *v3.Object {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteObjectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// object type
	ObjectType string `protobuf:"bytes,1,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`
	// object identifier
	ObjectId string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// delete object relations, both object and subject relations.
	WithRelations bool `protobuf:"varint,3,opt,name=with_relations,json=withRelations,proto3" json:"with_relations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteObjectRequest) Reset() {
	*x = DeleteObjectRequest{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectRequest) ProtoMessage() {}

func (x *DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteObjectRequest) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *DeleteObjectRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *DeleteObjectRequest) GetWithRelations() bool {
	if x != nil {
		return x.WithRelations
	}
	return false
}

type DeleteObjectResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// empty result
	Result        *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteObjectResponse) Reset() {
	*x = DeleteObjectResponse{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectResponse) ProtoMessage() {}

func (x *DeleteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteObjectResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteObjectResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

type SetRelationRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// relation instance
	Relation      *v3.Relation `protobuf:"bytes,1,opt,name=relation,proto3" json:"relation,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRelationRequest) Reset() {
	*x = SetRelationRequest{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRelationRequest) ProtoMessage() {}

func (x *SetRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRelationRequest.ProtoReflect.Descriptor instead.
func (*SetRelationRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{4}
}

func (x *SetRelationRequest) GetRelation() *v3.Relation {
	if x != nil {
		return x.Relation
	}
	return nil
}

type SetRelationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// relation instance
	Result        *v3.Relation `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRelationResponse) Reset() {
	*x = SetRelationResponse{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRelationResponse) ProtoMessage() {}

func (x *SetRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRelationResponse.ProtoReflect.Descriptor instead.
func (*SetRelationResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{5}
}

func (x *SetRelationResponse) GetResult() *v3.Relation {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteRelationRequest struct {
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

func (x *DeleteRelationRequest) Reset() {
	*x = DeleteRelationRequest{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationRequest) ProtoMessage() {}

func (x *DeleteRelationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationRequest.ProtoReflect.Descriptor instead.
func (*DeleteRelationRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRelationRequest) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *DeleteRelationRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *DeleteRelationRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *DeleteRelationRequest) GetSubjectType() string {
	if x != nil {
		return x.SubjectType
	}
	return ""
}

func (x *DeleteRelationRequest) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *DeleteRelationRequest) GetSubjectRelation() string {
	if x != nil {
		return x.SubjectRelation
	}
	return ""
}

type DeleteRelationResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// empty result
	Result        *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationResponse) Reset() {
	*x = DeleteRelationResponse{}
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationResponse) ProtoMessage() {}

func (x *DeleteRelationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_writer_v3_writer_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationResponse.ProtoReflect.Descriptor instead.
func (*DeleteRelationResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_writer_v3_writer_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRelationResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_aserto_directory_writer_v3_writer_proto protoreflect.FileDescriptor

const file_aserto_directory_writer_v3_writer_proto_rawDesc = "" +
	"\n" +
	"'aserto/directory/writer/v3/writer.proto\x12\x1aaserto.directory.writer.v3\x1a'aserto/directory/common/v3/common.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"S\n" +
	"\x10SetObjectRequest\x12?\n" +
	"\x06object\x18\x01 \x01(\v2\".aserto.directory.common.v3.ObjectB\x03\xe0A\x02R\x06object\"O\n" +
	"\x11SetObjectResponse\x12:\n" +
	"\x06result\x18\x01 \x01(\v2\".aserto.directory.common.v3.ObjectR\x06result\"\x89\x01\n" +
	"\x13DeleteObjectRequest\x12$\n" +
	"\vobject_type\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"objectType\x12 \n" +
	"\tobject_id\x18\x02 \x01(\tB\x03\xe0A\x02R\bobjectId\x12*\n" +
	"\x0ewith_relations\x18\x03 \x01(\bB\x03\xe0A\x01R\rwithRelations\"F\n" +
	"\x14DeleteObjectResponse\x12.\n" +
	"\x06result\x18\x01 \x01(\v2\x16.google.protobuf.EmptyR\x06result\"[\n" +
	"\x12SetRelationRequest\x12E\n" +
	"\brelation\x18\x01 \x01(\v2$.aserto.directory.common.v3.RelationB\x03\xe0A\x02R\brelation\"S\n" +
	"\x13SetRelationResponse\x12<\n" +
	"\x06result\x18\x01 \x01(\v2$.aserto.directory.common.v3.RelationR\x06result\"\xfc\x01\n" +
	"\x15DeleteRelationRequest\x12$\n" +
	"\vobject_type\x18\x01 \x01(\tB\x03\xe0A\x02R\n" +
	"objectType\x12 \n" +
	"\tobject_id\x18\x02 \x01(\tB\x03\xe0A\x02R\bobjectId\x12\x1f\n" +
	"\brelation\x18\x03 \x01(\tB\x03\xe0A\x02R\brelation\x12&\n" +
	"\fsubject_type\x18\x04 \x01(\tB\x03\xe0A\x02R\vsubjectType\x12\"\n" +
	"\n" +
	"subject_id\x18\x05 \x01(\tB\x03\xe0A\x02R\tsubjectId\x12.\n" +
	"\x10subject_relation\x18\x06 \x01(\tB\x03\xe0A\x01R\x0fsubjectRelation\"H\n" +
	"\x16DeleteRelationResponse\x12.\n" +
	"\x06result\x18\x01 \x01(\v2\x16.google.protobuf.EmptyR\x06result2\xd0\b\n" +
	"\x06Writer\x12\xfa\x01\n" +
	"\tSetObject\x12,.aserto.directory.writer.v3.SetObjectRequest\x1a-.aserto.directory.writer.v3.SetObjectResponse\"\x8f\x01\x92Ai\n" +
	"\tdirectory\x12\n" +
	"Set object\x1a\vSet object.*\x1edirectory.writer.v3.object.setb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1d:\x01*\"\x18/api/v3/directory/object\x12\xa3\x02\n" +
	"\fDeleteObject\x12/.aserto.directory.writer.v3.DeleteObjectRequest\x1a0.aserto.directory.writer.v3.DeleteObjectResponse\"\xaf\x01\x92Ar\n" +
	"\tdirectory\x12\rDelete object\x1a\x0eDelete object.*!directory.writer.v3.object.deleteb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x024*2/api/v3/directory/object/{object_type}/{object_id}\x12\x88\x02\n" +
	"\vSetRelation\x12..aserto.directory.writer.v3.SetRelationRequest\x1a/.aserto.directory.writer.v3.SetRelationResponse\"\x97\x01\x92Ao\n" +
	"\tdirectory\x12\fSet relation\x1a\rSet relation.* directory.writer.v3.relation.setb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1f:\x01*\"\x1a/api/v3/directory/relation\x12\x97\x02\n" +
	"\x0eDeleteRelation\x121.aserto.directory.writer.v3.DeleteRelationRequest\x1a2.aserto.directory.writer.v3.DeleteRelationResponse\"\x9d\x01\x92Ax\n" +
	"\tdirectory\x12\x0fDelete relation\x1a\x10Delete relation.*#directory.writer.v3.relation.deleteb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1c*\x1a/api/v3/directory/relationBFZDgithub.com/aserto-dev/go-directory/aserto/directory/writer/v3;writerb\x06proto3"

var (
	file_aserto_directory_writer_v3_writer_proto_rawDescOnce sync.Once
	file_aserto_directory_writer_v3_writer_proto_rawDescData []byte
)

func file_aserto_directory_writer_v3_writer_proto_rawDescGZIP() []byte {
	file_aserto_directory_writer_v3_writer_proto_rawDescOnce.Do(func() {
		file_aserto_directory_writer_v3_writer_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_directory_writer_v3_writer_proto_rawDesc), len(file_aserto_directory_writer_v3_writer_proto_rawDesc)))
	})
	return file_aserto_directory_writer_v3_writer_proto_rawDescData
}

var file_aserto_directory_writer_v3_writer_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aserto_directory_writer_v3_writer_proto_goTypes = []any{
	(*SetObjectRequest)(nil),       // 0: aserto.directory.writer.v3.SetObjectRequest
	(*SetObjectResponse)(nil),      // 1: aserto.directory.writer.v3.SetObjectResponse
	(*DeleteObjectRequest)(nil),    // 2: aserto.directory.writer.v3.DeleteObjectRequest
	(*DeleteObjectResponse)(nil),   // 3: aserto.directory.writer.v3.DeleteObjectResponse
	(*SetRelationRequest)(nil),     // 4: aserto.directory.writer.v3.SetRelationRequest
	(*SetRelationResponse)(nil),    // 5: aserto.directory.writer.v3.SetRelationResponse
	(*DeleteRelationRequest)(nil),  // 6: aserto.directory.writer.v3.DeleteRelationRequest
	(*DeleteRelationResponse)(nil), // 7: aserto.directory.writer.v3.DeleteRelationResponse
	(*v3.Object)(nil),              // 8: aserto.directory.common.v3.Object
	(*emptypb.Empty)(nil),          // 9: google.protobuf.Empty
	(*v3.Relation)(nil),            // 10: aserto.directory.common.v3.Relation
}
var file_aserto_directory_writer_v3_writer_proto_depIdxs = []int32{
	8,  // 0: aserto.directory.writer.v3.SetObjectRequest.object:type_name -> aserto.directory.common.v3.Object
	8,  // 1: aserto.directory.writer.v3.SetObjectResponse.result:type_name -> aserto.directory.common.v3.Object
	9,  // 2: aserto.directory.writer.v3.DeleteObjectResponse.result:type_name -> google.protobuf.Empty
	10, // 3: aserto.directory.writer.v3.SetRelationRequest.relation:type_name -> aserto.directory.common.v3.Relation
	10, // 4: aserto.directory.writer.v3.SetRelationResponse.result:type_name -> aserto.directory.common.v3.Relation
	9,  // 5: aserto.directory.writer.v3.DeleteRelationResponse.result:type_name -> google.protobuf.Empty
	0,  // 6: aserto.directory.writer.v3.Writer.SetObject:input_type -> aserto.directory.writer.v3.SetObjectRequest
	2,  // 7: aserto.directory.writer.v3.Writer.DeleteObject:input_type -> aserto.directory.writer.v3.DeleteObjectRequest
	4,  // 8: aserto.directory.writer.v3.Writer.SetRelation:input_type -> aserto.directory.writer.v3.SetRelationRequest
	6,  // 9: aserto.directory.writer.v3.Writer.DeleteRelation:input_type -> aserto.directory.writer.v3.DeleteRelationRequest
	1,  // 10: aserto.directory.writer.v3.Writer.SetObject:output_type -> aserto.directory.writer.v3.SetObjectResponse
	3,  // 11: aserto.directory.writer.v3.Writer.DeleteObject:output_type -> aserto.directory.writer.v3.DeleteObjectResponse
	5,  // 12: aserto.directory.writer.v3.Writer.SetRelation:output_type -> aserto.directory.writer.v3.SetRelationResponse
	7,  // 13: aserto.directory.writer.v3.Writer.DeleteRelation:output_type -> aserto.directory.writer.v3.DeleteRelationResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_aserto_directory_writer_v3_writer_proto_init() }
func file_aserto_directory_writer_v3_writer_proto_init() {
	if File_aserto_directory_writer_v3_writer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_directory_writer_v3_writer_proto_rawDesc), len(file_aserto_directory_writer_v3_writer_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_writer_v3_writer_proto_goTypes,
		DependencyIndexes: file_aserto_directory_writer_v3_writer_proto_depIdxs,
		MessageInfos:      file_aserto_directory_writer_v3_writer_proto_msgTypes,
	}.Build()
	File_aserto_directory_writer_v3_writer_proto = out.File
	file_aserto_directory_writer_v3_writer_proto_goTypes = nil
	file_aserto_directory_writer_v3_writer_proto_depIdxs = nil
}
