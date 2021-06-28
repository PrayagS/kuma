// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: mesh/v1alpha1/zoneingress_overview.proto

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ZoneIngressOverview defines the projected state of a ZoneIngress.
type ZoneIngressOverview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneIngress        *ZoneIngress        `protobuf:"bytes,1,opt,name=zone_ingress,json=zoneIngress,proto3" json:"zone_ingress,omitempty"`
	ZoneIngressInsight *ZoneIngressInsight `protobuf:"bytes,2,opt,name=zone_ingress_insight,json=zoneIngressInsight,proto3" json:"zone_ingress_insight,omitempty"`
}

func (x *ZoneIngressOverview) Reset() {
	*x = ZoneIngressOverview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mesh_v1alpha1_zoneingress_overview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneIngressOverview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneIngressOverview) ProtoMessage() {}

func (x *ZoneIngressOverview) ProtoReflect() protoreflect.Message {
	mi := &file_mesh_v1alpha1_zoneingress_overview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneIngressOverview.ProtoReflect.Descriptor instead.
func (*ZoneIngressOverview) Descriptor() ([]byte, []int) {
	return file_mesh_v1alpha1_zoneingress_overview_proto_rawDescGZIP(), []int{0}
}

func (x *ZoneIngressOverview) GetZoneIngress() *ZoneIngress {
	if x != nil {
		return x.ZoneIngress
	}
	return nil
}

func (x *ZoneIngressOverview) GetZoneIngressInsight() *ZoneIngressInsight {
	if x != nil {
		return x.ZoneIngressInsight
	}
	return nil
}

var File_mesh_v1alpha1_zoneingress_overview_proto protoreflect.FileDescriptor

var file_mesh_v1alpha1_zoneingress_overview_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x75, 0x6d, 0x61,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x20,
	0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x7a, 0x6f,
	0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x28, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x73,
	0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x13, 0x5a,
	0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x42, 0x0a, 0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x58, 0x0a, 0x14, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74, 0x52, 0x12, 0x7a, 0x6f,
	0x6e, 0x65, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x73, 0x69, 0x67, 0x68, 0x74,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x75, 0x6d, 0x61, 0x68, 0x71, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mesh_v1alpha1_zoneingress_overview_proto_rawDescOnce sync.Once
	file_mesh_v1alpha1_zoneingress_overview_proto_rawDescData = file_mesh_v1alpha1_zoneingress_overview_proto_rawDesc
)

func file_mesh_v1alpha1_zoneingress_overview_proto_rawDescGZIP() []byte {
	file_mesh_v1alpha1_zoneingress_overview_proto_rawDescOnce.Do(func() {
		file_mesh_v1alpha1_zoneingress_overview_proto_rawDescData = protoimpl.X.CompressGZIP(file_mesh_v1alpha1_zoneingress_overview_proto_rawDescData)
	})
	return file_mesh_v1alpha1_zoneingress_overview_proto_rawDescData
}

var file_mesh_v1alpha1_zoneingress_overview_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mesh_v1alpha1_zoneingress_overview_proto_goTypes = []interface{}{
	(*ZoneIngressOverview)(nil), // 0: kuma.mesh.v1alpha1.ZoneIngressOverview
	(*ZoneIngress)(nil),         // 1: kuma.mesh.v1alpha1.ZoneIngress
	(*ZoneIngressInsight)(nil),  // 2: kuma.mesh.v1alpha1.ZoneIngressInsight
}
var file_mesh_v1alpha1_zoneingress_overview_proto_depIdxs = []int32{
	1, // 0: kuma.mesh.v1alpha1.ZoneIngressOverview.zone_ingress:type_name -> kuma.mesh.v1alpha1.ZoneIngress
	2, // 1: kuma.mesh.v1alpha1.ZoneIngressOverview.zone_ingress_insight:type_name -> kuma.mesh.v1alpha1.ZoneIngressInsight
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mesh_v1alpha1_zoneingress_overview_proto_init() }
func file_mesh_v1alpha1_zoneingress_overview_proto_init() {
	if File_mesh_v1alpha1_zoneingress_overview_proto != nil {
		return
	}
	file_mesh_v1alpha1_zone_ingress_proto_init()
	file_mesh_v1alpha1_zone_ingress_insight_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mesh_v1alpha1_zoneingress_overview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneIngressOverview); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mesh_v1alpha1_zoneingress_overview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mesh_v1alpha1_zoneingress_overview_proto_goTypes,
		DependencyIndexes: file_mesh_v1alpha1_zoneingress_overview_proto_depIdxs,
		MessageInfos:      file_mesh_v1alpha1_zoneingress_overview_proto_msgTypes,
	}.Build()
	File_mesh_v1alpha1_zoneingress_overview_proto = out.File
	file_mesh_v1alpha1_zoneingress_overview_proto_rawDesc = nil
	file_mesh_v1alpha1_zoneingress_overview_proto_goTypes = nil
	file_mesh_v1alpha1_zoneingress_overview_proto_depIdxs = nil
}
