// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gactus.proto

package gactus

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

type Constant_ResponseCode int32

const (
	Constant_RESPONSE_OK                  Constant_ResponseCode = 0
	Constant_RESPONSE_PROCESS_NOT_FOUND   Constant_ResponseCode = 1
	Constant_RESPONSE_ERROR_SETUP_REQUEST Constant_ResponseCode = 2
)

var Constant_ResponseCode_name = map[int32]string{
	0: "RESPONSE_OK",
	1: "RESPONSE_PROCESS_NOT_FOUND",
	2: "RESPONSE_ERROR_SETUP_REQUEST",
}

var Constant_ResponseCode_value = map[string]int32{
	"RESPONSE_OK":                  0,
	"RESPONSE_PROCESS_NOT_FOUND":   1,
	"RESPONSE_ERROR_SETUP_REQUEST": 2,
}

func (x Constant_ResponseCode) String() string {
	return proto.EnumName(Constant_ResponseCode_name, int32(x))
}

func (Constant_ResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{0, 0}
}

type Constant_ContentType int32

const (
	Constant_CONTENT_TYPE_UNKNOWN               Constant_ContentType = 0
	Constant_CONTENT_TYPE_JSON                  Constant_ContentType = 1
	Constant_CONTENT_TYPE_FORM_DATA             Constant_ContentType = 2
	Constant_CONTENT_TYPE_X_WWW_FORM_URLENCODED Constant_ContentType = 3
)

var Constant_ContentType_name = map[int32]string{
	0: "CONTENT_TYPE_UNKNOWN",
	1: "CONTENT_TYPE_JSON",
	2: "CONTENT_TYPE_FORM_DATA",
	3: "CONTENT_TYPE_X_WWW_FORM_URLENCODED",
}

var Constant_ContentType_value = map[string]int32{
	"CONTENT_TYPE_UNKNOWN":               0,
	"CONTENT_TYPE_JSON":                  1,
	"CONTENT_TYPE_FORM_DATA":             2,
	"CONTENT_TYPE_X_WWW_FORM_URLENCODED": 3,
}

func (x Constant_ContentType) String() string {
	return proto.EnumName(Constant_ContentType_name, int32(x))
}

func (Constant_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{0, 1}
}

type Constant_HttpMethod int32

const (
	Constant_HTTP_METHOD_UNKNOWN Constant_HttpMethod = 0
	Constant_HTTP_METHOD_GET     Constant_HttpMethod = 1
	Constant_HTTP_METHOD_POST    Constant_HttpMethod = 2
)

var Constant_HttpMethod_name = map[int32]string{
	0: "HTTP_METHOD_UNKNOWN",
	1: "HTTP_METHOD_GET",
	2: "HTTP_METHOD_POST",
}

var Constant_HttpMethod_value = map[string]int32{
	"HTTP_METHOD_UNKNOWN": 0,
	"HTTP_METHOD_GET":     1,
	"HTTP_METHOD_POST":    2,
}

func (x Constant_HttpMethod) String() string {
	return proto.EnumName(Constant_HttpMethod_name, int32(x))
}

func (Constant_HttpMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{0, 2}
}

type Constant struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Constant) Reset()         { *m = Constant{} }
func (m *Constant) String() string { return proto.CompactTextString(m) }
func (*Constant) ProtoMessage()    {}
func (*Constant) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{0}
}

func (m *Constant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constant.Unmarshal(m, b)
}
func (m *Constant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constant.Marshal(b, m, deterministic)
}
func (m *Constant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constant.Merge(m, src)
}
func (m *Constant) XXX_Size() int {
	return xxx_messageInfo_Constant.Size(m)
}
func (m *Constant) XXX_DiscardUnknown() {
	xxx_messageInfo_Constant.DiscardUnknown(m)
}

var xxx_messageInfo_Constant proto.InternalMessageInfo

type Request struct {
	LogId                string               `protobuf:"bytes,1,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Command              string               `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	IsProto              bool                 `protobuf:"varint,3,opt,name=is_proto,json=isProto,proto3" json:"is_proto,omitempty"`
	Body                 []byte               `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	ContentType          Constant_ContentType `protobuf:"varint,5,opt,name=content_type,json=contentType,proto3,enum=gactus.Constant_ContentType" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

func (m *Request) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Request) GetIsProto() bool {
	if m != nil {
		return m.IsProto
	}
	return false
}

