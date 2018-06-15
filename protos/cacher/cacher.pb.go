// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cacher.proto

package cacher

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

type PushRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cacher_a4ea9ef9a8c3d1cf, []int{0}
}
func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (dst *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(dst, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
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
	return fileDescriptor_cacher_a4ea9ef9a8c3d1cf, []int{1}
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

type GetRequest struct {
	MAC                  string   `protobuf:"bytes,1,opt,name=MAC" json:"MAC,omitempty"`
	IP                   string   `protobuf:"bytes,2,opt,name=IP" json:"IP,omitempty"`
	ID                   string   `protobuf:"bytes,3,opt,name=ID" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cacher_a4ea9ef9a8c3d1cf, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetMAC() string {
	if m != nil {
		return m.MAC
	}
	return ""
}

func (m *GetRequest) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *GetRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Hardware struct {
	JSON                 string   `protobuf:"bytes,1,opt,name=JSON" json:"JSON,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hardware) Reset()         { *m = Hardware{} }
func (m *Hardware) String() string { return proto.CompactTextString(m) }
func (*Hardware) ProtoMessage()    {}
func (*Hardware) Descriptor() ([]byte, []int) {
	return fileDescriptor_cacher_a4ea9ef9a8c3d1cf, []int{3}
}
func (m *Hardware) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hardware.Unmarshal(m, b)
}
func (m *Hardware) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hardware.Marshal(b, m, deterministic)
}
func (dst *Hardware) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hardware.Merge(dst, src)
}
func (m *Hardware) XXX_Size() int {
	return xxx_messageInfo_Hardware.Size(m)
}
func (m *Hardware) XXX_DiscardUnknown() {
	xxx_messageInfo_Hardware.DiscardUnknown(m)
}

var xxx_messageInfo_Hardware proto.InternalMessageInfo

func (m *Hardware) GetJSON() string {
	if m != nil {
		return m.JSON
	}
	return ""
}

func init() {
	proto.RegisterType((*PushRequest)(nil), "cacher.PushRequest")
	proto.RegisterType((*Empty)(nil), "cacher.Empty")
	proto.RegisterType((*GetRequest)(nil), "cacher.GetRequest")
	proto.RegisterType((*Hardware)(nil), "cacher.Hardware")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cacher service

type CacherClient interface {
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*Empty, error)
	ByMAC(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	ByIP(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	ByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error)
	All(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Cacher_AllClient, error)
	Ingest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type cacherClient struct {
	cc *grpc.ClientConn
}

func NewCacherClient(cc *grpc.ClientConn) CacherClient {
	return &cacherClient{cc}
}

func (c *cacherClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/cacher.Cacher/Push", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) ByMAC(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := grpc.Invoke(ctx, "/cacher.Cacher/ByMAC", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) ByIP(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := grpc.Invoke(ctx, "/cacher.Cacher/ByIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) ByID(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Hardware, error) {
	out := new(Hardware)
	err := grpc.Invoke(ctx, "/cacher.Cacher/ByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacherClient) All(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Cacher_AllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Cacher_serviceDesc.Streams[0], c.cc, "/cacher.Cacher/All", opts...)
	if err != nil {
		return nil, err
	}
	x := &cacherAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Cacher_AllClient interface {
	Recv() (*Hardware, error)
	grpc.ClientStream
}

type cacherAllClient struct {
	grpc.ClientStream
}

func (x *cacherAllClient) Recv() (*Hardware, error) {
	m := new(Hardware)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cacherClient) Ingest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/cacher.Cacher/Ingest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cacher service

type CacherServer interface {
	Push(context.Context, *PushRequest) (*Empty, error)
	ByMAC(context.Context, *GetRequest) (*Hardware, error)
	ByIP(context.Context, *GetRequest) (*Hardware, error)
	ByID(context.Context, *GetRequest) (*Hardware, error)
	All(*Empty, Cacher_AllServer) error
	Ingest(context.Context, *Empty) (*Empty, error)
}

func RegisterCacherServer(s *grpc.Server, srv CacherServer) {
	s.RegisterService(&_Cacher_serviceDesc, srv)
}

func _Cacher_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_ByMAC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).ByMAC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/ByMAC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).ByMAC(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_ByIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).ByIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/ByIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).ByIP(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_ByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).ByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/ByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).ByID(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cacher_All_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CacherServer).All(m, &cacherAllServer{stream})
}

