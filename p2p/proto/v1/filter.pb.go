// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter.proto

package qitmeer_p2p_v1

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

type FilterAddRequest struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterAddRequest) Reset()         { *m = FilterAddRequest{} }
func (m *FilterAddRequest) String() string { return proto.CompactTextString(m) }
func (*FilterAddRequest) ProtoMessage()    {}
func (*FilterAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}

func (m *FilterAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAddRequest.Unmarshal(m, b)
}
func (m *FilterAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAddRequest.Marshal(b, m, deterministic)
}
func (m *FilterAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAddRequest.Merge(m, src)
}
func (m *FilterAddRequest) XXX_Size() int {
	return xxx_messageInfo_FilterAddRequest.Size(m)
}
func (m *FilterAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAddRequest proto.InternalMessageInfo

func (m *FilterAddRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type FilterClearRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterClearRequest) Reset()         { *m = FilterClearRequest{} }
func (m *FilterClearRequest) String() string { return proto.CompactTextString(m) }
func (*FilterClearRequest) ProtoMessage()    {}
func (*FilterClearRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{1}
}

func (m *FilterClearRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterClearRequest.Unmarshal(m, b)
}
func (m *FilterClearRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterClearRequest.Marshal(b, m, deterministic)
}
func (m *FilterClearRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterClearRequest.Merge(m, src)
}
func (m *FilterClearRequest) XXX_Size() int {
	return xxx_messageInfo_FilterClearRequest.Size(m)
}
func (m *FilterClearRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterClearRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterClearRequest proto.InternalMessageInfo

type FilterLoadRequest struct {
	Filter               []byte   `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	HashFuncs            int64    `protobuf:"varint,2,opt,name=hashFuncs,proto3" json:"hashFuncs,omitempty"`
	Tweak                int64    `protobuf:"varint,3,opt,name=tweak,proto3" json:"tweak,omitempty"`
	Flags                int64    `protobuf:"varint,4,opt,name=flags,proto3" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterLoadRequest) Reset()         { *m = FilterLoadRequest{} }
func (m *FilterLoadRequest) String() string { return proto.CompactTextString(m) }
func (*FilterLoadRequest) ProtoMessage()    {}
func (*FilterLoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{2}
}

func (m *FilterLoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterLoadRequest.Unmarshal(m, b)
}
func (m *FilterLoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterLoadRequest.Marshal(b, m, deterministic)
}
func (m *FilterLoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterLoadRequest.Merge(m, src)
}
func (m *FilterLoadRequest) XXX_Size() int {
	return xxx_messageInfo_FilterLoadRequest.Size(m)
}
func (m *FilterLoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterLoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterLoadRequest proto.InternalMessageInfo

func (m *FilterLoadRequest) GetFilter() []byte {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *FilterLoadRequest) GetHashFuncs() int64 {
	if m != nil {
		return m.HashFuncs
	}
	return 0
}

func (m *FilterLoadRequest) GetTweak() int64 {
	if m != nil {
		return m.Tweak
	}
	return 0
}

func (m *FilterLoadRequest) GetFlags() int64 {
	if m != nil {
		return m.Flags
	}
	return 0
}

type MemPoolRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemPoolRequest) Reset()         { *m = MemPoolRequest{} }
func (m *MemPoolRequest) String() string { return proto.CompactTextString(m) }
func (*MemPoolRequest) ProtoMessage()    {}
func (*MemPoolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{3}
}

func (m *MemPoolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemPoolRequest.Unmarshal(m, b)
}
func (m *MemPoolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemPoolRequest.Marshal(b, m, deterministic)
}
func (m *MemPoolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemPoolRequest.Merge(m, src)
}
func (m *MemPoolRequest) XXX_Size() int {
	return xxx_messageInfo_MemPoolRequest.Size(m)
}
func (m *MemPoolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemPoolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemPoolRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FilterAddRequest)(nil), "qitmeer.p2p.v1.FilterAddRequest")
	proto.RegisterType((*FilterClearRequest)(nil), "qitmeer.p2p.v1.FilterClearRequest")
	proto.RegisterType((*FilterLoadRequest)(nil), "qitmeer.p2p.v1.FilterLoadRequest")
	proto.RegisterType((*MemPoolRequest)(nil), "qitmeer.p2p.v1.MemPoolRequest")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_1f5303cab7a20d6f) }

var fileDescriptor_1f5303cab7a20d6f = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0xcc, 0x2c, 0xc9, 0x4d, 0x05,
	0x71, 0x8d, 0x0a, 0xf4, 0xca, 0x0c, 0x95, 0xd4, 0xb8, 0x04, 0xdc, 0xc0, 0xf2, 0x8e, 0x29, 0x29,
	0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x92, 0x08, 0x97, 0x10, 0x44, 0x9d, 0x73, 0x4e,
	0x6a, 0x62, 0x11, 0x54, 0xa5, 0x52, 0x29, 0x97, 0x20, 0x44, 0xd4, 0x27, 0x3f, 0x11, 0xae, 0x5d,
	0x8c, 0x8b, 0x0d, 0x62, 0x25, 0xd4, 0x00, 0x28, 0x4f, 0x48, 0x86, 0x8b, 0x33, 0x23, 0xb1, 0x38,
	0xc3, 0xad, 0x34, 0x2f, 0xb9, 0x58, 0x82, 0x49, 0x81, 0x51, 0x83, 0x39, 0x08, 0x21, 0x20, 0x24,
	0xc2, 0xc5, 0x5a, 0x52, 0x9e, 0x9a, 0x98, 0x2d, 0xc1, 0x0c, 0x96, 0x81, 0x70, 0x40, 0xa2, 0x69,
	0x39, 0x89, 0xe9, 0xc5, 0x12, 0x2c, 0x10, 0x51, 0x30, 0x47, 0x49, 0x80, 0x8b, 0xcf, 0x37, 0x35,
	0x37, 0x20, 0x3f, 0x3f, 0x07, 0x6a, 0x67, 0x12, 0x1b, 0xd8, 0x77, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x1c, 0x0b, 0xac, 0x42, 0xed, 0x00, 0x00, 0x00,
}
