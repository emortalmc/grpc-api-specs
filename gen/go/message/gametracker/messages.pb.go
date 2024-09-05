// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0
// source: game_tracker/messages.proto

package gametracker

import (
	gametracker "github.com/emortalmc/proto-specs/gen/go/model/gametracker"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameStartMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonData *CommonGameData        `protobuf:"bytes,1,opt,name=common_data,json=commonData,proto3" json:"common_data,omitempty"`
	StartTime  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	MapId      *string                `protobuf:"bytes,3,opt,name=map_id,json=mapId,proto3,oneof" json:"map_id,omitempty"`
	Content    []*anypb.Any           `protobuf:"bytes,1000,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *GameStartMessage) Reset() {
	*x = GameStartMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStartMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStartMessage) ProtoMessage() {}

func (x *GameStartMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStartMessage.ProtoReflect.Descriptor instead.
func (*GameStartMessage) Descriptor() ([]byte, []int) {
	return file_game_tracker_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GameStartMessage) GetCommonData() *CommonGameData {
	if x != nil {
		return x.CommonData
	}
	return nil
}

func (x *GameStartMessage) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GameStartMessage) GetMapId() string {
	if x != nil && x.MapId != nil {
		return *x.MapId
	}
	return ""
}

func (x *GameStartMessage) GetContent() []*anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

type GameUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonData *CommonGameData `protobuf:"bytes,1,opt,name=common_data,json=commonData,proto3" json:"common_data,omitempty"`
	Content    []*anypb.Any    `protobuf:"bytes,1000,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *GameUpdateMessage) Reset() {
	*x = GameUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameUpdateMessage) ProtoMessage() {}

func (x *GameUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameUpdateMessage.ProtoReflect.Descriptor instead.
func (*GameUpdateMessage) Descriptor() ([]byte, []int) {
	return file_game_tracker_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GameUpdateMessage) GetCommonData() *CommonGameData {
	if x != nil {
		return x.CommonData
	}
	return nil
}

func (x *GameUpdateMessage) GetContent() []*anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

type GameFinishMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonData *CommonGameData        `protobuf:"bytes,1,opt,name=common_data,json=commonData,proto3" json:"common_data,omitempty"`
	EndTime    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Content    []*anypb.Any           `protobuf:"bytes,1000,rep,name=content,proto3" json:"content,omitempty"`
}

func (x *GameFinishMessage) Reset() {
	*x = GameFinishMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameFinishMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameFinishMessage) ProtoMessage() {}

func (x *GameFinishMessage) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameFinishMessage.ProtoReflect.Descriptor instead.
func (*GameFinishMessage) Descriptor() ([]byte, []int) {
	return file_game_tracker_messages_proto_rawDescGZIP(), []int{2}
}

func (x *GameFinishMessage) GetCommonData() *CommonGameData {
	if x != nil {
		return x.CommonData
	}
	return nil
}

func (x *GameFinishMessage) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *GameFinishMessage) GetContent() []*anypb.Any {
	if x != nil {
		return x.Content
	}
	return nil
}

type CommonGameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameModeId string                         `protobuf:"bytes,1,opt,name=game_mode_id,json=gameModeId,proto3" json:"game_mode_id,omitempty"`
	GameId     string                         `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	ServerId   string                         `protobuf:"bytes,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Players    []*gametracker.BasicGamePlayer `protobuf:"bytes,6,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *CommonGameData) Reset() {
	*x = CommonGameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_tracker_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonGameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonGameData) ProtoMessage() {}

func (x *CommonGameData) ProtoReflect() protoreflect.Message {
	mi := &file_game_tracker_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonGameData.ProtoReflect.Descriptor instead.
func (*CommonGameData) Descriptor() ([]byte, []int) {
	return file_game_tracker_messages_proto_rawDescGZIP(), []int{3}
}

func (x *CommonGameData) GetGameModeId() string {
	if x != nil {
		return x.GameModeId
	}
	return ""
}

func (x *CommonGameData) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *CommonGameData) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *CommonGameData) GetPlayers() []*gametracker.BasicGamePlayer {
	if x != nil {
		return x.Players
	}
	return nil
}

