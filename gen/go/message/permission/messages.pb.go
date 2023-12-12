// Kafka Topic = "permission-manager"

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: permission/messages.proto

package permission

import (
	permission "github.com/emortalmc/proto-specs/gen/go/model/permission"
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

type RoleUpdateMessage_ChangeType int32

const (
	RoleUpdateMessage_CREATE RoleUpdateMessage_ChangeType = 0
	RoleUpdateMessage_DELETE RoleUpdateMessage_ChangeType = 1
	RoleUpdateMessage_MODIFY RoleUpdateMessage_ChangeType = 2
)

// Enum value maps for RoleUpdateMessage_ChangeType.
var (
	RoleUpdateMessage_ChangeType_name = map[int32]string{
		0: "CREATE",
		1: "DELETE",
		2: "MODIFY",
	}
	RoleUpdateMessage_ChangeType_value = map[string]int32{
		"CREATE": 0,
		"DELETE": 1,
		"MODIFY": 2,
	}
)

func (x RoleUpdateMessage_ChangeType) Enum() *RoleUpdateMessage_ChangeType {
	p := new(RoleUpdateMessage_ChangeType)
	*p = x
	return p
}

func (x RoleUpdateMessage_ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleUpdateMessage_ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_permission_messages_proto_enumTypes[0].Descriptor()
}

func (RoleUpdateMessage_ChangeType) Type() protoreflect.EnumType {
	return &file_permission_messages_proto_enumTypes[0]
}

func (x RoleUpdateMessage_ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleUpdateMessage_ChangeType.Descriptor instead.
func (RoleUpdateMessage_ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{0, 0}
}

type PlayerRolesUpdateMessage_ChangeType int32

const (
	PlayerRolesUpdateMessage_ADD    PlayerRolesUpdateMessage_ChangeType = 0
	PlayerRolesUpdateMessage_REMOVE PlayerRolesUpdateMessage_ChangeType = 1
)

// Enum value maps for PlayerRolesUpdateMessage_ChangeType.
var (
	PlayerRolesUpdateMessage_ChangeType_name = map[int32]string{
		0: "ADD",
		1: "REMOVE",
	}
	PlayerRolesUpdateMessage_ChangeType_value = map[string]int32{
		"ADD":    0,
		"REMOVE": 1,
	}
)

func (x PlayerRolesUpdateMessage_ChangeType) Enum() *PlayerRolesUpdateMessage_ChangeType {
	p := new(PlayerRolesUpdateMessage_ChangeType)
	*p = x
	return p
}

func (x PlayerRolesUpdateMessage_ChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerRolesUpdateMessage_ChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_permission_messages_proto_enumTypes[1].Descriptor()
}

func (PlayerRolesUpdateMessage_ChangeType) Type() protoreflect.EnumType {
	return &file_permission_messages_proto_enumTypes[1]
}

func (x PlayerRolesUpdateMessage_ChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerRolesUpdateMessage_ChangeType.Descriptor instead.
func (PlayerRolesUpdateMessage_ChangeType) EnumDescriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{1, 0}
}

type RoleUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangeType RoleUpdateMessage_ChangeType `protobuf:"varint,1,opt,name=change_type,json=changeType,proto3,enum=emortal.message.permission.RoleUpdateMessage_ChangeType" json:"change_type,omitempty"`
	// role either the deleted, created or updated role
	Role *permission.Role `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *RoleUpdateMessage) Reset() {
	*x = RoleUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUpdateMessage) ProtoMessage() {}

func (x *RoleUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUpdateMessage.ProtoReflect.Descriptor instead.
func (*RoleUpdateMessage) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{0}
}

func (x *RoleUpdateMessage) GetChangeType() RoleUpdateMessage_ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return RoleUpdateMessage_CREATE
}

func (x *RoleUpdateMessage) GetRole() *permission.Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type PlayerRolesUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChangeType PlayerRolesUpdateMessage_ChangeType `protobuf:"varint,1,opt,name=change_type,json=changeType,proto3,enum=emortal.message.permission.PlayerRolesUpdateMessage_ChangeType" json:"change_type,omitempty"`
	PlayerId   string                              `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	RoleId     string                              `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
}

func (x *PlayerRolesUpdateMessage) Reset() {
	*x = PlayerRolesUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_permission_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerRolesUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerRolesUpdateMessage) ProtoMessage() {}

func (x *PlayerRolesUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_permission_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerRolesUpdateMessage.ProtoReflect.Descriptor instead.
func (*PlayerRolesUpdateMessage) Descriptor() ([]byte, []int) {
	return file_permission_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerRolesUpdateMessage) GetChangeType() PlayerRolesUpdateMessage_ChangeType {
	if x != nil {
		return x.ChangeType
	}
	return PlayerRolesUpdateMessage_ADD
}

func (x *PlayerRolesUpdateMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerRolesUpdateMessage) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

var File_permission_messages_proto protoreflect.FileDescriptor

var file_permission_messages_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd4, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x32, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d,
	0x4f, 0x44, 0x49, 0x46, 0x59, 0x10, 0x02, 0x22, 0xd5, 0x01, 0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x21, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x42,
	0x62, 0x0a, 0x22, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_messages_proto_rawDescOnce sync.Once
	file_permission_messages_proto_rawDescData = file_permission_messages_proto_rawDesc
)

func file_permission_messages_proto_rawDescGZIP() []byte {
	file_permission_messages_proto_rawDescOnce.Do(func() {
		file_permission_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_messages_proto_rawDescData)
	})
	return file_permission_messages_proto_rawDescData
}

var file_permission_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_permission_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_permission_messages_proto_goTypes = []interface{}{
	(RoleUpdateMessage_ChangeType)(0),        // 0: emortal.message.permission.RoleUpdateMessage.ChangeType
	(PlayerRolesUpdateMessage_ChangeType)(0), // 1: emortal.message.permission.PlayerRolesUpdateMessage.ChangeType
	(*RoleUpdateMessage)(nil),                // 2: emortal.message.permission.RoleUpdateMessage
	(*PlayerRolesUpdateMessage)(nil),         // 3: emortal.message.permission.PlayerRolesUpdateMessage
	(*permission.Role)(nil),                  // 4: emortal.model.permission.Role
}
var file_permission_messages_proto_depIdxs = []int32{
	0, // 0: emortal.message.permission.RoleUpdateMessage.change_type:type_name -> emortal.message.permission.RoleUpdateMessage.ChangeType
	4, // 1: emortal.message.permission.RoleUpdateMessage.role:type_name -> emortal.model.permission.Role
	1, // 2: emortal.message.permission.PlayerRolesUpdateMessage.change_type:type_name -> emortal.message.permission.PlayerRolesUpdateMessage.ChangeType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_permission_messages_proto_init() }
func file_permission_messages_proto_init() {
	if File_permission_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_permission_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUpdateMessage); i {
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
		file_permission_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerRolesUpdateMessage); i {
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
			RawDescriptor: file_permission_messages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_messages_proto_goTypes,
		DependencyIndexes: file_permission_messages_proto_depIdxs,
		EnumInfos:         file_permission_messages_proto_enumTypes,
		MessageInfos:      file_permission_messages_proto_msgTypes,
	}.Build()
	File_permission_messages_proto = out.File
	file_permission_messages_proto_rawDesc = nil
	file_permission_messages_proto_goTypes = nil
	file_permission_messages_proto_depIdxs = nil
}
