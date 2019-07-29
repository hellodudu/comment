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
	proto.RegisterType((*PublishBroadCast)(nil), "ultimate.service.pubsub.PublishBroadCast")
	proto.RegisterType((*PublishSendWorldMessage)(nil), "ultimate.service.pubsub.PublishSendWorldMessage")
	proto.RegisterType((*PublishMatching)(nil), "ultimate.service.pubsub.PublishMatching")
	proto.RegisterType((*PublishAddRecord)(nil), "ultimate.service.pubsub.PublishAddRecord")
	proto.RegisterType((*PublishBattleResult)(nil), "ultimate.service.pubsub.PublishBattleResult")
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor_ce310d0bb9f289ed) }

var fileDescriptor_ce310d0bb9f289ed = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x69, 0xfa, 0x78, 0xaf, 0xdd, 0xf7, 0x9e, 0x3f, 0xd2, 0x43, 0xab, 0x22, 0xd4, 0x78,
	0x29, 0x22, 0x09, 0xe8, 0xc1, 0x73, 0xab, 0x07, 0x8b, 0x54, 0x24, 0x22, 0x05, 0x2f, 0x65, 0x92,
	0x1d, 0x92, 0xc5, 0x64, 0xb7, 0xec, 0xce, 0xea, 0xbf, 0x2f, 0xd9, 0x4d, 0x8b, 0x45, 0xbc, 0x4c,
	0xc8, 0x7c, 0x67, 0x3e, 0xdf, 0x9d, 0x19, 0x36, 0x58, 0xdb, 0xcc, 0xd8, 0x2c, 0xf1, 0x9f, 0x78,
	0xad, 0x15, 0xa9, 0x70, 0x68, 0x2b, 0x12, 0x35, 0x10, 0xc6, 0x06, 0xf5, 0xbb, 0xc8, 0x31, 0xf6,
	0xf2, 0xf1, 0x21, 0x68, 0x94, 0x90, 0xb8, 0xe8, 0x6b, 0xa3, 0x7b, 0x76, 0xf0, 0x64, 0xb3, 0x4a,
	0x98, 0x72, 0xa6, 0x15, 0xf0, 0x5b, 0x30, 0x14, 0x1e, 0xb1, 0x5e, 0x6d, 0x8a, 0x95, 0x84, 0x1a,
	0x47, 0x9d, 0x71, 0x67, 0xd2, 0x4f, 0xff, 0xd4, 0xa6, 0x78, 0x84, 0x1a, 0x37, 0x12, 0x07, 0x82,
	0x51, 0x30, 0xee, 0x4c, 0xfe, 0x39, 0xe9, 0x0e, 0x08, 0xa2, 0x15, 0x1b, 0xb6, 0xa4, 0x67, 0x94,
	0x7c, 0xa9, 0x74, 0xc5, 0x17, 0x68, 0x0c, 0x14, 0x18, 0xee, 0xb1, 0x40, 0x70, 0x87, 0xfa, 0x9f,
	0x06, 0x82, 0xef, 0x18, 0x04, 0x3f, 0x1b, 0x74, 0x77, 0x0d, 0xce, 0xd8, 0x7e, 0x6b, 0xb0, 0x00,
	0xca, 0x4b, 0x21, 0x8b, 0x2f, 0xe0, 0x6e, 0x03, 0x8e, 0x1e, 0xb6, 0xd3, 0x4c, 0x39, 0x4f, 0x31,
	0x57, 0x9a, 0x87, 0x37, 0xec, 0x97, 0xa3, 0x35, 0x55, 0x7f, 0xaf, 0xce, 0xe3, 0x6f, 0xcb, 0xf1,
	0xeb, 0x98, 0x36, 0xd1, 0xb7, 0xa4, 0xae, 0x21, 0x92, 0x6c, 0xb0, 0x59, 0x0d, 0x10, 0x55, 0x98,
	0xa2, 0xb1, 0x15, 0x85, 0x27, 0xac, 0x0f, 0x44, 0x90, 0xbf, 0xad, 0xb6, 0xd6, 0x3d, 0x9f, 0x98,
	0xf3, 0x46, 0x24, 0xd0, 0x05, 0x52, 0x23, 0x06, 0x5e, 0xf4, 0x89, 0x39, 0x0f, 0x4f, 0x19, 0x6b,
	0x3b, 0x3f, 0x84, 0x74, 0xd3, 0xf5, 0xd2, 0x96, 0xb5, 0x14, 0x72, 0x76, 0xf9, 0x7a, 0x51, 0x08,
	0x2a, 0x6d, 0x16, 0xe7, 0xaa, 0x4e, 0x4a, 0xac, 0x2a, 0xc5, 0x2d, 0xb7, 0xc9, 0x4b, 0xfb, 0xe0,
	0xc4, 0x5d, 0xac, 0x3d, 0x75, 0xf6, 0xdb, 0xfd, 0x5d, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x81, 0x42, 0x95, 0x02, 0x02, 0x00, 0x00,
}
