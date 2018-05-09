// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Block struct {
	Index                int64    `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Hash                 string   `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	PrevHash             string   `protobuf:"bytes,3,opt,name=prev_hash,json=prevHash" json:"prev_hash,omitempty"`
	Timestamp            int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_master_a848202f980010d4, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Block) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Block) GetPrevHash() string {
	if m != nil {
		return m.PrevHash
	}
	return ""
}

func (m *Block) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_master_a848202f980010d4, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Master service

type MasterClient interface {
	Candidate(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Candidate(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.Master/Candidate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Candidate(context.Context, *Block) (*Empty, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Candidate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Candidate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Master/Candidate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Candidate(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Candidate",
			Handler:    _Master_Candidate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}

func init() { proto.RegisterFile("master.proto", fileDescriptor_master_a848202f980010d4) }

var fileDescriptor_master_a848202f980010d4 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2c, 0x2e,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xca, 0xe1, 0x62,
	0x75, 0xca, 0xc9, 0x4f, 0xce, 0x16, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0x32, 0x12, 0x8b, 0x33, 0x24,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x69, 0x2e, 0xce, 0x82, 0xa2, 0xd4, 0xb2,
	0x78, 0xb0, 0x04, 0x33, 0x58, 0x82, 0x03, 0x24, 0xe0, 0x01, 0x92, 0x94, 0xe1, 0xe2, 0x2c, 0xc9,
	0xcc, 0x4d, 0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0x90, 0x60, 0x01, 0x1b, 0x85, 0x10, 0x50, 0x62, 0xe7,
	0x62, 0x75, 0xcd, 0x2d, 0x28, 0xa9, 0x34, 0xd2, 0xe5, 0x62, 0xf3, 0x05, 0x3b, 0x45, 0x48, 0x99,
	0x8b, 0xd3, 0x39, 0x31, 0x2f, 0x25, 0x33, 0x25, 0xb1, 0x24, 0x55, 0x88, 0x53, 0xaf, 0x20, 0x49,
	0x0f, 0xec, 0x1e, 0x29, 0x30, 0x13, 0xac, 0x58, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x60, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x70, 0x70, 0xe8, 0xa2, 0xc0, 0x00, 0x00, 0x00,
}
