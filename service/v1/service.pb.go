// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package v1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetFlashMapHighscoreForPlayerRequest struct {
	PlayerId             string   `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	WithCheckpoints      bool     `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
	MapName              string   `protobuf:"bytes,3,opt,name=mapName,proto3" json:"mapName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFlashMapHighscoreForPlayerRequest) Reset()         { *m = GetFlashMapHighscoreForPlayerRequest{} }
func (m *GetFlashMapHighscoreForPlayerRequest) String() string { return proto.CompactTextString(m) }
func (*GetFlashMapHighscoreForPlayerRequest) ProtoMessage()    {}
func (*GetFlashMapHighscoreForPlayerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetFlashMapHighscoreForPlayerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest.Unmarshal(m, b)
}
func (m *GetFlashMapHighscoreForPlayerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest.Marshal(b, m, deterministic)
}
func (m *GetFlashMapHighscoreForPlayerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest.Merge(m, src)
}
func (m *GetFlashMapHighscoreForPlayerRequest) XXX_Size() int {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest.Size(m)
}
func (m *GetFlashMapHighscoreForPlayerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlashMapHighscoreForPlayerRequest proto.InternalMessageInfo

func (m *GetFlashMapHighscoreForPlayerRequest) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *GetFlashMapHighscoreForPlayerRequest) GetWithCheckpoints() bool {
	if m != nil {
		return m.WithCheckpoints
	}
	return false
}

func (m *GetFlashMapHighscoreForPlayerRequest) GetMapName() string {
	if m != nil {
		return m.MapName
	}
	return ""
}

type GetFlashMapHighscoreForPlayerResponse struct {
	Highscore            *FlashMapStatistic `protobuf:"bytes,1,opt,name=highscore,proto3" json:"highscore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetFlashMapHighscoreForPlayerResponse) Reset()         { *m = GetFlashMapHighscoreForPlayerResponse{} }
