// Code generated by protoc-gen-go. DO NOT EDIT.
// source: world_message.proto

package world_message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type MWU_WorldLogon struct {
	WorldId              uint32   `protobuf:"varint,1,opt,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	WorldName            string   `protobuf:"bytes,2,opt,name=world_name,json=worldName,proto3" json:"world_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MWU_WorldLogon) Reset()         { *m = MWU_WorldLogon{} }
func (m *MWU_WorldLogon) String() string { return proto.CompactTextString(m) }
func (*MWU_WorldLogon) ProtoMessage()    {}
func (*MWU_WorldLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{0}
}

func (m *MWU_WorldLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MWU_WorldLogon.Unmarshal(m, b)
}
func (m *MWU_WorldLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MWU_WorldLogon.Marshal(b, m, deterministic)
}
func (m *MWU_WorldLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MWU_WorldLogon.Merge(m, src)
}
func (m *MWU_WorldLogon) XXX_Size() int {
	return xxx_messageInfo_MWU_WorldLogon.Size(m)
}
func (m *MWU_WorldLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MWU_WorldLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MWU_WorldLogon proto.InternalMessageInfo

func (m *MWU_WorldLogon) GetWorldId() uint32 {
	if m != nil {
		return m.WorldId
	}
	return 0
}

func (m *MWU_WorldLogon) GetWorldName() string {
	if m != nil {
		return m.WorldName
	}
	return ""
}

type MUW_WorldLogon struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MUW_WorldLogon) Reset()         { *m = MUW_WorldLogon{} }
func (m *MUW_WorldLogon) String() string { return proto.CompactTextString(m) }
func (*MUW_WorldLogon) ProtoMessage()    {}
func (*MUW_WorldLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{1}
}

func (m *MUW_WorldLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MUW_WorldLogon.Unmarshal(m, b)
}
func (m *MUW_WorldLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MUW_WorldLogon.Marshal(b, m, deterministic)
}
func (m *MUW_WorldLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MUW_WorldLogon.Merge(m, src)
}
func (m *MUW_WorldLogon) XXX_Size() int {
	return xxx_messageInfo_MUW_WorldLogon.Size(m)
}
func (m *MUW_WorldLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MUW_WorldLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MUW_WorldLogon proto.InternalMessageInfo

type MWU_TestConnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MWU_TestConnect) Reset()         { *m = MWU_TestConnect{} }
func (m *MWU_TestConnect) String() string { return proto.CompactTextString(m) }
func (*MWU_TestConnect) ProtoMessage()    {}
func (*MWU_TestConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{2}
}

func (m *MWU_TestConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MWU_TestConnect.Unmarshal(m, b)
}
func (m *MWU_TestConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MWU_TestConnect.Marshal(b, m, deterministic)
}
func (m *MWU_TestConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MWU_TestConnect.Merge(m, src)
}
func (m *MWU_TestConnect) XXX_Size() int {
	return xxx_messageInfo_MWU_TestConnect.Size(m)
}
func (m *MWU_TestConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_MWU_TestConnect.DiscardUnknown(m)
}

var xxx_messageInfo_MWU_TestConnect proto.InternalMessageInfo

type MUW_TestConnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MUW_TestConnect) Reset()         { *m = MUW_TestConnect{} }
func (m *MUW_TestConnect) String() string { return proto.CompactTextString(m) }
func (*MUW_TestConnect) ProtoMessage()    {}
func (*MUW_TestConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{3}
}

func (m *MUW_TestConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MUW_TestConnect.Unmarshal(m, b)
}
func (m *MUW_TestConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MUW_TestConnect.Marshal(b, m, deterministic)
}
func (m *MUW_TestConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MUW_TestConnect.Merge(m, src)
}
func (m *MUW_TestConnect) XXX_Size() int {
	return xxx_messageInfo_MUW_TestConnect.Size(m)
}
func (m *MUW_TestConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_MUW_TestConnect.DiscardUnknown(m)
}

var xxx_messageInfo_MUW_TestConnect proto.InternalMessageInfo

