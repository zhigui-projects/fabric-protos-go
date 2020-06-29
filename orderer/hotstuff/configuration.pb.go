// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/hotstuff/configuration.proto

package hotstuff

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

// ConfigMetadata is serialized and set as the value of ConsensusType.Metadata in
// a channel configuration when the ConsensusType.Type is set "hotstuff".
type ConfigMetadata struct {
	Consenters           []*Consenter `protobuf:"bytes,1,rep,name=consenters,proto3" json:"consenters,omitempty"`
	Options              *Options     `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfigMetadata) Reset()         { *m = ConfigMetadata{} }
func (m *ConfigMetadata) String() string { return proto.CompactTextString(m) }
func (*ConfigMetadata) ProtoMessage()    {}
func (*ConfigMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1f9c5da6505029, []int{0}
}

func (m *ConfigMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigMetadata.Unmarshal(m, b)
}
func (m *ConfigMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigMetadata.Marshal(b, m, deterministic)
}
func (m *ConfigMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigMetadata.Merge(m, src)
}
func (m *ConfigMetadata) XXX_Size() int {
	return xxx_messageInfo_ConfigMetadata.Size(m)
}
func (m *ConfigMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigMetadata proto.InternalMessageInfo

func (m *ConfigMetadata) GetConsenters() []*Consenter {
	if m != nil {
		return m.Consenters
	}
	return nil
}

func (m *ConfigMetadata) GetOptions() *Options {
	if m != nil {
		return m.Options
	}
	return nil
}

// Consenter represents a consenting node (i.e. replica).
type Consenter struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	SignCert             []byte   `protobuf:"bytes,3,opt,name=sign_cert,json=signCert,proto3" json:"sign_cert,omitempty"`
	ClientTlsCert        []byte   `protobuf:"bytes,4,opt,name=client_tls_cert,json=clientTlsCert,proto3" json:"client_tls_cert,omitempty"`
	ServerTlsCert        []byte   `protobuf:"bytes,5,opt,name=server_tls_cert,json=serverTlsCert,proto3" json:"server_tls_cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consenter) Reset()         { *m = Consenter{} }
func (m *Consenter) String() string { return proto.CompactTextString(m) }
func (*Consenter) ProtoMessage()    {}
func (*Consenter) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1f9c5da6505029, []int{1}
}

func (m *Consenter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consenter.Unmarshal(m, b)
}
func (m *Consenter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consenter.Marshal(b, m, deterministic)
}
func (m *Consenter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consenter.Merge(m, src)
}
func (m *Consenter) XXX_Size() int {
	return xxx_messageInfo_Consenter.Size(m)
}
func (m *Consenter) XXX_DiscardUnknown() {
	xxx_messageInfo_Consenter.DiscardUnknown(m)
}

var xxx_messageInfo_Consenter proto.InternalMessageInfo

func (m *Consenter) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Consenter) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Consenter) GetSignCert() []byte {
	if m != nil {
		return m.SignCert
	}
	return nil
}

func (m *Consenter) GetClientTlsCert() []byte {
	if m != nil {
		return m.ClientTlsCert
	}
	return nil
}

func (m *Consenter) GetServerTlsCert() []byte {
	if m != nil {
		return m.ServerTlsCert
	}
	return nil
}