func (m *GetFlashMapHighscoreForPlayerResponse) String() string { return proto.CompactTextString(m) }
func (*GetFlashMapHighscoreForPlayerResponse) ProtoMessage()    {}
func (*GetFlashMapHighscoreForPlayerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetFlashMapHighscoreForPlayerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse.Unmarshal(m, b)
}
func (m *GetFlashMapHighscoreForPlayerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse.Marshal(b, m, deterministic)
}
func (m *GetFlashMapHighscoreForPlayerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse.Merge(m, src)
}
func (m *GetFlashMapHighscoreForPlayerResponse) XXX_Size() int {
	return xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse.Size(m)
}
func (m *GetFlashMapHighscoreForPlayerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFlashMapHighscoreForPlayerResponse proto.InternalMessageInfo

func (m *GetFlashMapHighscoreForPlayerResponse) GetHighscore() *FlashMapStatistic {
	if m != nil {
		return m.Highscore
	}
	return nil
}

type GetGlobalFlashMapHighscoreRequest struct {
	MapName              string   `protobuf:"bytes,1,opt,name=mapName,proto3" json:"mapName,omitempty"`
	WithCheckpoints      bool     `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGlobalFlashMapHighscoreRequest) Reset()         { *m = GetGlobalFlashMapHighscoreRequest{} }
func (m *GetGlobalFlashMapHighscoreRequest) String() string { return proto.CompactTextString(m) }
func (*GetGlobalFlashMapHighscoreRequest) ProtoMessage()    {}
func (*GetGlobalFlashMapHighscoreRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *GetGlobalFlashMapHighscoreRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreRequest.Unmarshal(m, b)
}
func (m *GetGlobalFlashMapHighscoreRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreRequest.Marshal(b, m, deterministic)
}
func (m *GetGlobalFlashMapHighscoreRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGlobalFlashMapHighscoreRequest.Merge(m, src)
}
func (m *GetGlobalFlashMapHighscoreRequest) XXX_Size() int {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreRequest.Size(m)
}
func (m *GetGlobalFlashMapHighscoreRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGlobalFlashMapHighscoreRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGlobalFlashMapHighscoreRequest proto.InternalMessageInfo

func (m *GetGlobalFlashMapHighscoreRequest) GetMapName() string {
	if m != nil {
		return m.MapName
	}
	return ""
}

func (m *GetGlobalFlashMapHighscoreRequest) GetWithCheckpoints() bool {
	if m != nil {
		return m.WithCheckpoints
	}
	return false
}

type GetGlobalFlashMapHighscoreResponse struct {
	Highscore            *FlashMapStatistic `protobuf:"bytes,1,opt,name=highscore,proto3" json:"highscore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetGlobalFlashMapHighscoreResponse) Reset()         { *m = GetGlobalFlashMapHighscoreResponse{} }
func (m *GetGlobalFlashMapHighscoreResponse) String() string { return proto.CompactTextString(m) }
func (*GetGlobalFlashMapHighscoreResponse) ProtoMessage()    {}
func (*GetGlobalFlashMapHighscoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *GetGlobalFlashMapHighscoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreResponse.Unmarshal(m, b)
}
func (m *GetGlobalFlashMapHighscoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreResponse.Marshal(b, m, deterministic)
}
func (m *GetGlobalFlashMapHighscoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGlobalFlashMapHighscoreResponse.Merge(m, src)
}
func (m *GetGlobalFlashMapHighscoreResponse) XXX_Size() int {
	return xxx_messageInfo_GetGlobalFlashMapHighscoreResponse.Size(m)
}
func (m *GetGlobalFlashMapHighscoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGlobalFlashMapHighscoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGlobalFlashMapHighscoreResponse proto.InternalMessageInfo

func (m *GetGlobalFlashMapHighscoreResponse) GetHighscore() *FlashMapStatistic {
	if m != nil {
		return m.Highscore
	}
	return nil
}

type GetTopFlashMapHighscoresRequest struct {
	Limit                uint32   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	WithCheckpoints      bool     `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
	MapName              string   `protobuf:"bytes,3,opt,name=mapName,proto3" json:"mapName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopFlashMapHighscoresRequest) Reset()         { *m = GetTopFlashMapHighscoresRequest{} }
func (m *GetTopFlashMapHighscoresRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopFlashMapHighscoresRequest) ProtoMessage()    {}
func (*GetTopFlashMapHighscoresRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *GetTopFlashMapHighscoresRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopFlashMapHighscoresRequest.Unmarshal(m, b)
}
func (m *GetTopFlashMapHighscoresRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopFlashMapHighscoresRequest.Marshal(b, m, deterministic)
}
func (m *GetTopFlashMapHighscoresRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopFlashMapHighscoresRequest.Merge(m, src)
}
func (m *GetTopFlashMapHighscoresRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopFlashMapHighscoresRequest.Size(m)
}
func (m *GetTopFlashMapHighscoresRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopFlashMapHighscoresRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopFlashMapHighscoresRequest proto.InternalMessageInfo

func (m *GetTopFlashMapHighscoresRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTopFlashMapHighscoresRequest) GetWithCheckpoints() bool {
	if m != nil {
		return m.WithCheckpoints
	}
	return false
}

func (m *GetTopFlashMapHighscoresRequest) GetMapName() string {
	if m != nil {
		return m.MapName
	}
	return ""
}

type GetTopFlashMapHighscoresResponse struct {
	Highscores           []*FlashMapStatisticCollection `protobuf:"bytes,1,rep,name=highscores,proto3" json:"highscores,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *GetTopFlashMapHighscoresResponse) Reset()         { *m = GetTopFlashMapHighscoresResponse{} }
func (m *GetTopFlashMapHighscoresResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopFlashMapHighscoresResponse) ProtoMessage()    {}
func (*GetTopFlashMapHighscoresResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *GetTopFlashMapHighscoresResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopFlashMapHighscoresResponse.Unmarshal(m, b)
}
func (m *GetTopFlashMapHighscoresResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopFlashMapHighscoresResponse.Marshal(b, m, deterministic)
}
func (m *GetTopFlashMapHighscoresResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopFlashMapHighscoresResponse.Merge(m, src)
}
func (m *GetTopFlashMapHighscoresResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopFlashMapHighscoresResponse.Size(m)
}
func (m *GetTopFlashMapHighscoresResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopFlashMapHighscoresResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopFlashMapHighscoresResponse proto.InternalMessageInfo

func (m *GetTopFlashMapHighscoresResponse) GetHighscores() []*FlashMapStatisticCollection {
	if m != nil {
		return m.Highscores
	}
	return nil
}

type GetTopPlayersByPointsRequest struct {
	WithMapStats         bool     `protobuf:"varint,1,opt,name=withMapStats,proto3" json:"withMapStats,omitempty"`
	WithCheckpoints      bool     `protobuf:"varint,2,opt,name=withCheckpoints,proto3" json:"withCheckpoints,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopPlayersByPointsRequest) Reset()         { *m = GetTopPlayersByPointsRequest{} }
func (m *GetTopPlayersByPointsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTopPlayersByPointsRequest) ProtoMessage()    {}
func (*GetTopPlayersByPointsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *GetTopPlayersByPointsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopPlayersByPointsRequest.Unmarshal(m, b)
}
func (m *GetTopPlayersByPointsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopPlayersByPointsRequest.Marshal(b, m, deterministic)
}
func (m *GetTopPlayersByPointsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopPlayersByPointsRequest.Merge(m, src)
}
func (m *GetTopPlayersByPointsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTopPlayersByPointsRequest.Size(m)
}
func (m *GetTopPlayersByPointsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopPlayersByPointsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopPlayersByPointsRequest proto.InternalMessageInfo

func (m *GetTopPlayersByPointsRequest) GetWithMapStats() bool {
	if m != nil {
		return m.WithMapStats
	}
	return false
}

func (m *GetTopPlayersByPointsRequest) GetWithCheckpoints() bool {
	if m != nil {
		return m.WithCheckpoints
	}
	return false
}

func (m *GetTopPlayersByPointsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetTopPlayersByPointsResponse struct {
	TopPlayers           []*FlashAllStatistics `protobuf:"bytes,1,rep,name=topPlayers,proto3" json:"topPlayers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetTopPlayersByPointsResponse) Reset()         { *m = GetTopPlayersByPointsResponse{} }
func (m *GetTopPlayersByPointsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTopPlayersByPointsResponse) ProtoMessage()    {}
func (*GetTopPlayersByPointsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *GetTopPlayersByPointsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopPlayersByPointsResponse.Unmarshal(m, b)
}
func (m *GetTopPlayersByPointsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopPlayersByPointsResponse.Marshal(b, m, deterministic)
}
func (m *GetTopPlayersByPointsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopPlayersByPointsResponse.Merge(m, src)
}
func (m *GetTopPlayersByPointsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTopPlayersByPointsResponse.Size(m)
}
func (m *GetTopPlayersByPointsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopPlayersByPointsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopPlayersByPointsResponse proto.InternalMessageInfo

func (m *GetTopPlayersByPointsResponse) GetTopPlayers() []*FlashAllStatistics {
	if m != nil {
		return m.TopPlayers
	}
	return nil
}

type UpdateFlashStatisticsRequests struct {
	Stats                *FlashAllStatistics `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateFlashStatisticsRequests) Reset()         { *m = UpdateFlashStatisticsRequests{} }
func (m *UpdateFlashStatisticsRequests) String() string { return proto.CompactTextString(m) }
func (*UpdateFlashStatisticsRequests) ProtoMessage()    {}
func (*UpdateFlashStatisticsRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *UpdateFlashStatisticsRequests) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFlashStatisticsRequests.Unmarshal(m, b)
}
func (m *UpdateFlashStatisticsRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFlashStatisticsRequests.Marshal(b, m, deterministic)
}
func (m *UpdateFlashStatisticsRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFlashStatisticsRequests.Merge(m, src)
}
func (m *UpdateFlashStatisticsRequests) XXX_Size() int {
	return xxx_messageInfo_UpdateFlashStatisticsRequests.Size(m)
}
func (m *UpdateFlashStatisticsRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFlashStatisticsRequests.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFlashStatisticsRequests proto.InternalMessageInfo

func (m *UpdateFlashStatisticsRequests) GetStats() *FlashAllStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

type UpdateFlashStatisticsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFlashStatisticsResponse) Reset()         { *m = UpdateFlashStatisticsResponse{} }
func (m *UpdateFlashStatisticsResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateFlashStatisticsResponse) ProtoMessage()    {}
func (*UpdateFlashStatisticsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *UpdateFlashStatisticsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFlashStatisticsResponse.Unmarshal(m, b)
}
func (m *UpdateFlashStatisticsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFlashStatisticsResponse.Marshal(b, m, deterministic)
}
func (m *UpdateFlashStatisticsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFlashStatisticsResponse.Merge(m, src)
}
func (m *UpdateFlashStatisticsResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateFlashStatisticsResponse.Size(m)
}
func (m *UpdateFlashStatisticsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFlashStatisticsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFlashStatisticsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetFlashMapHighscoreForPlayerRequest)(nil), "stats.v1.GetFlashMapHighscoreForPlayerRequest")
	proto.RegisterType((*GetFlashMapHighscoreForPlayerResponse)(nil), "stats.v1.GetFlashMapHighscoreForPlayerResponse")
	proto.RegisterType((*GetGlobalFlashMapHighscoreRequest)(nil), "stats.v1.GetGlobalFlashMapHighscoreRequest")
	proto.RegisterType((*GetGlobalFlashMapHighscoreResponse)(nil), "stats.v1.GetGlobalFlashMapHighscoreResponse")
	proto.RegisterType((*GetTopFlashMapHighscoresRequest)(nil), "stats.v1.GetTopFlashMapHighscoresRequest")
	proto.RegisterType((*GetTopFlashMapHighscoresResponse)(nil), "stats.v1.GetTopFlashMapHighscoresResponse")
	proto.RegisterType((*GetTopPlayersByPointsRequest)(nil), "stats.v1.GetTopPlayersByPointsRequest")
	proto.RegisterType((*GetTopPlayersByPointsResponse)(nil), "stats.v1.GetTopPlayersByPointsResponse")
	proto.RegisterType((*UpdateFlashStatisticsRequests)(nil), "stats.v1.UpdateFlashStatisticsRequests")
	proto.RegisterType((*UpdateFlashStatisticsResponse)(nil), "stats.v1.UpdateFlashStatisticsResponse")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0x75, 0x83, 0xee, 0xed, 0x26, 0x24, 0x0b, 0xa4, 0x28, 0x6c, 0x5a, 0xb1, 0x18,
	0x94, 0x3f, 0x0a, 0x5a, 0x39, 0x21, 0x71, 0x61, 0x13, 0x2b, 0x1c, 0x40, 0x53, 0x0a, 0x17, 0x24,
	0x84, 0xdc, 0xcc, 0x2c, 0x16, 0x6e, 0x9d, 0xc5, 0x2f, 0x9d, 0x26, 0x2e, 0x88, 0x03, 0x5f, 0x91,
	0xaf, 0x83, 0x66, 0xc7, 0x59, 0xda, 0x35, 0x69, 0x10, 0x3b, 0xda, 0x79, 0xfa, 0x3c, 0xcf, 0xcf,
	0x7d, 0x9d, 0xc0, 0xa6, 0xe6, 0xd9, 0x54, 0xc4, 0x3c, 0x4c, 0x33, 0x85, 0x8a, 0xb4, 0x35, 0x32,
	0xd4, 0xe1, 0x74, 0x2f, 0xe8, 0x7c, 0x95, 0x4c, 0x27, 0x76, 0x9b, 0xfe, 0xf6, 0xe0, 0xfe, 0x80,
	0xe3, 0xe1, 0xc5, 0xd6, 0x3b, 0x96, 0xbe, 0x11, 0x27, 0x89, 0x8e, 0x55, 0xc6, 0x0f, 0x55, 0x76,
	0x24, 0xd9, 0x39, 0xcf, 0x22, 0x7e, 0xfa, 0x9d, 0x6b, 0x24, 0x01, 0xb4, 0x53, 0xb3, 0xf1, 0xf6,
	0xd8, 0xf7, 0xba, 0x5e, 0x6f, 0x3d, 0x2a, 0xd6, 0xa4, 0x07, 0xb7, 0xce, 0x04, 0x26, 0x07, 0x09,
	0x8f, 0xbf, 0xa5, 0x4a, 0x4c, 0x50, 0xfb, 0x2b, 0x5d, 0xaf, 0xd7, 0x8e, 0xe6, 0xb7, 0x89, 0x0f,
	0x37, 0xc7, 0x2c, 0x7d, 0xcf, 0xc6, 0xdc, 0x6f, 0x19, 0x13, 0xb7, 0xa4, 0x23, 0xd8, 0x5d, 0xd2,
	0x43, 0xa7, 0x6a, 0xa2, 0x39, 0x79, 0x01, 0xeb, 0x89, 0x7b, 0x6a, 0x9a, 0x74, 0xfa, 0x77, 0x43,
	0x07, 0x17, 0x3a, 0x83, 0x21, 0x32, 0x14, 0x1a, 0x45, 0x1c, 0x5d, 0xaa, 0xe9, 0x09, 0xdc, 0x1b,
	0x70, 0x1c, 0x48, 0x35, 0x62, 0xf2, 0x4a, 0x92, 0x03, 0x2d, 0x55, 0xf4, 0x66, 0x2a, 0x36, 0xc7,
	0xa4, 0x5f, 0x80, 0xd6, 0x05, 0xfd, 0x3f, 0xc9, 0x0f, 0xd8, 0x19, 0x70, 0xfc, 0xa0, 0xd2, 0x2b,
	0xee, 0xda, 0x71, 0xdc, 0x86, 0x35, 0x29, 0xc6, 0x02, 0x8d, 0xf3, 0x66, 0x64, 0x17, 0xd7, 0xf2,
	0x57, 0x09, 0xe8, 0x56, 0x87, 0xe7, 0x6c, 0xaf, 0x01, 0x8a, 0xb6, 0xda, 0xf7, 0xba, 0xad, 0x5e,
	0xa7, 0xbf, 0x5b, 0x03, 0x77, 0xa0, 0xa4, 0xe4, 0x31, 0x0a, 0x35, 0x89, 0x4a, 0x3f, 0xa4, 0xbf,
	0x3c, 0xd8, 0xb2, 0x59, 0x76, 0x0a, 0xf4, 0xfe, 0xf9, 0x91, 0xa9, 0xe7, 0x28, 0x29, 0x6c, 0x5c,
	0x14, 0xcf, 0xad, 0xb4, 0x81, 0x6d, 0x47, 0x33, 0x7b, 0xff, 0xc0, 0x5c, 0x9c, 0x59, 0xab, 0x74,
	0x66, 0xf4, 0x33, 0x6c, 0x57, 0x74, 0xc8, 0x61, 0x5f, 0x02, 0x60, 0xf1, 0x34, 0x87, 0xdd, 0x9a,
	0x83, 0x7d, 0x25, 0x65, 0x01, 0xab, 0xa3, 0x92, 0x9e, 0x0e, 0x61, 0xfb, 0x63, 0x7a, 0xcc, 0x90,
	0x1b, 0x5d, 0x49, 0x64, 0x11, 0x35, 0xe9, 0xc3, 0x9a, 0xf1, 0x32, 0xad, 0x97, 0x39, 0x5b, 0x29,
	0xdd, 0xa9, 0x34, 0xb5, 0x9d, 0xfb, 0x7f, 0x56, 0x61, 0xc3, 0x1c, 0xcf, 0xd0, 0xbe, 0x26, 0xc8,
	0x4f, 0xcf, 0x60, 0x56, 0xdf, 0x40, 0x12, 0x5e, 0x06, 0x37, 0x79, 0x65, 0x04, 0xcf, 0x1a, 0xeb,
	0xf3, 0x73, 0x3c, 0x83, 0xa0, 0xfa, 0xda, 0x90, 0x27, 0x33, 0x76, 0xf5, 0xb7, 0x38, 0x78, 0xda,
	0x4c, 0x9c, 0x07, 0x9f, 0x82, 0x5f, 0x35, 0xd1, 0xe4, 0xd1, 0x8c, 0x53, 0xdd, 0x95, 0x0b, 0x1e,
	0x37, 0x91, 0xe6, 0x91, 0x63, 0xc3, 0xea, 0x34, 0x73, 0x93, 0x45, 0x1e, 0xcc, 0x3b, 0x2d, 0x1e,
	0xff, 0xe0, 0xe1, 0x52, 0x5d, 0x1e, 0x27, 0xe0, 0xce, 0xc2, 0x79, 0x20, 0x25, 0x87, 0xda, 0x29,
	0x0c, 0x96, 0x0b, 0x6d, 0xd4, 0xfe, 0xea, 0xa7, 0x95, 0xe9, 0xde, 0xe8, 0x86, 0xf9, 0xbe, 0x3c,
	0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x27, 0x81, 0x62, 0xf0, 0x87, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	GetFlashMapHighscoreForPlayer(ctx context.Context, in *GetFlashMapHighscoreForPlayerRequest, opts ...grpc.CallOption) (*GetFlashMapHighscoreForPlayerResponse, error)
	GetGlobalFlashMapHighscore(ctx context.Context, in *GetGlobalFlashMapHighscoreRequest, opts ...grpc.CallOption) (*GetGlobalFlashMapHighscoreResponse, error)
	GetTopFlashMapHighscores(ctx context.Context, in *GetTopFlashMapHighscoresRequest, opts ...grpc.CallOption) (*GetTopFlashMapHighscoresResponse, error)
	GetTopFlashPlayersByPoints(ctx context.Context, in *GetTopPlayersByPointsRequest, opts ...grpc.CallOption) (*GetTopPlayersByPointsResponse, error)
	UpdateFlashStatistics(ctx context.Context, in *UpdateFlashStatisticsRequests, opts ...grpc.CallOption) (*UpdateFlashStatisticsResponse, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) GetFlashMapHighscoreForPlayer(ctx context.Context, in *GetFlashMapHighscoreForPlayerRequest, opts ...grpc.CallOption) (*GetFlashMapHighscoreForPlayerResponse, error) {
	out := new(GetFlashMapHighscoreForPlayerResponse)
	err := c.cc.Invoke(ctx, "/stats.v1.StatsService/GetFlashMapHighscoreForPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetGlobalFlashMapHighscore(ctx context.Context, in *GetGlobalFlashMapHighscoreRequest, opts ...grpc.CallOption) (*GetGlobalFlashMapHighscoreResponse, error) {
	out := new(GetGlobalFlashMapHighscoreResponse)
	err := c.cc.Invoke(ctx, "/stats.v1.StatsService/GetGlobalFlashMapHighscore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetTopFlashMapHighscores(ctx context.Context, in *GetTopFlashMapHighscoresRequest, opts ...grpc.CallOption) (*GetTopFlashMapHighscoresResponse, error) {
	out := new(GetTopFlashMapHighscoresResponse)
	err := c.cc.Invoke(ctx, "/stats.v1.StatsService/GetTopFlashMapHighscores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) GetTopFlashPlayersByPoints(ctx context.Context, in *GetTopPlayersByPointsRequest, opts ...grpc.CallOption) (*GetTopPlayersByPointsResponse, error) {
	out := new(GetTopPlayersByPointsResponse)
	err := c.cc.Invoke(ctx, "/stats.v1.StatsService/GetTopFlashPlayersByPoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) UpdateFlashStatistics(ctx context.Context, in *UpdateFlashStatisticsRequests, opts ...grpc.CallOption) (*UpdateFlashStatisticsResponse, error) {
	out := new(UpdateFlashStatisticsResponse)
	err := c.cc.Invoke(ctx, "/stats.v1.StatsService/UpdateFlashStatistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	GetFlashMapHighscoreForPlayer(context.Context, *GetFlashMapHighscoreForPlayerRequest) (*GetFlashMapHighscoreForPlayerResponse, error)
	GetGlobalFlashMapHighscore(context.Context, *GetGlobalFlashMapHighscoreRequest) (*GetGlobalFlashMapHighscoreResponse, error)
	GetTopFlashMapHighscores(context.Context, *GetTopFlashMapHighscoresRequest) (*GetTopFlashMapHighscoresResponse, error)
	GetTopFlashPlayersByPoints(context.Context, *GetTopPlayersByPointsRequest) (*GetTopPlayersByPointsResponse, error)
	UpdateFlashStatistics(context.Context, *UpdateFlashStatisticsRequests) (*UpdateFlashStatisticsResponse, error)
}

// UnimplementedStatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (*UnimplementedStatsServiceServer) GetFlashMapHighscoreForPlayer(ctx context.Context, req *GetFlashMapHighscoreForPlayerRequest) (*GetFlashMapHighscoreForPlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFlashMapHighscoreForPlayer not implemented")
}
func (*UnimplementedStatsServiceServer) GetGlobalFlashMapHighscore(ctx context.Context, req *GetGlobalFlashMapHighscoreRequest) (*GetGlobalFlashMapHighscoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalFlashMapHighscore not implemented")
}
func (*UnimplementedStatsServiceServer) GetTopFlashMapHighscores(ctx context.Context, req *GetTopFlashMapHighscoresRequest) (*GetTopFlashMapHighscoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopFlashMapHighscores not implemented")
}
func (*UnimplementedStatsServiceServer) GetTopFlashPlayersByPoints(ctx context.Context, req *GetTopPlayersByPointsRequest) (*GetTopPlayersByPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopFlashPlayersByPoints not implemented")
}
func (*UnimplementedStatsServiceServer) UpdateFlashStatistics(ctx context.Context, req *UpdateFlashStatisticsRequests) (*UpdateFlashStatisticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFlashStatistics not implemented")
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_GetFlashMapHighscoreForPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFlashMapHighscoreForPlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetFlashMapHighscoreForPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.StatsService/GetFlashMapHighscoreForPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetFlashMapHighscoreForPlayer(ctx, req.(*GetFlashMapHighscoreForPlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetGlobalFlashMapHighscore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGlobalFlashMapHighscoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetGlobalFlashMapHighscore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.StatsService/GetGlobalFlashMapHighscore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetGlobalFlashMapHighscore(ctx, req.(*GetGlobalFlashMapHighscoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetTopFlashMapHighscores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopFlashMapHighscoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetTopFlashMapHighscores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.StatsService/GetTopFlashMapHighscores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetTopFlashMapHighscores(ctx, req.(*GetTopFlashMapHighscoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_GetTopFlashPlayersByPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopPlayersByPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).GetTopFlashPlayersByPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.StatsService/GetTopFlashPlayersByPoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).GetTopFlashPlayersByPoints(ctx, req.(*GetTopPlayersByPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_UpdateFlashStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFlashStatisticsRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).UpdateFlashStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.StatsService/UpdateFlashStatistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).UpdateFlashStatistics(ctx, req.(*UpdateFlashStatisticsRequests))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stats.v1.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFlashMapHighscoreForPlayer",
			Handler:    _StatsService_GetFlashMapHighscoreForPlayer_Handler,
		},
		{
			MethodName: "GetGlobalFlashMapHighscore",
			Handler:    _StatsService_GetGlobalFlashMapHighscore_Handler,
		},
		{
			MethodName: "GetTopFlashMapHighscores",
			Handler:    _StatsService_GetTopFlashMapHighscores_Handler,
		},
		{
			MethodName: "GetTopFlashPlayersByPoints",
			Handler:    _StatsService_GetTopFlashPlayersByPoints_Handler,
		},
		{
			MethodName: "UpdateFlashStatistics",
			Handler:    _StatsService_UpdateFlashStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}