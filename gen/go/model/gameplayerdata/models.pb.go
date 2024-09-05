// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0
// source: game_player_data/models.proto

package gameplayerdata

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

type GameDataGameMode int32

const (
	GameDataGameMode_TOWER_DEFENCE GameDataGameMode = 0
	GameDataGameMode_BLOCK_SUMO    GameDataGameMode = 1
	GameDataGameMode_MINESWEEPER   GameDataGameMode = 2
	GameDataGameMode_MARATHON      GameDataGameMode = 3
)

// Enum value maps for GameDataGameMode.
var (
	GameDataGameMode_name = map[int32]string{
		0: "TOWER_DEFENCE",
		1: "BLOCK_SUMO",
		2: "MINESWEEPER",
		3: "MARATHON",
	}
	GameDataGameMode_value = map[string]int32{
		"TOWER_DEFENCE": 0,
		"BLOCK_SUMO":    1,
		"MINESWEEPER":   2,
		"MARATHON":      3,
	}
)

func (x GameDataGameMode) Enum() *GameDataGameMode {
	p := new(GameDataGameMode)
	*p = x
	return p
}

func (x GameDataGameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameDataGameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_player_data_models_proto_enumTypes[0].Descriptor()
}

func (GameDataGameMode) Type() protoreflect.EnumType {
	return &file_game_player_data_models_proto_enumTypes[0]
}

func (x GameDataGameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameDataGameMode.Descriptor instead.
func (GameDataGameMode) EnumDescriptor() ([]byte, []int) {
	return file_game_player_data_models_proto_rawDescGZIP(), []int{0}
}

type V1BlockSumoPlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockSlot  uint32 `protobuf:"varint,1,opt,name=block_slot,json=blockSlot,proto3" json:"block_slot,omitempty"`
	ShearsSlot uint32 `protobuf:"varint,2,opt,name=shears_slot,json=shearsSlot,proto3" json:"shears_slot,omitempty"`
}

func (x *V1BlockSumoPlayerData) Reset() {
	*x = V1BlockSumoPlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_data_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1BlockSumoPlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1BlockSumoPlayerData) ProtoMessage() {}

func (x *V1BlockSumoPlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_data_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1BlockSumoPlayerData.ProtoReflect.Descriptor instead.
func (*V1BlockSumoPlayerData) Descriptor() ([]byte, []int) {
	return file_game_player_data_models_proto_rawDescGZIP(), []int{0}
}

func (x *V1BlockSumoPlayerData) GetBlockSlot() uint32 {
	if x != nil {
		return x.BlockSlot
	}
	return 0
}

func (x *V1BlockSumoPlayerData) GetShearsSlot() uint32 {
	if x != nil {
		return x.ShearsSlot
	}
	return 0
}

type V1MarathonData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *V1MarathonData) Reset() {
	*x = V1MarathonData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_data_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1MarathonData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1MarathonData) ProtoMessage() {}

func (x *V1MarathonData) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_data_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1MarathonData.ProtoReflect.Descriptor instead.
func (*V1MarathonData) Descriptor() ([]byte, []int) {
	return file_game_player_data_models_proto_rawDescGZIP(), []int{1}
}

// not currently used
type V1TowerDefencePlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *V1TowerDefencePlayerData) Reset() {
	*x = V1TowerDefencePlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_data_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1TowerDefencePlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1TowerDefencePlayerData) ProtoMessage() {}

func (x *V1TowerDefencePlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_data_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1TowerDefencePlayerData.ProtoReflect.Descriptor instead.
func (*V1TowerDefencePlayerData) Descriptor() ([]byte, []int) {
	return file_game_player_data_models_proto_rawDescGZIP(), []int{2}
}

// not currently used
type V1MinesweeperPlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length   uint32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Width    uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Mines    uint32 `protobuf:"varint,3,opt,name=mines,proto3" json:"mines,omitempty"`
	Theme    string `protobuf:"bytes,4,opt,name=theme,proto3" json:"theme,omitempty"`
	Solvable bool   `protobuf:"varint,5,opt,name=solvable,proto3" json:"solvable,omitempty"`
}

func (x *V1MinesweeperPlayerData) Reset() {
	*x = V1MinesweeperPlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_player_data_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *V1MinesweeperPlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*V1MinesweeperPlayerData) ProtoMessage() {}

