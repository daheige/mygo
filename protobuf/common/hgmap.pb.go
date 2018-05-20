// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hgmap.proto

package dahei

/*
protoc --go_out=. hgmap.proto
*/

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

type FOO int32

const (
	FOO_X FOO = 17
	FOO_Y FOO = 18
)

var FOO_name = map[int32]string{
	17: "X",
	18: "Y",
}
var FOO_value = map[string]int32{
	"X": 17,
	"Y": 18,
}

func (x FOO) Enum() *FOO {
	p := new(FOO)
	*p = x
	return p
}
func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (x *FOO) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FOO_value, data, "FOO")
	if err != nil {
		return err
	}
	*x = FOO(value)
	return nil
}
func (FOO) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_hgmap_763fef4e1f3fc853, []int{0}
}

type Person struct {
	Id                   *int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_hgmap_763fef4e1f3fc853, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (dst *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(dst, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type UserInfo struct {
	Id                   *int32    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name                 *string   `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Options              *int32    `protobuf:"varint,3,opt,name=options" json:"options,omitempty"`
	Type                 *int32    `protobuf:"varint,4,opt,name=type,def=77" json:"type,omitempty"`
	SubUser              []*Person `protobuf:"bytes,5,rep,name=subUser" json:"subUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_hgmap_763fef4e1f3fc853, []int{1}
}
func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (dst *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(dst, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

const Default_UserInfo_Type int32 = 77

func (m *UserInfo) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserInfo) GetOptions() int32 {
	if m != nil && m.Options != nil {
		return *m.Options
	}
	return 0
}

func (m *UserInfo) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_UserInfo_Type
}

func (m *UserInfo) GetSubUser() []*Person {
	if m != nil {
		return m.SubUser
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "dahei.Person")
	proto.RegisterType((*UserInfo)(nil), "dahei.UserInfo")
	proto.RegisterEnum("dahei.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("hgmap.proto", fileDescriptor_hgmap_763fef4e1f3fc853) }

var fileDescriptor_hgmap_763fef4e1f3fc853 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcf, 0x4d,
	0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x49, 0xcc, 0x48, 0xcd, 0x54, 0x52,
	0xe2, 0x62, 0x0b, 0x48, 0x2d, 0x2a, 0xce, 0xcf, 0x13, 0xe2, 0xe2, 0x62, 0xca, 0x4c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x15, 0xe2, 0xe1, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52, 0x60,
	0xd4, 0xe0, 0x54, 0x4a, 0xe7, 0xe2, 0x08, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x4b, 0xcb, 0x87, 0xab,
	0x62, 0x42, 0x51, 0xc5, 0xa4, 0xc1, 0x29, 0xc4, 0xcf, 0xc5, 0x9e, 0x5f, 0x50, 0x92, 0x99, 0x9f,
	0x57, 0x2c, 0xc1, 0x0c, 0x36, 0x44, 0x80, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x05, 0xc4,
	0xb3, 0x62, 0x32, 0x37, 0x17, 0x92, 0xe3, 0x62, 0x2f, 0x2e, 0x4d, 0x02, 0x99, 0x25, 0xc1, 0xaa,
	0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xab, 0x07, 0x76, 0x85, 0x1e, 0xc4, 0x09, 0x5a, 0xc2, 0x5c, 0xcc,
	0x6e, 0xfe, 0xfe, 0x42, 0xac, 0x5c, 0x8c, 0x11, 0x02, 0x82, 0x20, 0x2a, 0x52, 0x40, 0x08, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xe5, 0x20, 0x81, 0x76, 0xb6, 0x00, 0x00, 0x00,
}
