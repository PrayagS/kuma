// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mesh/v1alpha1/dataplane_overview.proto

package v1alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// DataplaneOverview defines the projected state of a Dataplane.
type DataplaneOverview struct {
	Dataplane            *Dataplane        `protobuf:"bytes,1,opt,name=dataplane,proto3" json:"dataplane,omitempty"`
	DataplaneInsight     *DataplaneInsight `protobuf:"bytes,2,opt,name=dataplane_insight,json=dataplaneInsight,proto3" json:"dataplane_insight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DataplaneOverview) Reset()         { *m = DataplaneOverview{} }
func (m *DataplaneOverview) String() string { return proto.CompactTextString(m) }
func (*DataplaneOverview) ProtoMessage()    {}
func (*DataplaneOverview) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea8390996584fcf0, []int{0}
}

func (m *DataplaneOverview) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneOverview.Unmarshal(m, b)
}
func (m *DataplaneOverview) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneOverview.Marshal(b, m, deterministic)
}
func (m *DataplaneOverview) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneOverview.Merge(m, src)
}
func (m *DataplaneOverview) XXX_Size() int {
	return xxx_messageInfo_DataplaneOverview.Size(m)
}
func (m *DataplaneOverview) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneOverview.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneOverview proto.InternalMessageInfo

func (m *DataplaneOverview) GetDataplane() *Dataplane {
	if m != nil {
		return m.Dataplane
	}
	return nil
}

func (m *DataplaneOverview) GetDataplaneInsight() *DataplaneInsight {
	if m != nil {
		return m.DataplaneInsight
	}
	return nil
}

func init() {
	proto.RegisterType((*DataplaneOverview)(nil), "kuma.mesh.v1alpha1.DataplaneOverview")
}

func init() {
	proto.RegisterFile("mesh/v1alpha1/dataplane_overview.proto", fileDescriptor_ea8390996584fcf0)
}

var fileDescriptor_ea8390996584fcf0 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0x4d, 0x2d, 0xce,
	0xd0, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x4f, 0x49, 0x2c, 0x49, 0x2c, 0xc8,
	0x49, 0xcc, 0x4b, 0x8d, 0xcf, 0x2f, 0x4b, 0x2d, 0x2a, 0xcb, 0x4c, 0x2d, 0xd7, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xca, 0x2e, 0xcd, 0x4d, 0xd4, 0x03, 0x29, 0xd6, 0x83, 0x29, 0x96, 0x92,
	0xc5, 0xa1, 0x17, 0xa2, 0x45, 0x4a, 0x15, 0x97, 0xd1, 0x99, 0x79, 0xc5, 0x99, 0xe9, 0x19, 0x25,
	0x50, 0x65, 0xe2, 0x65, 0x89, 0x39, 0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x42,
	0x69, 0x3d, 0x23, 0x97, 0xa0, 0x0b, 0x4c, 0x93, 0x3f, 0xd4, 0x39, 0x42, 0xee, 0x5c, 0x9c, 0x70,
	0x93, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x64, 0xf5, 0x30, 0x1d, 0xa7, 0x07, 0xd7, 0xe9,
	0xc4, 0xb5, 0xeb, 0xe5, 0x01, 0x66, 0xd6, 0x2e, 0x46, 0x26, 0x01, 0xc6, 0x20, 0x84, 0x5e, 0xa1,
	0x40, 0x2e, 0x41, 0x0c, 0x27, 0x49, 0x30, 0x81, 0x0d, 0x54, 0xc1, 0x6b, 0xa0, 0x27, 0x44, 0x6d,
	0x90, 0x40, 0x0a, 0x9a, 0x88, 0x13, 0x57, 0x14, 0x07, 0x4c, 0x79, 0x12, 0x1b, 0xd8, 0x13, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x9c, 0x83, 0x82, 0x61, 0x01, 0x00, 0x00,
}
