// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0
// source: kurushimi/models.proto

package matchmaker

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type QueuedPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of type UUID (minecraft player id)
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ticket_id of type ObjectID, the ID of the ticket that the player belongs to
	TicketId string  `protobuf:"bytes,2,opt,name=ticket_id,json=ticketId,proto3" json:"ticket_id,omitempty"`
	MapId    *string `protobuf:"bytes,3,opt,name=map_id,json=mapId,proto3,oneof" json:"map_id,omitempty"`
}

func (x *QueuedPlayer) Reset() {
	*x = QueuedPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueuedPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueuedPlayer) ProtoMessage() {}

func (x *QueuedPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueuedPlayer.ProtoReflect.Descriptor instead.
func (*QueuedPlayer) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{0}
}

func (x *QueuedPlayer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QueuedPlayer) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *QueuedPlayer) GetMapId() string {
	if x != nil && x.MapId != nil {
		return *x.MapId
	}
	return ""
}

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// party_id of type ObjectID (mongo id)
	// Whilst technically optional, it is very rare. The party_id is only null on join lobby selection
	// as the player does not yet have a party.
	PartyId   *string                `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3,oneof" json:"party_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// player_ids of type UUID (minecraft player id)
	PlayerIds  []string `protobuf:"bytes,4,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
	GameModeId string   `protobuf:"bytes,5,opt,name=game_mode_id,json=gameModeId,proto3" json:"game_mode_id,omitempty"`
	// auto_teleport if true, the proxy will listen for match messages and teleport the player to the match.
	// if false, the proxy will not auto-teleport the player and you should listen for match messages yourself.
	//
	// e.g. this is used by the proxy for lobby matchmaking as it is handled differently
	// to when the player is already connected to the server.
	//
	// NOTE: This field is optional in the request, but not optional afterwards and is resolved to true/false from the optional.
	AutoTeleport bool `protobuf:"varint,6,opt,name=auto_teleport,json=autoTeleport,proto3" json:"auto_teleport,omitempty"`
	// dequeue_on_disconnect: is determined by the party settings. True if the player is not in a party, or if the party
	// is not set to stay in queue on disconnect.
	// NOTE: The ticket will be deleted anyway if the party is deleted.
	DequeueOnDisconnect bool `protobuf:"varint,7,opt,name=dequeue_on_disconnect,json=dequeueOnDisconnect,proto3" json:"dequeue_on_disconnect,omitempty"`
	InPendingMatch      bool `protobuf:"varint,8,opt,name=in_pending_match,json=inPendingMatch,proto3" json:"in_pending_match,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

func (x *Ticket) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{1}
}

func (x *Ticket) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ticket) GetPartyId() string {
	if x != nil && x.PartyId != nil {
		return *x.PartyId
	}
	return ""
}

func (x *Ticket) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Ticket) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

func (x *Ticket) GetGameModeId() string {
	if x != nil {
		return x.GameModeId
	}
	return ""
}

func (x *Ticket) GetAutoTeleport() bool {
	if x != nil {
		return x.AutoTeleport
	}
	return false
}

func (x *Ticket) GetDequeueOnDisconnect() bool {
	if x != nil {
		return x.DequeueOnDisconnect
	}
	return false
}

func (x *Ticket) GetInPendingMatch() bool {
	if x != nil {
		return x.InPendingMatch
	}
	return false
}

type Match struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GameModeId string `protobuf:"bytes,2,opt,name=game_mode_id,json=gameModeId,proto3" json:"game_mode_id,omitempty"`
	// map_id selected by the matchmaker based on the game mode config
	// Only present if the game mode has a map pool.
	// Clients should still always account for it being null for testing purposes (random map when in game dev mode)
	MapId *string `protobuf:"bytes,3,opt,name=map_id,json=mapId,proto3,oneof" json:"map_id,omitempty"`
	// tickets is a source of truth for all the grouped players in the match.
	Tickets    []*Ticket   `protobuf:"bytes,4,rep,name=tickets,proto3" json:"tickets,omitempty"`
	Assignment *Assignment `protobuf:"bytes,5,opt,name=assignment,proto3,oneof" json:"assignment,omitempty"`
}

func (x *Match) Reset() {
	*x = Match{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Match) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Match) ProtoMessage() {}

func (x *Match) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Match.ProtoReflect.Descriptor instead.
func (*Match) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{2}
}

func (x *Match) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Match) GetGameModeId() string {
	if x != nil {
		return x.GameModeId
	}
	return ""
}

func (x *Match) GetMapId() string {
	if x != nil && x.MapId != nil {
		return *x.MapId
	}
	return ""
}

func (x *Match) GetTickets() []*Ticket {
	if x != nil {
		return x.Tickets
	}
	return nil
}

func (x *Match) GetAssignment() *Assignment {
	if x != nil {
		return x.Assignment
	}
	return nil
}

type Assignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId      string `protobuf:"bytes,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerAddress string `protobuf:"bytes,2,opt,name=server_address,json=serverAddress,proto3" json:"server_address,omitempty"`
	ServerPort    uint32 `protobuf:"varint,3,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	// protocol_version e.g. 754
	// Not present for proxies
	ProtocolVersion *int64 `protobuf:"varint,4,opt,name=protocol_version,json=protocolVersion,proto3,oneof" json:"protocol_version,omitempty"`
	// version_name e.g. 1.19.4
	// Not present for proxies
	VersionName *string `protobuf:"bytes,5,opt,name=version_name,json=versionName,proto3,oneof" json:"version_name,omitempty"`
}

func (x *Assignment) Reset() {
	*x = Assignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assignment) ProtoMessage() {}

func (x *Assignment) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assignment.ProtoReflect.Descriptor instead.
func (*Assignment) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{3}
}

func (x *Assignment) GetServerId() string {
	if x != nil {
		return x.ServerId
	}
	return ""
}

func (x *Assignment) GetServerAddress() string {
	if x != nil {
		return x.ServerAddress
	}
	return ""
}

func (x *Assignment) GetServerPort() uint32 {
	if x != nil {
		return x.ServerPort
	}
	return 0
}

func (x *Assignment) GetProtocolVersion() int64 {
	if x != nil && x.ProtocolVersion != nil {
		return *x.ProtocolVersion
	}
	return 0
}

func (x *Assignment) GetVersionName() string {
	if x != nil && x.VersionName != nil {
		return *x.VersionName
	}
	return ""
}

type PendingMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	GameModeId string `protobuf:"bytes,2,opt,name=game_mode_id,json=gameModeId,proto3" json:"game_mode_id,omitempty"`
	// ticket_ids of type ObjectID (mongo id)
	TicketIds    []string               `protobuf:"bytes,3,rep,name=ticket_ids,json=ticketIds,proto3" json:"ticket_ids,omitempty"`
	TeleportTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=teleport_time,json=teleportTime,proto3" json:"teleport_time,omitempty"`
	// player_count is the number of players in the tickets of the PendingMatch.
	// This is updated each time the matchfunction is run.
	PlayerCount int64 `protobuf:"varint,5,opt,name=player_count,json=playerCount,proto3" json:"player_count,omitempty"`
}

