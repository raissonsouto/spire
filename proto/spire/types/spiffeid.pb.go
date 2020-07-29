// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/types/spiffeid.proto

package types

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

// A SPIFFE ID, consisting of the trust domain name and a path portions of
// the SPIFFE ID URI.
type SPIFFEID struct {
	// Trust domain portion the SPIFFE ID (e.g. "example.org")
	TrustDomain string `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	// The path component of the SPIFFE ID (e.g. "/foo/bar/baz"). The path
	// SHOULD have a leading slash. Consumers MUST normalize the path before
	// making any sort of comparison between IDs.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SPIFFEID) Reset()         { *m = SPIFFEID{} }
func (m *SPIFFEID) String() string { return proto.CompactTextString(m) }
func (*SPIFFEID) ProtoMessage()    {}
func (*SPIFFEID) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f81ccd53595df62, []int{0}
}

func (m *SPIFFEID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPIFFEID.Unmarshal(m, b)
}
func (m *SPIFFEID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPIFFEID.Marshal(b, m, deterministic)
}
func (m *SPIFFEID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPIFFEID.Merge(m, src)
}
func (m *SPIFFEID) XXX_Size() int {
	return xxx_messageInfo_SPIFFEID.Size(m)
}
func (m *SPIFFEID) XXX_DiscardUnknown() {
	xxx_messageInfo_SPIFFEID.DiscardUnknown(m)
}

var xxx_messageInfo_SPIFFEID proto.InternalMessageInfo

func (m *SPIFFEID) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *SPIFFEID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*SPIFFEID)(nil), "spire.types.SPIFFEID")
}

func init() {
	proto.RegisterFile("spire/types/spiffeid.proto", fileDescriptor_2f81ccd53595df62)
}

var fileDescriptor_2f81ccd53595df62 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0xc8, 0x2c,
	0x4a, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x2e, 0xc8, 0x4c, 0x4b, 0x4b, 0xcd, 0x4c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xcb, 0xe9, 0x81, 0xe5, 0x94, 0x1c, 0xb9,
	0x38, 0x82, 0x03, 0x3c, 0xdd, 0xdc, 0x5c, 0x3d, 0x5d, 0x84, 0x14, 0xb9, 0x78, 0x4a, 0x8a, 0x4a,
	0x8b, 0x4b, 0xe2, 0x53, 0xf2, 0x73, 0x13, 0x33, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xb8, 0xc1, 0x62, 0x2e, 0x60, 0x21, 0x21, 0x21, 0x2e, 0x96, 0x82, 0xc4, 0x92, 0x0c, 0x09, 0x26,
	0xb0, 0x14, 0x98, 0xed, 0xa4, 0x1d, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0x0b, 0xb5, 0x4c, 0x1f, 0x62, 0x3f, 0xd8, 0x42, 0x7d, 0x24, 0xb7, 0x24, 0xb1, 0x81, 0x85,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xd7, 0x11, 0x52, 0xa1, 0x00, 0x00, 0x00,
}