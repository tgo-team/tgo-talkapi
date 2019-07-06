// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contacts.proto

//包名，通过protoc生成时go文件时

package contacts

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

type ContactsVo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OpenId               string   `protobuf:"bytes,2,opt,name=openId,proto3" json:"openId,omitempty"`
	TalkId               uint64   `protobuf:"varint,3,opt,name=talkId,proto3" json:"talkId,omitempty"`
	RelationOpenId       string   `protobuf:"bytes,4,opt,name=relationOpenId,proto3" json:"relationOpenId,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	UpdatedAt            string   `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	CreatedAt            string   `protobuf:"bytes,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	RelationNickname     string   `protobuf:"bytes,9,opt,name=relationNickname,proto3" json:"relationNickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactsVo) Reset()         { *m = ContactsVo{} }
func (m *ContactsVo) String() string { return proto.CompactTextString(m) }
func (*ContactsVo) ProtoMessage()    {}
func (*ContactsVo) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{0}
}

func (m *ContactsVo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactsVo.Unmarshal(m, b)
}
func (m *ContactsVo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactsVo.Marshal(b, m, deterministic)
}
func (m *ContactsVo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactsVo.Merge(m, src)
}
func (m *ContactsVo) XXX_Size() int {
	return xxx_messageInfo_ContactsVo.Size(m)
}
func (m *ContactsVo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactsVo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactsVo proto.InternalMessageInfo

func (m *ContactsVo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ContactsVo) GetOpenId() string {
	if m != nil {
		return m.OpenId
	}
	return ""
}

func (m *ContactsVo) GetTalkId() uint64 {
	if m != nil {
		return m.TalkId
	}
	return 0
}

func (m *ContactsVo) GetRelationOpenId() string {
	if m != nil {
		return m.RelationOpenId
	}
	return ""
}

func (m *ContactsVo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ContactsVo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ContactsVo) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

func (m *ContactsVo) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *ContactsVo) GetRelationNickname() string {
	if m != nil {
		return m.RelationNickname
	}
	return ""
}

