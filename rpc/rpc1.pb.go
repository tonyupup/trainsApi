// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc1.proto

package rpc

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

type TrainCode struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainCode) Reset()         { *m = TrainCode{} }
func (m *TrainCode) String() string { return proto.CompactTextString(m) }
func (*TrainCode) ProtoMessage()    {}
func (*TrainCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_7964eb4b8215a43a, []int{0}
}

func (m *TrainCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainCode.Unmarshal(m, b)
}
func (m *TrainCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainCode.Marshal(b, m, deterministic)
}
func (m *TrainCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainCode.Merge(m, src)
}
func (m *TrainCode) XXX_Size() int {
	return xxx_messageInfo_TrainCode.Size(m)
}
func (m *TrainCode) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainCode.DiscardUnknown(m)
}

var xxx_messageInfo_TrainCode proto.InternalMessageInfo

func (m *TrainCode) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type TrainExists struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TrainExists) Reset()         { *m = TrainExists{} }
func (m *TrainExists) String() string { return proto.CompactTextString(m) }
func (*TrainExists) ProtoMessage()    {}
func (*TrainExists) Descriptor() ([]byte, []int) {
	return fileDescriptor_7964eb4b8215a43a, []int{1}
}

func (m *TrainExists) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainExists.Unmarshal(m, b)
}
func (m *TrainExists) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainExists.Marshal(b, m, deterministic)
}
func (m *TrainExists) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainExists.Merge(m, src)
}
func (m *TrainExists) XXX_Size() int {
	return xxx_messageInfo_TrainExists.Size(m)
}
func (m *TrainExists) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainExists.DiscardUnknown(m)
}

var xxx_messageInfo_TrainExists proto.InternalMessageInfo

func (m *TrainExists) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func init() {
	proto.RegisterType((*TrainCode)(nil), "TrainCode")
	proto.RegisterType((*TrainExists)(nil), "TrainExists")
}

func init() { proto.RegisterFile("rpc1.proto", fileDescriptor_7964eb4b8215a43a) }

var fileDescriptor_7964eb4b8215a43a = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2a, 0x48, 0x36,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe7, 0xe2, 0x0c, 0x29, 0x4a, 0xcc, 0xcc, 0x73,
	0xce, 0x4f, 0x49, 0x15, 0x12, 0xe2, 0x62, 0x49, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x54, 0xb9, 0xb8, 0xc1, 0x0a, 0x5c, 0x2b, 0x32, 0x8b, 0x4b, 0x8a,
	0x85, 0xc4, 0xb8, 0xd8, 0x52, 0xc1, 0x2c, 0xb0, 0x22, 0x8e, 0x20, 0x28, 0xcf, 0x48, 0x9f, 0x8b,
	0x23, 0x38, 0x1e, 0xac, 0xb0, 0x58, 0x48, 0x99, 0x8b, 0xd5, 0xb5, 0x22, 0xb3, 0xa4, 0x58, 0x88,
	0x4b, 0x0f, 0x6e, 0xb6, 0x14, 0x8f, 0x1e, 0x92, 0x31, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xfb, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x76, 0x14, 0xb2, 0x8d, 0x00, 0x00, 0x00,
}