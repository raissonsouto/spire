//* Resolves the derived selectors for a given Node Agent. This mapping
//will be stored, and used to further derive which workloads the Node
//Agent is authorized to run.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/plugin/server/noderesolver/v0/noderesolver.proto

package noderesolverv0

import (
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
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

//* Represents a request with a list of BaseSPIFFEIDs.
type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* A list of BaseSPIFFE Ids.
	BaseSpiffeIdList []string `protobuf:"bytes,1,rep,name=baseSpiffeIdList,proto3" json:"baseSpiffeIdList,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveRequest) GetBaseSpiffeIdList() []string {
	if x != nil {
		return x.BaseSpiffeIdList
	}
	return nil
}

//* Represents a response with a map of SPIFFE ID to a list of Selectors.
type ResolveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* Map[SPIFFE_ID] => Selectors.
	Map map[string]*common.Selectors `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResolveResponse) Reset() {
	*x = ResolveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveResponse) ProtoMessage() {}

func (x *ResolveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveResponse.ProtoReflect.Descriptor instead.
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveResponse) GetMap() map[string]*common.Selectors {
	if x != nil {
		return x.Map
	}
	return nil
}

var File_spire_plugin_server_noderesolver_v0_noderesolver_proto protoreflect.FileDescriptor

var file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2f, 0x76, 0x30, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x1a, 0x20, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x53, 0x70, 0x69, 0x66, 0x66, 0x65,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x62, 0x61,
	0x73, 0x65, 0x53, 0x70, 0x69, 0x66, 0x66, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xa9,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6e,
	0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x4f, 0x0a, 0x08, 0x4d, 0x61, 0x70,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xb4, 0x02, 0x0a, 0x0c, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x60, 0x0a, 0x07, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x30, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x76, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescOnce sync.Once
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescData = file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDesc
)

func file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescGZIP() []byte {
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescOnce.Do(func() {
		file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescData)
	})
	return file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDescData
}

var file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spire_plugin_server_noderesolver_v0_noderesolver_proto_goTypes = []interface{}{
	(*ResolveRequest)(nil),               // 0: spire.server.noderesolver.ResolveRequest
	(*ResolveResponse)(nil),              // 1: spire.server.noderesolver.ResolveResponse
	nil,                                  // 2: spire.server.noderesolver.ResolveResponse.MapEntry
	(*common.Selectors)(nil),             // 3: spire.common.Selectors
	(*plugin.ConfigureRequest)(nil),      // 4: spire.common.plugin.ConfigureRequest
	(*plugin.GetPluginInfoRequest)(nil),  // 5: spire.common.plugin.GetPluginInfoRequest
	(*plugin.ConfigureResponse)(nil),     // 6: spire.common.plugin.ConfigureResponse
	(*plugin.GetPluginInfoResponse)(nil), // 7: spire.common.plugin.GetPluginInfoResponse
}
var file_spire_plugin_server_noderesolver_v0_noderesolver_proto_depIdxs = []int32{
	2, // 0: spire.server.noderesolver.ResolveResponse.map:type_name -> spire.server.noderesolver.ResolveResponse.MapEntry
	3, // 1: spire.server.noderesolver.ResolveResponse.MapEntry.value:type_name -> spire.common.Selectors
	0, // 2: spire.server.noderesolver.NodeResolver.Resolve:input_type -> spire.server.noderesolver.ResolveRequest
	4, // 3: spire.server.noderesolver.NodeResolver.Configure:input_type -> spire.common.plugin.ConfigureRequest
	5, // 4: spire.server.noderesolver.NodeResolver.GetPluginInfo:input_type -> spire.common.plugin.GetPluginInfoRequest
	1, // 5: spire.server.noderesolver.NodeResolver.Resolve:output_type -> spire.server.noderesolver.ResolveResponse
	6, // 6: spire.server.noderesolver.NodeResolver.Configure:output_type -> spire.common.plugin.ConfigureResponse
	7, // 7: spire.server.noderesolver.NodeResolver.GetPluginInfo:output_type -> spire.common.plugin.GetPluginInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_spire_plugin_server_noderesolver_v0_noderesolver_proto_init() }
func file_spire_plugin_server_noderesolver_v0_noderesolver_proto_init() {
	if File_spire_plugin_server_noderesolver_v0_noderesolver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveResponse); i {
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
			RawDescriptor: file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_plugin_server_noderesolver_v0_noderesolver_proto_goTypes,
		DependencyIndexes: file_spire_plugin_server_noderesolver_v0_noderesolver_proto_depIdxs,
		MessageInfos:      file_spire_plugin_server_noderesolver_v0_noderesolver_proto_msgTypes,
	}.Build()
	File_spire_plugin_server_noderesolver_v0_noderesolver_proto = out.File
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_rawDesc = nil
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_goTypes = nil
	file_spire_plugin_server_noderesolver_v0_noderesolver_proto_depIdxs = nil
}
