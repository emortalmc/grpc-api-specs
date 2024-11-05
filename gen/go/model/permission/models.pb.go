// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: permission/models.proto

package permission

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

type PermissionNode_PermissionState int32

const (
	PermissionNode_ALLOW PermissionNode_PermissionState = 0
	PermissionNode_DENY  PermissionNode_PermissionState = 1
)

// Enum value maps for PermissionNode_PermissionState.
var (
	PermissionNode_PermissionState_name = map[int32]string{
		0: "ALLOW",
		1: "DENY",
	}
	PermissionNode_PermissionState_value = map[string]int32{
		"ALLOW": 0,
		"DENY":  1,
	}
)

func (x PermissionNode_PermissionState) Enum() *PermissionNode_PermissionState {
	p := new(PermissionNode_PermissionState)
	*p = x
	return p
}

func (x PermissionNode_PermissionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionNode_PermissionState) Descriptor() protoreflect.EnumDescriptor {
	return file_permission_models_proto_enumTypes[0].Descriptor()
}

func (PermissionNode_PermissionState) Type() protoreflect.EnumType {
	return &file_permission_models_proto_enumTypes[0]
}

func (x PermissionNode_PermissionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionNode_PermissionState.Descriptor instead.
func (PermissionNode_PermissionState) EnumDescriptor() ([]byte, []int) {
	return file_permission_models_proto_rawDescGZIP(), []int{1, 0}
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Priority    uint32            `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	DisplayName *string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	Permissions []*PermissionNode `protobuf:"bytes,5,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_permission_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_permission_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_permission_models_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Role) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Role) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *Role) GetPermissions() []*PermissionNode {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type PermissionNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node  string                         `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	State PermissionNode_PermissionState `protobuf:"varint,2,opt,name=state,proto3,enum=emortal.model.permission.PermissionNode_PermissionState" json:"state,omitempty"`
}

func (x *PermissionNode) Reset() {
	*x = PermissionNode{}
	mi := &file_permission_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionNode) ProtoMessage() {}

func (x *PermissionNode) ProtoReflect() protoreflect.Message {
	mi := &file_permission_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionNode.ProtoReflect.Descriptor instead.
func (*PermissionNode) Descriptor() ([]byte, []int) {
	return file_permission_models_proto_rawDescGZIP(), []int{1}
}

func (x *PermissionNode) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *PermissionNode) GetState() PermissionNode_PermissionState {
	if x != nil {
		return x.State
	}
	return PermissionNode_ALLOW
}

var File_permission_models_proto protoreflect.FileDescriptor

var file_permission_models_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xbd, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x26, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x4a, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x4a, 0x04, 0x08,
	0x03, 0x10, 0x04, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e,
	0x6f, 0x64, 0x65, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x0f, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59,
	0x10, 0x01, 0x42, 0x5e, 0x0a, 0x20, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_permission_models_proto_rawDescOnce sync.Once
	file_permission_models_proto_rawDescData = file_permission_models_proto_rawDesc
)

func file_permission_models_proto_rawDescGZIP() []byte {
	file_permission_models_proto_rawDescOnce.Do(func() {
		file_permission_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_permission_models_proto_rawDescData)
	})
	return file_permission_models_proto_rawDescData
}

var file_permission_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_permission_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_permission_models_proto_goTypes = []any{
	(PermissionNode_PermissionState)(0), // 0: emortal.model.permission.PermissionNode.PermissionState
	(*Role)(nil),                        // 1: emortal.model.permission.Role
	(*PermissionNode)(nil),              // 2: emortal.model.permission.PermissionNode
}
var file_permission_models_proto_depIdxs = []int32{
	2, // 0: emortal.model.permission.Role.permissions:type_name -> emortal.model.permission.PermissionNode
	0, // 1: emortal.model.permission.PermissionNode.state:type_name -> emortal.model.permission.PermissionNode.PermissionState
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_permission_models_proto_init() }
func file_permission_models_proto_init() {
	if File_permission_models_proto != nil {
		return
	}
	file_permission_models_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_permission_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_permission_models_proto_goTypes,
		DependencyIndexes: file_permission_models_proto_depIdxs,
		EnumInfos:         file_permission_models_proto_enumTypes,
		MessageInfos:      file_permission_models_proto_msgTypes,
	}.Build()
	File_permission_models_proto = out.File
	file_permission_models_proto_rawDesc = nil
	file_permission_models_proto_goTypes = nil
	file_permission_models_proto_depIdxs = nil
}