func (x *V1MinesweeperPlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_game_player_data_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use V1MinesweeperPlayerData.ProtoReflect.Descriptor instead.
func (*V1MinesweeperPlayerData) Descriptor() ([]byte, []int) {
	return file_game_player_data_models_proto_rawDescGZIP(), []int{3}
}

func (x *V1MinesweeperPlayerData) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *V1MinesweeperPlayerData) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *V1MinesweeperPlayerData) GetMines() uint32 {
	if x != nil {
		return x.Mines
	}
	return 0
}

func (x *V1MinesweeperPlayerData) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

func (x *V1MinesweeperPlayerData) GetSolvable() bool {
	if x != nil {
		return x.Solvable
	}
	return false
}

var File_game_player_data_models_proto protoreflect.FileDescriptor

var file_game_player_data_models_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x22, 0x57, 0x0a,
	0x15, 0x56, 0x31, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x75, 0x6d, 0x6f, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x65, 0x61, 0x72, 0x73, 0x5f,
	0x73, 0x6c, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x65, 0x61,
	0x72, 0x73, 0x53, 0x6c, 0x6f, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x31, 0x4d, 0x61, 0x72, 0x61,
	0x74, 0x68, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x18, 0x56, 0x31, 0x54, 0x6f,
	0x77, 0x65, 0x72, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x8f, 0x01, 0x0a, 0x17, 0x56, 0x31, 0x4d, 0x69, 0x6e, 0x65, 0x73,
	0x77, 0x65, 0x65, 0x70, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d,
	0x69, 0x6e, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f,
	0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x6f,
	0x6c, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x2a, 0x54, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x4f,
	0x57, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x46, 0x45, 0x4e, 0x43, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a,
	0x0a, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x4d, 0x4f, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x4d, 0x49, 0x4e, 0x45, 0x53, 0x57, 0x45, 0x45, 0x50, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x4d, 0x41, 0x52, 0x41, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x03, 0x42, 0x60, 0x0a, 0x1e,
	0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x64, 0x61, 0x74, 0x61, 0x50, 0x01,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f,
	0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x67, 0x61, 0x6d, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_player_data_models_proto_rawDescOnce sync.Once
	file_game_player_data_models_proto_rawDescData = file_game_player_data_models_proto_rawDesc
)

func file_game_player_data_models_proto_rawDescGZIP() []byte {
	file_game_player_data_models_proto_rawDescOnce.Do(func() {
		file_game_player_data_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_player_data_models_proto_rawDescData)
	})
	return file_game_player_data_models_proto_rawDescData
}

var file_game_player_data_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_player_data_models_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_game_player_data_models_proto_goTypes = []interface{}{
	(GameDataGameMode)(0),            // 0: emortal.model.gameplayerdata.GameDataGameMode
	(*V1BlockSumoPlayerData)(nil),    // 1: emortal.model.gameplayerdata.V1BlockSumoPlayerData
	(*V1MarathonData)(nil),           // 2: emortal.model.gameplayerdata.V1MarathonData
	(*V1TowerDefencePlayerData)(nil), // 3: emortal.model.gameplayerdata.V1TowerDefencePlayerData
	(*V1MinesweeperPlayerData)(nil),  // 4: emortal.model.gameplayerdata.V1MinesweeperPlayerData
}
var file_game_player_data_models_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_player_data_models_proto_init() }
func file_game_player_data_models_proto_init() {
	if File_game_player_data_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_player_data_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1BlockSumoPlayerData); i {
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
		file_game_player_data_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1MarathonData); i {
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
		file_game_player_data_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1TowerDefencePlayerData); i {
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
		file_game_player_data_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*V1MinesweeperPlayerData); i {
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
			RawDescriptor: file_game_player_data_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_player_data_models_proto_goTypes,
		DependencyIndexes: file_game_player_data_models_proto_depIdxs,
		EnumInfos:         file_game_player_data_models_proto_enumTypes,
		MessageInfos:      file_game_player_data_models_proto_msgTypes,
	}.Build()
	File_game_player_data_models_proto = out.File
	file_game_player_data_models_proto_rawDesc = nil
	file_game_player_data_models_proto_goTypes = nil
	file_game_player_data_models_proto_depIdxs = nil
}
