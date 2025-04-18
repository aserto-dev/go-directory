// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/directory/schema/v3/group.proto

package schema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Properties of "group" objects.
type GroupProperties struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the IDP connection the group instance is associated with.
	ConnectionId  string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GroupProperties) Reset() {
	*x = GroupProperties{}
	mi := &file_aserto_directory_schema_v3_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupProperties) ProtoMessage() {}

func (x *GroupProperties) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_schema_v3_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupProperties.ProtoReflect.Descriptor instead.
func (*GroupProperties) Descriptor() ([]byte, []int) {
	return file_aserto_directory_schema_v3_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupProperties) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

var File_aserto_directory_schema_v3_group_proto protoreflect.FileDescriptor

const file_aserto_directory_schema_v3_group_proto_rawDesc = "" +
	"\n" +
	"&aserto/directory/schema/v3/group.proto\x12\x1aaserto.directory.schema.v3\"6\n" +
	"\x0fGroupProperties\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionIdBFZDgithub.com/aserto-dev/go-directory/aserto/directory/schema/v3;schemab\x06proto3"

var (
	file_aserto_directory_schema_v3_group_proto_rawDescOnce sync.Once
	file_aserto_directory_schema_v3_group_proto_rawDescData []byte
)

func file_aserto_directory_schema_v3_group_proto_rawDescGZIP() []byte {
	file_aserto_directory_schema_v3_group_proto_rawDescOnce.Do(func() {
		file_aserto_directory_schema_v3_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_directory_schema_v3_group_proto_rawDesc), len(file_aserto_directory_schema_v3_group_proto_rawDesc)))
	})
	return file_aserto_directory_schema_v3_group_proto_rawDescData
}

var file_aserto_directory_schema_v3_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_directory_schema_v3_group_proto_goTypes = []any{
	(*GroupProperties)(nil), // 0: aserto.directory.schema.v3.GroupProperties
}
var file_aserto_directory_schema_v3_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aserto_directory_schema_v3_group_proto_init() }
func file_aserto_directory_schema_v3_group_proto_init() {
	if File_aserto_directory_schema_v3_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_directory_schema_v3_group_proto_rawDesc), len(file_aserto_directory_schema_v3_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_directory_schema_v3_group_proto_goTypes,
		DependencyIndexes: file_aserto_directory_schema_v3_group_proto_depIdxs,
		MessageInfos:      file_aserto_directory_schema_v3_group_proto_msgTypes,
	}.Build()
	File_aserto_directory_schema_v3_group_proto = out.File
	file_aserto_directory_schema_v3_group_proto_goTypes = nil
	file_aserto_directory_schema_v3_group_proto_depIdxs = nil
}
