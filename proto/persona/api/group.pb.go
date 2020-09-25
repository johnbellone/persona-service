// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: persona/api/group.proto

package api

import (
	proto "github.com/golang/protobuf/proto"
	_type "github.com/johnbellone/persona-service/proto/persona/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type GroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Realm string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Types that are assignable to Param:
	//	*GroupRequest_Name
	//	*GroupRequest_Uuid
	Param isGroupRequest_Param `protobuf_oneof:"param"`
}

func (x *GroupRequest) Reset() {
	*x = GroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persona_api_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupRequest) ProtoMessage() {}

func (x *GroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_persona_api_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupRequest.ProtoReflect.Descriptor instead.
func (*GroupRequest) Descriptor() ([]byte, []int) {
	return file_persona_api_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupRequest) GetRealm() string {
	if x != nil {
		return x.Realm
	}
	return ""
}

func (m *GroupRequest) GetParam() isGroupRequest_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *GroupRequest) GetName() string {
	if x, ok := x.GetParam().(*GroupRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (x *GroupRequest) GetUuid() string {
	if x, ok := x.GetParam().(*GroupRequest_Uuid); ok {
		return x.Uuid
	}
	return ""
}

type isGroupRequest_Param interface {
	isGroupRequest_Param()
}

type GroupRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

type GroupRequest_Uuid struct {
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3,oneof"`
}

func (*GroupRequest_Name) isGroupRequest_Param() {}

func (*GroupRequest_Uuid) isGroupRequest_Param() {}

var File_persona_api_group_proto protoreflect.FileDescriptor

var file_persona_api_group_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x0c, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65,
	0x61, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x61, 0x6c, 0x6d,
	0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x32, 0xf4, 0x02, 0x0a, 0x0c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x1c, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x1c, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12, 0x16, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x7d, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x7d,
	0x12, 0x5a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x2a, 0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x7d, 0x42, 0x3a, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x68, 0x6e, 0x62,
	0x65, 0x6c, 0x6c, 0x6f, 0x6e, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_persona_api_group_proto_rawDescOnce sync.Once
	file_persona_api_group_proto_rawDescData = file_persona_api_group_proto_rawDesc
)

func file_persona_api_group_proto_rawDescGZIP() []byte {
	file_persona_api_group_proto_rawDescOnce.Do(func() {
		file_persona_api_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_persona_api_group_proto_rawDescData)
	})
	return file_persona_api_group_proto_rawDescData
}

var file_persona_api_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_persona_api_group_proto_goTypes = []interface{}{
	(*GroupRequest)(nil),  // 0: persona.api.v1.GroupRequest
	(*status.Status)(nil), // 1: google.rpc.Status
	(*_type.Group)(nil),   // 2: persona.type.Group
}
var file_persona_api_group_proto_depIdxs = []int32{
	0, // 0: persona.api.v1.GroupService.Create:input_type -> persona.api.v1.GroupRequest
	0, // 1: persona.api.v1.GroupService.Get:input_type -> persona.api.v1.GroupRequest
	0, // 2: persona.api.v1.GroupService.Update:input_type -> persona.api.v1.GroupRequest
	0, // 3: persona.api.v1.GroupService.Delete:input_type -> persona.api.v1.GroupRequest
	1, // 4: persona.api.v1.GroupService.Create:output_type -> google.rpc.Status
	2, // 5: persona.api.v1.GroupService.Get:output_type -> persona.type.Group
	1, // 6: persona.api.v1.GroupService.Update:output_type -> google.rpc.Status
	1, // 7: persona.api.v1.GroupService.Delete:output_type -> google.rpc.Status
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_persona_api_group_proto_init() }
func file_persona_api_group_proto_init() {
	if File_persona_api_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_persona_api_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupRequest); i {
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
	file_persona_api_group_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*GroupRequest_Name)(nil),
		(*GroupRequest_Uuid)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_persona_api_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_persona_api_group_proto_goTypes,
		DependencyIndexes: file_persona_api_group_proto_depIdxs,
		MessageInfos:      file_persona_api_group_proto_msgTypes,
	}.Build()
	File_persona_api_group_proto = out.File
	file_persona_api_group_proto_rawDesc = nil
	file_persona_api_group_proto_goTypes = nil
	file_persona_api_group_proto_depIdxs = nil
}
