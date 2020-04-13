// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friendship.proto

package puppet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FriendshipType int32

const (
	FriendshipType_FRIENDSHIP_TYPE_UNSPECIFIED FriendshipType = 0
	FriendshipType_FRIENDSHIP_TYPE_CONFIRM     FriendshipType = 1
	FriendshipType_FRIENDSHIP_TYPE_RECEIVE     FriendshipType = 2
	FriendshipType_FRIENDSHIP_TYPE_VERIFY      FriendshipType = 3
)

var FriendshipType_name = map[int32]string{
	0: "FRIENDSHIP_TYPE_UNSPECIFIED",
	1: "FRIENDSHIP_TYPE_CONFIRM",
	2: "FRIENDSHIP_TYPE_RECEIVE",
	3: "FRIENDSHIP_TYPE_VERIFY",
}

var FriendshipType_value = map[string]int32{
	"FRIENDSHIP_TYPE_UNSPECIFIED": 0,
	"FRIENDSHIP_TYPE_CONFIRM":     1,
	"FRIENDSHIP_TYPE_RECEIVE":     2,
	"FRIENDSHIP_TYPE_VERIFY":      3,
}

func (x FriendshipType) String() string {
	return proto.EnumName(FriendshipType_name, int32(x))
}

func (FriendshipType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{0}
}

type FriendshipSceneType int32

const (
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_UNSPECIFIED FriendshipSceneType = 0
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_QQ          FriendshipSceneType = 1
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_EMAIL       FriendshipSceneType = 2
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_WEIXIN      FriendshipSceneType = 3
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_QQTBD       FriendshipSceneType = 12
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_ROOM        FriendshipSceneType = 14
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_PHONE       FriendshipSceneType = 15
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_CARD        FriendshipSceneType = 17
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_LOCATION    FriendshipSceneType = 18
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_BOTTLE      FriendshipSceneType = 25
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_SHAKING     FriendshipSceneType = 29
	FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_QRCODE      FriendshipSceneType = 30
)

var FriendshipSceneType_name = map[int32]string{
	0:  "FRIENDSHIP_SCENE_TYPE_UNSPECIFIED",
	1:  "FRIENDSHIP_SCENE_TYPE_QQ",
	2:  "FRIENDSHIP_SCENE_TYPE_EMAIL",
	3:  "FRIENDSHIP_SCENE_TYPE_WEIXIN",
	12: "FRIENDSHIP_SCENE_TYPE_QQTBD",
	14: "FRIENDSHIP_SCENE_TYPE_ROOM",
	15: "FRIENDSHIP_SCENE_TYPE_PHONE",
	17: "FRIENDSHIP_SCENE_TYPE_CARD",
	18: "FRIENDSHIP_SCENE_TYPE_LOCATION",
	25: "FRIENDSHIP_SCENE_TYPE_BOTTLE",
	29: "FRIENDSHIP_SCENE_TYPE_SHAKING",
	30: "FRIENDSHIP_SCENE_TYPE_QRCODE",
}

var FriendshipSceneType_value = map[string]int32{
	"FRIENDSHIP_SCENE_TYPE_UNSPECIFIED": 0,
	"FRIENDSHIP_SCENE_TYPE_QQ":          1,
	"FRIENDSHIP_SCENE_TYPE_EMAIL":       2,
	"FRIENDSHIP_SCENE_TYPE_WEIXIN":      3,
	"FRIENDSHIP_SCENE_TYPE_QQTBD":       12,
	"FRIENDSHIP_SCENE_TYPE_ROOM":        14,
	"FRIENDSHIP_SCENE_TYPE_PHONE":       15,
	"FRIENDSHIP_SCENE_TYPE_CARD":        17,
	"FRIENDSHIP_SCENE_TYPE_LOCATION":    18,
	"FRIENDSHIP_SCENE_TYPE_BOTTLE":      25,
	"FRIENDSHIP_SCENE_TYPE_SHAKING":     29,
	"FRIENDSHIP_SCENE_TYPE_QRCODE":      30,
}

func (x FriendshipSceneType) String() string {
	return proto.EnumName(FriendshipSceneType_name, int32(x))
}

func (FriendshipSceneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{1}
}

