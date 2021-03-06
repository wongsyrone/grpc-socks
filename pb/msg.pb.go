// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/msg.proto

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

type Payload struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_1cc3c4c731d1ef0f, []int{0}
}
func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (dst *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(dst, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type IPAddr struct {
	// option map_entry = true;
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Zone                 string   `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	CreateAt             int64    `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	CreateAt1            int64    `protobuf:"varint,5,opt,name=create_at1,json=createAt1,proto3" json:"create_at1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IPAddr) Reset()         { *m = IPAddr{} }
func (m *IPAddr) String() string { return proto.CompactTextString(m) }
func (*IPAddr) ProtoMessage()    {}
func (*IPAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_msg_1cc3c4c731d1ef0f, []int{1}
}
func (m *IPAddr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPAddr.Unmarshal(m, b)
}
func (m *IPAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPAddr.Marshal(b, m, deterministic)
}
func (dst *IPAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPAddr.Merge(dst, src)
}
func (m *IPAddr) XXX_Size() int {
	return xxx_messageInfo_IPAddr.Size(m)
}
func (m *IPAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_IPAddr.DiscardUnknown(m)
}

var xxx_messageInfo_IPAddr proto.InternalMessageInfo

func (m *IPAddr) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPAddr) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *IPAddr) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *IPAddr) GetCreateAt() int64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

func (m *IPAddr) GetCreateAt1() int64 {
	if m != nil {
		return m.CreateAt1
	}
	return 0
}

func init() {
	proto.RegisterType((*Payload)(nil), "pb.Payload")
	proto.RegisterType((*IPAddr)(nil), "pb.IPAddr")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProxyClient is the client API for Proxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyClient interface {
	Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error)
	ResolveIP(ctx context.Context, in *IPAddr, opts ...grpc.CallOption) (*IPAddr, error)
	Pump(ctx context.Context, opts ...grpc.CallOption) (Proxy_PumpClient, error)
	PipelineUDP(ctx context.Context, opts ...grpc.CallOption) (Proxy_PipelineUDPClient, error)
}

type proxyClient struct {
	cc *grpc.ClientConn
}

func NewProxyClient(cc *grpc.ClientConn) ProxyClient {
	return &proxyClient{cc}
}

func (c *proxyClient) Echo(ctx context.Context, in *Payload, opts ...grpc.CallOption) (*Payload, error) {
	out := new(Payload)
	err := c.cc.Invoke(ctx, "/pb.Proxy/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) ResolveIP(ctx context.Context, in *IPAddr, opts ...grpc.CallOption) (*IPAddr, error) {
	out := new(IPAddr)
	err := c.cc.Invoke(ctx, "/pb.Proxy/ResolveIP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyClient) Pump(ctx context.Context, opts ...grpc.CallOption) (Proxy_PumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Proxy_serviceDesc.Streams[0], "/pb.Proxy/Pump", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyPumpClient{stream}
	return x, nil
}

type Proxy_PumpClient interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type proxyPumpClient struct {
	grpc.ClientStream
}

func (x *proxyPumpClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyPumpClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyClient) PipelineUDP(ctx context.Context, opts ...grpc.CallOption) (Proxy_PipelineUDPClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Proxy_serviceDesc.Streams[1], "/pb.Proxy/PipelineUDP", opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyPipelineUDPClient{stream}
	return x, nil
}

type Proxy_PipelineUDPClient interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ClientStream
}

type proxyPipelineUDPClient struct {
	grpc.ClientStream
}

func (x *proxyPipelineUDPClient) Send(m *Payload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *proxyPipelineUDPClient) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProxyServer is the server API for Proxy service.
type ProxyServer interface {
	Echo(context.Context, *Payload) (*Payload, error)
	ResolveIP(context.Context, *IPAddr) (*IPAddr, error)
	Pump(Proxy_PumpServer) error
	PipelineUDP(Proxy_PipelineUDPServer) error
}

func RegisterProxyServer(s *grpc.Server, srv ProxyServer) {
	s.RegisterService(&_Proxy_serviceDesc, srv)
}

func _Proxy_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Proxy/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).Echo(ctx, req.(*Payload))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_ResolveIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IPAddr)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServer).ResolveIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Proxy/ResolveIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServer).ResolveIP(ctx, req.(*IPAddr))
	}
	return interceptor(ctx, in, info, handler)
}

func _Proxy_Pump_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).Pump(&proxyPumpServer{stream})
}

type Proxy_PumpServer interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type proxyPumpServer struct {
	grpc.ServerStream
}

func (x *proxyPumpServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyPumpServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Proxy_PipelineUDP_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProxyServer).PipelineUDP(&proxyPipelineUDPServer{stream})
}

type Proxy_PipelineUDPServer interface {
	Send(*Payload) error
	Recv() (*Payload, error)
	grpc.ServerStream
}

type proxyPipelineUDPServer struct {
	grpc.ServerStream
}

func (x *proxyPipelineUDPServer) Send(m *Payload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *proxyPipelineUDPServer) Recv() (*Payload, error) {
	m := new(Payload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Proxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Proxy",
	HandlerType: (*ProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Proxy_Echo_Handler,
		},
		{
			MethodName: "ResolveIP",
			Handler:    _Proxy_ResolveIP_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pump",
			Handler:       _Proxy_Pump_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "PipelineUDP",
			Handler:       _Proxy_PipelineUDP_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/msg.proto",
}

func init() { proto.RegisterFile("pb/msg.proto", fileDescriptor_msg_1cc3c4c731d1ef0f) }

var fileDescriptor_msg_1cc3c4c731d1ef0f = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x4f, 0x4e, 0x02, 0x31,
	0x14, 0xc6, 0xe9, 0x30, 0x80, 0xf3, 0x60, 0xf5, 0x56, 0x0d, 0x09, 0x91, 0x34, 0xd1, 0xcc, 0xc6,
	0x11, 0xf4, 0x04, 0x18, 0x5d, 0xb0, 0x6b, 0x9a, 0xb8, 0x36, 0x1d, 0xfa, 0xa2, 0x24, 0x03, 0x6d,
	0xda, 0x6a, 0xc4, 0x23, 0x78, 0x13, 0x6f, 0x69, 0xa6, 0x08, 0xb2, 0xd2, 0xdd, 0xf7, 0xe7, 0xf7,
	0x35, 0x6d, 0x61, 0xe4, 0xea, 0xeb, 0x4d, 0x78, 0xae, 0x9c, 0xb7, 0xd1, 0x62, 0xe6, 0x6a, 0x31,
	0x81, 0x81, 0xd4, 0xbb, 0xc6, 0x6a, 0x83, 0x08, 0xb9, 0xd1, 0x51, 0x73, 0x36, 0x65, 0xe5, 0x48,
	0x25, 0x2d, 0x3e, 0x19, 0xf4, 0x97, 0x72, 0x61, 0x8c, 0x47, 0x0e, 0x03, 0x6d, 0x8c, 0xa7, 0x10,
	0x12, 0x51, 0xa8, 0x83, 0x3d, 0x0e, 0xb3, 0xdf, 0x61, 0x9b, 0x7d, 0xd8, 0x2d, 0xf1, 0x6e, 0x42,
	0x93, 0xc6, 0x73, 0x28, 0x56, 0x9e, 0x74, 0xa4, 0x27, 0x1d, 0x79, 0x3e, 0x65, 0x65, 0xf7, 0x2e,
	0x9b, 0x65, 0xea, 0x6c, 0x1f, 0x2e, 0x22, 0x4e, 0x00, 0x8e, 0xc0, 0x9c, 0xf7, 0x5a, 0x42, 0x15,
	0x87, 0x76, 0x7e, 0xf3, 0xc5, 0xa0, 0x27, 0xbd, 0x7d, 0xdf, 0xa1, 0x80, 0xfc, 0x61, 0xf5, 0x62,
	0x71, 0x58, 0xb9, 0xba, 0xfa, 0xb9, 0xff, 0xf8, 0xd4, 0x88, 0x0e, 0x5e, 0x40, 0xa1, 0x28, 0xd8,
	0xe6, 0x8d, 0x96, 0x12, 0xa1, 0xed, 0xf6, 0x0f, 0x19, 0x9f, 0x68, 0xd1, 0xc1, 0x4b, 0xc8, 0xe5,
	0xeb, 0xc6, 0xfd, 0x75, 0x54, 0xc9, 0x66, 0x0c, 0xaf, 0x60, 0x28, 0xd7, 0x8e, 0x9a, 0xf5, 0x96,
	0x1e, 0xef, 0xe5, 0x7f, 0x78, 0xdd, 0x4f, 0x5f, 0x7c, 0xfb, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf2,
	0xb8, 0xd1, 0x42, 0x72, 0x01, 0x00, 0x00,
}