var File_game_tracker_messages_proto protoreflect.FileDescriptor

var file_game_tracker_messages_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x1a, 0x19, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x11, 0x47, 0x61,
	0x6d, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0xca, 0x01, 0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x4d, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xaf, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x47, 0x61, 0x6d, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x64,
	0x0a, 0x23, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_tracker_messages_proto_rawDescOnce sync.Once
	file_game_tracker_messages_proto_rawDescData = file_game_tracker_messages_proto_rawDesc
)

func file_game_tracker_messages_proto_rawDescGZIP() []byte {
	file_game_tracker_messages_proto_rawDescOnce.Do(func() {
		file_game_tracker_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_tracker_messages_proto_rawDescData)
	})
	return file_game_tracker_messages_proto_rawDescData
}

var file_game_tracker_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_tracker_messages_proto_goTypes = []interface{}{
	(*GameStartMessage)(nil),            // 0: emortal.message.game_tracker.GameStartMessage
	(*GameUpdateMessage)(nil),           // 1: emortal.message.game_tracker.GameUpdateMessage
	(*GameFinishMessage)(nil),           // 2: emortal.message.game_tracker.GameFinishMessage
	(*CommonGameData)(nil),              // 3: emortal.message.game_tracker.CommonGameData
	(*timestamppb.Timestamp)(nil),       // 4: google.protobuf.Timestamp
	(*anypb.Any)(nil),                   // 5: google.protobuf.Any
	(*gametracker.BasicGamePlayer)(nil), // 6: emortal.model.game_tracker.BasicGamePlayer
}
var file_game_tracker_messages_proto_depIdxs = []int32{
	3, // 0: emortal.message.game_tracker.GameStartMessage.common_data:type_name -> emortal.message.game_tracker.CommonGameData
	4, // 1: emortal.message.game_tracker.GameStartMessage.start_time:type_name -> google.protobuf.Timestamp
	5, // 2: emortal.message.game_tracker.GameStartMessage.content:type_name -> google.protobuf.Any
	3, // 3: emortal.message.game_tracker.GameUpdateMessage.common_data:type_name -> emortal.message.game_tracker.CommonGameData
	5, // 4: emortal.message.game_tracker.GameUpdateMessage.content:type_name -> google.protobuf.Any
	3, // 5: emortal.message.game_tracker.GameFinishMessage.common_data:type_name -> emortal.message.game_tracker.CommonGameData
	4, // 6: emortal.message.game_tracker.GameFinishMessage.end_time:type_name -> google.protobuf.Timestamp
	5, // 7: emortal.message.game_tracker.GameFinishMessage.content:type_name -> google.protobuf.Any
	6, // 8: emortal.message.game_tracker.CommonGameData.players:type_name -> emortal.model.game_tracker.BasicGamePlayer
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_game_tracker_messages_proto_init() }
func file_game_tracker_messages_proto_init() {
	if File_game_tracker_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_tracker_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStartMessage); i {
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
		file_game_tracker_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameUpdateMessage); i {
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
		file_game_tracker_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameFinishMessage); i {
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
		file_game_tracker_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonGameData); i {
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
	file_game_tracker_messages_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_tracker_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_tracker_messages_proto_goTypes,
		DependencyIndexes: file_game_tracker_messages_proto_depIdxs,
		MessageInfos:      file_game_tracker_messages_proto_msgTypes,
	}.Build()
	File_game_tracker_messages_proto = out.File
	file_game_tracker_messages_proto_rawDesc = nil
	file_game_tracker_messages_proto_goTypes = nil
	file_game_tracker_messages_proto_depIdxs = nil
}
