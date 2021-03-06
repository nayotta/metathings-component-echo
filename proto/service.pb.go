// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package ai_metathings_component_service_echo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type EchoRequest struct {
	Text                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

type EchoResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type StreamingEchoRequest struct {
	Text                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *StreamingEchoRequest) Reset()         { *m = StreamingEchoRequest{} }
func (m *StreamingEchoRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingEchoRequest) ProtoMessage()    {}
func (*StreamingEchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *StreamingEchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingEchoRequest.Unmarshal(m, b)
}
func (m *StreamingEchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingEchoRequest.Marshal(b, m, deterministic)
}
func (m *StreamingEchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingEchoRequest.Merge(m, src)
}
func (m *StreamingEchoRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingEchoRequest.Size(m)
}
func (m *StreamingEchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingEchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingEchoRequest proto.InternalMessageInfo

func (m *StreamingEchoRequest) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

type StreamingEchoResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingEchoResponse) Reset()         { *m = StreamingEchoResponse{} }
func (m *StreamingEchoResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingEchoResponse) ProtoMessage()    {}
func (*StreamingEchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *StreamingEchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingEchoResponse.Unmarshal(m, b)
}
func (m *StreamingEchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingEchoResponse.Marshal(b, m, deterministic)
}
func (m *StreamingEchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingEchoResponse.Merge(m, src)
}
func (m *StreamingEchoResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingEchoResponse.Size(m)
}
func (m *StreamingEchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingEchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingEchoResponse proto.InternalMessageInfo

func (m *StreamingEchoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "ai.metathings.component.service.echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "ai.metathings.component.service.echo.EchoResponse")
	proto.RegisterType((*StreamingEchoRequest)(nil), "ai.metathings.component.service.echo.StreamingEchoRequest")
	proto.RegisterType((*StreamingEchoResponse)(nil), "ai.metathings.component.service.echo.StreamingEchoResponse")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x8d, 0x14, 0xc1, 0xd4, 0x5e, 0x82, 0x82, 0x2c, 0x22, 0xb2, 0x78, 0x28, 0x08, 0xd3,
	0xba, 0xde, 0xf4, 0xe0, 0x49, 0xf0, 0xdc, 0x05, 0xef, 0x69, 0x18, 0xb3, 0x81, 0x6e, 0x26, 0x26,
	0xb3, 0xea, 0xaf, 0xd0, 0xbf, 0x2c, 0x6e, 0x5c, 0xa8, 0xa2, 0x50, 0xf5, 0x18, 0x78, 0xdf, 0x37,
	0x2f, 0x4f, 0x4e, 0x12, 0xc6, 0x47, 0x67, 0x10, 0x42, 0x24, 0x26, 0x75, 0xaa, 0x1d, 0xb4, 0xc8,
	0x9a, 0x1b, 0xe7, 0x6d, 0x02, 0x43, 0x6d, 0x20, 0x8f, 0x9e, 0x61, 0x88, 0xa1, 0x69, 0xa8, 0x38,
	0xb6, 0x44, 0x76, 0x85, 0xb3, 0x9e, 0x59, 0x76, 0xf7, 0xb3, 0xa7, 0xa8, 0x43, 0xc0, 0x98, 0xb2,
	0xa5, 0xbc, 0x96, 0xe3, 0x1b, 0xd3, 0xd0, 0x02, 0x1f, 0x3a, 0x4c, 0xac, 0xe6, 0x72, 0xc4, 0xf8,
	0xcc, 0x87, 0xe2, 0x44, 0x4c, 0xc7, 0xd5, 0x11, 0x64, 0x1a, 0x06, 0x1a, 0x6a, 0x8e, 0xce, 0xdb,
	0x3b, 0xbd, 0xea, 0x70, 0xd1, 0x27, 0xcb, 0x52, 0xee, 0x65, 0x41, 0x0a, 0xe4, 0x13, 0x2a, 0xb5,
	0x66, 0xd8, 0xfd, 0xc8, 0xdc, 0xca, 0xfd, 0x9a, 0x23, 0xea, 0xd6, 0x79, 0xfb, 0xbf, 0x6b, 0x67,
	0xf2, 0xe0, 0x8b, 0xe9, 0xe7, 0xb3, 0xd5, 0xeb, 0x76, 0xfe, 0x5c, 0x9d, 0x07, 0x51, 0x24, 0x47,
	0xef, 0x4f, 0x75, 0x0e, 0x9b, 0x4c, 0x07, 0x6b, 0x4d, 0x8b, 0xea, 0x37, 0x48, 0xae, 0x54, 0x6e,
	0xa9, 0x17, 0x21, 0x27, 0x9f, 0xea, 0xaa, 0xcb, 0xcd, 0x3c, 0xdf, 0xad, 0x55, 0x5c, 0xfd, 0x89,
	0x1d, 0xca, 0x4c, 0xc5, 0x5c, 0x2c, 0x77, 0xfa, 0x69, 0x2f, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb6, 0x1b, 0x43, 0xaa, 0x4b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	StreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_StreamingEchoClient, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/ai.metathings.component.service.echo.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) StreamingEcho(ctx context.Context, opts ...grpc.CallOption) (EchoService_StreamingEchoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoService_serviceDesc.Streams[0], "/ai.metathings.component.service.echo.EchoService/StreamingEcho", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServiceStreamingEchoClient{stream}
	return x, nil
}

type EchoService_StreamingEchoClient interface {
	Send(*StreamingEchoRequest) error
	Recv() (*StreamingEchoResponse, error)
	grpc.ClientStream
}

type echoServiceStreamingEchoClient struct {
	grpc.ClientStream
}

func (x *echoServiceStreamingEchoClient) Send(m *StreamingEchoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoServiceStreamingEchoClient) Recv() (*StreamingEchoResponse, error) {
	m := new(StreamingEchoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	StreamingEcho(EchoService_StreamingEchoServer) error
}

// UnimplementedEchoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (*UnimplementedEchoServiceServer) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedEchoServiceServer) StreamingEcho(srv EchoService_StreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingEcho not implemented")
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ai.metathings.component.service.echo.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_StreamingEcho_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServiceServer).StreamingEcho(&echoServiceStreamingEchoServer{stream})
}

type EchoService_StreamingEchoServer interface {
	Send(*StreamingEchoResponse) error
	Recv() (*StreamingEchoRequest, error)
	grpc.ServerStream
}

type echoServiceStreamingEchoServer struct {
	grpc.ServerStream
}

func (x *echoServiceStreamingEchoServer) Send(m *StreamingEchoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoServiceStreamingEchoServer) Recv() (*StreamingEchoRequest, error) {
	m := new(StreamingEchoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ai.metathings.component.service.echo.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingEcho",
			Handler:       _EchoService_StreamingEcho_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}
