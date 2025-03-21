// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: aserto/directory/exporter/v3/exporter.proto

package exporter

import (
	v3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
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

type Option int32

const (
	// nothing selected (default initialization value)
	Option_OPTION_UNKNOWN Option = 0
	// object instances
	Option_OPTION_DATA_OBJECTS Option = 8
	// relation instances
	Option_OPTION_DATA_RELATIONS Option = 16
	// all data = OPTION_DATA_OBJECTS | OPTION_DATA_RELATIONS
	Option_OPTION_DATA Option = 24
	// stats
	Option_OPTION_STATS Option = 64
)

// Enum value maps for Option.
var (
	Option_name = map[int32]string{
		0:  "OPTION_UNKNOWN",
		8:  "OPTION_DATA_OBJECTS",
		16: "OPTION_DATA_RELATIONS",
		24: "OPTION_DATA",
		64: "OPTION_STATS",
	}
	Option_value = map[string]int32{
		"OPTION_UNKNOWN":        0,
		"OPTION_DATA_OBJECTS":   8,
		"OPTION_DATA_RELATIONS": 16,
		"OPTION_DATA":           24,
		"OPTION_STATS":          64,
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
	return file_aserto_directory_exporter_v3_exporter_proto_enumTypes[0].Descriptor()
}

func (Option) Type() protoreflect.EnumType {
	return &file_aserto_directory_exporter_v3_exporter_proto_enumTypes[0]
}

func (x Option) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Option.Descriptor instead.
func (Option) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_exporter_v3_exporter_proto_rawDescGZIP(), []int{0}
}

type ExportRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// data export options mask
	Options uint32 `protobuf:"varint,1,opt,name=options,proto3" json:"options,omitempty"`
	// start export from timestamp (UTC)
	StartFrom     *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=start_from,json=startFrom,proto3" json:"start_from,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	mi := &file_aserto_directory_exporter_v3_exporter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_exporter_v3_exporter_proto_msgTypes[0]
	if x != nil {
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
	return file_aserto_directory_exporter_v3_exporter_proto_rawDescGZIP(), []int{0}
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

type ExportResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Msg:
	//
	//	*ExportResponse_Object
	//	*ExportResponse_Relation
	//	*ExportResponse_Stats
	Msg           isExportResponse_Msg `protobuf_oneof:"msg"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	mi := &file_aserto_directory_exporter_v3_exporter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_exporter_v3_exporter_proto_msgTypes[1]
	if x != nil {
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
	return file_aserto_directory_exporter_v3_exporter_proto_rawDescGZIP(), []int{1}
}

func (x *ExportResponse) GetMsg() isExportResponse_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *ExportResponse) GetObject() *v3.Object {
	if x != nil {
		if x, ok := x.Msg.(*ExportResponse_Object); ok {
			return x.Object
		}
	}
	return nil
}

func (x *ExportResponse) GetRelation() *v3.Relation {
	if x != nil {
		if x, ok := x.Msg.(*ExportResponse_Relation); ok {
			return x.Relation
		}
	}
	return nil
}

func (x *ExportResponse) GetStats() *structpb.Struct {
	if x != nil {
		if x, ok := x.Msg.(*ExportResponse_Stats); ok {
			return x.Stats
		}
	}
	return nil
}

type isExportResponse_Msg interface {
	isExportResponse_Msg()
}

type ExportResponse_Object struct {
	// object instance (data)
	Object *v3.Object `protobuf:"bytes,2,opt,name=object,proto3,oneof"`
}

type ExportResponse_Relation struct {
	// relation instance (data)
	Relation *v3.Relation `protobuf:"bytes,4,opt,name=relation,proto3,oneof"`
}

type ExportResponse_Stats struct {
	// object and/or relation stats (no data)
	Stats *structpb.Struct `protobuf:"bytes,8,opt,name=stats,proto3,oneof"`
}

func (*ExportResponse_Object) isExportResponse_Msg() {}

func (*ExportResponse_Relation) isExportResponse_Msg() {}

func (*ExportResponse_Stats) isExportResponse_Msg() {}

var File_aserto_directory_exporter_v3_exporter_proto protoreflect.FileDescriptor

var file_aserto_directory_exporter_v3_exporter_proto_rawDesc = string([]byte{
	0x0a, 0x2b, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x1a, 0x27, 0x61, 0x73, 0x65,
	0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x22, 0xca, 0x01, 0x0a, 0x0e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42,
	0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x73, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x53, 0x10, 0x08, 0x12, 0x19, 0x0a,
	0x15, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x52, 0x45, 0x4c,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x18, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x10, 0x40, 0x32, 0x73, 0x0a, 0x08, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x67, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x2b, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33,
	0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x45, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x73, 0x65, 0x72, 0x74, 0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2f, 0x76, 0x33, 0x3b, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_aserto_directory_exporter_v3_exporter_proto_rawDescOnce sync.Once
	file_aserto_directory_exporter_v3_exporter_proto_rawDescData []byte
)

func file_aserto_directory_exporter_v3_exporter_proto_rawDescGZIP() []byte {
	file_aserto_directory_exporter_v3_exporter_proto_rawDescOnce.Do(func() {
		file_aserto_directory_exporter_v3_exporter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_directory_exporter_v3_exporter_proto_rawDesc), len(file_aserto_directory_exporter_v3_exporter_proto_rawDesc)))
	})
	return file_aserto_directory_exporter_v3_exporter_proto_rawDescData
}

var file_aserto_directory_exporter_v3_exporter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_exporter_v3_exporter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aserto_directory_exporter_v3_exporter_proto_goTypes = []any{
	(Option)(0),                   // 0: aserto.directory.exporter.v3.Option
	(*ExportRequest)(nil),         // 1: aserto.directory.exporter.v3.ExportRequest
	(*ExportResponse)(nil),        // 2: aserto.directory.exporter.v3.ExportResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*v3.Object)(nil),             // 4: aserto.directory.common.v3.Object
	(*v3.Relation)(nil),           // 5: aserto.directory.common.v3.Relation
	(*structpb.Struct)(nil),       // 6: google.protobuf.Struct
}
var file_aserto_directory_exporter_v3_exporter_proto_depIdxs = []int32{
	3, // 0: aserto.directory.exporter.v3.ExportRequest.start_from:type_name -> google.protobuf.Timestamp
	4, // 1: aserto.directory.exporter.v3.ExportResponse.object:type_name -> aserto.directory.common.v3.Object
	5, // 2: aserto.directory.exporter.v3.ExportResponse.relation:type_name -> aserto.directory.common.v3.Relation
	6, // 3: aserto.directory.exporter.v3.ExportResponse.stats:type_name -> google.protobuf.Struct
	1, // 4: aserto.directory.exporter.v3.Exporter.Export:input_type -> aserto.directory.exporter.v3.ExportRequest
	2, // 5: aserto.directory.exporter.v3.Exporter.Export:output_type -> aserto.directory.exporter.v3.ExportResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aserto_directory_exporter_v3_exporter_proto_init() }
func file_aserto_directory_exporter_v3_exporter_proto_init() {
	if File_aserto_directory_exporter_v3_exporter_proto != nil {
		return
	}
	file_aserto_directory_exporter_v3_exporter_proto_msgTypes[1].OneofWrappers = []any{
		(*ExportResponse_Object)(nil),
		(*ExportResponse_Relation)(nil),
		(*ExportResponse_Stats)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_directory_exporter_v3_exporter_proto_rawDesc), len(file_aserto_directory_exporter_v3_exporter_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_exporter_v3_exporter_proto_goTypes,
		DependencyIndexes: file_aserto_directory_exporter_v3_exporter_proto_depIdxs,
		EnumInfos:         file_aserto_directory_exporter_v3_exporter_proto_enumTypes,
		MessageInfos:      file_aserto_directory_exporter_v3_exporter_proto_msgTypes,
	}.Build()
	File_aserto_directory_exporter_v3_exporter_proto = out.File
	file_aserto_directory_exporter_v3_exporter_proto_goTypes = nil
	file_aserto_directory_exporter_v3_exporter_proto_depIdxs = nil
}
