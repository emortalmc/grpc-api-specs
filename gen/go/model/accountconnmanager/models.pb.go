// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: account_conn_manager/models.proto

package accountconnmanager

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

type ConnectionType int32

const (
	ConnectionType_MINECRAFT ConnectionType = 0
	ConnectionType_DISCORD   ConnectionType = 1
	ConnectionType_GITHUB    ConnectionType = 2
)

// Enum value maps for ConnectionType.
var (
	ConnectionType_name = map[int32]string{
		0: "MINECRAFT",
		1: "DISCORD",
		2: "GITHUB",
	}
	ConnectionType_value = map[string]int32{
		"MINECRAFT": 0,
		"DISCORD":   1,
		"GITHUB":    2,
	}
)

func (x ConnectionType) Enum() *ConnectionType {
	p := new(ConnectionType)
	*p = x
	return p
}

func (x ConnectionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnectionType) Descriptor() protoreflect.EnumDescriptor {
	return file_account_conn_manager_models_proto_enumTypes[0].Descriptor()
}

func (ConnectionType) Type() protoreflect.EnumType {
	return &file_account_conn_manager_models_proto_enumTypes[0]
}

func (x ConnectionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnectionType.Descriptor instead.
func (ConnectionType) EnumDescriptor() ([]byte, []int) {
	return file_account_conn_manager_models_proto_rawDescGZIP(), []int{0}
}

type ConnectionUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// minecraft_id of type UUID
	MinecraftId *string `protobuf:"bytes,1,opt,name=minecraft_id,json=minecraftId,proto3,oneof" json:"minecraft_id,omitempty"`
	// discord_id of type discord snowflake
	DiscordId *int64 `protobuf:"varint,2,opt,name=discord_id,json=discordId,proto3,oneof" json:"discord_id,omitempty"`
	// github_id is a plain incremental integer
	GithubId *int64 `protobuf:"varint,3,opt,name=github_id,json=githubId,proto3,oneof" json:"github_id,omitempty"`
}

func (x *ConnectionUser) Reset() {
	*x = ConnectionUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionUser) ProtoMessage() {}

func (x *ConnectionUser) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionUser.ProtoReflect.Descriptor instead.
func (*ConnectionUser) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_models_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionUser) GetMinecraftId() string {
	if x != nil && x.MinecraftId != nil {
		return *x.MinecraftId
	}
	return ""
}

func (x *ConnectionUser) GetDiscordId() int64 {
	if x != nil && x.DiscordId != nil {
		return *x.DiscordId
	}
	return 0
}

func (x *ConnectionUser) GetGithubId() int64 {
	if x != nil && x.GithubId != nil {
		return *x.GithubId
	}
	return 0
}

var File_account_conn_manager_models_proto protoreflect.FileDescriptor

var file_account_conn_manager_models_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x22, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22, 0xac, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x6d, 0x69,
	0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x49, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x02, 0x52, 0x08, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x6d, 0x69, 0x6e,
	0x65, 0x63, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x5f, 0x69, 0x64, 0x2a, 0x38, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x4d, 0x49, 0x4e, 0x45,
	0x43, 0x52, 0x41, 0x46, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x49, 0x53, 0x43, 0x4f,
	0x52, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x49, 0x54, 0x48, 0x55, 0x42, 0x10, 0x02,
	0x42, 0x6e, 0x0a, 0x28, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x40,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_conn_manager_models_proto_rawDescOnce sync.Once
	file_account_conn_manager_models_proto_rawDescData = file_account_conn_manager_models_proto_rawDesc
)

func file_account_conn_manager_models_proto_rawDescGZIP() []byte {
	file_account_conn_manager_models_proto_rawDescOnce.Do(func() {
		file_account_conn_manager_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_conn_manager_models_proto_rawDescData)
	})
	return file_account_conn_manager_models_proto_rawDescData
}

var file_account_conn_manager_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_conn_manager_models_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_conn_manager_models_proto_goTypes = []interface{}{
	(ConnectionType)(0),    // 0: emortal.model.account_conn_manager.ConnectionType
	(*ConnectionUser)(nil), // 1: emortal.model.account_conn_manager.ConnectionUser
}
var file_account_conn_manager_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_conn_manager_models_proto_init() }
func file_account_conn_manager_models_proto_init() {
	if File_account_conn_manager_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_conn_manager_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionUser); i {
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
	file_account_conn_manager_models_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_conn_manager_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_conn_manager_models_proto_goTypes,
		DependencyIndexes: file_account_conn_manager_models_proto_depIdxs,
		EnumInfos:         file_account_conn_manager_models_proto_enumTypes,
		MessageInfos:      file_account_conn_manager_models_proto_msgTypes,
	}.Build()
	File_account_conn_manager_models_proto = out.File
	file_account_conn_manager_models_proto_rawDesc = nil
	file_account_conn_manager_models_proto_goTypes = nil
	file_account_conn_manager_models_proto_depIdxs = nil
}
