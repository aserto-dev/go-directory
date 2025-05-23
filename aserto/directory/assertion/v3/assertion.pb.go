// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: aserto/directory/assertion/v3/assertion.proto

package assertion

import (
	v3 "github.com/aserto-dev/go-directory/aserto/directory/common/v3"
	v31 "github.com/aserto-dev/go-directory/aserto/directory/reader/v3"
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

type GetAssertionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// assertion identifier
	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssertionRequest) Reset() {
	*x = GetAssertionRequest{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssertionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssertionRequest) ProtoMessage() {}

func (x *GetAssertionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssertionRequest.ProtoReflect.Descriptor instead.
func (*GetAssertionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{0}
}

func (x *GetAssertionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAssertionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *Assert                `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssertionResponse) Reset() {
	*x = GetAssertionResponse{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssertionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssertionResponse) ProtoMessage() {}

func (x *GetAssertionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssertionResponse.ProtoReflect.Descriptor instead.
func (*GetAssertionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssertionResponse) GetResult() *Assert {
	if x != nil {
		return x.Result
	}
	return nil
}

type ListAssertionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// pagination request
	Page          *v3.PaginationRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssertionsRequest) Reset() {
	*x = ListAssertionsRequest{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssertionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssertionsRequest) ProtoMessage() {}

func (x *ListAssertionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssertionsRequest.ProtoReflect.Descriptor instead.
func (*ListAssertionsRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{2}
}

func (x *ListAssertionsRequest) GetPage() *v3.PaginationRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListAssertionsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Results       []*Assert              `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Page          *v3.PaginationResponse `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssertionsResponse) Reset() {
	*x = ListAssertionsResponse{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssertionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssertionsResponse) ProtoMessage() {}

func (x *ListAssertionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssertionsResponse.ProtoReflect.Descriptor instead.
func (*ListAssertionsResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{3}
}

func (x *ListAssertionsResponse) GetResults() []*Assert {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListAssertionsResponse) GetPage() *v3.PaginationResponse {
	if x != nil {
		return x.Page
	}
	return nil
}

type SetAssertionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Assert        *Assert                `protobuf:"bytes,1,opt,name=assert,proto3" json:"assert,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAssertionRequest) Reset() {
	*x = SetAssertionRequest{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAssertionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAssertionRequest) ProtoMessage() {}

func (x *SetAssertionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAssertionRequest.ProtoReflect.Descriptor instead.
func (*SetAssertionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{4}
}

func (x *SetAssertionRequest) GetAssert() *Assert {
	if x != nil {
		return x.Assert
	}
	return nil
}

type SetAssertionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        *Assert                `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAssertionResponse) Reset() {
	*x = SetAssertionResponse{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAssertionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAssertionResponse) ProtoMessage() {}

func (x *SetAssertionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAssertionResponse.ProtoReflect.Descriptor instead.
func (*SetAssertionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{5}
}

func (x *SetAssertionResponse) GetResult() *Assert {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteAssertionRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// assertion identifier
	Id            uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssertionRequest) Reset() {
	*x = DeleteAssertionRequest{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssertionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssertionRequest) ProtoMessage() {}

func (x *DeleteAssertionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssertionRequest.ProtoReflect.Descriptor instead.
func (*DeleteAssertionRequest) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAssertionRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteAssertionResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// empty result
	Result        *emptypb.Empty `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssertionResponse) Reset() {
	*x = DeleteAssertionResponse{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssertionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssertionResponse) ProtoMessage() {}

func (x *DeleteAssertionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssertionResponse.ProtoReflect.Descriptor instead.
func (*DeleteAssertionResponse) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAssertionResponse) GetResult() *emptypb.Empty {
	if x != nil {
		return x.Result
	}
	return nil
}

type Assert struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// assertion identifier
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// expected outcome of assertion
	Expected bool `protobuf:"varint,2,opt,name=expected,proto3" json:"expected,omitempty"`
	// assertion request
	//
	// Types that are valid to be assigned to Msg:
	//
	//	*Assert_Check
	//	*Assert_CheckRelation
	//	*Assert_CheckPermission
	Msg isAssert_Msg `protobuf_oneof:"msg"`
	// description
	Description   string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Assert) Reset() {
	*x = Assert{}
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Assert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assert) ProtoMessage() {}

func (x *Assert) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_assertion_v3_assertion_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assert.ProtoReflect.Descriptor instead.
func (*Assert) Descriptor() ([]byte, []int) {
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP(), []int{8}
}

func (x *Assert) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Assert) GetExpected() bool {
	if x != nil {
		return x.Expected
	}
	return false
}

func (x *Assert) GetMsg() isAssert_Msg {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Assert) GetCheck() *v31.CheckRequest {
	if x != nil {
		if x, ok := x.Msg.(*Assert_Check); ok {
			return x.Check
		}
	}
	return nil
}

func (x *Assert) GetCheckRelation() *v31.CheckRelationRequest {
	if x != nil {
		if x, ok := x.Msg.(*Assert_CheckRelation); ok {
			return x.CheckRelation
		}
	}
	return nil
}

func (x *Assert) GetCheckPermission() *v31.CheckPermissionRequest {
	if x != nil {
		if x, ok := x.Msg.(*Assert_CheckPermission); ok {
			return x.CheckPermission
		}
	}
	return nil
}

func (x *Assert) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type isAssert_Msg interface {
	isAssert_Msg()
}

type Assert_Check struct {
	Check *v31.CheckRequest `protobuf:"bytes,3,opt,name=check,proto3,oneof"`
}

type Assert_CheckRelation struct {
	CheckRelation *v31.CheckRelationRequest `protobuf:"bytes,4,opt,name=check_relation,json=checkRelation,proto3,oneof"`
}

type Assert_CheckPermission struct {
	CheckPermission *v31.CheckPermissionRequest `protobuf:"bytes,5,opt,name=check_permission,json=checkPermission,proto3,oneof"`
}

func (*Assert_Check) isAssert_Msg() {}

func (*Assert_CheckRelation) isAssert_Msg() {}

func (*Assert_CheckPermission) isAssert_Msg() {}

var File_aserto_directory_assertion_v3_assertion_proto protoreflect.FileDescriptor

const file_aserto_directory_assertion_v3_assertion_proto_rawDesc = "" +
	"\n" +
	"-aserto/directory/assertion/v3/assertion.proto\x12\x1daserto.directory.assertion.v3\x1a'aserto/directory/common/v3/common.proto\x1a'aserto/directory/reader/v3/reader.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"*\n" +
	"\x13GetAssertionRequest\x12\x13\n" +
	"\x02id\x18\x01 \x01(\rB\x03\xe0A\x02R\x02id\"U\n" +
	"\x14GetAssertionResponse\x12=\n" +
	"\x06result\x18\x01 \x01(\v2%.aserto.directory.assertion.v3.AssertR\x06result\"_\n" +
	"\x15ListAssertionsRequest\x12F\n" +
	"\x04page\x18\x01 \x01(\v2-.aserto.directory.common.v3.PaginationRequestB\x03\xe0A\x01R\x04page\"\x9d\x01\n" +
	"\x16ListAssertionsResponse\x12?\n" +
	"\aresults\x18\x01 \x03(\v2%.aserto.directory.assertion.v3.AssertR\aresults\x12B\n" +
	"\x04page\x18\x02 \x01(\v2..aserto.directory.common.v3.PaginationResponseR\x04page\"Y\n" +
	"\x13SetAssertionRequest\x12B\n" +
	"\x06assert\x18\x01 \x01(\v2%.aserto.directory.assertion.v3.AssertB\x03\xe0A\x02R\x06assert\"U\n" +
	"\x14SetAssertionResponse\x12=\n" +
	"\x06result\x18\x01 \x01(\v2%.aserto.directory.assertion.v3.AssertR\x06result\"-\n" +
	"\x16DeleteAssertionRequest\x12\x13\n" +
	"\x02id\x18\x01 \x01(\rB\x03\xe0A\x02R\x02id\"I\n" +
	"\x17DeleteAssertionResponse\x12.\n" +
	"\x06result\x18\x01 \x01(\v2\x16.google.protobuf.EmptyR\x06result\"\xe0\x02\n" +
	"\x06Assert\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\rR\x02id\x12\x1f\n" +
	"\bexpected\x18\x02 \x01(\bB\x03\xe0A\x02R\bexpected\x12@\n" +
	"\x05check\x18\x03 \x01(\v2(.aserto.directory.reader.v3.CheckRequestH\x00R\x05check\x12Y\n" +
	"\x0echeck_relation\x18\x04 \x01(\v20.aserto.directory.reader.v3.CheckRelationRequestH\x00R\rcheckRelation\x12_\n" +
	"\x10check_permission\x18\x05 \x01(\v22.aserto.directory.reader.v3.CheckPermissionRequestH\x00R\x0fcheckPermission\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescriptionB\x05\n" +
	"\x03msg2\xca\t\n" +
	"\tAssertion\x12\xb8\x02\n" +
	"\fGetAssertion\x122.aserto.directory.assertion.v3.GetAssertionRequest\x1a3.aserto.directory.assertion.v3.GetAssertionResponse\"\xbe\x01\x92A\x92\x01\n" +
	"\tdirectory\x12\x16Get assertion instance\x1a\"Returns single assertion instance.*$directory.assertion.v3.assertion.getb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\"\x12 /api/v3/directory/assertion/{id}\x12\xb7\x02\n" +
	"\x0eListAssertions\x124.aserto.directory.assertion.v3.ListAssertionsRequest\x1a5.aserto.directory.assertion.v3.ListAssertionsResponse\"\xb7\x01\x92A\x8f\x01\n" +
	"\tdirectory\x12\x0fList assertions\x1a$Returns list of assertion instances.*&directory.assertion.v3.assertions.listb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\x1e\x12\x1c/api/v3/directory/assertions\x12\x98\x02\n" +
	"\fSetAssertion\x122.aserto.directory.assertion.v3.SetAssertionRequest\x1a3.aserto.directory.assertion.v3.SetAssertionResponse\"\x9e\x01\x92Au\n" +
	"\tdirectory\x12\rSet assertion\x1a\x0eSet assertion.*$directory.assertion.v3.assertion.setb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02 :\x01*\"\x1b/api/v3/directory/assertion\x12\xac\x02\n" +
	"\x0fDeleteAssertion\x125.aserto.directory.assertion.v3.DeleteAssertionRequest\x1a6.aserto.directory.assertion.v3.DeleteAssertionResponse\"\xa9\x01\x92A~\n" +
	"\tdirectory\x12\x10Delete assertion\x1a\x11Delete assertion.*'directory.assertion.v3.assertion.deleteb#\n" +
	"\x13\n" +
	"\x0fDirectoryAPIKey\x12\x00\n" +
	"\f\n" +
	"\bTenantID\x12\x00\x82\xd3\xe4\x93\x02\"* /api/v3/directory/assertion/{id}BLZJgithub.com/aserto-dev/go-directory/aserto/directory/assertion/v3;assertionb\x06proto3"

var (
	file_aserto_directory_assertion_v3_assertion_proto_rawDescOnce sync.Once
	file_aserto_directory_assertion_v3_assertion_proto_rawDescData []byte
)

func file_aserto_directory_assertion_v3_assertion_proto_rawDescGZIP() []byte {
	file_aserto_directory_assertion_v3_assertion_proto_rawDescOnce.Do(func() {
		file_aserto_directory_assertion_v3_assertion_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_aserto_directory_assertion_v3_assertion_proto_rawDesc), len(file_aserto_directory_assertion_v3_assertion_proto_rawDesc)))
	})
	return file_aserto_directory_assertion_v3_assertion_proto_rawDescData
}