// Defines the hotstuff parameters when 'hotstuff' is specified as the 'OrdererType'
type Options struct {
	// Number of peers
	N int64 `protobuf:"varint,1,opt,name=n,proto3" json:"n,omitempty"`
	// Fault tolerance
	F int64 `protobuf:"varint,2,opt,name=f,proto3" json:"f,omitempty"`
	// Timeout of requests (seconds)
	RequestTimeoutSec    int64    `protobuf:"varint,3,opt,name=request_timeout_sec,json=requestTimeoutSec,proto3" json:"request_timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Options) Reset()         { *m = Options{} }
func (m *Options) String() string { return proto.CompactTextString(m) }
func (*Options) ProtoMessage()    {}
func (*Options) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e1f9c5da6505029, []int{2}
}

func (m *Options) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Options.Unmarshal(m, b)
}
func (m *Options) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Options.Marshal(b, m, deterministic)
}
func (m *Options) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Options.Merge(m, src)
}
func (m *Options) XXX_Size() int {
	return xxx_messageInfo_Options.Size(m)
}
func (m *Options) XXX_DiscardUnknown() {
	xxx_messageInfo_Options.DiscardUnknown(m)
}

var xxx_messageInfo_Options proto.InternalMessageInfo

func (m *Options) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *Options) GetF() int64 {
	if m != nil {
		return m.F
	}
	return 0
}

func (m *Options) GetRequestTimeoutSec() int64 {
	if m != nil {
		return m.RequestTimeoutSec
	}
	return 0
}

func init() {
	proto.RegisterType((*ConfigMetadata)(nil), "hotstuff.ConfigMetadata")
	proto.RegisterType((*Consenter)(nil), "hotstuff.Consenter")
	proto.RegisterType((*Options)(nil), "hotstuff.Options")
}

func init() {
	proto.RegisterFile("orderer/hotstuff/configuration.proto", fileDescriptor_4e1f9c5da6505029)
}

var fileDescriptor_4e1f9c5da6505029 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0xab, 0xd3, 0x40,
	0x14, 0xc5, 0x19, 0xf3, 0xf4, 0xb5, 0xd3, 0x56, 0xe9, 0x74, 0x13, 0x70, 0x13, 0x8a, 0x48, 0x40,
	0x3a, 0x81, 0x76, 0xe3, 0xda, 0xae, 0x45, 0x18, 0xeb, 0xc6, 0x4d, 0x48, 0xa7, 0x37, 0x7f, 0x24,
	0x9d, 0x89, 0x77, 0x6e, 0x04, 0xbf, 0x8d, 0x1f, 0x55, 0x32, 0x93, 0xb4, 0xe5, 0xed, 0x6e, 0xce,
	0xf9, 0x9d, 0x9b, 0xc3, 0x5c, 0xfe, 0xc1, 0xe2, 0x05, 0x10, 0x30, 0xab, 0x2d, 0x39, 0xea, 0xcb,
	0x32, 0xd3, 0xd6, 0x94, 0x4d, 0xd5, 0x63, 0x41, 0x8d, 0x35, 0xb2, 0x43, 0x4b, 0x56, 0xcc, 0x26,
	0x77, 0x8b, 0xfc, 0xed, 0xd1, 0x03, 0x5f, 0x81, 0x8a, 0x4b, 0x41, 0x85, 0x38, 0x70, 0xae, 0xad,
	0x71, 0x60, 0x08, 0xd0, 0xc5, 0x2c, 0x89, 0xd2, 0xc5, 0x7e, 0x23, 0xa7, 0x80, 0x3c, 0x4e, 0x9e,
	0x7a, 0xc0, 0xc4, 0x27, 0xfe, 0x6c, 0xbb, 0xe1, 0x07, 0x2e, 0x7e, 0x95, 0xb0, 0x74, 0xb1, 0x5f,
	0xdf, 0x13, 0xdf, 0x82, 0xa1, 0x26, 0x62, 0xfb, 0x8f, 0xf1, 0xf9, 0x6d, 0x8d, 0x10, 0xfc, 0xa9,
	0xb6, 0x8e, 0x62, 0x96, 0xb0, 0x74, 0xae, 0xfc, 0x3c, 0x68, 0x9d, 0x45, 0xf2, 0xbb, 0x56, 0xca,
	0xcf, 0xe2, 0x3d, 0x9f, 0xbb, 0xa6, 0x32, 0xb9, 0x06, 0xa4, 0x38, 0x4a, 0x58, 0xba, 0x54, 0xb3,
	0x41, 0x38, 0x02, 0x92, 0xf8, 0xc8, 0xdf, 0xe9, 0xb6, 0x01, 0x43, 0x39, 0xb5, 0x2e, 0x20, 0x4f,
	0x1e, 0x59, 0x05, 0xf9, 0xd4, 0xba, 0x89, 0x73, 0x80, 0x7f, 0x00, 0xef, 0xdc, 0xeb, 0xc0, 0x05,
	0x79, 0xe4, 0xb6, 0x3f, 0xf8, 0xf3, 0x58, 0x5b, 0x2c, 0x39, 0x33, 0xbe, 0x5c, 0xa4, 0x98, 0x19,
	0xbe, 0x4a, 0x5f, 0x2b, 0x52, 0xac, 0x14, 0x92, 0x6f, 0x10, 0x7e, 0xf7, 0xe0, 0x28, 0xa7, 0xe6,
	0x0a, 0xb6, 0xa7, 0xdc, 0x81, 0xf6, 0xed, 0x22, 0xb5, 0x1e, 0xad, 0x53, 0x70, 0xbe, 0x83, 0xfe,
	0xf2, 0x8b, 0x4b, 0x8b, 0x95, 0xac, 0xff, 0x76, 0x80, 0x2d, 0x5c, 0x2a, 0x40, 0x59, 0x16, 0x67,
	0x6c, 0x74, 0xb8, 0x8b, 0x93, 0xe3, 0xf5, 0x6e, 0x8f, 0xf7, 0xf3, 0x73, 0xd5, 0x50, 0xdd, 0x9f,
	0xa5, 0xb6, 0xd7, 0xec, 0x21, 0x96, 0x85, 0xd8, 0x2e, 0xc4, 0x76, 0x95, 0xcd, 0x5e, 0xde, 0xfd,
	0xfc, 0xc6, 0x7b, 0x87, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x88, 0xb9, 0x71, 0x12, 0x02,
	0x00, 0x00,
}
