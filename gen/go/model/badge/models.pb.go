// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: badges/models.proto

package badge

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

type Badge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Priority     int64  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Required     bool   `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	FriendlyName string `protobuf:"bytes,4,opt,name=friendly_name,json=friendlyName,proto3" json:"friendly_name,omitempty"`
	// chat_string normally one character but might be more.
	// should be parsed with MiniMessage as it might contain a colour code.
	ChatString string         `protobuf:"bytes,5,opt,name=chat_string,json=chatString,proto3" json:"chat_string,omitempty"`
	HoverText  string         `protobuf:"bytes,6,opt,name=hover_text,json=hoverText,proto3" json:"hover_text,omitempty"`
	GuiItem    *Badge_GuiItem `protobuf:"bytes,7,opt,name=gui_item,json=guiItem,proto3" json:"gui_item,omitempty"`
}

func (x *Badge) Reset() {
	*x = Badge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badges_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Badge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Badge) ProtoMessage() {}

func (x *Badge) ProtoReflect() protoreflect.Message {
	mi := &file_badges_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Badge.ProtoReflect.Descriptor instead.
func (*Badge) Descriptor() ([]byte, []int) {
	return file_badges_models_proto_rawDescGZIP(), []int{0}
}

func (x *Badge) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Badge) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Badge) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *Badge) GetFriendlyName() string {
	if x != nil {
		return x.FriendlyName
	}
	return ""
}

func (x *Badge) GetChatString() string {
	if x != nil {
		return x.ChatString
	}
	return ""
}

func (x *Badge) GetHoverText() string {
	if x != nil {
		return x.HoverText
	}
	return ""
}

func (x *Badge) GetGuiItem() *Badge_GuiItem {
	if x != nil {
		return x.GuiItem
	}
	return nil
}

type Badge_GuiItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Material    string   `protobuf:"bytes,1,opt,name=material,proto3" json:"material,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Lore        []string `protobuf:"bytes,3,rep,name=lore,proto3" json:"lore,omitempty"`
}

func (x *Badge_GuiItem) Reset() {
	*x = Badge_GuiItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_badges_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Badge_GuiItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Badge_GuiItem) ProtoMessage() {}

func (x *Badge_GuiItem) ProtoReflect() protoreflect.Message {
	mi := &file_badges_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Badge_GuiItem.ProtoReflect.Descriptor instead.
func (*Badge_GuiItem) Descriptor() ([]byte, []int) {
	return file_badges_models_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Badge_GuiItem) GetMaterial() string {
	if x != nil {
		return x.Material
	}
	return ""
}

func (x *Badge_GuiItem) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Badge_GuiItem) GetLore() []string {
	if x != nil {
		return x.Lore
	}
	return nil
}

var File_badges_models_proto protoreflect.FileDescriptor

var file_badges_models_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x05, 0x42,
	0x61, 0x64, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x6c, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x76, 0x65, 0x72, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x3d, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x2e,
	0x47, 0x75, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x07, 0x67, 0x75, 0x69, 0x49, 0x74, 0x65, 0x6d,
	0x1a, 0x5c, 0x0a, 0x07, 0x47, 0x75, 0x69, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x61, 0x74, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x72, 0x65, 0x42, 0x54,
	0x0a, 0x1b, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x64, 0x67, 0x65, 0x50, 0x01, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_badges_models_proto_rawDescOnce sync.Once
	file_badges_models_proto_rawDescData = file_badges_models_proto_rawDesc
)

func file_badges_models_proto_rawDescGZIP() []byte {
	file_badges_models_proto_rawDescOnce.Do(func() {
		file_badges_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_badges_models_proto_rawDescData)
	})
	return file_badges_models_proto_rawDescData
}

var file_badges_models_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_badges_models_proto_goTypes = []interface{}{
	(*Badge)(nil),         // 0: emortal.model.badge.Badge
	(*Badge_GuiItem)(nil), // 1: emortal.model.badge.Badge.GuiItem
}
var file_badges_models_proto_depIdxs = []int32{
	1, // 0: emortal.model.badge.Badge.gui_item:type_name -> emortal.model.badge.Badge.GuiItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_badges_models_proto_init() }
func file_badges_models_proto_init() {
	if File_badges_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_badges_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Badge); i {
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
		file_badges_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Badge_GuiItem); i {
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
			RawDescriptor: file_badges_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_badges_models_proto_goTypes,
		DependencyIndexes: file_badges_models_proto_depIdxs,
		MessageInfos:      file_badges_models_proto_msgTypes,
	}.Build()
	File_badges_models_proto = out.File
	file_badges_models_proto_rawDesc = nil
	file_badges_models_proto_goTypes = nil
	file_badges_models_proto_depIdxs = nil
}