var file_aserto_directory_assertion_v3_assertion_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_aserto_directory_assertion_v3_assertion_proto_goTypes = []any{
	(*GetAssertionRequest)(nil),        // 0: aserto.directory.assertion.v3.GetAssertionRequest
	(*GetAssertionResponse)(nil),       // 1: aserto.directory.assertion.v3.GetAssertionResponse
	(*ListAssertionsRequest)(nil),      // 2: aserto.directory.assertion.v3.ListAssertionsRequest
	(*ListAssertionsResponse)(nil),     // 3: aserto.directory.assertion.v3.ListAssertionsResponse
	(*SetAssertionRequest)(nil),        // 4: aserto.directory.assertion.v3.SetAssertionRequest
	(*SetAssertionResponse)(nil),       // 5: aserto.directory.assertion.v3.SetAssertionResponse
	(*DeleteAssertionRequest)(nil),     // 6: aserto.directory.assertion.v3.DeleteAssertionRequest
	(*DeleteAssertionResponse)(nil),    // 7: aserto.directory.assertion.v3.DeleteAssertionResponse
	(*Assert)(nil),                     // 8: aserto.directory.assertion.v3.Assert
	(*v3.PaginationRequest)(nil),       // 9: aserto.directory.common.v3.PaginationRequest
	(*v3.PaginationResponse)(nil),      // 10: aserto.directory.common.v3.PaginationResponse
	(*emptypb.Empty)(nil),              // 11: google.protobuf.Empty
	(*v31.CheckRequest)(nil),           // 12: aserto.directory.reader.v3.CheckRequest
	(*v31.CheckRelationRequest)(nil),   // 13: aserto.directory.reader.v3.CheckRelationRequest
	(*v31.CheckPermissionRequest)(nil), // 14: aserto.directory.reader.v3.CheckPermissionRequest
}
var file_aserto_directory_assertion_v3_assertion_proto_depIdxs = []int32{
	8,  // 0: aserto.directory.assertion.v3.GetAssertionResponse.result:type_name -> aserto.directory.assertion.v3.Assert
	9,  // 1: aserto.directory.assertion.v3.ListAssertionsRequest.page:type_name -> aserto.directory.common.v3.PaginationRequest
	8,  // 2: aserto.directory.assertion.v3.ListAssertionsResponse.results:type_name -> aserto.directory.assertion.v3.Assert
	10, // 3: aserto.directory.assertion.v3.ListAssertionsResponse.page:type_name -> aserto.directory.common.v3.PaginationResponse
	8,  // 4: aserto.directory.assertion.v3.SetAssertionRequest.assert:type_name -> aserto.directory.assertion.v3.Assert
	8,  // 5: aserto.directory.assertion.v3.SetAssertionResponse.result:type_name -> aserto.directory.assertion.v3.Assert
	11, // 6: aserto.directory.assertion.v3.DeleteAssertionResponse.result:type_name -> google.protobuf.Empty
	12, // 7: aserto.directory.assertion.v3.Assert.check:type_name -> aserto.directory.reader.v3.CheckRequest
	13, // 8: aserto.directory.assertion.v3.Assert.check_relation:type_name -> aserto.directory.reader.v3.CheckRelationRequest
	14, // 9: aserto.directory.assertion.v3.Assert.check_permission:type_name -> aserto.directory.reader.v3.CheckPermissionRequest
	0,  // 10: aserto.directory.assertion.v3.Assertion.GetAssertion:input_type -> aserto.directory.assertion.v3.GetAssertionRequest
	2,  // 11: aserto.directory.assertion.v3.Assertion.ListAssertions:input_type -> aserto.directory.assertion.v3.ListAssertionsRequest
	4,  // 12: aserto.directory.assertion.v3.Assertion.SetAssertion:input_type -> aserto.directory.assertion.v3.SetAssertionRequest
	6,  // 13: aserto.directory.assertion.v3.Assertion.DeleteAssertion:input_type -> aserto.directory.assertion.v3.DeleteAssertionRequest
	1,  // 14: aserto.directory.assertion.v3.Assertion.GetAssertion:output_type -> aserto.directory.assertion.v3.GetAssertionResponse
	3,  // 15: aserto.directory.assertion.v3.Assertion.ListAssertions:output_type -> aserto.directory.assertion.v3.ListAssertionsResponse
	5,  // 16: aserto.directory.assertion.v3.Assertion.SetAssertion:output_type -> aserto.directory.assertion.v3.SetAssertionResponse
	7,  // 17: aserto.directory.assertion.v3.Assertion.DeleteAssertion:output_type -> aserto.directory.assertion.v3.DeleteAssertionResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_aserto_directory_assertion_v3_assertion_proto_init() }
func file_aserto_directory_assertion_v3_assertion_proto_init() {
	if File_aserto_directory_assertion_v3_assertion_proto != nil {
		return
	}
	file_aserto_directory_assertion_v3_assertion_proto_msgTypes[8].OneofWrappers = []any{
		(*Assert_Check)(nil),
		(*Assert_CheckRelation)(nil),
		(*Assert_CheckPermission)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_aserto_directory_assertion_v3_assertion_proto_rawDesc), len(file_aserto_directory_assertion_v3_assertion_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aserto_directory_assertion_v3_assertion_proto_goTypes,
		DependencyIndexes: file_aserto_directory_assertion_v3_assertion_proto_depIdxs,
		MessageInfos:      file_aserto_directory_assertion_v3_assertion_proto_msgTypes,
	}.Build()
	File_aserto_directory_assertion_v3_assertion_proto = out.File
	file_aserto_directory_assertion_v3_assertion_proto_goTypes = nil
	file_aserto_directory_assertion_v3_assertion_proto_depIdxs = nil
}