//同步联系人请求
type SyncContactsReq struct {
	SyncKey              string   `protobuf:"bytes,1,opt,name=syncKey,proto3" json:"syncKey,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncContactsReq) Reset()         { *m = SyncContactsReq{} }
func (m *SyncContactsReq) String() string { return proto.CompactTextString(m) }
func (*SyncContactsReq) ProtoMessage()    {}
func (*SyncContactsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{1}
}

func (m *SyncContactsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncContactsReq.Unmarshal(m, b)
}
func (m *SyncContactsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncContactsReq.Marshal(b, m, deterministic)
}
func (m *SyncContactsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncContactsReq.Merge(m, src)
}
func (m *SyncContactsReq) XXX_Size() int {
	return xxx_messageInfo_SyncContactsReq.Size(m)
}
func (m *SyncContactsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncContactsReq.DiscardUnknown(m)
}

var xxx_messageInfo_SyncContactsReq proto.InternalMessageInfo

func (m *SyncContactsReq) GetSyncKey() string {
	if m != nil {
		return m.SyncKey
	}
	return ""
}

func (m *SyncContactsReq) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type SyncContactsResp struct {
	SyncKey              string        `protobuf:"bytes,1,opt,name=syncKey,proto3" json:"syncKey,omitempty"`
	Contacts             []*ContactsVo `protobuf:"bytes,2,rep,name=contacts,proto3" json:"contacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SyncContactsResp) Reset()         { *m = SyncContactsResp{} }
func (m *SyncContactsResp) String() string { return proto.CompactTextString(m) }
func (*SyncContactsResp) ProtoMessage()    {}
func (*SyncContactsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72e48e1bf84d56b8, []int{2}
}

func (m *SyncContactsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncContactsResp.Unmarshal(m, b)
}
func (m *SyncContactsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncContactsResp.Marshal(b, m, deterministic)
}
func (m *SyncContactsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncContactsResp.Merge(m, src)
}
func (m *SyncContactsResp) XXX_Size() int {
	return xxx_messageInfo_SyncContactsResp.Size(m)
}
func (m *SyncContactsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncContactsResp.DiscardUnknown(m)
}

var xxx_messageInfo_SyncContactsResp proto.InternalMessageInfo

func (m *SyncContactsResp) GetSyncKey() string {
	if m != nil {
		return m.SyncKey
	}
	return ""
}

func (m *SyncContactsResp) GetContacts() []*ContactsVo {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func init() {
	proto.RegisterType((*ContactsVo)(nil), "contacts.ContactsVo")
	proto.RegisterType((*SyncContactsReq)(nil), "contacts.SyncContactsReq")
	proto.RegisterType((*SyncContactsResp)(nil), "contacts.SyncContactsResp")
}

func init() { proto.RegisterFile("contacts.proto", fileDescriptor_72e48e1bf84d56b8) }

var fileDescriptor_72e48e1bf84d56b8 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xe9, 0xdf, 0x99, 0xde, 0x0f, 0xfa, 0x0d, 0x97, 0x41, 0xb2, 0x70, 0x51, 0xba, 0x90,
	0xe2, 0x62, 0x10, 0x7d, 0x82, 0xc1, 0xd5, 0x20, 0x28, 0x44, 0x70, 0x29, 0xc4, 0x24, 0x8b, 0xd0,
	0x36, 0xa9, 0x6d, 0x5c, 0xf4, 0x0d, 0x7c, 0x6c, 0x69, 0xd2, 0x76, 0x50, 0xc1, 0x5d, 0x7e, 0xe7,
	0xe4, 0xde, 0x0b, 0xe7, 0x40, 0xce, 0x8d, 0xb6, 0x8c, 0xdb, 0xe1, 0xd0, 0xf5, 0xc6, 0x1a, 0xdc,
	0x2e, 0x5c, 0x7e, 0x86, 0x00, 0xf7, 0x33, 0xbc, 0x18, 0xcc, 0x21, 0x54, 0x82, 0x04, 0x45, 0x50,
	0x45, 0x34, 0x54, 0x02, 0x2f, 0x20, 0x35, 0x9d, 0xd4, 0x27, 0x41, 0xc2, 0x22, 0xa8, 0x32, 0x3a,
	0xd3, 0xa4, 0x5b, 0xd6, 0xd4, 0x27, 0x41, 0xa2, 0x22, 0xa8, 0x62, 0x3a, 0x13, 0x5e, 0x41, 0xde,
	0xcb, 0x86, 0x59, 0x65, 0xf4, 0x93, 0x9f, 0x8b, 0xdd, 0xdc, 0x0f, 0x15, 0x11, 0x62, 0x3b, 0x76,
	0x92, 0x24, 0x45, 0x50, 0x25, 0xd4, 0xbd, 0xa7, 0x9d, 0xbd, 0x6c, 0x59, 0x5f, 0x93, 0xd4, 0xdf,
	0xf2, 0x84, 0x97, 0x90, 0x7d, 0x74, 0x82, 0x59, 0x29, 0x8e, 0x96, 0x6c, 0x9c, 0x75, 0x16, 0x26,
	0x97, 0xf7, 0x72, 0x76, 0xb7, 0xde, 0x5d, 0x05, 0xbc, 0x86, 0xdd, 0x72, 0xf9, 0x51, 0xf1, 0x5a,
	0xb3, 0x56, 0x92, 0xcc, 0x7d, 0xfa, 0xa5, 0x97, 0x47, 0xf8, 0xff, 0x3c, 0x6a, 0xbe, 0xa4, 0x41,
	0xe5, 0x3b, 0x12, 0xd8, 0x0c, 0xa3, 0xe6, 0x0f, 0x72, 0x74, 0x99, 0x64, 0x74, 0x41, 0xdc, 0x43,
	0xd2, 0xa8, 0x56, 0x59, 0x97, 0x4b, 0x4c, 0x3d, 0x94, 0xaf, 0xb0, 0xfb, 0xbe, 0x62, 0xe8, 0xfe,
	0xd8, 0x71, 0x03, 0x6b, 0x0f, 0x24, 0x2c, 0xa2, 0xea, 0xdf, 0xed, 0xfe, 0xb0, 0x16, 0x75, 0x2e,
	0x85, 0xae, 0xbf, 0xde, 0x52, 0x57, 0xdf, 0xdd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x33,
	0x4c, 0xab, 0xd0, 0x01, 0x00, 0x00,
}