type Cacher_AllServer interface {
	Send(*Hardware) error
	grpc.ServerStream
}

type cacherAllServer struct {
	grpc.ServerStream
}

func (x *cacherAllServer) Send(m *Hardware) error {
	return x.ServerStream.SendMsg(m)
}

func _Cacher_Ingest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacherServer).Ingest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cacher.Cacher/Ingest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacherServer).Ingest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cacher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cacher.Cacher",
	HandlerType: (*CacherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Cacher_Push_Handler,
		},
		{
			MethodName: "ByMAC",
			Handler:    _Cacher_ByMAC_Handler,
		},
		{
			MethodName: "ByIP",
			Handler:    _Cacher_ByIP_Handler,
		},
		{
			MethodName: "ByID",
			Handler:    _Cacher_ByID_Handler,
		},
		{
			MethodName: "Ingest",
			Handler:    _Cacher_Ingest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "All",
			Handler:       _Cacher_All_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cacher.proto",
}

func init() { proto.RegisterFile("cacher.proto", fileDescriptor_cacher_a4ea9ef9a8c3d1cf) }

var fileDescriptor_cacher_a4ea9ef9a8c3d1cf = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4e, 0x4c, 0xce,
	0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x14, 0xb9, 0xb8,
	0x03, 0x4a, 0x8b, 0x33, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x52,
	0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x76, 0x2e, 0x56,
	0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x25, 0x3b, 0x2e, 0x2e, 0xf7, 0xd4, 0x12, 0x98, 0x52, 0x01, 0x2e,
	0x66, 0x5f, 0x47, 0x67, 0xa8, 0x4a, 0x10, 0x53, 0x88, 0x8f, 0x8b, 0xc9, 0x33, 0x40, 0x82, 0x09,
	0x2c, 0xc0, 0xe4, 0x19, 0x00, 0xe6, 0xbb, 0x48, 0x30, 0x43, 0xf9, 0x2e, 0x4a, 0x72, 0x5c, 0x1c,
	0x1e, 0x89, 0x45, 0x29, 0xe5, 0x89, 0x45, 0xa9, 0x20, 0x8b, 0xbc, 0x82, 0xfd, 0xfd, 0x60, 0x16,
	0x81, 0xd8, 0x46, 0x13, 0x99, 0xb8, 0xd8, 0x9c, 0xc1, 0xce, 0x12, 0xd2, 0xe2, 0x62, 0x01, 0x39,
	0x4b, 0x48, 0x58, 0x0f, 0xea, 0x6a, 0x24, 0x47, 0x4a, 0xf1, 0xc2, 0x04, 0xc1, 0xce, 0x12, 0xd2,
	0xe5, 0x62, 0x75, 0xaa, 0x04, 0xd9, 0x2f, 0x04, 0x13, 0x47, 0xb8, 0x52, 0x4a, 0x00, 0x26, 0x06,
	0xb7, 0x59, 0x87, 0x8b, 0xc5, 0xa9, 0xd2, 0x33, 0x80, 0x24, 0xd5, 0x2e, 0x44, 0xaa, 0xd6, 0xe0,
	0x62, 0x76, 0xcc, 0xc9, 0x11, 0x42, 0x75, 0x20, 0xa6, 0x3a, 0x03, 0x46, 0x21, 0x35, 0x2e, 0x36,
	0xcf, 0xbc, 0x74, 0x50, 0x38, 0xa2, 0x29, 0x46, 0xe5, 0x26, 0xb1, 0x81, 0xa3, 0xcb, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x65, 0x77, 0x4f, 0xd3, 0xbe, 0x01, 0x00, 0x00,
}
