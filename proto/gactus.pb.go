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
	return fileDescriptor_gactus_19973b659dc869c0, []int{0, 0}
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{0, 1}
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{0, 2}
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{0}
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
	HttpAddress          string               `protobuf:"bytes,1,opt,name=http_address,json=httpAddress,proto3" json:"http_address,omitempty"`
	LogId                string               `protobuf:"bytes,2,opt,name=log_id,json=logId,proto3" json:"log_id,omitempty"`
	Command              string               `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	IsProto              bool                 `protobuf:"varint,4,opt,name=is_proto,json=isProto,proto3" json:"is_proto,omitempty"`
	Header               map[string]string    `protobuf:"bytes,5,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Query                map[string]string    `protobuf:"bytes,6,rep,name=query,proto3" json:"query,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte               `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ContentType          Constant_ContentType `protobuf:"varint,8,opt,name=content_type,json=contentType,proto3,enum=gactus.Constant_ContentType" json:"content_type,omitempty"`
	RawContentType       string               `protobuf:"bytes,9,opt,name=raw_content_type,json=rawContentType,proto3" json:"raw_content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{1}
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

func (m *Request) GetHttpAddress() string {
	if m != nil {
		return m.HttpAddress
	}
	return ""
}

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

func (m *Request) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Request) GetQuery() map[string]string {
	if m != nil {
		return m.Query
	}
	return nil
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

func (m *Request) GetRawContentType() string {
	if m != nil {
		return m.RawContentType
	}
	return ""
}

type Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	DebugMessage         string   `protobuf:"bytes,3,opt,name=debug_message,json=debugMessage,proto3" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{2}
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

func (m *Response) GetDebugMessage() string {
	if m != nil {
		return m.DebugMessage
	}
	return ""
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{3}
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

type ConnectionConfig struct {
	MinConns             uint32   `protobuf:"varint,1,opt,name=min_conns,json=minConns,proto3" json:"min_conns,omitempty"`
	MaxConns             uint32   `protobuf:"varint,2,opt,name=max_conns,json=maxConns,proto3" json:"max_conns,omitempty"`
	IdleConnTimeout      uint32   `protobuf:"varint,3,opt,name=idle_conn_timeout,json=idleConnTimeout,proto3" json:"idle_conn_timeout,omitempty"`
	WaitConnTimeout      uint32   `protobuf:"varint,4,opt,name=wait_conn_timeout,json=waitConnTimeout,proto3" json:"wait_conn_timeout,omitempty"`
	ClearPeriod          uint32   `protobuf:"varint,5,opt,name=clear_period,json=clearPeriod,proto3" json:"clear_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionConfig) Reset()         { *m = ConnectionConfig{} }
func (m *ConnectionConfig) String() string { return proto.CompactTextString(m) }
func (*ConnectionConfig) ProtoMessage()    {}
func (*ConnectionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{4}
}
func (m *ConnectionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionConfig.Unmarshal(m, b)
}
func (m *ConnectionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionConfig.Marshal(b, m, deterministic)
}
func (dst *ConnectionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionConfig.Merge(dst, src)
}
func (m *ConnectionConfig) XXX_Size() int {
	return xxx_messageInfo_ConnectionConfig.Size(m)
}
func (m *ConnectionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionConfig proto.InternalMessageInfo

func (m *ConnectionConfig) GetMinConns() uint32 {
	if m != nil {
		return m.MinConns
	}
	return 0
}

func (m *ConnectionConfig) GetMaxConns() uint32 {
	if m != nil {
		return m.MaxConns
	}
	return 0
}

func (m *ConnectionConfig) GetIdleConnTimeout() uint32 {
	if m != nil {
		return m.IdleConnTimeout
	}
	return 0
}

func (m *ConnectionConfig) GetWaitConnTimeout() uint32 {
	if m != nil {
		return m.WaitConnTimeout
	}
	return 0
}

func (m *ConnectionConfig) GetClearPeriod() uint32 {
	if m != nil {
		return m.ClearPeriod
	}
	return 0
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{5}
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{6}
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

type AddressConfig struct {
	Address              string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ConnConfig           *ConnectionConfig `protobuf:"bytes,2,opt,name=conn_config,json=connConfig,proto3" json:"conn_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AddressConfig) Reset()         { *m = AddressConfig{} }
func (m *AddressConfig) String() string { return proto.CompactTextString(m) }
func (*AddressConfig) ProtoMessage()    {}
func (*AddressConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{7}
}
func (m *AddressConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressConfig.Unmarshal(m, b)
}
func (m *AddressConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressConfig.Marshal(b, m, deterministic)
}
func (dst *AddressConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressConfig.Merge(dst, src)
}
func (m *AddressConfig) XXX_Size() int {
	return xxx_messageInfo_AddressConfig.Size(m)
}
func (m *AddressConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AddressConfig proto.InternalMessageInfo

func (m *AddressConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddressConfig) GetConnConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnConfig
	}
	return nil
}

type RegisterServiceRequest struct {
	Addresses            []string             `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
	ProcessorRegistries  []*ProcessorRegistry `protobuf:"bytes,2,rep,name=processor_registries,json=processorRegistries,proto3" json:"processor_registries,omitempty"`
	ConnConfig           *ConnectionConfig    `protobuf:"bytes,3,opt,name=conn_config,json=connConfig,proto3" json:"conn_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisterServiceRequest) Reset()         { *m = RegisterServiceRequest{} }
func (m *RegisterServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterServiceRequest) ProtoMessage()    {}
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{8}
}
func (m *RegisterServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterServiceRequest.Unmarshal(m, b)
}
func (m *RegisterServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterServiceRequest.Marshal(b, m, deterministic)
}
func (dst *RegisterServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterServiceRequest.Merge(dst, src)
}
func (m *RegisterServiceRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterServiceRequest.Size(m)
}
func (m *RegisterServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterServiceRequest proto.InternalMessageInfo

func (m *RegisterServiceRequest) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *RegisterServiceRequest) GetProcessorRegistries() []*ProcessorRegistry {
	if m != nil {
		return m.ProcessorRegistries
	}
	return nil
}

func (m *RegisterServiceRequest) GetConnConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnConfig
	}
	return nil
}

type RegisterServiceResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DebugMessage         string   `protobuf:"bytes,2,opt,name=debug_message,json=debugMessage,proto3" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterServiceResponse) Reset()         { *m = RegisterServiceResponse{} }
func (m *RegisterServiceResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterServiceResponse) ProtoMessage()    {}
func (*RegisterServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{9}
}
func (m *RegisterServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterServiceResponse.Unmarshal(m, b)
}
func (m *RegisterServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterServiceResponse.Marshal(b, m, deterministic)
}
func (dst *RegisterServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterServiceResponse.Merge(dst, src)
}
func (m *RegisterServiceResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterServiceResponse.Size(m)
}
func (m *RegisterServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterServiceResponse proto.InternalMessageInfo

func (m *RegisterServiceResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RegisterServiceResponse) GetDebugMessage() string {
	if m != nil {
		return m.DebugMessage
	}
	return ""
}

type UpdateRegistriesRequest struct {
	CommandAddressPairs  []*CommandAddressPair `protobuf:"bytes,1,rep,name=command_address_pairs,json=commandAddressPairs,proto3" json:"command_address_pairs,omitempty"`
	AddrConfigs          []*AddressConfig      `protobuf:"bytes,2,rep,name=addr_configs,json=addrConfigs,proto3" json:"addr_configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRegistriesRequest) Reset()         { *m = UpdateRegistriesRequest{} }
func (m *UpdateRegistriesRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRegistriesRequest) ProtoMessage()    {}
func (*UpdateRegistriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{10}
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

func (m *UpdateRegistriesRequest) GetCommandAddressPairs() []*CommandAddressPair {
	if m != nil {
		return m.CommandAddressPairs
	}
	return nil
}

func (m *UpdateRegistriesRequest) GetAddrConfigs() []*AddressConfig {
	if m != nil {
		return m.AddrConfigs
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
	return fileDescriptor_gactus_19973b659dc869c0, []int{11}
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

type HealthCheckRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{12}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type HealthCheckResponse struct {
	DebugMessage         string   `protobuf:"bytes,1,opt,name=debug_message,json=debugMessage,proto3" json:"debug_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gactus_19973b659dc869c0, []int{13}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetDebugMessage() string {
	if m != nil {
		return m.DebugMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Constant)(nil), "gactus.Constant")
	proto.RegisterType((*Request)(nil), "gactus.Request")
	proto.RegisterMapType((map[string]string)(nil), "gactus.Request.HeaderEntry")
	proto.RegisterMapType((map[string]string)(nil), "gactus.Request.QueryEntry")
	proto.RegisterType((*Response)(nil), "gactus.Response")
	proto.RegisterType((*HttpConfig)(nil), "gactus.HttpConfig")
	proto.RegisterType((*ConnectionConfig)(nil), "gactus.ConnectionConfig")
	proto.RegisterType((*ProcessorRegistry)(nil), "gactus.ProcessorRegistry")
	proto.RegisterType((*CommandAddressPair)(nil), "gactus.CommandAddressPair")
	proto.RegisterType((*AddressConfig)(nil), "gactus.AddressConfig")
	proto.RegisterType((*RegisterServiceRequest)(nil), "gactus.RegisterServiceRequest")
	proto.RegisterType((*RegisterServiceResponse)(nil), "gactus.RegisterServiceResponse")
	proto.RegisterType((*UpdateRegistriesRequest)(nil), "gactus.UpdateRegistriesRequest")
	proto.RegisterType((*UpdateRegistriesResponse)(nil), "gactus.UpdateRegistriesResponse")
	proto.RegisterType((*HealthCheckRequest)(nil), "gactus.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "gactus.HealthCheckResponse")
	proto.RegisterEnum("gactus.Constant_ResponseCode", Constant_ResponseCode_name, Constant_ResponseCode_value)
	proto.RegisterEnum("gactus.Constant_ContentType", Constant_ContentType_name, Constant_ContentType_value)
	proto.RegisterEnum("gactus.Constant_HttpMethod", Constant_HttpMethod_name, Constant_HttpMethod_value)
}

func init() { proto.RegisterFile("proto/gactus.proto", fileDescriptor_gactus_19973b659dc869c0) }

var fileDescriptor_gactus_19973b659dc869c0 = []byte{
	// 1011 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x41, 0x53, 0xdb, 0x46,
	0x14, 0x8e, 0xec, 0x00, 0xe6, 0xd9, 0x80, 0x58, 0x20, 0x28, 0x84, 0x66, 0x8c, 0x3a, 0xd3, 0xf1,
	0xf4, 0x40, 0x3b, 0x70, 0x21, 0xb9, 0x64, 0x3c, 0xb2, 0xa8, 0x29, 0x58, 0x52, 0xd6, 0xf2, 0x38,
	0x3d, 0x69, 0x84, 0xb4, 0xd8, 0x9a, 0xd8, 0x92, 0x22, 0xad, 0x43, 0x7c, 0xed, 0x5f, 0xe9, 0x0f,
	0xe9, 0xb1, 0xf7, 0xf6, 0x7f, 0xf4, 0x37, 0x74, 0x76, 0xb5, 0xb2, 0x25, 0x0c, 0x9d, 0xe1, 0xb6,
	0xfb, 0x7d, 0xdf, 0xbe, 0xdd, 0xf7, 0xbd, 0xf7, 0x24, 0x40, 0x71, 0x12, 0xd1, 0xe8, 0xa7, 0x91,
	0xeb, 0xd1, 0x59, 0x7a, 0xca, 0x37, 0x68, 0x3d, 0xdb, 0xa9, 0x7f, 0x57, 0xa1, 0xa6, 0x45, 0x61,
	0x4a, 0xdd, 0x90, 0xaa, 0xff, 0x4a, 0xd0, 0xc0, 0x24, 0x8d, 0xa3, 0x30, 0x25, 0x5a, 0xe4, 0x13,
	0xb4, 0x03, 0x75, 0xac, 0xf7, 0x2d, 0xd3, 0xe8, 0xeb, 0x8e, 0x79, 0x2d, 0xbf, 0x40, 0x6f, 0xe1,
	0x68, 0x01, 0x68, 0x66, 0xaf, 0xd7, 0x36, 0x3a, 0x8e, 0x61, 0xda, 0xce, 0xa5, 0x39, 0x30, 0x3a,
	0xb2, 0x8f, 0x9a, 0x70, 0xbc, 0xe0, 0x75, 0x8c, 0x4d, 0xec, 0xf4, 0x75, 0x7b, 0x60, 0x39, 0x58,
	0xff, 0x38, 0xd0, 0xfb, 0xb6, 0x4c, 0xd0, 0x09, 0x7c, 0xf7, 0x40, 0x31, 0x30, 0xac, 0xb6, 0x76,
	0xbd, 0x90, 0xdc, 0x3d, 0x22, 0xc9, 0x83, 0x64, 0xa0, 0x3c, 0x42, 0x2a, 0xbc, 0x7d, 0x2a, 0x8a,
	0xd0, 0x8c, 0x4b, 0x61, 0x34, 0xac, 0xb7, 0x6d, 0xdd, 0xd1, 0x6e, 0xae, 0x74, 0xc3, 0x76, 0x2e,
	0xdb, 0x57, 0x37, 0x7a, 0x47, 0x0e, 0xd4, 0xdf, 0x25, 0xa8, 0x6b, 0x51, 0x48, 0x49, 0x48, 0xed,
	0x79, 0x4c, 0x90, 0x02, 0xfb, 0x9a, 0x69, 0xd8, 0x4c, 0x63, 0xff, 0x66, 0xe9, 0xce, 0xc0, 0xb8,
	0x36, 0xcc, 0xa1, 0x21, 0xbf, 0x40, 0x07, 0xb0, 0x5b, 0x62, 0x7e, 0xed, 0x9b, 0x86, 0x2c, 0xa1,
	0x23, 0x78, 0x55, 0x82, 0x2f, 0x4d, 0xdc, 0x73, 0x3a, 0x6d, 0xbb, 0x2d, 0x57, 0xd0, 0x0f, 0xa0,
	0x96, 0xb8, 0x4f, 0xce, 0x70, 0x38, 0xcc, 0x14, 0x03, 0x7c, 0xa3, 0x1b, 0x9a, 0xd9, 0xd1, 0x3b,
	0x72, 0x55, 0xb5, 0x00, 0xba, 0x94, 0xc6, 0x3d, 0x42, 0xc7, 0x91, 0x8f, 0x0e, 0x61, 0xaf, 0x6b,
	0xdb, 0x96, 0xd3, 0xd3, 0xed, 0xae, 0xd9, 0x29, 0xbc, 0x60, 0x0f, 0x76, 0x8a, 0xc4, 0x2f, 0xba,
	0x2d, 0x4b, 0x68, 0x1f, 0xe4, 0x22, 0x68, 0x99, 0x7d, 0x5b, 0xae, 0xa8, 0xff, 0x54, 0x61, 0x03,
	0x93, 0x2f, 0x33, 0x92, 0x52, 0x74, 0x02, 0x8d, 0x31, 0xa5, 0xb1, 0xe3, 0xfa, 0x7e, 0x42, 0xd2,
	0x54, 0x91, 0x9a, 0x52, 0x6b, 0x13, 0xd7, 0x19, 0xd6, 0xce, 0x20, 0x74, 0x00, 0xeb, 0x93, 0x68,
	0xe4, 0x04, 0xbe, 0x52, 0xe1, 0xe4, 0xda, 0x24, 0x1a, 0x5d, 0xf9, 0x48, 0x81, 0x0d, 0x2f, 0x9a,
	0x4e, 0xdd, 0xd0, 0x57, 0xaa, 0x1c, 0xcf, 0xb7, 0xe8, 0x35, 0xd4, 0x82, 0xd4, 0xe1, 0x8d, 0xa4,
	0xbc, 0x6c, 0x4a, 0xad, 0x1a, 0xde, 0x08, 0x52, 0x8b, 0xf7, 0xd5, 0x39, 0xac, 0x8f, 0x89, 0xeb,
	0x93, 0x44, 0x59, 0x6b, 0x56, 0x5b, 0xf5, 0xb3, 0x37, 0xa7, 0xa2, 0xed, 0xc4, 0x7b, 0x4e, 0xbb,
	0x9c, 0xd5, 0x43, 0x9a, 0xcc, 0xb1, 0x90, 0xa2, 0x9f, 0x61, 0xed, 0xcb, 0x8c, 0x24, 0x73, 0x65,
	0x9d, 0x9f, 0x39, 0x7a, 0x78, 0xe6, 0x23, 0x23, 0xb3, 0x23, 0x99, 0x10, 0x21, 0x78, 0x79, 0x1b,
	0xf9, 0x73, 0x65, 0xa3, 0x29, 0xb5, 0x1a, 0x98, 0xaf, 0xd1, 0x07, 0x68, 0x78, 0x59, 0x2d, 0x1d,
	0x3a, 0x8f, 0x89, 0x52, 0x6b, 0x4a, 0xad, 0xed, 0xb3, 0xe3, 0x3c, 0x58, 0xde, 0xe5, 0xa7, 0x85,
	0x82, 0xe3, 0xba, 0x57, 0xa8, 0x7e, 0x0b, 0xe4, 0xc4, 0xbd, 0x77, 0x4a, 0x41, 0x36, 0x79, 0xe6,
	0xdb, 0x89, 0x7b, 0x5f, 0x38, 0x76, 0xf4, 0x0e, 0xea, 0x85, 0x3c, 0x90, 0x0c, 0xd5, 0xcf, 0x64,
	0x2e, 0xac, 0x65, 0x4b, 0xb4, 0x0f, 0x6b, 0x5f, 0xdd, 0xc9, 0x8c, 0xe4, 0x8e, 0xf2, 0xcd, 0xfb,
	0xca, 0x85, 0x74, 0x74, 0x01, 0xb0, 0x4c, 0xe7, 0x39, 0x27, 0xd5, 0x21, 0xd4, 0xf2, 0xe1, 0x64,
	0xf9, 0x7b, 0x91, 0x4f, 0xf8, 0xc1, 0x2d, 0xcc, 0xd7, 0x0b, 0x4f, 0x2a, 0x05, 0x4f, 0xbe, 0x87,
	0x2d, 0x9f, 0xdc, 0xce, 0x46, 0xce, 0x94, 0xa4, 0xa9, 0x3b, 0x22, 0xa2, 0x92, 0x0d, 0x0e, 0xf6,
	0x32, 0x4c, 0x1d, 0x64, 0x0d, 0xa8, 0x45, 0xe1, 0x5d, 0x30, 0x62, 0x15, 0x9c, 0xf2, 0x56, 0xe4,
	0xc1, 0xb7, 0x97, 0x15, 0x5c, 0x18, 0xb8, 0xec, 0x56, 0x2c, 0xa4, 0xec, 0xee, 0xd8, 0xa5, 0x63,
	0xf1, 0x68, 0xbe, 0x56, 0xff, 0x92, 0x40, 0xd6, 0xa2, 0x30, 0x24, 0x1e, 0x0d, 0xa2, 0x50, 0x44,
	0x7f, 0x03, 0x9b, 0xd3, 0x20, 0x64, 0x1e, 0x87, 0xa9, 0x78, 0x7d, 0x6d, 0x1a, 0x30, 0x36, 0x4c,
	0x39, 0xe9, 0x7e, 0x13, 0x64, 0x45, 0x90, 0xee, 0xb7, 0x8c, 0xfc, 0x11, 0x76, 0x03, 0x7f, 0x42,
	0x38, 0xeb, 0xd0, 0x60, 0x4a, 0xa2, 0x19, 0xe5, 0xe9, 0x6c, 0xe1, 0x1d, 0x46, 0x30, 0x95, 0x9d,
	0xc1, 0x4c, 0x7b, 0xef, 0x06, 0xb4, 0xac, 0x7d, 0x99, 0x69, 0x19, 0x51, 0xd4, 0x9e, 0x40, 0xc3,
	0x9b, 0x10, 0x37, 0x71, 0x62, 0x92, 0x04, 0x91, 0xaf, 0xac, 0x71, 0x59, 0x9d, 0x63, 0x16, 0x87,
	0xd4, 0x5b, 0xd8, 0xb5, 0x92, 0xc8, 0x23, 0x69, 0x1a, 0x25, 0x98, 0x8c, 0x82, 0x94, 0x95, 0xae,
	0x30, 0x1e, 0x52, 0x79, 0x3c, 0xce, 0x81, 0x8f, 0x17, 0xbb, 0xfd, 0x2e, 0x18, 0xf1, 0x44, 0xea,
	0x67, 0x28, 0xb7, 0x71, 0x69, 0x35, 0x86, 0xf1, 0x62, 0xad, 0x76, 0x01, 0x69, 0xd9, 0x79, 0x31,
	0x96, 0x96, 0x1b, 0x24, 0xff, 0x73, 0x89, 0x02, 0x1b, 0xf9, 0x48, 0x67, 0xa6, 0xe7, 0x5b, 0xd5,
	0x87, 0x2d, 0x11, 0x42, 0x78, 0x5e, 0x90, 0x4a, 0x25, 0x29, 0x7a, 0x07, 0x75, 0x6e, 0x51, 0xe9,
	0xa5, 0x4a, 0xa1, 0xe0, 0xa5, 0xe2, 0x61, 0x60, 0x62, 0xf1, 0xde, 0x3f, 0x25, 0x78, 0x95, 0x79,
	0x41, 0x92, 0x3e, 0x49, 0xbe, 0x06, 0x1e, 0xc9, 0x3f, 0x39, 0xc7, 0xb0, 0x29, 0x2e, 0x20, 0xec,
	0xc6, 0x6a, 0x6b, 0x13, 0x2f, 0x01, 0x74, 0x03, 0xfb, 0x71, 0x6e, 0xa6, 0x93, 0x64, 0x6e, 0x06,
	0x84, 0x65, 0xc1, 0x66, 0xff, 0x75, 0x7e, 0xf9, 0x8a, 0xe1, 0x78, 0x2f, 0x7e, 0x00, 0x05, 0x64,
	0x25, 0x83, 0xea, 0x33, 0x32, 0xf8, 0x04, 0x87, 0x2b, 0x09, 0x88, 0xf1, 0x7a, 0xda, 0xb1, 0x95,
	0x81, 0xaa, 0x3c, 0x32, 0x50, 0x7f, 0x48, 0x70, 0x38, 0x88, 0x7d, 0x97, 0x92, 0xe5, 0x4b, 0x73,
	0x73, 0x0c, 0x38, 0x10, 0x25, 0xcc, 0x3f, 0xc9, 0x4e, 0xec, 0x06, 0x49, 0x66, 0x54, 0xe1, 0xdb,
	0xb7, 0xda, 0x0c, 0x78, 0xcf, 0x5b, 0xc1, 0x52, 0x74, 0x01, 0x0d, 0x16, 0x47, 0x18, 0x90, 0xdb,
	0x78, 0x90, 0x87, 0x29, 0x75, 0x02, 0xae, 0x33, 0x69, 0xb6, 0x4e, 0xd5, 0x0f, 0xa0, 0xac, 0x3e,
	0x52, 0x18, 0xb0, 0x92, 0xa6, 0xf4, 0x48, 0x9a, 0xa7, 0x80, 0xba, 0xc4, 0x9d, 0xd0, 0xb1, 0x36,
	0x26, 0xde, 0xe7, 0x3c, 0xc1, 0x27, 0xbd, 0x53, 0xdf, 0xc3, 0x5e, 0x49, 0xff, 0x8c, 0xbb, 0x6e,
	0xd7, 0xf9, 0xdf, 0xe6, 0xfc, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x50, 0xa3, 0xdf, 0xcc,
	0x08, 0x00, 0x00,
}
