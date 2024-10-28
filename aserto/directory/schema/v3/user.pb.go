// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: aserto/directory/schema/v3/user.proto

package schema

import (
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

type UserStatus int32

const (
	// User status undefined
	UserStatus_USER_STATUS_UNKNOWN UserStatus = 0
	// Staged status, is when the user object is first created, before the activation flow is initiated, or if there is a pending admin action.
	UserStatus_USER_STATUS_STAGED UserStatus = 1
	// Provisioned status, is when the user object is provisioned, but the user has not provided verification by clicking through the activation email or provided a password.
	UserStatus_USER_STATUS_PROVISIONED UserStatus = 2
	// Active status, is when:
	//   - An admin adds a user and sets the user password without requiring email verification.
	//   - An admin adds a user, sets the user password, and requires the user to set their password when they first sign-in.
	//   - A user self-registers into a custom app or IDP and email verification is not required.
	//   - An admin explicitly activates the user.
	UserStatus_USER_STATUS_ACTIVE UserStatus = 3
	// Recovery status, when the user requests a password reset or an admin initiates one on their behalf.
	UserStatus_USER_STATUS_RECOVERY UserStatus = 4
	// Password expired, status when the users' password has expired and the account requires an update to the password before a user is granted access.
	UserStatus_USER_STATUS_PASSWORD_EXPIRED UserStatus = 5
	// Locked out status, is when the user exceeds the number of login attempts defined in the login policy.
	UserStatus_USER_STATUS_LOCKED_OUT UserStatus = 6
	// Suspended status, when an admin explicitly suspends the user account.
	UserStatus_USER_STATUS_SUSPENDED UserStatus = 7
	// Deprovisioned status, is when an administrator explicitly deactivates or deprovisions/deletes the account.
	UserStatus_USER_STATUS_DEPROVISIONED UserStatus = 8
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "USER_STATUS_UNKNOWN",
		1: "USER_STATUS_STAGED",
		2: "USER_STATUS_PROVISIONED",
		3: "USER_STATUS_ACTIVE",
		4: "USER_STATUS_RECOVERY",
		5: "USER_STATUS_PASSWORD_EXPIRED",
		6: "USER_STATUS_LOCKED_OUT",
		7: "USER_STATUS_SUSPENDED",
		8: "USER_STATUS_DEPROVISIONED",
	}
	UserStatus_value = map[string]int32{
		"USER_STATUS_UNKNOWN":          0,
		"USER_STATUS_STAGED":           1,
		"USER_STATUS_PROVISIONED":      2,
		"USER_STATUS_ACTIVE":           3,
		"USER_STATUS_RECOVERY":         4,
		"USER_STATUS_PASSWORD_EXPIRED": 5,
		"USER_STATUS_LOCKED_OUT":       6,
		"USER_STATUS_SUSPENDED":        7,
		"USER_STATUS_DEPROVISIONED":    8,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_aserto_directory_schema_v3_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_aserto_directory_schema_v3_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_aserto_directory_schema_v3_user_proto_rawDescGZIP(), []int{0}
}

// Properties of "user" objects.
type UserProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// main email address of user
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// URL to user's picture
	Picture string `protobuf:"bytes,2,opt,name=picture,proto3" json:"picture,omitempty"`
	// user lifecycle status
	Status UserStatus `protobuf:"varint,3,opt,name=status,proto3,enum=aserto.directory.schema.v3.UserStatus" json:"status,omitempty"`
	// enabled (false prevents the user from accessing anything)
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// ID of the IDP connection the user instance is associated with.
	ConnectionId string `protobuf:"bytes,5,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
}

func (x *UserProperties) Reset() {
	*x = UserProperties{}
	mi := &file_aserto_directory_schema_v3_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProperties) ProtoMessage() {}

func (x *UserProperties) ProtoReflect() protoreflect.Message {
	mi := &file_aserto_directory_schema_v3_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProperties.ProtoReflect.Descriptor instead.
func (*UserProperties) Descriptor() ([]byte, []int) {
	return file_aserto_directory_schema_v3_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserProperties) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProperties) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *UserProperties) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_USER_STATUS_UNKNOWN
}

func (x *UserProperties) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UserProperties) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

var File_aserto_directory_schema_v3_user_proto protoreflect.FileDescriptor

var file_aserto_directory_schema_v3_user_proto_rawDesc = []byte{
	0x0a, 0x25, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x33, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x76, 0x33, 0x22, 0xbf, 0x01, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x76, 0x33, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x2a, 0x84, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a,
	0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x54, 0x41,
	0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x56, 0x45,
	0x52, 0x59, 0x10, 0x04, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54,
	0x10, 0x06, 0x12, 0x19, 0x0a, 0x15, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x07, 0x12, 0x1d, 0x0a,
	0x19, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x50,
	0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x08, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74,
	0x6f, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2f, 0x61, 0x73, 0x65, 0x72, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x33, 0x3b, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aserto_directory_schema_v3_user_proto_rawDescOnce sync.Once
	file_aserto_directory_schema_v3_user_proto_rawDescData = file_aserto_directory_schema_v3_user_proto_rawDesc
)

func file_aserto_directory_schema_v3_user_proto_rawDescGZIP() []byte {
	file_aserto_directory_schema_v3_user_proto_rawDescOnce.Do(func() {
		file_aserto_directory_schema_v3_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_aserto_directory_schema_v3_user_proto_rawDescData)
	})
	return file_aserto_directory_schema_v3_user_proto_rawDescData
}

var file_aserto_directory_schema_v3_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_aserto_directory_schema_v3_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_aserto_directory_schema_v3_user_proto_goTypes = []any{
	(UserStatus)(0),        // 0: aserto.directory.schema.v3.UserStatus
	(*UserProperties)(nil), // 1: aserto.directory.schema.v3.UserProperties
}
var file_aserto_directory_schema_v3_user_proto_depIdxs = []int32{
	0, // 0: aserto.directory.schema.v3.UserProperties.status:type_name -> aserto.directory.schema.v3.UserStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aserto_directory_schema_v3_user_proto_init() }
func file_aserto_directory_schema_v3_user_proto_init() {
	if File_aserto_directory_schema_v3_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aserto_directory_schema_v3_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aserto_directory_schema_v3_user_proto_goTypes,
		DependencyIndexes: file_aserto_directory_schema_v3_user_proto_depIdxs,
		EnumInfos:         file_aserto_directory_schema_v3_user_proto_enumTypes,
		MessageInfos:      file_aserto_directory_schema_v3_user_proto_msgTypes,
	}.Build()
	File_aserto_directory_schema_v3_user_proto = out.File
	file_aserto_directory_schema_v3_user_proto_rawDesc = nil
	file_aserto_directory_schema_v3_user_proto_goTypes = nil
	file_aserto_directory_schema_v3_user_proto_depIdxs = nil
}
