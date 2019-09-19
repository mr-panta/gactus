// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gactus.proto

package gactus

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

type Constant_ResponseCode int32

const (
	Constant_RESPONSE_OK Constant_ResponseCode = 0
	// Internal response code
	Constant_RESPONSE_COMMAND_NOT_FOUND     Constant_ResponseCode = 100
	Constant_RESPONSE_ERROR_SETUP_REQUEST   Constant_ResponseCode = 101
	Constant_RESPONSE_ERROR_UNPACK_REQUEST  Constant_ResponseCode = 102
	Constant_RESPONSE_ERROR_SETUP_RESPONSE  Constant_ResponseCode = 103
	Constant_RESPONSE_ERROR_UNPACK_RESPONSE Constant_ResponseCode = 104
	Constant_RESPONSE_CREATE_CLIENT_FAILED  Constant_ResponseCode = 105
)

var Constant_ResponseCode_name = map[int32]string{
	0:   "RESPONSE_OK",
	100: "RESPONSE_COMMAND_NOT_FOUND",
	101: "RESPONSE_ERROR_SETUP_REQUEST",
	102: "RESPONSE_ERROR_UNPACK_REQUEST",
	103: "RESPONSE_ERROR_SETUP_RESPONSE",
	104: "RESPONSE_ERROR_UNPACK_RESPONSE",
	105: "RESPONSE_CREATE_CLIENT_FAILED",
}
var Constant_ResponseCode_value = map[string]int32{
	"RESPONSE_OK":                    0,
	"RESPONSE_COMMAND_NOT_FOUND":     100,
	"RESPONSE_ERROR_SETUP_REQUEST":   101,
	"RESPONSE_ERROR_UNPACK_REQUEST":  102,
	"RESPONSE_ERROR_SETUP_RESPONSE":  103,
	"RESPONSE_ERROR_UNPACK_RESPONSE": 104,
	"RESPONSE_CREATE_CLIENT_FAILED":  105,
}

func (x Constant_ResponseCode) String() string {
	return proto.EnumName(Constant_ResponseCode_name, int32(x))
}
func (Constant_ResponseCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_gactus_768e06ca764c40a1, []int{0, 0}
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{0, 1}
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{0, 2}
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{0}
}
func (m *Constant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constant.Unmarshal(m, b)
}
func (m *Constant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constant.Marshal(b, m, deterministic)
}
func (dst *Constant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constant.Merge(dst, src)
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{3}
}
func (m *HttpConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpConfig.Unmarshal(m, b)
}
func (m *HttpConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpConfig.Marshal(b, m, deterministic)
}
func (dst *HttpConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpConfig.Merge(dst, src)
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{4}
}
func (m *ProcessorRegistry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessorRegistry.Unmarshal(m, b)
}
func (m *ProcessorRegistry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessorRegistry.Marshal(b, m, deterministic)
}
func (dst *ProcessorRegistry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessorRegistry.Merge(dst, src)
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
	Address              string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ProcessorRegistries  []*ProcessorRegistry `protobuf:"bytes,2,rep,name=processor_registries,json=processorRegistries,proto3" json:"processor_registries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisterProcessorsRequest) Reset()         { *m = RegisterProcessorsRequest{} }
func (m *RegisterProcessorsRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterProcessorsRequest) ProtoMessage()    {}
func (*RegisterProcessorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_768e06ca764c40a1, []int{5}
}
func (m *RegisterProcessorsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterProcessorsRequest.Unmarshal(m, b)
}
func (m *RegisterProcessorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterProcessorsRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterProcessorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterProcessorsRequest.Merge(dst, src)
}
func (m *RegisterProcessorsRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterProcessorsRequest.Size(m)
}
func (m *RegisterProcessorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterProcessorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterProcessorsRequest proto.InternalMessageInfo

func (m *RegisterProcessorsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegisterProcessorsRequest) GetProcessorRegistries() []*ProcessorRegistry {
	if m != nil {
		return m.ProcessorRegistries
	}
	return nil
}

type CommandAddressPair struct {
	Command              string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandAddressPair) Reset()         { *m = CommandAddressPair{} }
func (m *CommandAddressPair) String() string { return proto.CompactTextString(m) }
func (*CommandAddressPair) ProtoMessage()    {}
func (*CommandAddressPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_768e06ca764c40a1, []int{6}
}
func (m *CommandAddressPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandAddressPair.Unmarshal(m, b)
}
func (m *CommandAddressPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandAddressPair.Marshal(b, m, deterministic)
}
func (dst *CommandAddressPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandAddressPair.Merge(dst, src)
}
func (m *CommandAddressPair) XXX_Size() int {
	return xxx_messageInfo_CommandAddressPair.Size(m)
}
func (m *CommandAddressPair) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandAddressPair.DiscardUnknown(m)
}

var xxx_messageInfo_CommandAddressPair proto.InternalMessageInfo

func (m *CommandAddressPair) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandAddressPair) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
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
	return fileDescriptor_gactus_768e06ca764c40a1, []int{7}
}
func (m *RegisterProcessorsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterProcessorsResponse.Unmarshal(m, b)
}
func (m *RegisterProcessorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterProcessorsResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterProcessorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterProcessorsResponse.Merge(dst, src)
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

type UpdateRegistriesRequest struct {
	Pairs                []*CommandAddressPair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRegistriesRequest) Reset()         { *m = UpdateRegistriesRequest{} }
func (m *UpdateRegistriesRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRegistriesRequest) ProtoMessage()    {}
func (*UpdateRegistriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_768e06ca764c40a1, []int{8}
}
func (m *UpdateRegistriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRegistriesRequest.Unmarshal(m, b)
}
func (m *UpdateRegistriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRegistriesRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRegistriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRegistriesRequest.Merge(dst, src)
}
func (m *UpdateRegistriesRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRegistriesRequest.Size(m)
}
func (m *UpdateRegistriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRegistriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRegistriesRequest proto.InternalMessageInfo

func (m *UpdateRegistriesRequest) GetPairs() []*CommandAddressPair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type UpdateRegistriesResponse struct {
	DebugMessage         string   `protobuf:"bytes,1,opt,name=debug_message,json=debugMessage,proto3" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRegistriesResponse) Reset()         { *m = UpdateRegistriesResponse{} }
