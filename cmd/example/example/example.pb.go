// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/example/example/example.proto

package example

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

type AddRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{0}
}
func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (dst *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(dst, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *AddRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type AddResponse struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{1}
}
func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (dst *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(dst, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type SubtractRequest struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractRequest) Reset()         { *m = SubtractRequest{} }
func (m *SubtractRequest) String() string { return proto.CompactTextString(m) }
func (*SubtractRequest) ProtoMessage()    {}
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{2}
}
func (m *SubtractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractRequest.Unmarshal(m, b)
}
func (m *SubtractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractRequest.Marshal(b, m, deterministic)
}
func (dst *SubtractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractRequest.Merge(dst, src)
}
func (m *SubtractRequest) XXX_Size() int {
	return xxx_messageInfo_SubtractRequest.Size(m)
}
func (m *SubtractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractRequest proto.InternalMessageInfo

func (m *SubtractRequest) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *SubtractRequest) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type SubtractResponse struct {
	C                    int32    `protobuf:"varint,1,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubtractResponse) Reset()         { *m = SubtractResponse{} }
func (m *SubtractResponse) String() string { return proto.CompactTextString(m) }
func (*SubtractResponse) ProtoMessage()    {}
func (*SubtractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{3}
}
func (m *SubtractResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubtractResponse.Unmarshal(m, b)
}
func (m *SubtractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubtractResponse.Marshal(b, m, deterministic)
}
func (dst *SubtractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubtractResponse.Merge(dst, src)
}
func (m *SubtractResponse) XXX_Size() int {
	return xxx_messageInfo_SubtractResponse.Size(m)
}
func (m *SubtractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubtractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubtractResponse proto.InternalMessageInfo

func (m *SubtractResponse) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

type GactusFile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GactusFile) Reset()         { *m = GactusFile{} }
func (m *GactusFile) String() string { return proto.CompactTextString(m) }
func (*GactusFile) ProtoMessage()    {}
func (*GactusFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{4}
}
func (m *GactusFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GactusFile.Unmarshal(m, b)
}
func (m *GactusFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GactusFile.Marshal(b, m, deterministic)
}
func (dst *GactusFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GactusFile.Merge(dst, src)
}
func (m *GactusFile) XXX_Size() int {
	return xxx_messageInfo_GactusFile.Size(m)
}
func (m *GactusFile) XXX_DiscardUnknown() {
	xxx_messageInfo_GactusFile.DiscardUnknown(m)
}

var xxx_messageInfo_GactusFile proto.InternalMessageInfo

func (m *GactusFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GactusFile) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ChangeProfileRequest struct {
	Pictures             []*GactusFile `protobuf:"bytes,1,rep,name=pictures,proto3" json:"pictures,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChangeProfileRequest) Reset()         { *m = ChangeProfileRequest{} }
func (m *ChangeProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeProfileRequest) ProtoMessage()    {}
func (*ChangeProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{5}
}
func (m *ChangeProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeProfileRequest.Unmarshal(m, b)
}
func (m *ChangeProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeProfileRequest.Marshal(b, m, deterministic)
}
func (dst *ChangeProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeProfileRequest.Merge(dst, src)
}
func (m *ChangeProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeProfileRequest.Size(m)
}
func (m *ChangeProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeProfileRequest proto.InternalMessageInfo

func (m *ChangeProfileRequest) GetPictures() []*GactusFile {
	if m != nil {
		return m.Pictures
	}
	return nil
}

type ChangeProfileResponse struct {
	FileSize             uint32   `protobuf:"varint,1,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeProfileResponse) Reset()         { *m = ChangeProfileResponse{} }
func (m *ChangeProfileResponse) String() string { return proto.CompactTextString(m) }
func (*ChangeProfileResponse) ProtoMessage()    {}
func (*ChangeProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_example_32b3ef89e4a56f9c, []int{6}
}
func (m *ChangeProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeProfileResponse.Unmarshal(m, b)
}
func (m *ChangeProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeProfileResponse.Marshal(b, m, deterministic)
}
func (dst *ChangeProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeProfileResponse.Merge(dst, src)
}
func (m *ChangeProfileResponse) XXX_Size() int {
	return xxx_messageInfo_ChangeProfileResponse.Size(m)
}
func (m *ChangeProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeProfileResponse proto.InternalMessageInfo

func (m *ChangeProfileResponse) GetFileSize() uint32 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func init() {
	proto.RegisterType((*AddRequest)(nil), "example.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "example.AddResponse")
	proto.RegisterType((*SubtractRequest)(nil), "example.SubtractRequest")
	proto.RegisterType((*SubtractResponse)(nil), "example.SubtractResponse")
	proto.RegisterType((*GactusFile)(nil), "example.GactusFile")
	proto.RegisterType((*ChangeProfileRequest)(nil), "example.ChangeProfileRequest")
	proto.RegisterType((*ChangeProfileResponse)(nil), "example.ChangeProfileResponse")
}

func init() {
	proto.RegisterFile("cmd/example/example/example.proto", fileDescriptor_example_32b3ef89e4a56f9c)
}

var fileDescriptor_example_32b3ef89e4a56f9c = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x86, 0x89, 0x5f, 0x6d, 0xa7, 0x15, 0x25, 0x2a, 0x2c, 0xf4, 0xb2, 0xe6, 0xb4, 0x17, 0x5b,
	0x50, 0x4f, 0xde, 0x44, 0xb0, 0x57, 0x49, 0x7f, 0x80, 0x64, 0xb3, 0xa3, 0x06, 0x76, 0x93, 0x75,
	0x33, 0x0b, 0xd2, 0x5f, 0x2f, 0x49, 0x1b, 0x0b, 0x8b, 0x87, 0x9e, 0x92, 0x27, 0xf3, 0xf2, 0xe4,
	0x65, 0xe0, 0x56, 0x37, 0xd5, 0x12, 0x7f, 0x54, 0xd3, 0xd6, 0x38, 0x3c, 0x17, 0x6d, 0xe7, 0xc8,
	0xf1, 0xd1, 0x0e, 0x45, 0x01, 0xf0, 0x5c, 0x55, 0x12, 0xbf, 0x7b, 0xf4, 0xc4, 0x67, 0xc0, 0x54,
	0xc6, 0x72, 0x56, 0x9c, 0x4a, 0xa6, 0x02, 0x95, 0xd9, 0xd1, 0x96, 0x4a, 0x31, 0x87, 0x69, 0x4c,
	0xfa, 0xd6, 0x59, 0x8f, 0x61, 0xa8, 0x53, 0x54, 0x8b, 0x3b, 0xb8, 0x58, 0xf7, 0x25, 0x75, 0x4a,
	0xd3, 0x21, 0xae, 0x1c, 0x2e, 0xf7, 0xf1, 0x7f, 0x85, 0x4f, 0x00, 0x2b, 0xa5, 0xa9, 0xf7, 0xaf,
	0xa6, 0x46, 0xce, 0xe1, 0xc4, 0xaa, 0x06, 0xe3, 0x78, 0x22, 0xe3, 0x9d, 0x67, 0x30, 0xd2, 0xce,
	0x12, 0x5a, 0x8a, 0xde, 0x99, 0x4c, 0x28, 0x56, 0x70, 0xfd, 0xf2, 0xa5, 0xec, 0x27, 0xbe, 0x75,
	0xee, 0xc3, 0xd4, 0x98, 0x1a, 0x2d, 0x61, 0xdc, 0x1a, 0x4d, 0x7d, 0x87, 0x3e, 0x63, 0xf9, 0x71,
	0x31, 0xbd, 0xbf, 0x5a, 0xa4, 0xb5, 0xec, 0x3f, 0x93, 0x7f, 0x21, 0xf1, 0x08, 0x37, 0x03, 0xd1,
	0xae, 0xeb, 0x1c, 0x26, 0x81, 0xdf, 0xbd, 0xd9, 0x6c, 0x4b, 0x9d, 0xcb, 0x71, 0x78, 0x58, 0x9b,
	0x0d, 0x96, 0x67, 0x71, 0xc5, 0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xd5, 0xad, 0x83,
	0x87, 0x01, 0x00, 0x00,
}