type FriendshipPayloadRequest struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Payload              *wrappers.StringValue `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FriendshipPayloadRequest) Reset()         { *m = FriendshipPayloadRequest{} }
func (m *FriendshipPayloadRequest) String() string { return proto.CompactTextString(m) }
func (*FriendshipPayloadRequest) ProtoMessage()    {}
func (*FriendshipPayloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{0}
}

func (m *FriendshipPayloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipPayloadRequest.Unmarshal(m, b)
}
func (m *FriendshipPayloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipPayloadRequest.Marshal(b, m, deterministic)
}
func (m *FriendshipPayloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipPayloadRequest.Merge(m, src)
}
func (m *FriendshipPayloadRequest) XXX_Size() int {
	return xxx_messageInfo_FriendshipPayloadRequest.Size(m)
}
func (m *FriendshipPayloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipPayloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipPayloadRequest proto.InternalMessageInfo

func (m *FriendshipPayloadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FriendshipPayloadRequest) GetPayload() *wrappers.StringValue {
	if m != nil {
		return m.Payload
	}
	return nil
}

type FriendshipPayloadResponse struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContactId            string              `protobuf:"bytes,2,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	Hello                string              `protobuf:"bytes,3,opt,name=hello,proto3" json:"hello,omitempty"`
	Type                 FriendshipType      `protobuf:"varint,4,opt,name=type,proto3,enum=wechaty.puppet.FriendshipType" json:"type,omitempty"`
	Stranger             string              `protobuf:"bytes,5,opt,name=stranger,proto3" json:"stranger,omitempty"`
	Ticket               string              `protobuf:"bytes,6,opt,name=ticket,proto3" json:"ticket,omitempty"`
	Scene                FriendshipSceneType `protobuf:"varint,7,opt,name=scene,proto3,enum=wechaty.puppet.FriendshipSceneType" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *FriendshipPayloadResponse) Reset()         { *m = FriendshipPayloadResponse{} }
func (m *FriendshipPayloadResponse) String() string { return proto.CompactTextString(m) }
func (*FriendshipPayloadResponse) ProtoMessage()    {}
func (*FriendshipPayloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{1}
}

func (m *FriendshipPayloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipPayloadResponse.Unmarshal(m, b)
}
func (m *FriendshipPayloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipPayloadResponse.Marshal(b, m, deterministic)
}
func (m *FriendshipPayloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipPayloadResponse.Merge(m, src)
}
func (m *FriendshipPayloadResponse) XXX_Size() int {
	return xxx_messageInfo_FriendshipPayloadResponse.Size(m)
}
func (m *FriendshipPayloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipPayloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipPayloadResponse proto.InternalMessageInfo

func (m *FriendshipPayloadResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FriendshipPayloadResponse) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *FriendshipPayloadResponse) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

func (m *FriendshipPayloadResponse) GetType() FriendshipType {
	if m != nil {
		return m.Type
	}
	return FriendshipType_FRIENDSHIP_TYPE_UNSPECIFIED
}

func (m *FriendshipPayloadResponse) GetStranger() string {
	if m != nil {
		return m.Stranger
	}
	return ""
}

func (m *FriendshipPayloadResponse) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *FriendshipPayloadResponse) GetScene() FriendshipSceneType {
	if m != nil {
		return m.Scene
	}
	return FriendshipSceneType_FRIENDSHIP_SCENE_TYPE_UNSPECIFIED
}

type FriendshipSearchPhoneRequest struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipSearchPhoneRequest) Reset()         { *m = FriendshipSearchPhoneRequest{} }
func (m *FriendshipSearchPhoneRequest) String() string { return proto.CompactTextString(m) }
func (*FriendshipSearchPhoneRequest) ProtoMessage()    {}
func (*FriendshipSearchPhoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{2}
}

func (m *FriendshipSearchPhoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipSearchPhoneRequest.Unmarshal(m, b)
}
func (m *FriendshipSearchPhoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipSearchPhoneRequest.Marshal(b, m, deterministic)
}
func (m *FriendshipSearchPhoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipSearchPhoneRequest.Merge(m, src)
}
func (m *FriendshipSearchPhoneRequest) XXX_Size() int {
	return xxx_messageInfo_FriendshipSearchPhoneRequest.Size(m)
}
func (m *FriendshipSearchPhoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipSearchPhoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipSearchPhoneRequest proto.InternalMessageInfo

func (m *FriendshipSearchPhoneRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type FriendshipSearchPhoneResponse struct {
	// nullable
	ContactId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FriendshipSearchPhoneResponse) Reset()         { *m = FriendshipSearchPhoneResponse{} }
func (m *FriendshipSearchPhoneResponse) String() string { return proto.CompactTextString(m) }
func (*FriendshipSearchPhoneResponse) ProtoMessage()    {}
func (*FriendshipSearchPhoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{3}
}

func (m *FriendshipSearchPhoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipSearchPhoneResponse.Unmarshal(m, b)
}
func (m *FriendshipSearchPhoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipSearchPhoneResponse.Marshal(b, m, deterministic)
}
func (m *FriendshipSearchPhoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipSearchPhoneResponse.Merge(m, src)
}
func (m *FriendshipSearchPhoneResponse) XXX_Size() int {
	return xxx_messageInfo_FriendshipSearchPhoneResponse.Size(m)
}
func (m *FriendshipSearchPhoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipSearchPhoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipSearchPhoneResponse proto.InternalMessageInfo

func (m *FriendshipSearchPhoneResponse) GetContactId() *wrappers.StringValue {
	if m != nil {
		return m.ContactId
	}
	return nil
}

type FriendshipSearchWeixinRequest struct {
	Weixin               string   `protobuf:"bytes,1,opt,name=weixin,proto3" json:"weixin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipSearchWeixinRequest) Reset()         { *m = FriendshipSearchWeixinRequest{} }
