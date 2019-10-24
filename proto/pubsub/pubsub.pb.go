// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub/pubsub.proto

package pubsub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	arena "github.com/hellodudu/Ultimate/proto/arena"
	math "math"
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

/////////////////////////////////////////////////
// pub/sub
/////////////////////////////////////////////////
type PublishBroadCast struct {
	MsgName              string   `protobuf:"bytes,1,opt,name=msg_name,json=msgName,proto3" json:"msg_name,omitempty"`
	MsgData              []byte   `protobuf:"bytes,2,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishBroadCast) Reset()         { *m = PublishBroadCast{} }
func (m *PublishBroadCast) String() string { return proto.CompactTextString(m) }
func (*PublishBroadCast) ProtoMessage()    {}
func (*PublishBroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{0}
}

func (m *PublishBroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishBroadCast.Unmarshal(m, b)
}
func (m *PublishBroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishBroadCast.Marshal(b, m, deterministic)
}
func (m *PublishBroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishBroadCast.Merge(m, src)
}
func (m *PublishBroadCast) XXX_Size() int {
	return xxx_messageInfo_PublishBroadCast.Size(m)
}
func (m *PublishBroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishBroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_PublishBroadCast proto.InternalMessageInfo

func (m *PublishBroadCast) GetMsgName() string {
	if m != nil {
		return m.MsgName
	}
	return ""
}

func (m *PublishBroadCast) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type PublishSendWorldMessage struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgName              string   `protobuf:"bytes,2,opt,name=msg_name,json=msgName,proto3" json:"msg_name,omitempty"`
	MsgData              []byte   `protobuf:"bytes,3,opt,name=msg_data,json=msgData,proto3" json:"msg_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishSendWorldMessage) Reset()         { *m = PublishSendWorldMessage{} }
func (m *PublishSendWorldMessage) String() string { return proto.CompactTextString(m) }
func (*PublishSendWorldMessage) ProtoMessage()    {}
func (*PublishSendWorldMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{1}
}

func (m *PublishSendWorldMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishSendWorldMessage.Unmarshal(m, b)
}
func (m *PublishSendWorldMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishSendWorldMessage.Marshal(b, m, deterministic)
}
func (m *PublishSendWorldMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishSendWorldMessage.Merge(m, src)
}
func (m *PublishSendWorldMessage) XXX_Size() int {
	return xxx_messageInfo_PublishSendWorldMessage.Size(m)
}
func (m *PublishSendWorldMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishSendWorldMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PublishSendWorldMessage proto.InternalMessageInfo

func (m *PublishSendWorldMessage) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PublishSendWorldMessage) GetMsgName() string {
	if m != nil {
		return m.MsgName
	}
	return ""
}

func (m *PublishSendWorldMessage) GetMsgData() []byte {
	if m != nil {
		return m.MsgData
	}
	return nil
}

type PublishMatching struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishMatching) Reset()         { *m = PublishMatching{} }
func (m *PublishMatching) String() string { return proto.CompactTextString(m) }
func (*PublishMatching) ProtoMessage()    {}
func (*PublishMatching) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{2}
}

func (m *PublishMatching) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishMatching.Unmarshal(m, b)
}
func (m *PublishMatching) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishMatching.Marshal(b, m, deterministic)
}
func (m *PublishMatching) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishMatching.Merge(m, src)
}
func (m *PublishMatching) XXX_Size() int {
	return xxx_messageInfo_PublishMatching.Size(m)
}
func (m *PublishMatching) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishMatching.DiscardUnknown(m)
}

var xxx_messageInfo_PublishMatching proto.InternalMessageInfo

func (m *PublishMatching) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PublishAddRecord struct {
	Data                 *arena.ArenaRecord `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PublishAddRecord) Reset()         { *m = PublishAddRecord{} }
func (m *PublishAddRecord) String() string { return proto.CompactTextString(m) }
func (*PublishAddRecord) ProtoMessage()    {}
func (*PublishAddRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{3}
}

func (m *PublishAddRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishAddRecord.Unmarshal(m, b)
}
func (m *PublishAddRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishAddRecord.Marshal(b, m, deterministic)
}
func (m *PublishAddRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishAddRecord.Merge(m, src)
}
func (m *PublishAddRecord) XXX_Size() int {
	return xxx_messageInfo_PublishAddRecord.Size(m)
}
func (m *PublishAddRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishAddRecord.DiscardUnknown(m)
}

var xxx_messageInfo_PublishAddRecord proto.InternalMessageInfo

func (m *PublishAddRecord) GetData() *arena.ArenaRecord {
	if m != nil {
		return m.Data
	}
	return nil
}

type PublishBattleResult struct {
	AttackId             int64    `protobuf:"varint,1,opt,name=attack_id,json=attackId,proto3" json:"attack_id,omitempty"`
	TargetId             int64    `protobuf:"varint,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	AttackWin            bool     `protobuf:"varint,3,opt,name=attack_win,json=attackWin,proto3" json:"attack_win,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishBattleResult) Reset()         { *m = PublishBattleResult{} }
func (m *PublishBattleResult) String() string { return proto.CompactTextString(m) }
func (*PublishBattleResult) ProtoMessage()    {}
func (*PublishBattleResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce310d0bb9f289ed, []int{4}
}

func (m *PublishBattleResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishBattleResult.Unmarshal(m, b)
}
func (m *PublishBattleResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishBattleResult.Marshal(b, m, deterministic)
}
func (m *PublishBattleResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishBattleResult.Merge(m, src)
}
func (m *PublishBattleResult) XXX_Size() int {
	return xxx_messageInfo_PublishBattleResult.Size(m)
}
func (m *PublishBattleResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishBattleResult.DiscardUnknown(m)
}