func (m *Request) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Request) GetContentType() Constant_ContentType {
	if m != nil {
		return m.ContentType
	}
	return Constant_CONTENT_TYPE_UNKNOWN
}

type Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type HttpConfig struct {
	Method               Constant_HttpMethod `protobuf:"varint,1,opt,name=method,proto3,enum=gactus.Constant_HttpMethod" json:"method,omitempty"`
	Path                 string              `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HttpConfig) Reset()         { *m = HttpConfig{} }
func (m *HttpConfig) String() string { return proto.CompactTextString(m) }
func (*HttpConfig) ProtoMessage()    {}
func (*HttpConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{3}
}

func (m *HttpConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConfig.Unmarshal(m, b)
}
func (m *HttpConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConfig.Marshal(b, m, deterministic)
}
func (m *HttpConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConfig.Merge(m, src)
}
func (m *HttpConfig) XXX_Size() int {
	return xxx_messageInfo_HttpConfig.Size(m)
}
func (m *HttpConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpConfig.DiscardUnknown(m)
}

var xxx_messageInfo_HttpConfig proto.InternalMessageInfo

func (m *HttpConfig) GetMethod() Constant_HttpMethod {
	if m != nil {
		return m.Method
	}
	return Constant_HTTP_METHOD_UNKNOWN
}

func (m *HttpConfig) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ProcessorRegistry struct {
	Command              string      `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	HttpConfig           *HttpConfig `protobuf:"bytes,2,opt,name=http_config,json=httpConfig,proto3" json:"http_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ProcessorRegistry) Reset()         { *m = ProcessorRegistry{} }
func (m *ProcessorRegistry) String() string { return proto.CompactTextString(m) }
func (*ProcessorRegistry) ProtoMessage()    {}
func (*ProcessorRegistry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{4}
}

func (m *ProcessorRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessorRegistry.Unmarshal(m, b)
}
func (m *ProcessorRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessorRegistry.Marshal(b, m, deterministic)
}
func (m *ProcessorRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessorRegistry.Merge(m, src)
}
func (m *ProcessorRegistry) XXX_Size() int {
	return xxx_messageInfo_ProcessorRegistry.Size(m)
}
func (m *ProcessorRegistry) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessorRegistry.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessorRegistry proto.InternalMessageInfo

func (m *ProcessorRegistry) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *ProcessorRegistry) GetHttpConfig() *HttpConfig {
	if m != nil {
		return m.HttpConfig
	}
	return nil
}

type RegisterProcessorsRequest struct {
	Addr                 string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	ProcessorRegistries  []*ProcessorRegistry `protobuf:"bytes,2,rep,name=processor_registries,json=processorRegistries,proto3" json:"processor_registries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisterProcessorsRequest) Reset()         { *m = RegisterProcessorsRequest{} }
func (m *RegisterProcessorsRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterProcessorsRequest) ProtoMessage()    {}
func (*RegisterProcessorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{5}
}

func (m *RegisterProcessorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterProcessorsRequest.Unmarshal(m, b)
}
func (m *RegisterProcessorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterProcessorsRequest.Marshal(b, m, deterministic)
}
func (m *RegisterProcessorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterProcessorsRequest.Merge(m, src)
}
func (m *RegisterProcessorsRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterProcessorsRequest.Size(m)
}
func (m *RegisterProcessorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterProcessorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterProcessorsRequest proto.InternalMessageInfo

func (m *RegisterProcessorsRequest) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *RegisterProcessorsRequest) GetProcessorRegistries() []*ProcessorRegistry {
	if m != nil {
		return m.ProcessorRegistries
	}
	return nil
}

type RegisterProcessorsResponse struct {
	DebugMessage         string   `protobuf:"bytes,1,opt,name=debug_message,json=debugMessage,proto3" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterProcessorsResponse) Reset()         { *m = RegisterProcessorsResponse{} }
func (m *RegisterProcessorsResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterProcessorsResponse) ProtoMessage()    {}
func (*RegisterProcessorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b4b60700f4b1236, []int{6}
}

func (m *RegisterProcessorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterProcessorsResponse.Unmarshal(m, b)
}
func (m *RegisterProcessorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterProcessorsResponse.Marshal(b, m, deterministic)
}
func (m *RegisterProcessorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterProcessorsResponse.Merge(m, src)
}
func (m *RegisterProcessorsResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterProcessorsResponse.Size(m)
}
func (m *RegisterProcessorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterProcessorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterProcessorsResponse proto.InternalMessageInfo

func (m *RegisterProcessorsResponse) GetDebugMessage() string {
	if m != nil {
		return m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("gactus.Constant_ResponseCode", Constant_ResponseCode_name, Constant_ResponseCode_value)
	proto.RegisterEnum("gactus.Constant_ContentType", Constant_ContentType_name, Constant_ContentType_value)
	proto.RegisterEnum("gactus.Constant_HttpMethod", Constant_HttpMethod_name, Constant_HttpMethod_value)
	proto.RegisterType((*Constant)(nil), "gactus.Constant")
	proto.RegisterType((*Request)(nil), "gactus.Request")
	proto.RegisterType((*Response)(nil), "gactus.Response")
	proto.RegisterType((*HttpConfig)(nil), "gactus.HttpConfig")
	proto.RegisterType((*ProcessorRegistry)(nil), "gactus.ProcessorRegistry")
	proto.RegisterType((*RegisterProcessorsRequest)(nil), "gactus.RegisterProcessorsRequest")
	proto.RegisterType((*RegisterProcessorsResponse)(nil), "gactus.RegisterProcessorsResponse")
}

func init() { proto.RegisterFile("proto/gactus.proto", fileDescriptor_0b4b60700f4b1236) }

var fileDescriptor_0b4b60700f4b1236 = []byte{
	// 565 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x5d, 0x4f, 0xdb, 0x30,
	0x14, 0x25, 0x05, 0x4a, 0x77, 0x5b, 0x20, 0x18, 0xd8, 0x02, 0x43, 0x53, 0x95, 0x49, 0x53, 0x9f,
	0x98, 0x54, 0x7e, 0xc0, 0x84, 0x52, 0x33, 0x36, 0x68, 0x9c, 0x39, 0xae, 0xba, 0x3d, 0x59, 0x69,
	0xe2, 0xa5, 0x91, 0x68, 0x9c, 0xc5, 0xe6, 0xa1, 0x0f, 0x7b, 0xd9, 0xcf, 0xd9, 0xef, 0xd8, 0x0f,
	0x9b, 0x70, 0x9a, 0x7e, 0x8c, 0xbd, 0xdd, 0x7b, 0xee, 0xcd, 0x3d, 0xc7, 0xe7, 0x04, 0x50, 0x51,
	0x4a, 0x2d, 0xdf, 0xa7, 0x51, 0xac, 0x1f, 0xd5, 0xa5, 0x69, 0x50, 0xb3, 0xea, 0xdc, 0x3f, 0x0d,
	0x68, 0x79, 0x32, 0x57, 0x3a, 0xca, 0xb5, 0x1b, 0x41, 0x87, 0x0a, 0x55, 0xc8, 0x5c, 0x09, 0x4f,
	0x26, 0x02, 0x1d, 0x42, 0x9b, 0xe2, 0x30, 0x20, 0x7e, 0x88, 0x39, 0xb9, 0xb3, 0xb7, 0xd0, 0x1b,
	0x38, 0x5f, 0x02, 0x01, 0x25, 0x1e, 0x0e, 0x43, 0xee, 0x13, 0xc6, 0x6f, 0xc8, 0xc8, 0x1f, 0xd8,
	0x16, 0xea, 0xc2, 0xc5, 0x72, 0x8e, 0x29, 0x25, 0x94, 0x87, 0x98, 0x8d, 0x02, 0x4e, 0xf1, 0x97,
	0x11, 0x0e, 0x99, 0xdd, 0x70, 0x7f, 0x59, 0xd0, 0xf6, 0x64, 0xae, 0x45, 0xae, 0xd9, 0xbc, 0x10,
	0xc8, 0x81, 0x13, 0x8f, 0xf8, 0x0c, 0xfb, 0x8c, 0xb3, 0x6f, 0x01, 0xe6, 0x23, 0xff, 0xce, 0x27,
	0x63, 0xdf, 0xde, 0x42, 0xa7, 0x70, 0xb4, 0x31, 0xf9, 0x1c, 0x12, 0xdf, 0xb6, 0xd0, 0x39, 0xbc,
	0xdc, 0x80, 0x6f, 0x08, 0x1d, 0xf2, 0xc1, 0x35, 0xbb, 0xb6, 0x1b, 0xe8, 0x1d, 0xb8, 0x1b, 0xb3,
	0xaf, 0x7c, 0x3c, 0x1e, 0x57, 0x1b, 0x23, 0x7a, 0x8f, 0x7d, 0x8f, 0x0c, 0xf0, 0xc0, 0xde, 0x76,
	0x03, 0x80, 0x5b, 0xad, 0x8b, 0xa1, 0xd0, 0x53, 0x99, 0xa0, 0x57, 0x70, 0x7c, 0xcb, 0x58, 0xc0,
	0x87, 0x98, 0xdd, 0x92, 0xc1, 0x9a, 0x82, 0x63, 0x38, 0x5c, 0x1f, 0x7c, 0xc4, 0xcc, 0xb6, 0xd0,
	0x09, 0xd8, 0xeb, 0x60, 0x40, 0xcc, 0xb3, 0x7e, 0x5b, 0xb0, 0x47, 0xc5, 0x8f, 0x47, 0xa1, 0x34,
	0x3a, 0x85, 0xe6, 0x83, 0x4c, 0x79, 0x96, 0x38, 0x56, 0xd7, 0xea, 0xbd, 0xa0, 0xbb, 0x0f, 0x32,
	0xfd, 0x94, 0x20, 0x07, 0xf6, 0x62, 0x39, 0x9b, 0x45, 0x79, 0xe2, 0x34, 0x0c, 0x5e, 0xb7, 0xe8,
	0x0c, 0x5a, 0x99, 0xe2, 0x26, 0x17, 0x67, 0xbb, 0x6b, 0xf5, 0x5a, 0x74, 0x2f, 0x53, 0x81, 0x89,
	0x09, 0xc1, 0xce, 0x44, 0x26, 0x73, 0x67, 0xa7, 0x6b, 0xf5, 0x3a, 0xd4, 0xd4, 0xe8, 0x03, 0x74,
	0xe2, 0xca, 0x41, 0xae, 0xe7, 0x85, 0x70, 0x76, 0xbb, 0x56, 0xef, 0xa0, 0x7f, 0x71, 0xb9, 0xc8,
	0xb7, 0x4e, 0xf3, 0x72, 0xcd, 0x66, 0xda, 0x8e, 0x57, 0x8d, 0xdb, 0x87, 0x56, 0x1d, 0xf3, 0x13,
	0x41, 0x2c, 0x13, 0x61, 0xa4, 0xee, 0x53, 0x53, 0x2f, 0x49, 0x1b, 0x2b, 0x52, 0x77, 0x54, 0x59,
	0xe6, 0xc9, 0xfc, 0x7b, 0x96, 0xa2, 0x2b, 0x68, 0xce, 0x8c, 0x79, 0xe6, 0xbb, 0x83, 0xfe, 0xeb,
	0x67, 0xe4, 0x2b, 0x7f, 0xe9, 0x62, 0xf5, 0xe9, 0x6c, 0x11, 0xe9, 0xe9, 0xe2, 0xf5, 0xa6, 0x76,
	0x27, 0x70, 0x14, 0x94, 0x32, 0x16, 0x4a, 0xc9, 0x92, 0x8a, 0x34, 0x53, 0xba, 0x9c, 0xaf, 0x3b,
	0x65, 0x6d, 0x3a, 0x75, 0x05, 0xed, 0xa9, 0xd6, 0x05, 0x8f, 0x8d, 0x0c, 0x73, 0xa9, 0xdd, 0x47,
	0x35, 0xf9, 0x4a, 0x20, 0x85, 0xe9, 0xb2, 0x76, 0x7f, 0xc2, 0x59, 0x75, 0x5a, 0x94, 0x4b, 0x2e,
	0x55, 0x87, 0x85, 0x60, 0x27, 0x4a, 0x92, 0x72, 0x41, 0x64, 0x6a, 0x74, 0x0f, 0x27, 0x45, 0xbd,
	0xc8, 0xcb, 0x4a, 0x55, 0x26, 0x94, 0xd3, 0xe8, 0x6e, 0xf7, 0xda, 0xfd, 0xb3, 0x9a, 0xee, 0x99,
	0x70, 0x7a, 0x5c, 0xfc, 0x03, 0x65, 0x42, 0xb9, 0xd7, 0x70, 0xfe, 0x3f, 0xfa, 0x85, 0xff, 0x6f,
	0x61, 0x3f, 0x11, 0x93, 0xc7, 0x94, 0xcf, 0x84, 0x52, 0x51, 0x2a, 0x16, 0x42, 0x3a, 0x06, 0x1c,
	0x56, 0xd8, 0xa4, 0x69, 0xfe, 0x8d, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0x00, 0xf8,
	0x8c, 0xc9, 0x03, 0x00, 0x00,
}