type MWU_HeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MWU_HeartBeat) Reset()         { *m = MWU_HeartBeat{} }
func (m *MWU_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MWU_HeartBeat) ProtoMessage()    {}
func (*MWU_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{4}
}

func (m *MWU_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MWU_HeartBeat.Unmarshal(m, b)
}
func (m *MWU_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MWU_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MWU_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MWU_HeartBeat.Merge(m, src)
}
func (m *MWU_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MWU_HeartBeat.Size(m)
}
func (m *MWU_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MWU_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MWU_HeartBeat proto.InternalMessageInfo

type MUW_HeartBeat struct {
	BattleTime           uint32   `protobuf:"varint,1,opt,name=battle_time,json=battleTime,proto3" json:"battle_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MUW_HeartBeat) Reset()         { *m = MUW_HeartBeat{} }
func (m *MUW_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MUW_HeartBeat) ProtoMessage()    {}
func (*MUW_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_33865c398800497b, []int{5}
}

func (m *MUW_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MUW_HeartBeat.Unmarshal(m, b)
}
func (m *MUW_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MUW_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MUW_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MUW_HeartBeat.Merge(m, src)
}
func (m *MUW_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MUW_HeartBeat.Size(m)
}
func (m *MUW_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MUW_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MUW_HeartBeat proto.InternalMessageInfo

func (m *MUW_HeartBeat) GetBattleTime() uint32 {
	if m != nil {
		return m.BattleTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MWU_WorldLogon)(nil), "world_message.MWU_WorldLogon")
	proto.RegisterType((*MUW_WorldLogon)(nil), "world_message.MUW_WorldLogon")
	proto.RegisterType((*MWU_TestConnect)(nil), "world_message.MWU_TestConnect")
	proto.RegisterType((*MUW_TestConnect)(nil), "world_message.MUW_TestConnect")
	proto.RegisterType((*MWU_HeartBeat)(nil), "world_message.MWU_HeartBeat")
	proto.RegisterType((*MUW_HeartBeat)(nil), "world_message.MUW_HeartBeat")
}

func init() { proto.RegisterFile("world_message.proto", fileDescriptor_33865c398800497b) }

var fileDescriptor_33865c398800497b = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xcf, 0x2f, 0xca,
	0x49, 0x89, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x45, 0x11, 0x54, 0xf2, 0xe2, 0xe2, 0xf3, 0x0d, 0x0f, 0x8d, 0x0f, 0x07, 0x09, 0xfa, 0xe4,
	0xa7, 0xe7, 0xe7, 0x09, 0x49, 0x72, 0x71, 0x40, 0x94, 0x64, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x06, 0xb1, 0x83, 0xf9, 0x9e, 0x29, 0x42, 0xb2, 0x5c, 0x5c, 0x10, 0xa9, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x4e, 0xb0, 0x88, 0x5f, 0x62, 0x6e, 0xaa, 0x92,
	0x00, 0x17, 0x9f, 0x6f, 0x68, 0x38, 0x92, 0x59, 0x4a, 0x82, 0x5c, 0xfc, 0x20, 0xd3, 0x43, 0x52,
	0x8b, 0x4b, 0x9c, 0xf3, 0xf3, 0xf2, 0x52, 0x93, 0x4b, 0xc0, 0x42, 0xa1, 0xe1, 0x28, 0x42, 0xfc,
	0x5c, 0xbc, 0x20, 0x55, 0x1e, 0xa9, 0x89, 0x45, 0x25, 0x4e, 0xa9, 0x89, 0x25, 0x4a, 0x06, 0x5c,
	0xbc, 0x20, 0x35, 0x70, 0x01, 0x21, 0x79, 0x2e, 0xee, 0xa4, 0xc4, 0x92, 0x92, 0x9c, 0xd4, 0xf8,
	0x92, 0xcc, 0xdc, 0x54, 0xa8, 0xb3, 0xb8, 0x20, 0x42, 0x21, 0x99, 0xb9, 0xa9, 0x49, 0x6c, 0x60,
	0xcf, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x3c, 0x9e, 0x95, 0xf3, 0x00, 0x00, 0x00,
}