var xxx_messageInfo_PublishBattleResult proto.InternalMessageInfo

func (m *PublishBattleResult) GetAttackId() int64 {
	if m != nil {
		return m.AttackId
	}
	return 0
}

func (m *PublishBattleResult) GetTargetId() int64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *PublishBattleResult) GetAttackWin() bool {
	if m != nil {
		return m.AttackWin
	}
	return false
}

func init() {
	proto.RegisterType((*PublishBroadCast)(nil), "ultimate_service_pubsub.PublishBroadCast")
	proto.RegisterType((*PublishSendWorldMessage)(nil), "ultimate_service_pubsub.PublishSendWorldMessage")
	proto.RegisterType((*PublishMatching)(nil), "ultimate_service_pubsub.PublishMatching")
	proto.RegisterType((*PublishAddRecord)(nil), "ultimate_service_pubsub.PublishAddRecord")
	proto.RegisterType((*PublishBattleResult)(nil), "ultimate_service_pubsub.PublishBattleResult")
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor_ce310d0bb9f289ed) }

var fileDescriptor_ce310d0bb9f289ed = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x69, 0xfa, 0x78, 0xaf, 0xdd, 0xf7, 0x9e, 0x3f, 0xd2, 0x43, 0xab, 0x22, 0xd4, 0x78,
	0x29, 0x22, 0x0d, 0xe8, 0xc1, 0x73, 0xab, 0x07, 0x8b, 0x54, 0x24, 0x22, 0x05, 0x2f, 0x61, 0x92,
	0x1d, 0x92, 0xc5, 0x64, 0xb7, 0xec, 0xce, 0xea, 0xbf, 0x2f, 0xd9, 0x4d, 0x8b, 0x45, 0xbc, 0x4c,
	0xc8, 0x7c, 0x67, 0x3e, 0xdf, 0x9d, 0x19, 0x36, 0x58, 0xdb, 0xcc, 0xd8, 0x2c, 0xf6, 0x9f, 0xe9,
	0x5a, 0x2b, 0x52, 0xe1, 0xd0, 0x56, 0x24, 0x6a, 0x20, 0x4c, 0x0d, 0xea, 0x77, 0x91, 0x63, 0xea,
	0xe5, 0xe3, 0x43, 0xd0, 0x28, 0x21, 0x76, 0xd1, 0xd7, 0x46, 0xf7, 0xec, 0xe0, 0xc9, 0x66, 0x95,
	0x30, 0xe5, 0x5c, 0x2b, 0xe0, 0xb7, 0x60, 0x28, 0x3c, 0x62, 0xbd, 0xda, 0x14, 0xa9, 0x84, 0x1a,
	0x47, 0x9d, 0x71, 0x67, 0xd2, 0x4f, 0xfe, 0xd4, 0xa6, 0x78, 0x84, 0x1a, 0x37, 0x12, 0x07, 0x82,
	0x51, 0x30, 0xee, 0x4c, 0xfe, 0x39, 0xe9, 0x0e, 0x08, 0xa2, 0x94, 0x0d, 0x5b, 0xd2, 0x33, 0x4a,
	0xbe, 0x52, 0xba, 0xe2, 0x4b, 0x34, 0x06, 0x0a, 0x0c, 0xf7, 0x58, 0x20, 0xb8, 0x43, 0xfd, 0x4f,
	0x02, 0xc1, 0x77, 0x0c, 0x82, 0x9f, 0x0d, 0xba, 0xbb, 0x06, 0x67, 0x6c, 0xbf, 0x35, 0x58, 0x02,
	0xe5, 0xa5, 0x90, 0xc5, 0x17, 0x70, 0xb7, 0x01, 0x47, 0x0f, 0xdb, 0x69, 0x66, 0x9c, 0x27, 0x98,
	0x2b, 0xcd, 0xc3, 0x1b, 0xf6, 0xcb, 0xd1, 0x9a, 0xaa, 0xbf, 0x57, 0xe7, 0xd3, 0x6f, 0xcb, 0xf1,
	0xeb, 0x98, 0x35, 0xd1, 0xb7, 0x24, 0xae, 0x21, 0x92, 0x6c, 0xb0, 0x59, 0x0d, 0x10, 0x55, 0x98,
	0xa0, 0xb1, 0x15, 0x85, 0x27, 0xac, 0x0f, 0x44, 0x90, 0xbf, 0xa5, 0x5b, 0xeb, 0x9e, 0x4f, 0x2c,
	0x78, 0x23, 0x12, 0xe8, 0x02, 0xa9, 0x11, 0x03, 0x2f, 0xfa, 0xc4, 0x82, 0x87, 0xa7, 0x8c, 0xb5,
	0x9d, 0x1f, 0x42, 0xba, 0xe9, 0x7a, 0x49, 0xcb, 0x5a, 0x09, 0x39, 0xbf, 0x7c, 0xbd, 0x28, 0x04,
	0x95, 0x36, 0x9b, 0xe6, 0xaa, 0x8e, 0x4b, 0xac, 0x2a, 0xc5, 0x2d, 0xb7, 0xf1, 0x4b, 0xfb, 0xe0,
	0xd8, 0x5d, 0xac, 0x3d, 0x75, 0xf6, 0xdb, 0xfd, 0x5d, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x0e, 0x32, 0xdd, 0x02, 0x02, 0x00, 0x00,
}