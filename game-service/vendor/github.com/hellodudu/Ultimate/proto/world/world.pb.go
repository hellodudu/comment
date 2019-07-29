// Code generated by protoc-gen-go. DO NOT EDIT.
// source: world/world.proto

package world

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
	return fileDescriptor_2154e1c0b2e786c2, []int{0}
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
	return fileDescriptor_2154e1c0b2e786c2, []int{1}
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
	return fileDescriptor_2154e1c0b2e786c2, []int{2}
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
	return fileDescriptor_2154e1c0b2e786c2, []int{3}
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
	return fileDescriptor_2154e1c0b2e786c2, []int{4}
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
	return fileDescriptor_2154e1c0b2e786c2, []int{5}
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

type MWU_WorldConnected struct {
	WorldId              []uint32 `protobuf:"varint,1,rep,packed,name=world_id,json=worldId,proto3" json:"world_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MWU_WorldConnected) Reset()         { *m = MWU_WorldConnected{} }
func (m *MWU_WorldConnected) String() string { return proto.CompactTextString(m) }
func (*MWU_WorldConnected) ProtoMessage()    {}
func (*MWU_WorldConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_2154e1c0b2e786c2, []int{6}
}

func (m *MWU_WorldConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MWU_WorldConnected.Unmarshal(m, b)
}
func (m *MWU_WorldConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MWU_WorldConnected.Marshal(b, m, deterministic)
}
func (m *MWU_WorldConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MWU_WorldConnected.Merge(m, src)
}
func (m *MWU_WorldConnected) XXX_Size() int {
	return xxx_messageInfo_MWU_WorldConnected.Size(m)
}
func (m *MWU_WorldConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_MWU_WorldConnected.DiscardUnknown(m)
}

var xxx_messageInfo_MWU_WorldConnected proto.InternalMessageInfo

func (m *MWU_WorldConnected) GetWorldId() []uint32 {
	if m != nil {
		return m.WorldId
	}
	return nil
}

func init() {
	proto.RegisterType((*MWU_WorldLogon)(nil), "ultimate.service.world.MWU_WorldLogon")
	proto.RegisterType((*MUW_WorldLogon)(nil), "ultimate.service.world.MUW_WorldLogon")
	proto.RegisterType((*MWU_TestConnect)(nil), "ultimate.service.world.MWU_TestConnect")
	proto.RegisterType((*MUW_TestConnect)(nil), "ultimate.service.world.MUW_TestConnect")
	proto.RegisterType((*MWU_HeartBeat)(nil), "ultimate.service.world.MWU_HeartBeat")
	proto.RegisterType((*MUW_HeartBeat)(nil), "ultimate.service.world.MUW_HeartBeat")
	proto.RegisterType((*MWU_WorldConnected)(nil), "ultimate.service.world.MWU_WorldConnected")
}

func init() { proto.RegisterFile("world/world.proto", fileDescriptor_2154e1c0b2e786c2) }

var fileDescriptor_2154e1c0b2e786c2 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x51, 0x4b, 0xc3, 0x30,
	0x14, 0x85, 0xa9, 0x82, 0xba, 0x2b, 0xdd, 0x5c, 0x1e, 0x64, 0x3e, 0x88, 0x23, 0x4f, 0x13, 0xa1,
	0x11, 0xfc, 0x07, 0xf3, 0x45, 0xc5, 0xf9, 0x30, 0x56, 0x0a, 0xbe, 0x94, 0xb4, 0xb9, 0x6c, 0x81,
	0xa4, 0x91, 0xee, 0x56, 0xff, 0xbe, 0xe4, 0xb6, 0x8c, 0xe9, 0x4b, 0xe0, 0x7c, 0xe7, 0x72, 0xc8,
	0x39, 0x30, 0xfd, 0x09, 0xad, 0x33, 0x8a, 0xdf, 0xec, 0xab, 0x0d, 0x14, 0xc4, 0x75, 0xe7, 0xc8,
	0x7a, 0x4d, 0x98, 0xed, 0xb1, 0xfd, 0xb6, 0x35, 0x66, 0xec, 0xca, 0x37, 0x18, 0xaf, 0x8a, 0xbc,
	0x2c, 0xa2, 0x78, 0x0f, 0xdb, 0xd0, 0x88, 0x1b, 0xb8, 0x60, 0xab, 0xb4, 0x66, 0x96, 0xcc, 0x93,
	0x45, 0xba, 0x3e, 0x67, 0xfd, 0x6a, 0xc4, 0x2d, 0x40, 0x6f, 0x35, 0xda, 0xe3, 0xec, 0x64, 0x9e,
	0x2c, 0x46, 0xeb, 0x11, 0x93, 0x0f, 0xed, 0x51, 0x5e, 0xc1, 0x78, 0x95, 0x17, 0x47, 0x59, 0x72,
	0x0a, 0x93, 0x98, 0xbe, 0xc1, 0x3d, 0x3d, 0x87, 0xa6, 0xc1, 0x9a, 0x18, 0xe5, 0xc5, 0x1f, 0x34,
	0x81, 0x34, 0x5e, 0xbd, 0xa0, 0x6e, 0x69, 0x89, 0x9a, 0xe4, 0x23, 0xa4, 0xf1, 0xe6, 0x00, 0xc4,
	0x1d, 0x5c, 0x56, 0x9a, 0xc8, 0x61, 0x49, 0xd6, 0xe3, 0xf0, 0x2d, 0xe8, 0xd1, 0xc6, 0x7a, 0x94,
	0x0a, 0xc4, 0xa1, 0xc6, 0x10, 0x8b, 0xe6, 0x5f, 0x95, 0xd3, 0xa3, 0x2a, 0xcb, 0x87, 0xcf, 0xfb,
	0xad, 0xa5, 0x5d, 0x57, 0x65, 0x75, 0xf0, 0x6a, 0x87, 0xce, 0x05, 0xd3, 0x99, 0x4e, 0xe5, 0xc3,
	0x4c, 0x8a, 0x67, 0xeb, 0x27, 0xac, 0xce, 0x58, 0x3c, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x82,
	0x4a, 0x04, 0x5a, 0x58, 0x01, 0x00, 0x00,
}
