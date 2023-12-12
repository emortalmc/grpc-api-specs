// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: message_handler/grpc.proto

package messagehandler

import (
	messagehandler "github.com/emortalmc/proto-specs/gen/go/model/messagehandler"
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

type PrivateMessageErrorResponse_Reason int32

const (
	PrivateMessageErrorResponse_PLAYER_NOT_ONLINE PrivateMessageErrorResponse_Reason = 0
	PrivateMessageErrorResponse_PRIVACY_BLOCKED   PrivateMessageErrorResponse_Reason = 1
	PrivateMessageErrorResponse_YOU_BLOCKED       PrivateMessageErrorResponse_Reason = 2
)

// Enum value maps for PrivateMessageErrorResponse_Reason.
var (
	PrivateMessageErrorResponse_Reason_name = map[int32]string{
		0: "PLAYER_NOT_ONLINE",
		1: "PRIVACY_BLOCKED",
		2: "YOU_BLOCKED",
	}
	PrivateMessageErrorResponse_Reason_value = map[string]int32{
		"PLAYER_NOT_ONLINE": 0,
		"PRIVACY_BLOCKED":   1,
		"YOU_BLOCKED":       2,
	}
)

func (x PrivateMessageErrorResponse_Reason) Enum() *PrivateMessageErrorResponse_Reason {
	p := new(PrivateMessageErrorResponse_Reason)
	*p = x
	return p
}

func (x PrivateMessageErrorResponse_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PrivateMessageErrorResponse_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_message_handler_grpc_proto_enumTypes[0].Descriptor()
}

func (PrivateMessageErrorResponse_Reason) Type() protoreflect.EnumType {
	return &file_message_handler_grpc_proto_enumTypes[0]
}

func (x PrivateMessageErrorResponse_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PrivateMessageErrorResponse_Reason.Descriptor instead.
func (PrivateMessageErrorResponse_Reason) EnumDescriptor() ([]byte, []int) {
	return file_message_handler_grpc_proto_rawDescGZIP(), []int{2, 0}
}

type PrivateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *messagehandler.PrivateMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PrivateMessageRequest) Reset() {
	*x = PrivateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_handler_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateMessageRequest) ProtoMessage() {}

func (x *PrivateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_handler_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateMessageRequest.ProtoReflect.Descriptor instead.
func (*PrivateMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_handler_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateMessageRequest) GetMessage() *messagehandler.PrivateMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type PrivateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The message may have been modified for some safety reasons
	Message *messagehandler.PrivateMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PrivateMessageResponse) Reset() {
	*x = PrivateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_handler_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateMessageResponse) ProtoMessage() {}

func (x *PrivateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_handler_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateMessageResponse.ProtoReflect.Descriptor instead.
func (*PrivateMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_handler_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *PrivateMessageResponse) GetMessage() *messagehandler.PrivateMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type PrivateMessageErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason PrivateMessageErrorResponse_Reason `protobuf:"varint,1,opt,name=reason,proto3,enum=emortal.grpc.messagehandler.PrivateMessageErrorResponse_Reason" json:"reason,omitempty"`
}

func (x *PrivateMessageErrorResponse) Reset() {
	*x = PrivateMessageErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_handler_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateMessageErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateMessageErrorResponse) ProtoMessage() {}

func (x *PrivateMessageErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_handler_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateMessageErrorResponse.ProtoReflect.Descriptor instead.
func (*PrivateMessageErrorResponse) Descriptor() ([]byte, []int) {
	return file_message_handler_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateMessageErrorResponse) GetReason() PrivateMessageErrorResponse_Reason {
	if x != nil {
		return x.Reason
	}
	return PrivateMessageErrorResponse_PLAYER_NOT_ONLINE
}

var File_message_handler_grpc_proto protoreflect.FileDescriptor

var file_message_handler_grpc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x65, 0x6d,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x1a, 0x1c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x15, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x46, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x16, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbd, 0x01, 0x0a, 0x1b, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x15, 0x0a,
	0x11, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4f, 0x4e, 0x4c, 0x49,
	0x4e, 0x45, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x43, 0x59, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x59, 0x4f, 0x55,
	0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x32, 0x8f, 0x01, 0x0a, 0x0e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x7d, 0x0a,
	0x12, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x32, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x77, 0x0a, 0x23,
	0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x42, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_handler_grpc_proto_rawDescOnce sync.Once
	file_message_handler_grpc_proto_rawDescData = file_message_handler_grpc_proto_rawDesc
)

func file_message_handler_grpc_proto_rawDescGZIP() []byte {
	file_message_handler_grpc_proto_rawDescOnce.Do(func() {
		file_message_handler_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_handler_grpc_proto_rawDescData)
	})
	return file_message_handler_grpc_proto_rawDescData
}

var file_message_handler_grpc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_handler_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_handler_grpc_proto_goTypes = []interface{}{
	(PrivateMessageErrorResponse_Reason)(0), // 0: emortal.grpc.messagehandler.PrivateMessageErrorResponse.Reason
	(*PrivateMessageRequest)(nil),           // 1: emortal.grpc.messagehandler.PrivateMessageRequest
	(*PrivateMessageResponse)(nil),          // 2: emortal.grpc.messagehandler.PrivateMessageResponse
	(*PrivateMessageErrorResponse)(nil),     // 3: emortal.grpc.messagehandler.PrivateMessageErrorResponse
	(*messagehandler.PrivateMessage)(nil),   // 4: emortal.model.messagehandler.PrivateMessage
}
var file_message_handler_grpc_proto_depIdxs = []int32{
	4, // 0: emortal.grpc.messagehandler.PrivateMessageRequest.message:type_name -> emortal.model.messagehandler.PrivateMessage
	4, // 1: emortal.grpc.messagehandler.PrivateMessageResponse.message:type_name -> emortal.model.messagehandler.PrivateMessage
	0, // 2: emortal.grpc.messagehandler.PrivateMessageErrorResponse.reason:type_name -> emortal.grpc.messagehandler.PrivateMessageErrorResponse.Reason
	1, // 3: emortal.grpc.messagehandler.MessageHandler.SendPrivateMessage:input_type -> emortal.grpc.messagehandler.PrivateMessageRequest
	2, // 4: emortal.grpc.messagehandler.MessageHandler.SendPrivateMessage:output_type -> emortal.grpc.messagehandler.PrivateMessageResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_handler_grpc_proto_init() }
func file_message_handler_grpc_proto_init() {
	if File_message_handler_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_handler_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateMessageRequest); i {
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
		file_message_handler_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateMessageResponse); i {
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
		file_message_handler_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrivateMessageErrorResponse); i {
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
			RawDescriptor: file_message_handler_grpc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_handler_grpc_proto_goTypes,
		DependencyIndexes: file_message_handler_grpc_proto_depIdxs,
		EnumInfos:         file_message_handler_grpc_proto_enumTypes,
		MessageInfos:      file_message_handler_grpc_proto_msgTypes,
	}.Build()
	File_message_handler_grpc_proto = out.File
	file_message_handler_grpc_proto_rawDesc = nil
	file_message_handler_grpc_proto_goTypes = nil
	file_message_handler_grpc_proto_depIdxs = nil
}
