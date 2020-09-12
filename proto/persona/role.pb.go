// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persona/role.proto

package persona

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/rpc/status"
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

type RoleRequest struct {
	Realm string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Types that are valid to be assigned to Param:
	//	*RoleRequest_Name
	//	*RoleRequest_Slug
	Param                isRoleRequest_Param `protobuf_oneof:"param"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RoleRequest) Reset()         { *m = RoleRequest{} }
func (m *RoleRequest) String() string { return proto.CompactTextString(m) }
func (*RoleRequest) ProtoMessage()    {}
func (*RoleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e13756a0d1ef8e4, []int{0}
}

func (m *RoleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleRequest.Unmarshal(m, b)
}
func (m *RoleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleRequest.Marshal(b, m, deterministic)
}
func (m *RoleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleRequest.Merge(m, src)
}
func (m *RoleRequest) XXX_Size() int {
	return xxx_messageInfo_RoleRequest.Size(m)
}
func (m *RoleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoleRequest proto.InternalMessageInfo

func (m *RoleRequest) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

type isRoleRequest_Param interface {
	isRoleRequest_Param()
}

type RoleRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

type RoleRequest_Slug struct {
	Slug string `protobuf:"bytes,3,opt,name=slug,proto3,oneof"`
}

func (*RoleRequest_Name) isRoleRequest_Param() {}

func (*RoleRequest_Slug) isRoleRequest_Param() {}

func (m *RoleRequest) GetParam() isRoleRequest_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *RoleRequest) GetName() string {
	if x, ok := m.GetParam().(*RoleRequest_Name); ok {
		return x.Name
	}
	return ""
}

func (m *RoleRequest) GetSlug() string {
	if x, ok := m.GetParam().(*RoleRequest_Slug); ok {
		return x.Slug
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RoleRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RoleRequest_Name)(nil),
		(*RoleRequest_Slug)(nil),
	}
}

type Role struct {
	Realm                string               `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Metadata             *_struct.Struct      `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e13756a0d1ef8e4, []int{1}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Role) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Role) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Role) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*RoleRequest)(nil), "persona.v1.RoleRequest")
	proto.RegisterType((*Role)(nil), "persona.v1.Role")
}

func init() { proto.RegisterFile("persona/role.proto", fileDescriptor_2e13756a0d1ef8e4) }

var fileDescriptor_2e13756a0d1ef8e4 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0xed, 0x4f, 0x35, 0x45, 0x94, 0xb0, 0x4b, 0x87, 0x41, 0xd9, 0x32, 0xa7, 0x45, 0x30,
	0x61, 0x77, 0xc5, 0xb3, 0x56, 0x41, 0x0f, 0x82, 0x30, 0x5d, 0x41, 0xbc, 0x2c, 0x99, 0xe9, 0xdb,
	0x36, 0x92, 0x99, 0xc4, 0x24, 0xd3, 0x8b, 0x78, 0xf1, 0xea, 0xd1, 0x3f, 0xcd, 0x8b, 0x47, 0x0f,
	0xfe, 0x21, 0x92, 0x1f, 0x53, 0xc7, 0x6a, 0x41, 0xf7, 0xf8, 0xbe, 0xef, 0xfb, 0x3e, 0x93, 0x6f,
	0x5e, 0x06, 0x61, 0x05, 0xda, 0xc8, 0x9a, 0x51, 0x2d, 0x05, 0x10, 0xa5, 0xa5, 0x95, 0x18, 0x45,
	0x8d, 0x6c, 0x4e, 0xd2, 0xa3, 0x95, 0x94, 0x2b, 0x01, 0x94, 0x29, 0x4e, 0x2f, 0x39, 0x88, 0xe5,
	0x45, 0x01, 0x6b, 0xb6, 0xe1, 0x52, 0x07, 0x73, 0x7a, 0xb7, 0x63, 0x60, 0x75, 0x2d, 0x2d, 0xb3,
	0x5c, 0xd6, 0x66, 0xa7, 0xeb, 0xab, 0xa2, 0xb9, 0xa4, 0xc6, 0xea, 0xa6, 0xb4, 0xb1, 0x7b, 0xb4,
	0xdb, 0xb5, 0xbc, 0x02, 0x63, 0x59, 0xa5, 0xa2, 0x61, 0x1a, 0x0d, 0x5a, 0x95, 0xd4, 0x58, 0x66,
	0x9b, 0xc8, 0xcd, 0xde, 0xa0, 0x49, 0x2e, 0x05, 0xe4, 0xf0, 0xbe, 0x01, 0x63, 0xf1, 0x01, 0x1a,
	0x69, 0x60, 0xa2, 0x4a, 0x7a, 0xb3, 0xde, 0xf1, 0xcd, 0x3c, 0x14, 0xf8, 0x00, 0x0d, 0x6b, 0x56,
	0x41, 0xd2, 0x77, 0xe2, 0x8b, 0x6b, 0xb9, 0xaf, 0x9c, 0x6a, 0x44, 0xb3, 0x4a, 0x06, 0xad, 0xea,
	0xaa, 0xf9, 0x75, 0x34, 0x52, 0x4c, 0xb3, 0x2a, 0xfb, 0xdc, 0x47, 0x43, 0x87, 0xde, 0xc3, 0xc4,
	0x5d, 0x66, 0x24, 0xce, 0xd0, 0x64, 0x09, 0xa6, 0xd4, 0x5c, 0xb9, 0xe8, 0x01, 0x9c, 0x77, 0x25,
	0xfc, 0x18, 0x4d, 0x4a, 0x0d, 0xcc, 0xc2, 0x85, 0x4b, 0x98, 0x0c, 0x67, 0xbd, 0xe3, 0xc9, 0x69,
	0x4a, 0x42, 0x3a, 0xd2, 0xc6, 0x27, 0xe7, 0x6d, 0xfc, 0xf9, 0xe0, 0xfb, 0x93, 0x41, 0x8e, 0xc2,
	0x8c, 0x53, 0x1d, 0xa1, 0x51, 0xcb, 0x2d, 0x61, 0xf4, 0x8f, 0x84, 0x30, 0xe3, 0x09, 0x67, 0xe8,
	0x46, 0x05, 0x96, 0x2d, 0x99, 0x65, 0xc9, 0xd8, 0x8f, 0x4f, 0xff, 0x18, 0x5f, 0xf8, 0xed, 0xe4,
	0x5b, 0xe3, 0xe9, 0xb7, 0x7e, 0xb8, 0xe8, 0x05, 0xe8, 0x0d, 0x2f, 0x01, 0xbf, 0x44, 0xe3, 0xa7,
	0xfe, 0x50, 0x78, 0x4a, 0x7e, 0xbd, 0x12, 0xd2, 0xd9, 0x45, 0x8a, 0x5b, 0xaa, 0x56, 0x25, 0x59,
	0xf8, 0xa5, 0x65, 0x87, 0x9f, 0xbe, 0xfe, 0xf8, 0xd2, 0xbf, 0x9d, 0xdd, 0xf2, 0xcf, 0x64, 0x73,
	0xe2, 0x5f, 0x9b, 0xc1, 0xaf, 0xd0, 0xe0, 0x39, 0xd8, 0xfd, 0xa8, 0x3b, 0xbb, 0x8d, 0xec, 0x9e,
	0x07, 0x4d, 0xf1, 0xe1, 0x6f, 0x20, 0xfa, 0xc1, 0xef, 0xee, 0x23, 0x3e, 0x47, 0xe3, 0xd7, 0x3e,
	0xf1, 0xff, 0x1d, 0x2f, 0x52, 0xd3, 0xfd, 0xd4, 0x67, 0x20, 0xe0, 0x8a, 0xd4, 0xfb, 0x7f, 0xa7,
	0xce, 0x1f, 0xbd, 0x7d, 0xb8, 0xe2, 0x76, 0xdd, 0x14, 0xa4, 0x94, 0x15, 0x7d, 0x27, 0xd7, 0x75,
	0x01, 0x42, 0xc8, 0x1a, 0x68, 0xfc, 0xc6, 0x03, 0x13, 0x2e, 0x3d, 0xfc, 0x1e, 0xad, 0x5a, 0x8c,
	0x7d, 0x79, 0xf6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xac, 0x0d, 0xa4, 0xfb, 0xba, 0x03, 0x00, 0x00,
}
