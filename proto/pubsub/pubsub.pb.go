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

func init() {
	proto.RegisterType((*PublishBroadCast)(nil), "ultimate.service.pubsub.PublishBroadCast")
	proto.RegisterType((*PublishSendWorldMessage)(nil), "ultimate.service.pubsub.PublishSendWorldMessage")
	proto.RegisterType((*PublishMatching)(nil), "ultimate.service.pubsub.PublishMatching")
	proto.RegisterType((*PublishAddRecord)(nil), "ultimate.service.pubsub.PublishAddRecord")
}

func init() { proto.RegisterFile("pubsub/pubsub.proto", fileDescriptor_ce310d0bb9f289ed) }

var fileDescriptor_ce310d0bb9f289ed = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xfa, 0xe5, 0xab, 0xae, 0xbf, 0xe3, 0xa1, 0xd5, 0x53, 0x8d, 0x97, 0x22, 0x92,
	0x80, 0x1e, 0x3c, 0xb7, 0x7a, 0x10, 0xa4, 0x22, 0x11, 0x11, 0xbc, 0x94, 0x49, 0x66, 0xd8, 0x2c,
	0xec, 0x66, 0xcb, 0xfe, 0xf0, 0xef, 0x97, 0xec, 0x46, 0xb1, 0x88, 0x97, 0x59, 0x96, 0x37, 0xef,
	0xf3, 0x98, 0xc7, 0x4e, 0xd6, 0xbe, 0xb6, 0xbe, 0x2e, 0xe3, 0x53, 0xac, 0x8d, 0x76, 0x3a, 0x1b,
	0x7b, 0xe9, 0x84, 0x02, 0x47, 0x85, 0x25, 0xf3, 0x21, 0x1a, 0x2a, 0xa2, 0x7c, 0x76, 0x0c, 0x86,
	0x3a, 0x28, 0xc3, 0x8c, 0xbb, 0xf9, 0x03, 0x3b, 0x7a, 0xf6, 0xb5, 0x14, 0xb6, 0x5d, 0x18, 0x0d,
	0x78, 0x07, 0xd6, 0x65, 0xa7, 0x6c, 0x5b, 0x59, 0xbe, 0xea, 0x40, 0xd1, 0x24, 0x99, 0x26, 0xb3,
	0x9d, 0x6a, 0x4b, 0x59, 0xfe, 0x04, 0x8a, 0xbe, 0x24, 0x04, 0x07, 0x93, 0x74, 0x9a, 0xcc, 0xf6,
	0x82, 0x74, 0x0f, 0x0e, 0xf2, 0x15, 0x1b, 0x0f, 0xa4, 0x17, 0xea, 0xf0, 0x4d, 0x1b, 0x89, 0x4b,
	0xb2, 0x16, 0x38, 0x65, 0x07, 0x2c, 0x15, 0x18, 0x50, 0xfb, 0x55, 0x2a, 0x70, 0x23, 0x20, 0xfd,
	0x3b, 0x60, 0xb4, 0x19, 0x70, 0xce, 0x0e, 0x87, 0x80, 0x25, 0xb8, 0xa6, 0x15, 0x1d, 0xff, 0x01,
	0x1e, 0xf5, 0xe0, 0xfc, 0xf1, 0xfb, 0x9a, 0x39, 0x62, 0x45, 0x8d, 0x36, 0x98, 0xdd, 0xb2, 0x7f,
	0x81, 0xd6, 0x6f, 0xed, 0x5e, 0x5f, 0x14, 0xbf, 0xca, 0x89, 0x75, 0xcc, 0xfb, 0x19, 0x2d, 0x55,
	0x30, 0x2c, 0xae, 0xde, 0x2f, 0xb9, 0x70, 0xad, 0xaf, 0x8b, 0x46, 0xab, 0xb2, 0x25, 0x29, 0x35,
	0x7a, 0xf4, 0xe5, 0xeb, 0x00, 0x28, 0x43, 0x83, 0x43, 0xf5, 0xf5, 0xff, 0xf0, 0xbb, 0xf9, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0x73, 0xc9, 0x33, 0xca, 0x92, 0x01, 0x00, 0x00,
}
