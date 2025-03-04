// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: relationship/messages.proto

package relationship

import (
	relationship "github.com/emortalmc/proto-specs/gen/go/model/relationship"
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

type FriendRequestReceivedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *relationship.FriendRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *FriendRequestReceivedMessage) Reset() {
	*x = FriendRequestReceivedMessage{}
	mi := &file_relationship_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendRequestReceivedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestReceivedMessage) ProtoMessage() {}

func (x *FriendRequestReceivedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_relationship_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestReceivedMessage.ProtoReflect.Descriptor instead.
func (*FriendRequestReceivedMessage) Descriptor() ([]byte, []int) {
	return file_relationship_messages_proto_rawDescGZIP(), []int{0}
}

func (x *FriendRequestReceivedMessage) GetRequest() *relationship.FriendRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

// Called when a friend request is accepted
type FriendAddedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request that was accepted to add the friend
	SenderId       string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	SenderUsername string `protobuf:"bytes,2,opt,name=sender_username,json=senderUsername,proto3" json:"sender_username,omitempty"`
	RecipientId    string `protobuf:"bytes,3,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
}

func (x *FriendAddedMessage) Reset() {
	*x = FriendAddedMessage{}
	mi := &file_relationship_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendAddedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendAddedMessage) ProtoMessage() {}

func (x *FriendAddedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_relationship_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendAddedMessage.ProtoReflect.Descriptor instead.
func (*FriendAddedMessage) Descriptor() ([]byte, []int) {
	return file_relationship_messages_proto_rawDescGZIP(), []int{1}
}

func (x *FriendAddedMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *FriendAddedMessage) GetSenderUsername() string {
	if x != nil {
		return x.SenderUsername
	}
	return ""
}

func (x *FriendAddedMessage) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type FriendRemovedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderId    string `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	RecipientId string `protobuf:"bytes,2,opt,name=recipient_id,json=recipientId,proto3" json:"recipient_id,omitempty"`
}

func (x *FriendRemovedMessage) Reset() {
	*x = FriendRemovedMessage{}
	mi := &file_relationship_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendRemovedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRemovedMessage) ProtoMessage() {}

func (x *FriendRemovedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_relationship_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRemovedMessage.ProtoReflect.Descriptor instead.
func (*FriendRemovedMessage) Descriptor() ([]byte, []int) {
	return file_relationship_messages_proto_rawDescGZIP(), []int{2}
}

func (x *FriendRemovedMessage) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *FriendRemovedMessage) GetRecipientId() string {
	if x != nil {
		return x.RecipientId
	}
	return ""
}

type FriendConnectionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageTargetIds []string `protobuf:"bytes,1,rep,name=message_target_ids,json=messageTargetIds,proto3" json:"message_target_ids,omitempty"`
	PlayerId         string   `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Username         string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// joined if true, player joined. If false, player disconnected.
	Joined bool `protobuf:"varint,4,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *FriendConnectionMessage) Reset() {
	*x = FriendConnectionMessage{}
	mi := &file_relationship_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FriendConnectionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendConnectionMessage) ProtoMessage() {}

func (x *FriendConnectionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_relationship_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendConnectionMessage.ProtoReflect.Descriptor instead.
func (*FriendConnectionMessage) Descriptor() ([]byte, []int) {
	return file_relationship_messages_proto_rawDescGZIP(), []int{3}
}

func (x *FriendConnectionMessage) GetMessageTargetIds() []string {
	if x != nil {
		return x.MessageTargetIds
	}
	return nil
}

func (x *FriendConnectionMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *FriendConnectionMessage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FriendConnectionMessage) GetJoined() bool {
	if x != nil {
		return x.Joined
	}
	return false
}

var File_relationship_messages_proto protoreflect.FileDescriptor

var file_relationship_messages_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x65,
	0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x1a, 0x19, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x1c, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7d, 0x0a, 0x12, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x14, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x98, 0x01, 0x0a, 0x17, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x66, 0x0a,
	0x24, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x68, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relationship_messages_proto_rawDescOnce sync.Once
	file_relationship_messages_proto_rawDescData = file_relationship_messages_proto_rawDesc
)

func file_relationship_messages_proto_rawDescGZIP() []byte {
	file_relationship_messages_proto_rawDescOnce.Do(func() {
		file_relationship_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_relationship_messages_proto_rawDescData)
	})
	return file_relationship_messages_proto_rawDescData
}

var file_relationship_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_relationship_messages_proto_goTypes = []any{
	(*FriendRequestReceivedMessage)(nil), // 0: emortal.message.relationship.FriendRequestReceivedMessage
	(*FriendAddedMessage)(nil),           // 1: emortal.message.relationship.FriendAddedMessage
	(*FriendRemovedMessage)(nil),         // 2: emortal.message.relationship.FriendRemovedMessage
	(*FriendConnectionMessage)(nil),      // 3: emortal.message.relationship.FriendConnectionMessage
	(*relationship.FriendRequest)(nil),   // 4: emortal.model.relationship.FriendRequest
}
var file_relationship_messages_proto_depIdxs = []int32{
	4, // 0: emortal.message.relationship.FriendRequestReceivedMessage.request:type_name -> emortal.model.relationship.FriendRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_relationship_messages_proto_init() }
func file_relationship_messages_proto_init() {
	if File_relationship_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relationship_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relationship_messages_proto_goTypes,
		DependencyIndexes: file_relationship_messages_proto_depIdxs,
		MessageInfos:      file_relationship_messages_proto_msgTypes,
	}.Build()
	File_relationship_messages_proto = out.File
	file_relationship_messages_proto_rawDesc = nil
	file_relationship_messages_proto_goTypes = nil
	file_relationship_messages_proto_depIdxs = nil
}