func (m *UpdateRegistriesResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateRegistriesResponse) ProtoMessage()    {}
func (*UpdateRegistriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_768e06ca764c40a1, []int{9}
}
func (m *UpdateRegistriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRegistriesResponse.Unmarshal(m, b)
}
func (m *UpdateRegistriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRegistriesResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateRegistriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRegistriesResponse.Merge(dst, src)
}
func (m *UpdateRegistriesResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateRegistriesResponse.Size(m)
}
func (m *UpdateRegistriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRegistriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRegistriesResponse proto.InternalMessageInfo

func (m *UpdateRegistriesResponse) GetDebugMessage() string {
	if m != nil {
		return m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Constant)(nil), "gactus.Constant")
	proto.RegisterType((*Request)(nil), "gactus.Request")
	proto.RegisterType((*Response)(nil), "gactus.Response")
	proto.RegisterType((*HttpConfig)(nil), "gactus.HttpConfig")
	proto.RegisterType((*ProcessorRegistry)(nil), "gactus.ProcessorRegistry")
	proto.RegisterType((*RegisterProcessorsRequest)(nil), "gactus.RegisterProcessorsRequest")
	proto.RegisterType((*CommandAddressPair)(nil), "gactus.CommandAddressPair")
	proto.RegisterType((*RegisterProcessorsResponse)(nil), "gactus.RegisterProcessorsResponse")
	proto.RegisterType((*UpdateRegistriesRequest)(nil), "gactus.UpdateRegistriesRequest")
	proto.RegisterType((*UpdateRegistriesResponse)(nil), "gactus.UpdateRegistriesResponse")
	proto.RegisterEnum("gactus.Constant_ResponseCode", Constant_ResponseCode_name, Constant_ResponseCode_value)
	proto.RegisterEnum("gactus.Constant_ContentType", Constant_ContentType_name, Constant_ContentType_value)
	proto.RegisterEnum("gactus.Constant_HttpMethod", Constant_HttpMethod_name, Constant_HttpMethod_value)
}

func init() { proto.RegisterFile("proto/gactus.proto", fileDescriptor_gactus_768e06ca764c40a1) }

