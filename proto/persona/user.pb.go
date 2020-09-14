// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persona/user.proto

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

type User_State int32

const (
	User_INACTIVE  User_State = 0
	User_STAGED    User_State = 1
	User_PENDING   User_State = 2
	User_ACTIVE    User_State = 3
	User_RECOVERY  User_State = 4
	User_LOCKED    User_State = 5
	User_SUSPENDED User_State = 6
	User_EXPIRED   User_State = 7
)

var User_State_name = map[int32]string{
	0: "INACTIVE",
	1: "STAGED",
	2: "PENDING",
	3: "ACTIVE",
	4: "RECOVERY",
	5: "LOCKED",
	6: "SUSPENDED",
	7: "EXPIRED",
}

var User_State_value = map[string]int32{
	"INACTIVE":  0,
	"STAGED":    1,
	"PENDING":   2,
	"ACTIVE":    3,
	"RECOVERY":  4,
	"LOCKED":    5,
	"SUSPENDED": 6,
	"EXPIRED":   7,
}

func (x User_State) String() string {
	return proto.EnumName(User_State_name, int32(x))
}

func (User_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_addc16e92f4af9e4, []int{1, 0}
}

type UserRequest struct {
	Realm string `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	// Types that are valid to be assigned to Param:
	//	*UserRequest_Login
	//	*UserRequest_Email
	Param                isUserRequest_Param `protobuf_oneof:"param"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_addc16e92f4af9e4, []int{0}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

type isUserRequest_Param interface {
	isUserRequest_Param()
}

type UserRequest_Login struct {
	Login string `protobuf:"bytes,2,opt,name=login,proto3,oneof"`
}

type UserRequest_Email struct {
	Email string `protobuf:"bytes,3,opt,name=email,proto3,oneof"`
}

func (*UserRequest_Login) isUserRequest_Param() {}

func (*UserRequest_Email) isUserRequest_Param() {}

func (m *UserRequest) GetParam() isUserRequest_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *UserRequest) GetLogin() string {
	if x, ok := m.GetParam().(*UserRequest_Login); ok {
		return x.Login
	}
	return ""
}

func (m *UserRequest) GetEmail() string {
	if x, ok := m.GetParam().(*UserRequest_Email); ok {
		return x.Email
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UserRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UserRequest_Login)(nil),
		(*UserRequest_Email)(nil),
	}
}

