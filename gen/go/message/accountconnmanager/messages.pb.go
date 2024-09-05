// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0
// source: account_conn_manager/messages.proto

package accountconnmanager

import (
	accountconnmanager "github.com/emortalmc/proto-specs/gen/go/model/accountconnmanager"
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

// AccountConnectedMessage is sent when a connection is successful, not when the OAuth URL is created.
// NOTE: multiple connection types my be present (e.g. if it's the first link there might be MINECRAFT and DISCORD)
type AccountConnectedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User            *accountconnmanager.ConnectionUser  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	ConnectionTypes []accountconnmanager.ConnectionType `protobuf:"varint,2,rep,packed,name=connection_types,json=connectionTypes,proto3,enum=emortal.model.account_conn_manager.ConnectionType" json:"connection_types,omitempty"`
}

func (x *AccountConnectedMessage) Reset() {
	*x = AccountConnectedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountConnectedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountConnectedMessage) ProtoMessage() {}

func (x *AccountConnectedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountConnectedMessage.ProtoReflect.Descriptor instead.
func (*AccountConnectedMessage) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_messages_proto_rawDescGZIP(), []int{0}
}

func (x *AccountConnectedMessage) GetUser() *accountconnmanager.ConnectionUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AccountConnectedMessage) GetConnectionTypes() []accountconnmanager.ConnectionType {
	if x != nil {
		return x.ConnectionTypes
	}
	return nil
}

type AccountConnectionRemovedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User              *accountconnmanager.ConnectionUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	RemovedConnection accountconnmanager.ConnectionType  `protobuf:"varint,2,opt,name=removed_connection,json=removedConnection,proto3,enum=emortal.model.account_conn_manager.ConnectionType" json:"removed_connection,omitempty"`
}

func (x *AccountConnectionRemovedMessage) Reset() {
	*x = AccountConnectionRemovedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_conn_manager_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountConnectionRemovedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountConnectionRemovedMessage) ProtoMessage() {}

func (x *AccountConnectionRemovedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_account_conn_manager_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountConnectionRemovedMessage.ProtoReflect.Descriptor instead.
func (*AccountConnectionRemovedMessage) Descriptor() ([]byte, []int) {
	return file_account_conn_manager_messages_proto_rawDescGZIP(), []int{1}
}

func (x *AccountConnectionRemovedMessage) GetUser() *accountconnmanager.ConnectionUser {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *AccountConnectionRemovedMessage) GetRemovedConnection() accountconnmanager.ConnectionType {
	if x != nil {
		return x.RemovedConnection
	}
	return accountconnmanager.ConnectionType(0)
}

var File_account_conn_manager_messages_proto protoreflect.FileDescriptor

var file_account_conn_manager_messages_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x21, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0,
	0x01, 0x0a, 0x17, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x5d, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x22, 0xcc, 0x01, 0x0a, 0x1f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x61, 0x0a,
	0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x11, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x72, 0x0a, 0x2a, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x63, 0x6f, 0x6e, 0x6e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_conn_manager_messages_proto_rawDescOnce sync.Once
	file_account_conn_manager_messages_proto_rawDescData = file_account_conn_manager_messages_proto_rawDesc
)

func file_account_conn_manager_messages_proto_rawDescGZIP() []byte {
	file_account_conn_manager_messages_proto_rawDescOnce.Do(func() {
		file_account_conn_manager_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_conn_manager_messages_proto_rawDescData)
	})
	return file_account_conn_manager_messages_proto_rawDescData
}

var file_account_conn_manager_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_conn_manager_messages_proto_goTypes = []interface{}{
	(*AccountConnectedMessage)(nil),           // 0: emortal.message.account_conn_manager.AccountConnectedMessage
	(*AccountConnectionRemovedMessage)(nil),   // 1: emortal.message.account_conn_manager.AccountConnectionRemovedMessage
	(*accountconnmanager.ConnectionUser)(nil), // 2: emortal.model.account_conn_manager.ConnectionUser
	(accountconnmanager.ConnectionType)(0),    // 3: emortal.model.account_conn_manager.ConnectionType
}
var file_account_conn_manager_messages_proto_depIdxs = []int32{
	2, // 0: emortal.message.account_conn_manager.AccountConnectedMessage.user:type_name -> emortal.model.account_conn_manager.ConnectionUser
	3, // 1: emortal.message.account_conn_manager.AccountConnectedMessage.connection_types:type_name -> emortal.model.account_conn_manager.ConnectionType
	2, // 2: emortal.message.account_conn_manager.AccountConnectionRemovedMessage.user:type_name -> emortal.model.account_conn_manager.ConnectionUser
	3, // 3: emortal.message.account_conn_manager.AccountConnectionRemovedMessage.removed_connection:type_name -> emortal.model.account_conn_manager.ConnectionType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_account_conn_manager_messages_proto_init() }
func file_account_conn_manager_messages_proto_init() {
	if File_account_conn_manager_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_conn_manager_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountConnectedMessage); i {
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
		file_account_conn_manager_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountConnectionRemovedMessage); i {
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
			RawDescriptor: file_account_conn_manager_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_conn_manager_messages_proto_goTypes,
		DependencyIndexes: file_account_conn_manager_messages_proto_depIdxs,
		MessageInfos:      file_account_conn_manager_messages_proto_msgTypes,
	}.Build()
	File_account_conn_manager_messages_proto = out.File
	file_account_conn_manager_messages_proto_rawDesc = nil
	file_account_conn_manager_messages_proto_goTypes = nil
	file_account_conn_manager_messages_proto_depIdxs = nil
}