var fileDescriptor_gactus_768e06ca764c40a1 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0x8d, 0x21, 0x01, 0xbe, 0x81, 0x24, 0xce, 0x26, 0xf9, 0xe2, 0xf0, 0xe5, 0x8b, 0xa8, 0x2b,
	0x55, 0x5c, 0xa5, 0x15, 0x79, 0x80, 0xc8, 0x32, 0x9b, 0x92, 0x02, 0xb6, 0xbb, 0xd8, 0xa2, 0xbd,
	0x5a, 0x19, 0xbc, 0x31, 0x96, 0x02, 0xeb, 0x7a, 0x37, 0x17, 0xdc, 0xb6, 0x6f, 0xd3, 0xc7, 0xe9,
	0x83, 0xf4, 0x19, 0x2a, 0x6c, 0xcc, 0x4f, 0x92, 0x56, 0xea, 0xdd, 0xce, 0xcc, 0xf1, 0x39, 0x67,
	0x8e, 0x07, 0x50, 0x9c, 0x70, 0xc9, 0xdf, 0x86, 0xfe, 0x58, 0x3e, 0x8a, 0xab, 0xb4, 0x40, 0xa5,
	0xac, 0xd2, 0x7f, 0x14, 0xa1, 0x62, 0xf2, 0x99, 0x90, 0xfe, 0x4c, 0xea, 0x3f, 0x15, 0xa8, 0x11,
	0x26, 0x62, 0x3e, 0x13, 0xcc, 0xe4, 0x01, 0x43, 0x87, 0x50, 0x25, 0x78, 0xe0, 0xd8, 0xd6, 0x00,
	0x53, 0xbb, 0xab, 0xee, 0xa0, 0x4b, 0xa8, 0xaf, 0x1a, 0xa6, 0xdd, 0xef, 0x1b, 0x56, 0x9b, 0x5a,
	0xb6, 0x4b, 0x6f, 0x6d, 0xcf, 0x6a, 0xab, 0x01, 0x6a, 0xc0, 0xc5, 0x6a, 0x8e, 0x09, 0xb1, 0x09,
	0x1d, 0x60, 0xd7, 0x73, 0x28, 0xc1, 0x1f, 0x3d, 0x3c, 0x70, 0x55, 0x86, 0x5e, 0xc1, 0xff, 0x4f,
	0x10, 0x9e, 0xe5, 0x18, 0x66, 0x77, 0x05, 0xb9, 0x7f, 0x01, 0x92, 0x93, 0x64, 0x4d, 0x35, 0x44,
	0x3a, 0x5c, 0xfe, 0x8e, 0x65, 0x89, 0x99, 0x6c, 0xd1, 0x98, 0x04, 0x1b, 0x2e, 0xa6, 0x66, 0xef,
	0x0e, 0x5b, 0x2e, 0xbd, 0x35, 0xee, 0x7a, 0xb8, 0xad, 0x46, 0xfa, 0x57, 0x05, 0xaa, 0x26, 0x9f,
	0x49, 0x36, 0x93, 0xee, 0x3c, 0x66, 0x48, 0x83, 0x13, 0xd3, 0xb6, 0xdc, 0x05, 0xc6, 0xfd, 0xec,
	0x60, 0xea, 0x59, 0x5d, 0xcb, 0x1e, 0x5a, 0xea, 0x0e, 0x3a, 0x85, 0xa3, 0xad, 0xc9, 0x87, 0x81,
	0x6d, 0xa9, 0x0a, 0xaa, 0xc3, 0xbf, 0x5b, 0xed, 0x5b, 0x9b, 0xf4, 0x69, 0xdb, 0x70, 0x0d, 0xb5,
	0x80, 0xde, 0x80, 0xbe, 0x35, 0xfb, 0x44, 0x87, 0xc3, 0x61, 0x86, 0xf0, 0x48, 0x0f, 0x5b, 0xa6,
	0xdd, 0xc6, 0x6d, 0xb5, 0xa8, 0x3b, 0x00, 0x1d, 0x29, 0xe3, 0x3e, 0x93, 0x13, 0x1e, 0xa0, 0x33,
	0x38, 0xee, 0xb8, 0xae, 0x43, 0xfb, 0xd8, 0xed, 0xd8, 0xed, 0x0d, 0x07, 0xc7, 0x70, 0xb8, 0x39,
	0x78, 0x8f, 0x5d, 0x55, 0x41, 0x27, 0xa0, 0x6e, 0x36, 0x1d, 0x7b, 0xe0, 0xaa, 0x05, 0xfd, 0xbb,
	0x02, 0x65, 0xc2, 0xbe, 0x3c, 0x32, 0x21, 0xd1, 0x29, 0x94, 0x1e, 0x78, 0x48, 0xa3, 0x40, 0x53,
	0x1a, 0x4a, 0xf3, 0x1f, 0xb2, 0xf7, 0xc0, 0xc3, 0xbb, 0x00, 0x69, 0x50, 0x1e, 0xf3, 0xe9, 0xd4,
	0x9f, 0x05, 0x5a, 0x21, 0xed, 0xe7, 0x25, 0x3a, 0x87, 0x4a, 0x24, 0x68, 0x7a, 0x25, 0x5a, 0xb1,
	0xa1, 0x34, 0x2b, 0xa4, 0x1c, 0x09, 0x27, 0x3d, 0x1a, 0x04, 0xbb, 0x23, 0x1e, 0xcc, 0xb5, 0xdd,
	0x86, 0xd2, 0xac, 0x91, 0xf4, 0x8d, 0x6e, 0xa0, 0x36, 0xce, 0x12, 0xa4, 0x72, 0x1e, 0x33, 0x6d,
	0xaf, 0xa1, 0x34, 0x0f, 0x5a, 0x17, 0x57, 0xcb, 0x6b, 0xcb, 0x6f, 0xeb, 0x6a, 0x23, 0x66, 0x52,
	0x1d, 0xaf, 0x0b, 0xbd, 0x05, 0x95, 0xfc, 0xe6, 0x16, 0x02, 0x63, 0x1e, 0xb0, 0xd4, 0xea, 0x3e,
	0x49, 0xdf, 0x2b, 0xd1, 0xc2, 0x5a, 0x54, 0xf7, 0xb2, 0xc8, 0x4c, 0x3e, 0xbb, 0x8f, 0x42, 0x74,
	0x0d, 0xa5, 0x69, 0x1a, 0x5e, 0xfa, 0xdd, 0x41, 0xeb, 0xbf, 0x67, 0xe2, 0xeb, 0x7c, 0xc9, 0x12,
	0xba, 0xa0, 0x8d, 0x7d, 0x39, 0x59, 0x6e, 0x9f, 0xbe, 0xf5, 0x11, 0x1c, 0x39, 0x09, 0x1f, 0x33,
	0x21, 0x78, 0x42, 0x58, 0x18, 0x09, 0x99, 0xcc, 0x37, 0x93, 0x52, 0xb6, 0x93, 0xba, 0x86, 0xea,
	0x44, 0xca, 0x98, 0x8e, 0x53, 0x1b, 0x29, 0x53, 0xb5, 0x85, 0x72, 0xf1, 0xb5, 0x41, 0x02, 0x93,
	0xd5, 0x5b, 0xff, 0xa6, 0xc0, 0x79, 0xc6, 0xcd, 0x92, 0x95, 0x98, 0xc8, 0xff, 0x96, 0x06, 0x65,
	0x3f, 0x08, 0x12, 0x26, 0x44, 0x2e, 0xb6, 0x2c, 0x51, 0x0f, 0x4e, 0xe2, 0x1c, 0x4e, 0x93, 0xcc,
	0x5c, 0xc4, 0x84, 0x56, 0x68, 0x14, 0x9b, 0xd5, 0xd6, 0x79, 0xae, 0xfa, 0xcc, 0x3f, 0x39, 0x8e,
	0x9f, 0xb4, 0x22, 0x26, 0xf4, 0x0e, 0x20, 0x33, 0xdb, 0xc2, 0xc8, 0xf8, 0x1d, 0x3f, 0x4a, 0xfe,
	0xb0, 0xea, 0x86, 0xaf, 0xc2, 0x96, 0x2f, 0xdd, 0x80, 0xfa, 0x4b, 0xeb, 0x2c, 0x7f, 0xe8, 0x6b,
	0xd8, 0x0f, 0xd8, 0xe8, 0x31, 0xa4, 0x53, 0x26, 0x84, 0x1f, 0xb2, 0x25, 0x6f, 0x2d, 0x6d, 0xf6,
	0xb3, 0x9e, 0xde, 0x85, 0x33, 0x2f, 0x0e, 0x7c, 0xc9, 0xd6, 0x06, 0xf3, 0x3c, 0xde, 0xc1, 0x5e,
	0xec, 0x47, 0xc9, 0x22, 0x8d, 0xc5, 0x9a, 0xf5, 0xf5, 0x9f, 0x7d, 0x6a, 0x9e, 0x64, 0x40, 0xfd,
	0x06, 0xb4, 0xe7, 0x64, 0x7f, 0xe1, 0x66, 0x54, 0x4a, 0x4f, 0xff, 0xfa, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x41, 0x2d, 0x81, 0x00, 0x36, 0x05, 0x00, 0x00,
}