type User struct {
	Realm                string               `protobuf:"bytes,1,opt,name=realm,proto3" json:"realm,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Login                string               `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password             string               `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Status               User_State           `protobuf:"varint,7,opt,name=status,proto3,enum=persona.v1.User_State" json:"status,omitempty"`
	LastPasswordChange   *timestamp.Timestamp `protobuf:"bytes,8,opt,name=last_password_change,json=lastPasswordChange,proto3" json:"last_password_change,omitempty"`
	LastLogin            *timestamp.Timestamp `protobuf:"bytes,9,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	LastActive           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=last_active,json=lastActive,proto3" json:"last_active,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,12,opt,name=updated,proto3" json:"updated,omitempty"`
	Activated            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=activated,proto3" json:"activated,omitempty"`
	Metadata             *_struct.Struct      `protobuf:"bytes,14,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_addc16e92f4af9e4, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetRealm() string {
	if m != nil {
		return m.Realm
	}
	return ""
}

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetStatus() User_State {
	if m != nil {
		return m.Status
	}
	return User_INACTIVE
}

func (m *User) GetLastPasswordChange() *timestamp.Timestamp {
	if m != nil {
		return m.LastPasswordChange
	}
	return nil
}

func (m *User) GetLastLogin() *timestamp.Timestamp {
	if m != nil {
		return m.LastLogin
	}
	return nil
}

func (m *User) GetLastActive() *timestamp.Timestamp {
	if m != nil {
		return m.LastActive
	}
	return nil
}

func (m *User) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *User) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *User) GetActivated() *timestamp.Timestamp {
	if m != nil {
		return m.Activated
	}
	return nil
}

func (m *User) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("persona.v1.User_State", User_State_name, User_State_value)
	proto.RegisterType((*UserRequest)(nil), "persona.v1.UserRequest")
	proto.RegisterType((*User)(nil), "persona.v1.User")
}

func init() { proto.RegisterFile("persona/user.proto", fileDescriptor_addc16e92f4af9e4) }

var fileDescriptor_addc16e92f4af9e4 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x9b, 0x0f, 0x3b, 0xc9, 0xa4, 0x2d, 0xd6, 0xaa, 0x6d, 0xac, 0x08, 0xd4, 0x2a, 0xa7,
	0x0a, 0x09, 0x5b, 0x6d, 0x11, 0x12, 0x07, 0x10, 0x69, 0x62, 0x95, 0x88, 0xaa, 0x2d, 0x4e, 0x5a,
	0x01, 0x42, 0xaa, 0x36, 0xce, 0x36, 0x31, 0xb2, 0xbd, 0xc6, 0xbb, 0x0e, 0x07, 0xc4, 0x85, 0x57,
	0xe0, 0xd1, 0xb8, 0x70, 0xe4, 0xc0, 0x95, 0x77, 0x40, 0xfb, 0xe1, 0x02, 0x81, 0x4a, 0x81, 0x9b,
	0x67, 0xe6, 0xff, 0xff, 0xcd, 0xee, 0xac, 0x77, 0x01, 0xa5, 0x24, 0x63, 0x34, 0xc1, 0x6e, 0xce,
	0x48, 0xe6, 0xa4, 0x19, 0xe5, 0x14, 0x81, 0xce, 0x39, 0xf3, 0xbd, 0xf6, 0xf6, 0x94, 0xd2, 0x69,
	0x44, 0x5c, 0x9c, 0x86, 0xee, 0x55, 0x48, 0xa2, 0xc9, 0xe5, 0x98, 0xcc, 0xf0, 0x3c, 0xa4, 0x5a,
	0xdc, 0xbe, 0xfd, 0x8b, 0x00, 0x27, 0x09, 0xe5, 0x98, 0x87, 0x34, 0x61, 0x0b, 0x55, 0x19, 0x8d,
	0xf3, 0x2b, 0x97, 0xf1, 0x2c, 0x0f, 0xb8, 0xae, 0x6e, 0x2f, 0x56, 0x79, 0x18, 0x13, 0xc6, 0x71,
	0x9c, 0x6a, 0x41, 0x4b, 0x0b, 0xb2, 0x34, 0x70, 0x19, 0xc7, 0x3c, 0xd7, 0xdc, 0xce, 0x6b, 0x68,
	0x9e, 0x33, 0x92, 0xf9, 0xe4, 0x6d, 0x4e, 0x18, 0x47, 0x1b, 0x60, 0x64, 0x04, 0x47, 0xb1, 0x5d,
	0xda, 0x29, 0xed, 0x36, 0x7c, 0x15, 0xa0, 0x2d, 0x30, 0x22, 0x3a, 0x0d, 0x13, 0xbb, 0x2c, 0xb2,
	0x4f, 0x57, 0x7c, 0x15, 0x8a, 0x3c, 0x89, 0x71, 0x18, 0xd9, 0x95, 0x22, 0x2f, 0xc3, 0xc3, 0x1a,
	0x18, 0x29, 0xce, 0x70, 0xdc, 0xf9, 0x6e, 0x40, 0x55, 0xe0, 0x6f, 0xe0, 0x22, 0xa8, 0xe6, 0x79,
	0x38, 0x51, 0x58, 0x5f, 0x7e, 0x0b, 0xa5, 0xea, 0x55, 0x51, 0x4a, 0xd5, 0x09, 0x41, 0x35, 0xc1,
	0x31, 0xb1, 0xab, 0x4a, 0x29, 0xbe, 0x85, 0x52, 0x75, 0x37, 0x94, 0x52, 0x06, 0xa8, 0x0d, 0xf5,
	0x14, 0x33, 0xf6, 0x8e, 0x66, 0x13, 0xdb, 0x94, 0x85, 0xeb, 0x18, 0x39, 0x60, 0xaa, 0xcd, 0xdb,
	0xb5, 0x9d, 0xd2, 0xee, 0xfa, 0xfe, 0x96, 0xf3, 0xf3, 0x80, 0x1c, 0xb1, 0x4e, 0x67, 0xc8, 0x31,
	0x27, 0xbe, 0x56, 0xa1, 0xe7, 0xb0, 0x11, 0x61, 0xc6, 0x2f, 0x0b, 0xc0, 0x65, 0x30, 0xc3, 0xc9,
	0x94, 0xd8, 0xf5, 0x9d, 0xd2, 0x6e, 0x73, 0xbf, 0xed, 0xa8, 0xa1, 0x3a, 0xc5, 0xd4, 0x9d, 0x51,
	0x31, 0xf5, 0xc3, 0xca, 0xd7, 0x6e, 0xc5, 0x47, 0xc2, 0x7c, 0xa6, 0xbd, 0x3d, 0x69, 0x45, 0x8f,
	0x01, 0x24, 0x52, 0xed, 0xb1, 0xb1, 0x1c, 0xa8, 0x21, 0x2c, 0xc7, 0x72, 0x10, 0x4f, 0xa0, 0x29,
	0xfd, 0x38, 0xe0, 0xe1, 0x9c, 0xd8, 0xb0, 0x1c, 0x40, 0xf6, 0xec, 0x4a, 0x0b, 0x7a, 0x08, 0xb5,
	0x20, 0x23, 0x98, 0x93, 0x89, 0xdd, 0x5c, 0xce, 0x5d, 0xe8, 0x85, 0x35, 0x4f, 0x27, 0xd2, 0xba,
	0xba, 0xa4, 0x55, 0xeb, 0xd1, 0x23, 0x68, 0xc8, 0x25, 0x4b, 0xf3, 0xda, 0x92, 0xdb, 0xbe, 0x76,
	0xa0, 0x03, 0xa8, 0xc7, 0x84, 0xe3, 0x09, 0xe6, 0xd8, 0x5e, 0x97, 0xee, 0xd6, 0x1f, 0xee, 0xa1,
	0xbc, 0x11, 0xfe, 0xb5, 0xb0, 0x93, 0x82, 0x21, 0xcf, 0x13, 0xad, 0x42, 0x7d, 0x70, 0xd2, 0xed,
	0x8d, 0x06, 0x17, 0x9e, 0xb5, 0x82, 0x00, 0xcc, 0xe1, 0xa8, 0x7b, 0xe4, 0xf5, 0xad, 0x12, 0x6a,
	0x42, 0xed, 0xcc, 0x3b, 0xe9, 0x0f, 0x4e, 0x8e, 0xac, 0xb2, 0x28, 0x68, 0x51, 0x45, 0x58, 0x7c,
	0xaf, 0x77, 0x7a, 0xe1, 0xf9, 0x2f, 0xad, 0xaa, 0xa8, 0x1c, 0x9f, 0xf6, 0x9e, 0x79, 0x7d, 0xcb,
	0x40, 0x6b, 0xd0, 0x18, 0x9e, 0x0f, 0x85, 0xcb, 0xeb, 0x5b, 0xa6, 0x20, 0x78, 0x2f, 0xce, 0x06,
	0xbe, 0xd7, 0xb7, 0x6a, 0xfb, 0x5f, 0xca, 0xea, 0x3a, 0x0d, 0x49, 0x36, 0x0f, 0x03, 0x82, 0x8e,
	0xc1, 0xec, 0xc9, 0xd9, 0xa1, 0xd6, 0xe2, 0xaf, 0xa6, 0x6f, 0x5c, 0x1b, 0x15, 0xfb, 0xc8, 0xd2,
	0x40, 0xfe, 0x7e, 0x39, 0xeb, 0x6c, 0x7e, 0xfc, 0xfc, 0xed, 0x53, 0xf9, 0x56, 0x67, 0x4d, 0x3e,
	0x06, 0xf3, 0x3d, 0xf9, 0xa6, 0x30, 0x74, 0x0a, 0x95, 0x23, 0xc2, 0x6f, 0x46, 0x59, 0x8b, 0x85,
	0xce, 0x1d, 0x09, 0x6a, 0xa1, 0xcd, 0xdf, 0x40, 0xee, 0x7b, 0x79, 0x3b, 0x3f, 0xa0, 0x11, 0x98,
	0xe7, 0xf2, 0x7c, 0xfe, 0x6d, 0x79, 0x9a, 0xda, 0xbe, 0x99, 0xda, 0x27, 0x11, 0xf9, 0x4f, 0xea,
	0xdd, 0xbf, 0x53, 0x0f, 0x1f, 0xbc, 0xba, 0x3f, 0x0d, 0xf9, 0x2c, 0x1f, 0x3b, 0x01, 0x8d, 0xdd,
	0x37, 0x74, 0x96, 0x8c, 0x49, 0x14, 0xd1, 0x84, 0xb8, 0xba, 0xc7, 0x3d, 0xa6, 0x86, 0xae, 0x1e,
	0xc1, 0x22, 0x3b, 0x36, 0x65, 0x78, 0xf0, 0x23, 0x00, 0x00, 0xff, 0xff, 0x52, 0xb7, 0x3c, 0x2e,
	0xa0, 0x05, 0x00, 0x00,
}
