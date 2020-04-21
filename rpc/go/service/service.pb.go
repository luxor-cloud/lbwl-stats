// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: service.proto

package service

import (
	model "freggy.dev/stats/rpc/go/model"
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

type GetFlashMapHighscoreForPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId        string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	WithCheckpoints bool   `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
}

func (x *GetFlashMapHighscoreForPlayerRequest) Reset() {
	*x = GetFlashMapHighscoreForPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlashMapHighscoreForPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlashMapHighscoreForPlayerRequest) ProtoMessage() {}

func (x *GetFlashMapHighscoreForPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlashMapHighscoreForPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetFlashMapHighscoreForPlayerRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetFlashMapHighscoreForPlayerRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *GetFlashMapHighscoreForPlayerRequest) GetWithCheckpoints() bool {
	if x != nil {
		return x.WithCheckpoints
	}
	return false
}

type GetFlashMapHighscoreForPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Highscores *model.FlashMapStatisticCollection `protobuf:"bytes,1,opt,name=highscores,proto3" json:"highscores,omitempty"`
}

func (x *GetFlashMapHighscoreForPlayerResponse) Reset() {
	*x = GetFlashMapHighscoreForPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFlashMapHighscoreForPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFlashMapHighscoreForPlayerResponse) ProtoMessage() {}

func (x *GetFlashMapHighscoreForPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFlashMapHighscoreForPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetFlashMapHighscoreForPlayerResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetFlashMapHighscoreForPlayerResponse) GetHighscores() *model.FlashMapStatisticCollection {
	if x != nil {
		return x.Highscores
	}
	return nil
}

type GetGlobalFlashMapHighscoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapName         string `protobuf:"bytes,1,opt,name=mapName,proto3" json:"mapName,omitempty"`
	WithCheckpoints bool   `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
}

func (x *GetGlobalFlashMapHighscoreRequest) Reset() {
	*x = GetGlobalFlashMapHighscoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGlobalFlashMapHighscoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGlobalFlashMapHighscoreRequest) ProtoMessage() {}

func (x *GetGlobalFlashMapHighscoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGlobalFlashMapHighscoreRequest.ProtoReflect.Descriptor instead.
func (*GetGlobalFlashMapHighscoreRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetGlobalFlashMapHighscoreRequest) GetMapName() string {
	if x != nil {
		return x.MapName
	}
	return ""
}

func (x *GetGlobalFlashMapHighscoreRequest) GetWithCheckpoints() bool {
	if x != nil {
		return x.WithCheckpoints
	}
	return false
}

type GetGlobalFlashMapHighscoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Highscore *model.FlashMapStatistic `protobuf:"bytes,1,opt,name=highscore,proto3" json:"highscore,omitempty"`
}

func (x *GetGlobalFlashMapHighscoreResponse) Reset() {
	*x = GetGlobalFlashMapHighscoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGlobalFlashMapHighscoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGlobalFlashMapHighscoreResponse) ProtoMessage() {}

func (x *GetGlobalFlashMapHighscoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGlobalFlashMapHighscoreResponse.ProtoReflect.Descriptor instead.
func (*GetGlobalFlashMapHighscoreResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetGlobalFlashMapHighscoreResponse) GetHighscore() *model.FlashMapStatistic {
	if x != nil {
		return x.Highscore
	}
	return nil
}

type GetTopFlashMapHighscoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit           uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	WithCheckpoints bool   `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
}

func (x *GetTopFlashMapHighscoresRequest) Reset() {
	*x = GetTopFlashMapHighscoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopFlashMapHighscoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopFlashMapHighscoresRequest) ProtoMessage() {}

func (x *GetTopFlashMapHighscoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopFlashMapHighscoresRequest.ProtoReflect.Descriptor instead.
func (*GetTopFlashMapHighscoresRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTopFlashMapHighscoresRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTopFlashMapHighscoresRequest) GetWithCheckpoints() bool {
	if x != nil {
		return x.WithCheckpoints
	}
	return false
}

type GetTopFlashMapHighscoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Highscores *model.FlashMapStatisticCollection `protobuf:"bytes,1,opt,name=highscores,proto3" json:"highscores,omitempty"`
}

func (x *GetTopFlashMapHighscoresResponse) Reset() {
	*x = GetTopFlashMapHighscoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopFlashMapHighscoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopFlashMapHighscoresResponse) ProtoMessage() {}

func (x *GetTopFlashMapHighscoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopFlashMapHighscoresResponse.ProtoReflect.Descriptor instead.
func (*GetTopFlashMapHighscoresResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTopFlashMapHighscoresResponse) GetHighscores() *model.FlashMapStatisticCollection {
	if x != nil {
		return x.Highscores
	}
	return nil
}

type GetTopPlayersByPointsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerStats     bool `protobuf:"varint,1,opt,name=playerStats,proto3" json:"playerStats,omitempty"`
	MapStats        bool `protobuf:"varint,2,opt,name=mapStats,proto3" json:"mapStats,omitempty"`
	WithCheckpoints bool `protobuf:"varint,3,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
	Limit           bool `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetTopPlayersByPointsRequest) Reset() {
	*x = GetTopPlayersByPointsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopPlayersByPointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopPlayersByPointsRequest) ProtoMessage() {}

func (x *GetTopPlayersByPointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopPlayersByPointsRequest.ProtoReflect.Descriptor instead.
func (*GetTopPlayersByPointsRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTopPlayersByPointsRequest) GetPlayerStats() bool {
	if x != nil {
		return x.PlayerStats
	}
	return false
}

func (x *GetTopPlayersByPointsRequest) GetMapStats() bool {
	if x != nil {
		return x.MapStats
	}
	return false
}

func (x *GetTopPlayersByPointsRequest) GetWithCheckpoints() bool {
	if x != nil {
		return x.WithCheckpoints
	}
	return false
}

func (x *GetTopPlayersByPointsRequest) GetLimit() bool {
	if x != nil {
		return x.Limit
	}
	return false
}

type GetTopPlayersByPointsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopPlayers []*model.FlashAllStatistics `protobuf:"bytes,1,rep,name=topPlayers,proto3" json:"topPlayers,omitempty"`
}

func (x *GetTopPlayersByPointsResponse) Reset() {
	*x = GetTopPlayersByPointsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTopPlayersByPointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTopPlayersByPointsResponse) ProtoMessage() {}

func (x *GetTopPlayersByPointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTopPlayersByPointsResponse.ProtoReflect.Descriptor instead.
func (*GetTopPlayersByPointsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetTopPlayersByPointsResponse) GetTopPlayers() []*model.FlashAllStatistics {
	if x != nil {
		return x.TopPlayers
	}
	return nil
}

type UpdateFlashStatisticsRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string                    `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Stats    *model.FlashAllStatistics `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *UpdateFlashStatisticsRequests) Reset() {
	*x = UpdateFlashStatisticsRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlashStatisticsRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlashStatisticsRequests) ProtoMessage() {}

func (x *UpdateFlashStatisticsRequests) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlashStatisticsRequests.ProtoReflect.Descriptor instead.
func (*UpdateFlashStatisticsRequests) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateFlashStatisticsRequests) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *UpdateFlashStatisticsRequests) GetStats() *model.FlashAllStatistics {
	if x != nil {
		return x.Stats
	}
	return nil
}

type UpdateFlashStatisticsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFlashStatisticsResponse) Reset() {
	*x = UpdateFlashStatisticsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFlashStatisticsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFlashStatisticsResponse) ProtoMessage() {}

func (x *UpdateFlashStatisticsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFlashStatisticsResponse.ProtoReflect.Descriptor instead.
func (*UpdateFlashStatisticsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{9}
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x66, 0x6c, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x24, 0x47,
	0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x28, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x65, 0x0a, 0x25, 0x47, 0x65, 0x74,
	0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x22, 0x67, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x46, 0x6c, 0x61,
	0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x56, 0x0a, 0x22, 0x47, 0x65, 0x74,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69,
	0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x52, 0x09, 0x68, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x22, 0x61, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x6c, 0x61, 0x73, 0x68,
	0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x69,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x6c,
	0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x61, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x77, 0x69, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x54, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x46, 0x6c, 0x61,
	0x73, 0x68, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52,
	0x0a, 0x74, 0x6f, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x66, 0x0a, 0x1d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x41,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61,
	0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xc7, 0x04, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73,
	0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x46, 0x6f, 0x72,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x75, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48,
	0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6f, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69,
	0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61,
	0x70, 0x48, 0x69, 0x67, 0x68, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x70, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x70, 0x48, 0x69, 0x67, 0x68, 0x73,
	0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x42, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x26,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6c, 0x61, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x61, 0x73, 0x68, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25,
	0x0a, 0x18, 0x64, 0x65, 0x76, 0x2e, 0x66, 0x72, 0x65, 0x67, 0x67, 0x79, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x09, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_proto_goTypes = []interface{}{
	(*GetFlashMapHighscoreForPlayerRequest)(nil),  // 0: service.GetFlashMapHighscoreForPlayerRequest
	(*GetFlashMapHighscoreForPlayerResponse)(nil), // 1: service.GetFlashMapHighscoreForPlayerResponse
	(*GetGlobalFlashMapHighscoreRequest)(nil),     // 2: service.GetGlobalFlashMapHighscoreRequest
	(*GetGlobalFlashMapHighscoreResponse)(nil),    // 3: service.GetGlobalFlashMapHighscoreResponse
	(*GetTopFlashMapHighscoresRequest)(nil),       // 4: service.GetTopFlashMapHighscoresRequest
	(*GetTopFlashMapHighscoresResponse)(nil),      // 5: service.GetTopFlashMapHighscoresResponse
	(*GetTopPlayersByPointsRequest)(nil),          // 6: service.GetTopPlayersByPointsRequest
	(*GetTopPlayersByPointsResponse)(nil),         // 7: service.GetTopPlayersByPointsResponse
	(*UpdateFlashStatisticsRequests)(nil),         // 8: service.UpdateFlashStatisticsRequests
	(*UpdateFlashStatisticsResponse)(nil),         // 9: service.UpdateFlashStatisticsResponse
	(*model.FlashMapStatisticCollection)(nil),     // 10: FlashMapStatisticCollection
	(*model.FlashMapStatistic)(nil),               // 11: FlashMapStatistic
	(*model.FlashAllStatistics)(nil),              // 12: FlashAllStatistics
}
var file_service_proto_depIdxs = []int32{
	10, // 0: service.GetFlashMapHighscoreForPlayerResponse.highscores:type_name -> FlashMapStatisticCollection
	11, // 1: service.GetGlobalFlashMapHighscoreResponse.highscore:type_name -> FlashMapStatistic
	10, // 2: service.GetTopFlashMapHighscoresResponse.highscores:type_name -> FlashMapStatisticCollection
	12, // 3: service.GetTopPlayersByPointsResponse.topPlayers:type_name -> FlashAllStatistics
	12, // 4: service.UpdateFlashStatisticsRequests.stats:type_name -> FlashAllStatistics
	0,  // 5: service.StatsService.GetFlashMapHighscoreForPlayer:input_type -> service.GetFlashMapHighscoreForPlayerRequest
	2,  // 6: service.StatsService.GetGlobalFlashMapHighscore:input_type -> service.GetGlobalFlashMapHighscoreRequest
	4,  // 7: service.StatsService.GetTopFlashMapHighscores:input_type -> service.GetTopFlashMapHighscoresRequest
	6,  // 8: service.StatsService.GetTopPlayersByPoints:input_type -> service.GetTopPlayersByPointsRequest
	8,  // 9: service.StatsService.UpdateFlashStatistics:input_type -> service.UpdateFlashStatisticsRequests
	1,  // 10: service.StatsService.GetFlashMapHighscoreForPlayer:output_type -> service.GetFlashMapHighscoreForPlayerResponse
	3,  // 11: service.StatsService.GetGlobalFlashMapHighscore:output_type -> service.GetGlobalFlashMapHighscoreResponse
	5,  // 12: service.StatsService.GetTopFlashMapHighscores:output_type -> service.GetTopFlashMapHighscoresResponse
	7,  // 13: service.StatsService.GetTopPlayersByPoints:output_type -> service.GetTopPlayersByPointsResponse
	9,  // 14: service.StatsService.UpdateFlashStatistics:output_type -> service.UpdateFlashStatisticsResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlashMapHighscoreForPlayerRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFlashMapHighscoreForPlayerResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGlobalFlashMapHighscoreRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGlobalFlashMapHighscoreResponse); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopFlashMapHighscoresRequest); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopFlashMapHighscoresResponse); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopPlayersByPointsRequest); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTopPlayersByPointsResponse); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlashStatisticsRequests); i {
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
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFlashStatisticsResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
