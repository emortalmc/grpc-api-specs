// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: badges/messages.proto

package badge

import (
	badge "github.com/emortalmc/proto-specs/gen/go/model/badge"
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

type PlayerBadgeAddedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string       `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Badge    *badge.Badge `protobuf:"bytes,2,opt,name=badge,proto3" json:"badge,omitempty"`
}

func (x *PlayerBadgeAddedMessage) Reset() {
	*x = PlayerBadgeAddedMessage{}
	mi := &file_badges_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerBadgeAddedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerBadgeAddedMessage) ProtoMessage() {}

func (x *PlayerBadgeAddedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_badges_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerBadgeAddedMessage.ProtoReflect.Descriptor instead.
func (*PlayerBadgeAddedMessage) Descriptor() ([]byte, []int) {
	return file_badges_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerBadgeAddedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerBadgeAddedMessage) GetBadge() *badge.Badge {
	if x != nil {
		return x.Badge
	}
	return nil
}

type PlayerBadgeRemovedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string       `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Badge    *badge.Badge `protobuf:"bytes,2,opt,name=badge,proto3" json:"badge,omitempty"`
}

func (x *PlayerBadgeRemovedMessage) Reset() {
	*x = PlayerBadgeRemovedMessage{}
	mi := &file_badges_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerBadgeRemovedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerBadgeRemovedMessage) ProtoMessage() {}

func (x *PlayerBadgeRemovedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_badges_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerBadgeRemovedMessage.ProtoReflect.Descriptor instead.
func (*PlayerBadgeRemovedMessage) Descriptor() ([]byte, []int) {
	return file_badges_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerBadgeRemovedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerBadgeRemovedMessage) GetBadge() *badge.Badge {
	if x != nil {
		return x.Badge
	}
	return nil
}

type PlayerActiveBadgeChangedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// player_id of type UUID
	PlayerId string       `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Badge    *badge.Badge `protobuf:"bytes,2,opt,name=badge,proto3" json:"badge,omitempty"`
}

func (x *PlayerActiveBadgeChangedMessage) Reset() {
	*x = PlayerActiveBadgeChangedMessage{}
	mi := &file_badges_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerActiveBadgeChangedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerActiveBadgeChangedMessage) ProtoMessage() {}

func (x *PlayerActiveBadgeChangedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_badges_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerActiveBadgeChangedMessage.ProtoReflect.Descriptor instead.
func (*PlayerActiveBadgeChangedMessage) Descriptor() ([]byte, []int) {
	return file_badges_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PlayerActiveBadgeChangedMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerActiveBadgeChangedMessage) GetBadge() *badge.Badge {
	if x != nil {
		return x.Badge
	}
	return nil
}

var File_badges_messages_proto protoreflect.FileDescriptor

var file_badges_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x1a, 0x13,
	0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x41, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x22, 0x6a, 0x0a,
	0x19, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x61, 0x64, 0x67, 0x65, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x64,
	0x67, 0x65, 0x52, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x22, 0x70, 0x0a, 0x1f, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x42, 0x61, 0x64, 0x67, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x61, 0x64,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x52, 0x05, 0x62, 0x61, 0x64, 0x67, 0x65, 0x42, 0x58, 0x0a, 0x1d, 0x64,
	0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x62, 0x61, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_badges_messages_proto_rawDescOnce sync.Once
	file_badges_messages_proto_rawDescData = file_badges_messages_proto_rawDesc
)

func file_badges_messages_proto_rawDescGZIP() []byte {
	file_badges_messages_proto_rawDescOnce.Do(func() {
		file_badges_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_badges_messages_proto_rawDescData)
	})
	return file_badges_messages_proto_rawDescData
}

var file_badges_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_badges_messages_proto_goTypes = []any{
	(*PlayerBadgeAddedMessage)(nil),         // 0: emortal.message.badge.PlayerBadgeAddedMessage
	(*PlayerBadgeRemovedMessage)(nil),       // 1: emortal.message.badge.PlayerBadgeRemovedMessage
	(*PlayerActiveBadgeChangedMessage)(nil), // 2: emortal.message.badge.PlayerActiveBadgeChangedMessage
	(*badge.Badge)(nil),                     // 3: emortal.model.badge.Badge
}
var file_badges_messages_proto_depIdxs = []int32{
	3, // 0: emortal.message.badge.PlayerBadgeAddedMessage.badge:type_name -> emortal.model.badge.Badge
	3, // 1: emortal.message.badge.PlayerBadgeRemovedMessage.badge:type_name -> emortal.model.badge.Badge
	3, // 2: emortal.message.badge.PlayerActiveBadgeChangedMessage.badge:type_name -> emortal.model.badge.Badge
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_badges_messages_proto_init() }
func file_badges_messages_proto_init() {
	if File_badges_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_badges_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_badges_messages_proto_goTypes,
		DependencyIndexes: file_badges_messages_proto_depIdxs,
		MessageInfos:      file_badges_messages_proto_msgTypes,
	}.Build()
	File_badges_messages_proto = out.File
	file_badges_messages_proto_rawDesc = nil
	file_badges_messages_proto_goTypes = nil
	file_badges_messages_proto_depIdxs = nil
}