func (m *FriendshipSearchWeixinRequest) String() string { return proto.CompactTextString(m) }
func (*FriendshipSearchWeixinRequest) ProtoMessage()    {}
func (*FriendshipSearchWeixinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{4}
}

func (m *FriendshipSearchWeixinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipSearchWeixinRequest.Unmarshal(m, b)
}
func (m *FriendshipSearchWeixinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipSearchWeixinRequest.Marshal(b, m, deterministic)
}
func (m *FriendshipSearchWeixinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipSearchWeixinRequest.Merge(m, src)
}
func (m *FriendshipSearchWeixinRequest) XXX_Size() int {
	return xxx_messageInfo_FriendshipSearchWeixinRequest.Size(m)
}
func (m *FriendshipSearchWeixinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipSearchWeixinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipSearchWeixinRequest proto.InternalMessageInfo

func (m *FriendshipSearchWeixinRequest) GetWeixin() string {
	if m != nil {
		return m.Weixin
	}
	return ""
}

type FriendshipSearchWeixinResponse struct {
	// nullable
	ContactId            *wrappers.StringValue `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FriendshipSearchWeixinResponse) Reset()         { *m = FriendshipSearchWeixinResponse{} }
func (m *FriendshipSearchWeixinResponse) String() string { return proto.CompactTextString(m) }
func (*FriendshipSearchWeixinResponse) ProtoMessage()    {}
func (*FriendshipSearchWeixinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{5}
}

func (m *FriendshipSearchWeixinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipSearchWeixinResponse.Unmarshal(m, b)
}
func (m *FriendshipSearchWeixinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipSearchWeixinResponse.Marshal(b, m, deterministic)
}
func (m *FriendshipSearchWeixinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipSearchWeixinResponse.Merge(m, src)
}
func (m *FriendshipSearchWeixinResponse) XXX_Size() int {
	return xxx_messageInfo_FriendshipSearchWeixinResponse.Size(m)
}
func (m *FriendshipSearchWeixinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipSearchWeixinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipSearchWeixinResponse proto.InternalMessageInfo

func (m *FriendshipSearchWeixinResponse) GetContactId() *wrappers.StringValue {
	if m != nil {
		return m.ContactId
	}
	return nil
}

type FriendshipAddRequest struct {
	ContactId            string   `protobuf:"bytes,1,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	Hello                string   `protobuf:"bytes,2,opt,name=hello,proto3" json:"hello,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipAddRequest) Reset()         { *m = FriendshipAddRequest{} }
func (m *FriendshipAddRequest) String() string { return proto.CompactTextString(m) }
func (*FriendshipAddRequest) ProtoMessage()    {}
func (*FriendshipAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{6}
}

func (m *FriendshipAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipAddRequest.Unmarshal(m, b)
}
func (m *FriendshipAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipAddRequest.Marshal(b, m, deterministic)
}
func (m *FriendshipAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipAddRequest.Merge(m, src)
}
func (m *FriendshipAddRequest) XXX_Size() int {
	return xxx_messageInfo_FriendshipAddRequest.Size(m)
}
func (m *FriendshipAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipAddRequest proto.InternalMessageInfo

func (m *FriendshipAddRequest) GetContactId() string {
	if m != nil {
		return m.ContactId
	}
	return ""
}

func (m *FriendshipAddRequest) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

type FriendshipAddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipAddResponse) Reset()         { *m = FriendshipAddResponse{} }
func (m *FriendshipAddResponse) String() string { return proto.CompactTextString(m) }
func (*FriendshipAddResponse) ProtoMessage()    {}
func (*FriendshipAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{7}
}

func (m *FriendshipAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipAddResponse.Unmarshal(m, b)
}
func (m *FriendshipAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipAddResponse.Marshal(b, m, deterministic)
}
func (m *FriendshipAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipAddResponse.Merge(m, src)
}
func (m *FriendshipAddResponse) XXX_Size() int {
	return xxx_messageInfo_FriendshipAddResponse.Size(m)
}
func (m *FriendshipAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipAddResponse proto.InternalMessageInfo

type FriendshipAcceptRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipAcceptRequest) Reset()         { *m = FriendshipAcceptRequest{} }
func (m *FriendshipAcceptRequest) String() string { return proto.CompactTextString(m) }
func (*FriendshipAcceptRequest) ProtoMessage()    {}
func (*FriendshipAcceptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{8}
}

func (m *FriendshipAcceptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipAcceptRequest.Unmarshal(m, b)
}
func (m *FriendshipAcceptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipAcceptRequest.Marshal(b, m, deterministic)
}
func (m *FriendshipAcceptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipAcceptRequest.Merge(m, src)
}
func (m *FriendshipAcceptRequest) XXX_Size() int {
	return xxx_messageInfo_FriendshipAcceptRequest.Size(m)
}
func (m *FriendshipAcceptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipAcceptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipAcceptRequest proto.InternalMessageInfo

func (m *FriendshipAcceptRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FriendshipAcceptResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendshipAcceptResponse) Reset()         { *m = FriendshipAcceptResponse{} }
func (m *FriendshipAcceptResponse) String() string { return proto.CompactTextString(m) }
func (*FriendshipAcceptResponse) ProtoMessage()    {}
func (*FriendshipAcceptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4164b89dcff7420, []int{9}
}

func (m *FriendshipAcceptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendshipAcceptResponse.Unmarshal(m, b)
}
func (m *FriendshipAcceptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendshipAcceptResponse.Marshal(b, m, deterministic)
}
func (m *FriendshipAcceptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendshipAcceptResponse.Merge(m, src)
}
func (m *FriendshipAcceptResponse) XXX_Size() int {
	return xxx_messageInfo_FriendshipAcceptResponse.Size(m)
}
func (m *FriendshipAcceptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendshipAcceptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FriendshipAcceptResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("wechaty.puppet.FriendshipType", FriendshipType_name, FriendshipType_value)
	proto.RegisterEnum("wechaty.puppet.FriendshipSceneType", FriendshipSceneType_name, FriendshipSceneType_value)
	proto.RegisterType((*FriendshipPayloadRequest)(nil), "wechaty.puppet.FriendshipPayloadRequest")
	proto.RegisterType((*FriendshipPayloadResponse)(nil), "wechaty.puppet.FriendshipPayloadResponse")
	proto.RegisterType((*FriendshipSearchPhoneRequest)(nil), "wechaty.puppet.FriendshipSearchPhoneRequest")
	proto.RegisterType((*FriendshipSearchPhoneResponse)(nil), "wechaty.puppet.FriendshipSearchPhoneResponse")
	proto.RegisterType((*FriendshipSearchWeixinRequest)(nil), "wechaty.puppet.FriendshipSearchWeixinRequest")
	proto.RegisterType((*FriendshipSearchWeixinResponse)(nil), "wechaty.puppet.FriendshipSearchWeixinResponse")
	proto.RegisterType((*FriendshipAddRequest)(nil), "wechaty.puppet.FriendshipAddRequest")
	proto.RegisterType((*FriendshipAddResponse)(nil), "wechaty.puppet.FriendshipAddResponse")
	proto.RegisterType((*FriendshipAcceptRequest)(nil), "wechaty.puppet.FriendshipAcceptRequest")
	proto.RegisterType((*FriendshipAcceptResponse)(nil), "wechaty.puppet.FriendshipAcceptResponse")
}

func init() { proto.RegisterFile("friendship.proto", fileDescriptor_d4164b89dcff7420) }

var fileDescriptor_d4164b89dcff7420 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x51, 0x4f, 0xd3, 0x50,
	0x14, 0xc7, 0x6d, 0xc7, 0x86, 0x1c, 0xcd, 0xbc, 0x5e, 0x11, 0xca, 0xd8, 0xe6, 0xa8, 0x31, 0x99,
	0x3c, 0xb4, 0x09, 0x1a, 0x8d, 0xf1, 0x69, 0x74, 0x77, 0xd2, 0x00, 0xed, 0xd6, 0x4d, 0x10, 0xa3,
	0x21, 0x5d, 0x77, 0xe9, 0x1a, 0x67, 0x7b, 0x6d, 0xef, 0x82, 0x7b, 0xf6, 0xc1, 0xef, 0xe2, 0xa7,
	0x34, 0xb4, 0x1d, 0x8c, 0xd1, 0x26, 0x3e, 0xf8, 0x78, 0xfa, 0xff, 0x9f, 0xdf, 0x39, 0xf7, 0xf4,
	0xdc, 0x0b, 0xe8, 0x22, 0xf4, 0xa8, 0x3f, 0x8a, 0xc6, 0x1e, 0x53, 0x58, 0x18, 0xf0, 0x00, 0x97,
	0x2f, 0xa9, 0x33, 0xb6, 0xf9, 0x4c, 0x61, 0x53, 0xc6, 0x28, 0xaf, 0xd4, 0xdd, 0x20, 0x70, 0x27,
	0x54, 0x8d, 0xd5, 0xe1, 0xf4, 0x42, 0xbd, 0x0c, 0x6d, 0xc6, 0x68, 0x18, 0x25, 0x7e, 0x79, 0x08,
	0x52, 0xe7, 0x9a, 0xd1, 0xb5, 0x67, 0x93, 0xc0, 0x1e, 0x59, 0xf4, 0xc7, 0x94, 0x46, 0x1c, 0x97,
	0x41, 0xf4, 0x46, 0x92, 0xd0, 0x10, 0x9a, 0x6b, 0x96, 0xe8, 0x8d, 0xf0, 0x1b, 0x58, 0x65, 0x89,
	0x43, 0x12, 0x1b, 0x42, 0xf3, 0xc1, 0x5e, 0x55, 0x49, 0xe8, 0xca, 0x9c, 0xae, 0xf4, 0x79, 0xe8,
	0xf9, 0xee, 0x89, 0x3d, 0x99, 0x52, 0x6b, 0x6e, 0x96, 0x7f, 0x89, 0xb0, 0x95, 0x51, 0x24, 0x62,
	0x81, 0x1f, 0xd1, 0x3b, 0x55, 0x6a, 0x00, 0x4e, 0xe0, 0x73, 0xdb, 0xe1, 0xe7, 0x5e, 0x52, 0x68,
	0xcd, 0x5a, 0x4b, 0xbf, 0xe8, 0x23, 0xbc, 0x0e, 0xc5, 0x31, 0x9d, 0x4c, 0x02, 0xa9, 0x10, 0x2b,
	0x49, 0x80, 0xf7, 0x60, 0x85, 0xcf, 0x18, 0x95, 0x56, 0x1a, 0x42, 0xb3, 0xbc, 0x57, 0x57, 0x6e,
	0x4f, 0x41, 0xb9, 0xa9, 0x3e, 0x98, 0x31, 0x6a, 0xc5, 0x5e, 0x5c, 0x81, 0xfb, 0x11, 0x0f, 0x6d,
	0xdf, 0xa5, 0xa1, 0x54, 0x8c, 0x61, 0xd7, 0x31, 0xde, 0x80, 0x12, 0xf7, 0x9c, 0x6f, 0x94, 0x4b,
	0xa5, 0x58, 0x49, 0x23, 0xfc, 0x0e, 0x8a, 0x91, 0x43, 0x7d, 0x2a, 0xad, 0xc6, 0x85, 0x9e, 0xe7,
	0x17, 0xea, 0x5f, 0xd9, 0xe2, 0x6a, 0x49, 0x86, 0xfc, 0x1a, 0xaa, 0x0b, 0x2a, 0xb5, 0x43, 0x67,
	0xdc, 0x1d, 0x07, 0x3e, 0x9d, 0x4f, 0x7b, 0x1d, 0x8a, 0xec, 0x2a, 0x4e, 0x47, 0x91, 0x04, 0xf2,
	0x17, 0xa8, 0xe5, 0x64, 0xa5, 0xe3, 0x7b, 0x7f, 0x6b, 0x5c, 0xc2, 0x3f, 0xfc, 0x97, 0x9b, 0x61,
	0xca, 0x6f, 0xef, 0xd2, 0x4f, 0xa9, 0xf7, 0xd3, 0xf3, 0xe7, 0x4d, 0x6d, 0x40, 0xe9, 0x32, 0xfe,
	0x90, 0x76, 0x95, 0x46, 0xf2, 0x57, 0xa8, 0xe7, 0x25, 0xfe, 0x8f, 0xbe, 0x0e, 0x61, 0xfd, 0x06,
	0xdf, 0x1a, 0x5d, 0x6f, 0x64, 0xed, 0x0e, 0x34, 0x7b, 0x37, 0xc4, 0x85, 0xdd, 0x90, 0x37, 0xe1,
	0xe9, 0x12, 0x2c, 0x69, 0x51, 0x7e, 0x09, 0x9b, 0x0b, 0x82, 0xe3, 0x50, 0xc6, 0x73, 0x56, 0x5f,
	0xae, 0x2c, 0x5e, 0x93, 0xb9, 0x35, 0xc1, 0xec, 0xfe, 0x16, 0xa0, 0x7c, 0x7b, 0xc1, 0xf0, 0x33,
	0xd8, 0xee, 0x58, 0x3a, 0x31, 0xda, 0xfd, 0x03, 0xbd, 0x7b, 0x3e, 0x38, 0xeb, 0x92, 0xf3, 0x8f,
	0x46, 0xbf, 0x4b, 0x34, 0xbd, 0xa3, 0x93, 0x36, 0xba, 0x87, 0xb7, 0x61, 0x73, 0xd9, 0xa0, 0x99,
	0x46, 0x47, 0xb7, 0x8e, 0x91, 0x90, 0x25, 0x5a, 0x44, 0x23, 0xfa, 0x09, 0x41, 0x22, 0xae, 0xc0,
	0xc6, 0xb2, 0x78, 0x42, 0x2c, 0xbd, 0x73, 0x86, 0x0a, 0xbb, 0x7f, 0x0a, 0xf0, 0x24, 0x63, 0x03,
	0xf1, 0x0b, 0xd8, 0x59, 0xc8, 0xe9, 0x6b, 0xc4, 0x20, 0x59, 0x4d, 0x55, 0x41, 0xca, 0xb6, 0xf5,
	0x7a, 0x48, 0x58, 0x3a, 0xd3, 0x82, 0x4a, 0x8e, 0x5b, 0xfa, 0x11, 0x12, 0x71, 0x03, 0xaa, 0xd9,
	0x86, 0x53, 0xa2, 0x7f, 0xd2, 0x0d, 0x54, 0xc8, 0x47, 0xf4, 0x7a, 0x83, 0xfd, 0x36, 0x7a, 0x88,
	0xeb, 0x50, 0xc9, 0x36, 0x58, 0xa6, 0x79, 0x8c, 0xca, 0xf9, 0x80, 0xee, 0x81, 0x69, 0x10, 0xf4,
	0x28, 0x1f, 0xa0, 0xb5, 0xac, 0x36, 0x7a, 0x8c, 0x65, 0xa8, 0x67, 0xeb, 0x47, 0xa6, 0xd6, 0x1a,
	0xe8, 0xa6, 0x81, 0x70, 0xfe, 0x39, 0xf6, 0xcd, 0xc1, 0xe0, 0x88, 0xa0, 0x2d, 0xbc, 0x03, 0xb5,
	0x6c, 0x47, 0xff, 0xa0, 0x75, 0xa8, 0x1b, 0x1f, 0x50, 0x2d, 0x1f, 0xd2, 0xb3, 0x34, 0xb3, 0x4d,
	0x50, 0x7d, 0x7f, 0xf7, 0x73, 0xd3, 0xf5, 0xf8, 0x78, 0x3a, 0x54, 0x9c, 0xe0, 0xbb, 0x7a, 0xf5,
	0x8a, 0x78, 0x54, 0x75, 0x43, 0xe6, 0xa8, 0x6e, 0xa0, 0xa6, 0xcf, 0x8a, 0x9a, 0x3c, 0x2b, 0xc3,
	0x52, 0x7c, 0x61, 0x5e, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x62, 0x60, 0xb7, 0x24, 0xf0, 0x05,
	0x00, 0x00,
}