func (x *PendingMatch) Reset() {
	*x = PendingMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingMatch) ProtoMessage() {}

func (x *PendingMatch) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingMatch.ProtoReflect.Descriptor instead.
func (*PendingMatch) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{4}
}

func (x *PendingMatch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PendingMatch) GetGameModeId() string {
	if x != nil {
		return x.GameModeId
	}
	return ""
}

func (x *PendingMatch) GetTicketIds() []string {
	if x != nil {
		return x.TicketIds
	}
	return nil
}

func (x *PendingMatch) GetTeleportTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TeleportTime
	}
	return nil
}

func (x *PendingMatch) GetPlayerCount() int64 {
	if x != nil {
		return x.PlayerCount
	}
	return 0
}

type AllocationData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match *Match `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
}

func (x *AllocationData) Reset() {
	*x = AllocationData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kurushimi_models_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllocationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllocationData) ProtoMessage() {}

func (x *AllocationData) ProtoReflect() protoreflect.Message {
	mi := &file_kurushimi_models_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllocationData.ProtoReflect.Descriptor instead.
func (*AllocationData) Descriptor() ([]byte, []int) {
	return file_kurushimi_models_proto_rawDescGZIP(), []int{5}
}

func (x *AllocationData) GetMatch() *Match {
	if x != nil {
		return x.Match
	}
	return nil
}

var File_kurushimi_models_proto protoreflect.FileDescriptor

var file_kurushimi_models_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6b, 0x75, 0x72, 0x75, 0x73, 0x68, 0x69, 0x6d, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61,
	0x6c, 0x2e, 0x6b, 0x75, 0x72, 0x75, 0x73, 0x68, 0x69, 0x6d, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x06, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x22, 0xc4, 0x02, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1e, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x64, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6f, 0x6e,
	0x5f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x64, 0x65, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4f, 0x6e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x5f, 0x70, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x69, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x22, 0xf4, 0x01,
	0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x06, 0x6d, 0x61, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x61, 0x70,
	0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c,
	0x2e, 0x6b, 0x75, 0x72, 0x75, 0x73, 0x68, 0x69, 0x6d, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x48, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6b,
	0x75, 0x72, 0x75, 0x73, 0x68, 0x69, 0x6d, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x01, 0x52, 0x0a, 0x61, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6d,
	0x61, 0x70, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0xef, 0x01, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2e, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xc3, 0x01, 0x0a, 0x0c, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x0e,
	0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34,
	0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x6b, 0x75, 0x72, 0x75, 0x73, 0x68, 0x69, 0x6d,
	0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x42, 0x5e, 0x0a, 0x20, 0x64, 0x65, 0x76, 0x2e, 0x65, 0x6d, 0x6f, 0x72,
	0x74, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x6d, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x6d,
	0x61, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kurushimi_models_proto_rawDescOnce sync.Once
	file_kurushimi_models_proto_rawDescData = file_kurushimi_models_proto_rawDesc
)

func file_kurushimi_models_proto_rawDescGZIP() []byte {
	file_kurushimi_models_proto_rawDescOnce.Do(func() {
		file_kurushimi_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_kurushimi_models_proto_rawDescData)
	})
	return file_kurushimi_models_proto_rawDescData
}

var file_kurushimi_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_kurushimi_models_proto_goTypes = []interface{}{
	(*QueuedPlayer)(nil),          // 0: emortal.kurushimi.model.QueuedPlayer
	(*Ticket)(nil),                // 1: emortal.kurushimi.model.Ticket
	(*Match)(nil),                 // 2: emortal.kurushimi.model.Match
	(*Assignment)(nil),            // 3: emortal.kurushimi.model.Assignment
	(*PendingMatch)(nil),          // 4: emortal.kurushimi.model.PendingMatch
	(*AllocationData)(nil),        // 5: emortal.kurushimi.model.AllocationData
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_kurushimi_models_proto_depIdxs = []int32{
	6, // 0: emortal.kurushimi.model.Ticket.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: emortal.kurushimi.model.Match.tickets:type_name -> emortal.kurushimi.model.Ticket
	3, // 2: emortal.kurushimi.model.Match.assignment:type_name -> emortal.kurushimi.model.Assignment
	6, // 3: emortal.kurushimi.model.PendingMatch.teleport_time:type_name -> google.protobuf.Timestamp
	2, // 4: emortal.kurushimi.model.AllocationData.match:type_name -> emortal.kurushimi.model.Match
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_kurushimi_models_proto_init() }
func file_kurushimi_models_proto_init() {
	if File_kurushimi_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kurushimi_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueuedPlayer); i {
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
		file_kurushimi_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
		file_kurushimi_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Match); i {
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
		file_kurushimi_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assignment); i {
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
		file_kurushimi_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingMatch); i {
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
		file_kurushimi_models_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllocationData); i {
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
	file_kurushimi_models_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_kurushimi_models_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_kurushimi_models_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_kurushimi_models_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_kurushimi_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kurushimi_models_proto_goTypes,
		DependencyIndexes: file_kurushimi_models_proto_depIdxs,
		MessageInfos:      file_kurushimi_models_proto_msgTypes,
	}.Build()
	File_kurushimi_models_proto = out.File
	file_kurushimi_models_proto_rawDesc = nil
	file_kurushimi_models_proto_goTypes = nil
	file_kurushimi_models_proto_depIdxs = nil
}
