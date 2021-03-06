// Code generated by protoc-gen-go.
// source: example.proto
// DO NOT EDIT!

/*
Package staff is a generated protocol buffer package.

It is generated from these files:
	example.proto

It has these top-level messages:
	Staff
*/
package staff

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Staff struct {
	ID   string `protobuf:"bytes,1,opt,name=ID"     json:"id,omitempty"    xml:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name"   json:"name,omitempty"  xml:"name,omitempty"`
	Age  int64  `protobuf:"varint,3,opt,name=Age"   json:"age,omitempty"   xml:"age,omitempty"`
}

func (m *Staff) Reset()                    { *m = Staff{} }
func (m *Staff) String() string            { return proto.CompactTextString(m) }
func (*Staff) ProtoMessage()               {}
func (*Staff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Staff) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Staff) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Staff) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Staff)(nil), "staff.Staff")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2e, 0x49, 0x4c, 0x4b,
	0x53, 0xb2, 0xe5, 0x62, 0x0d, 0x06, 0x31, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x3c, 0x5d, 0x84, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25,
	0x98, 0xc0, 0x22, 0x60, 0xb6, 0x90, 0x00, 0x17, 0xb3, 0x63, 0x7a, 0xaa, 0x04, 0xb3, 0x02, 0xa3,
	0x06, 0x73, 0x10, 0x88, 0x99, 0xc4, 0x06, 0x36, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x72, 0x21, 0xc1, 0x5d, 0x00, 0x00, 0x00,
}
