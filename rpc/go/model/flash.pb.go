// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: flash.proto

package model

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FlashMapStatisticCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string               `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Maps     []*FlashMapStatistic `protobuf:"bytes,2,rep,name=maps,proto3" json:"maps,omitempty"`
}

func (x *FlashMapStatisticCollection) Reset() {
	*x = FlashMapStatisticCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashMapStatisticCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashMapStatisticCollection) ProtoMessage() {}

func (x *FlashMapStatisticCollection) ProtoReflect() protoreflect.Message {
	mi := &file_flash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashMapStatisticCollection.ProtoReflect.Descriptor instead.
func (*FlashMapStatisticCollection) Descriptor() ([]byte, []int) {
	return file_flash_proto_rawDescGZIP(), []int{0}
}

func (x *FlashMapStatisticCollection) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *FlashMapStatisticCollection) GetMaps() []*FlashMapStatistic {
	if x != nil {
		return x.Maps
	}
	return nil
}

type FlashAllStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId    string                `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	PlayerStats *FlashPlayerStatistic `protobuf:"bytes,2,opt,name=playerStats,proto3" json:"playerStats,omitempty"`
	MapStats    []*FlashMapStatistic  `protobuf:"bytes,3,rep,name=mapStats,proto3" json:"mapStats,omitempty"`
}

func (x *FlashAllStatistics) Reset() {
	*x = FlashAllStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashAllStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashAllStatistics) ProtoMessage() {}

func (x *FlashAllStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_flash_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashAllStatistics.ProtoReflect.Descriptor instead.
func (*FlashAllStatistics) Descriptor() ([]byte, []int) {
	return file_flash_proto_rawDescGZIP(), []int{1}
}

func (x *FlashAllStatistics) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *FlashAllStatistics) GetPlayerStats() *FlashPlayerStatistic {
	if x != nil {
		return x.PlayerStats
	}
	return nil
}

func (x *FlashAllStatistics) GetMapStats() []*FlashMapStatistic {
	if x != nil {
		return x.MapStats
	}
	return nil
}

type FlashPlayerStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wins          uint32 `protobuf:"varint,1,opt,name=wins,proto3" json:"wins,omitempty"`
	Deaths        uint32 `protobuf:"varint,2,opt,name=deaths,proto3" json:"deaths,omitempty"`
	GamesPlayed   uint32 `protobuf:"varint,3,opt,name=games_played,json=gamesPlayed,proto3" json:"games_played,omitempty"`
	InstantDeaths uint32 `protobuf:"varint,4,opt,name=instant_deaths,json=instantDeaths,proto3" json:"instant_deaths,omitempty"`
	Checkpoints   uint32 `protobuf:"varint,5,opt,name=checkpoints,proto3" json:"checkpoints,omitempty"`
	Points        uint32 `protobuf:"varint,6,opt,name=points,proto3" json:"points,omitempty"`
}

func (x *FlashPlayerStatistic) Reset() {
	*x = FlashPlayerStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashPlayerStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashPlayerStatistic) ProtoMessage() {}

func (x *FlashPlayerStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_flash_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashPlayerStatistic.ProtoReflect.Descriptor instead.
func (*FlashPlayerStatistic) Descriptor() ([]byte, []int) {
	return file_flash_proto_rawDescGZIP(), []int{2}
}

func (x *FlashPlayerStatistic) GetWins() uint32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *FlashPlayerStatistic) GetDeaths() uint32 {
	if x != nil {
		return x.Deaths
	}
	return 0
}

func (x *FlashPlayerStatistic) GetGamesPlayed() uint32 {
	if x != nil {
		return x.GamesPlayed
	}
	return 0
}

func (x *FlashPlayerStatistic) GetInstantDeaths() uint32 {
	if x != nil {
		return x.InstantDeaths
	}
	return 0
}

func (x *FlashPlayerStatistic) GetCheckpoints() uint32 {
	if x != nil {
		return x.Checkpoints
	}
	return 0
}

func (x *FlashPlayerStatistic) GetPoints() uint32 {
	if x != nil {
		return x.Points
	}
	return 0
}

type FlashMapStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name           string                      `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeNeeded     uint64                      `protobuf:"varint,2,opt,name=time_needed,json=timeNeeded,proto3" json:"time_needed,omitempty"`
	AccomplishedAt string                      `protobuf:"bytes,3,opt,name=accomplished_at,json=accomplishedAt,proto3" json:"accomplished_at,omitempty"`
	Checkpoints    []*FlashCheckpointStatistic `protobuf:"bytes,4,rep,name=checkpoints,proto3" json:"checkpoints,omitempty"`
}

func (x *FlashMapStatistic) Reset() {
	*x = FlashMapStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashMapStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashMapStatistic) ProtoMessage() {}

func (x *FlashMapStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_flash_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashMapStatistic.ProtoReflect.Descriptor instead.
func (*FlashMapStatistic) Descriptor() ([]byte, []int) {
	return file_flash_proto_rawDescGZIP(), []int{3}
}

func (x *FlashMapStatistic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FlashMapStatistic) GetTimeNeeded() uint64 {
	if x != nil {
		return x.TimeNeeded
	}
	return 0
}

func (x *FlashMapStatistic) GetAccomplishedAt() string {
	if x != nil {
		return x.AccomplishedAt
	}
	return ""
}

func (x *FlashMapStatistic) GetCheckpoints() []*FlashCheckpointStatistic {
	if x != nil {
		return x.Checkpoints
	}
	return nil
}

type FlashCheckpointStatistic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Checkpoint     int32  `protobuf:"fixed32,1,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
	TimeNeeded     uint64 `protobuf:"varint,2,opt,name=time_needed,json=timeNeeded,proto3" json:"time_needed,omitempty"`
	AccomplishedAt string `protobuf:"bytes,3,opt,name=accomplished_at,json=accomplishedAt,proto3" json:"accomplished_at,omitempty"`
}

func (x *FlashCheckpointStatistic) Reset() {
	*x = FlashCheckpointStatistic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flash_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlashCheckpointStatistic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlashCheckpointStatistic) ProtoMessage() {}

func (x *FlashCheckpointStatistic) ProtoReflect() protoreflect.Message {
	mi := &file_flash_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlashCheckpointStatistic.ProtoReflect.Descriptor instead.
func (*FlashCheckpointStatistic) Descriptor() ([]byte, []int) {
	return file_flash_proto_rawDescGZIP(), []int{4}
}

func (x *FlashCheckpointStatistic) GetCheckpoint() int32 {
	if x != nil {
		return x.Checkpoint
	}
	return 0
}

func (x *FlashCheckpointStatistic) GetTimeNeeded() uint64 {
	if x != nil {
		return x.TimeNeeded
	}
	return 0
}

func (x *FlashCheckpointStatistic) GetAccomplishedAt() string {
	if x != nil {
		return x.AccomplishedAt
	}
	return ""
}

var File_flash_proto protoreflect.FileDescriptor

var file_flash_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a,
	0x1b, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x6d, 0x61, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73,
	0x22, 0x99, 0x01, 0x0a, 0x12, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x08,
	0x6d, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x73, 0x22, 0xc6, 0x01, 0x0a,
	0x14, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x64, 0x65, 0x61, 0x74, 0x68,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f,
	0x64, 0x65, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x61, 0x74, 0x68, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x11, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d,
	0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x18, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x65, 0x65, 0x64,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x4e, 0x65,
	0x65, 0x64, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x42, 0x37, 0x0a,
	0x16, 0x64, 0x65, 0x76, 0x2e, 0x66, 0x72, 0x65, 0x67, 0x67, 0x79, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5a, 0x1d, 0x66, 0x72, 0x65, 0x67, 0x67, 0x79, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flash_proto_rawDescOnce sync.Once
	file_flash_proto_rawDescData = file_flash_proto_rawDesc
)

func file_flash_proto_rawDescGZIP() []byte {
	file_flash_proto_rawDescOnce.Do(func() {
		file_flash_proto_rawDescData = protoimpl.X.CompressGZIP(file_flash_proto_rawDescData)
	})
	return file_flash_proto_rawDescData
}

var file_flash_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_flash_proto_goTypes = []interface{}{
	(*FlashMapStatisticCollection)(nil), // 0: FlashMapStatisticCollection
	(*FlashAllStatistics)(nil),          // 1: FlashAllStatistics
	(*FlashPlayerStatistic)(nil),        // 2: FlashPlayerStatistic
	(*FlashMapStatistic)(nil),           // 3: FlashMapStatistic
	(*FlashCheckpointStatistic)(nil),    // 4: FlashCheckpointStatistic
}
var file_flash_proto_depIdxs = []int32{
	3, // 0: FlashMapStatisticCollection.maps:type_name -> FlashMapStatistic
	2, // 1: FlashAllStatistics.playerStats:type_name -> FlashPlayerStatistic
	3, // 2: FlashAllStatistics.mapStats:type_name -> FlashMapStatistic
	4, // 3: FlashMapStatistic.checkpoints:type_name -> FlashCheckpointStatistic
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_flash_proto_init() }
func file_flash_proto_init() {
	if File_flash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlashMapStatisticCollection); i {
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
		file_flash_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlashAllStatistics); i {
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
		file_flash_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlashPlayerStatistic); i {
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
		file_flash_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlashMapStatistic); i {
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
		file_flash_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlashCheckpointStatistic); i {
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
			RawDescriptor: file_flash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flash_proto_goTypes,
		DependencyIndexes: file_flash_proto_depIdxs,
		MessageInfos:      file_flash_proto_msgTypes,
	}.Build()
	File_flash_proto = out.File
	file_flash_proto_rawDesc = nil
	file_flash_proto_goTypes = nil
	file_flash_proto_depIdxs = nil
}
