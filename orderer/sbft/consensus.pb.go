// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/sbft/consensus.proto

package sbft

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Handshake struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a6e98e76fb74366, []int{0}
}

func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (m *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(m, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Handshake)(nil), "sbft.handshake")
}

func init() { proto.RegisterFile("orderer/sbft/consensus.proto", fileDescriptor_8a6e98e76fb74366) }

var fileDescriptor_8a6e98e76fb74366 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x4d, 0xca, 0x83, 0x30,
	0x10, 0x40, 0xbf, 0x0f, 0x4a, 0xc1, 0x74, 0x51, 0xb0, 0x3b, 0xe9, 0xca, 0x65, 0xc1, 0xa4, 0xf4,
	0xe7, 0x02, 0x76, 0xed, 0x05, 0xba, 0x4b, 0x74, 0x4c, 0x42, 0x35, 0x91, 0x99, 0xb8, 0xe8, 0xed,
	0x8b, 0xa6, 0x88, 0x5d, 0xe6, 0xf1, 0x5e, 0x66, 0x86, 0x1d, 0x3d, 0x36, 0x80, 0x80, 0x82, 0x54,
	0x1b, 0x44, 0xed, 0x1d, 0x81, 0xa3, 0x91, 0xf8, 0x80, 0x3e, 0xf8, 0x74, 0x33, 0xd1, 0xec, 0xd7,
	0x21, 0xdb, 0x0f, 0x1d, 0xa8, 0x36, 0x44, 0x27, 0xdf, 0xb1, 0xc4, 0x48, 0xd7, 0x90, 0x91, 0x2f,
	0xb8, 0x94, 0x2c, 0x59, 0xfe, 0x48, 0xef, 0xeb, 0xc7, 0x9e, 0x4f, 0x35, 0x5f, 0xd4, 0xec, 0x10,
	0x41, 0x35, 0x76, 0xc1, 0x3e, 0x8c, 0xb4, 0xae, 0x22, 0x9d, 0xff, 0x9d, 0xff, 0x4b, 0xc5, 0x4e,
	0x1e, 0x35, 0x37, 0xef, 0x01, 0xb0, 0x83, 0x46, 0x03, 0xf2, 0x56, 0x2a, 0xb4, 0x75, 0x1c, 0x48,
	0xfc, 0xbb, 0xce, 0xdc, 0x3f, 0x6f, 0xda, 0x06, 0x33, 0x2a, 0x5e, 0xfb, 0x5e, 0xac, 0x12, 0x11,
	0x93, 0x22, 0x26, 0x85, 0xf6, 0x62, 0x7d, 0x84, 0xda, 0xce, 0xfc, 0xfa, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x34, 0x90, 0xd7, 0x56, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsensusClient is the client API for Consensus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsensusClient interface {
	Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error)
}

type consensusClient struct {
	cc *grpc.ClientConn
}

func NewConsensusClient(cc *grpc.ClientConn) ConsensusClient {
	return &consensusClient{cc}
}

func (c *consensusClient) Consensus(ctx context.Context, in *Handshake, opts ...grpc.CallOption) (Consensus_ConsensusClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Consensus_serviceDesc.Streams[0], "/sbft.consensus/consensus", opts...)
	if err != nil {
		return nil, err
	}
	x := &consensusConsensusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Consensus_ConsensusClient interface {
	Recv() (*MultiChainMsg, error)
	grpc.ClientStream
}

type consensusConsensusClient struct {
	grpc.ClientStream
}

func (x *consensusConsensusClient) Recv() (*MultiChainMsg, error) {
	m := new(MultiChainMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConsensusServer is the server API for Consensus service.
type ConsensusServer interface {
	Consensus(*Handshake, Consensus_ConsensusServer) error
}

// UnimplementedConsensusServer can be embedded to have forward compatible implementations.
type UnimplementedConsensusServer struct {
}

func (*UnimplementedConsensusServer) Consensus(req *Handshake, srv Consensus_ConsensusServer) error {
	return status.Errorf(codes.Unimplemented, "method Consensus not implemented")
}

func RegisterConsensusServer(s *grpc.Server, srv ConsensusServer) {
	s.RegisterService(&_Consensus_serviceDesc, srv)
}

func _Consensus_Consensus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Handshake)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConsensusServer).Consensus(m, &consensusConsensusServer{stream})
}

type Consensus_ConsensusServer interface {
	Send(*MultiChainMsg) error
	grpc.ServerStream
}

type consensusConsensusServer struct {
	grpc.ServerStream
}

func (x *consensusConsensusServer) Send(m *MultiChainMsg) error {
	return x.ServerStream.SendMsg(m)
}

var _Consensus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sbft.consensus",
	HandlerType: (*ConsensusServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "consensus",
			Handler:       _Consensus_Consensus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "orderer/sbft/consensus.proto",
}
