// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/asim/go-micro/plugins/broker/grpc/v3/proto/broker.proto

/*
Package broker is a generated protocol buffer package.

It is generated from these files:
	github.com/asim/go-micro/plugins/broker/grpc/v3/proto/broker.proto

It has these top-level messages:
	Message
	Empty
*/
package broker

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type Message struct {
	Topic  string            `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Id     string            `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Header map[string]string `protobuf:"bytes,3,rep,name=header" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body   []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Message)(nil), "broker.Message")
	proto.RegisterType((*Empty)(nil), "broker.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Broker service

type BrokerClient interface {
	Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
}

type brokerClient struct {
	cc *grpc.ClientConn
}

func NewBrokerClient(cc *grpc.ClientConn) BrokerClient {
	return &brokerClient{cc}
}

func (c *brokerClient) Publish(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/broker.Broker/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Broker service

type BrokerServer interface {
	Publish(context.Context, *Message) (*Empty, error)
}

func RegisterBrokerServer(s *grpc.Server, srv BrokerServer) {
	s.RegisterService(&_Broker_serviceDesc, srv)
}

func _Broker_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/broker.Broker/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServer).Publish(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Broker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "broker.Broker",
	HandlerType: (*BrokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Broker_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/asim/go-micro/plugins/broker/grpc/v3/proto/broker.proto",
}

func init() {
	proto.RegisterFile("github.com/asim/go-micro/plugins/broker/grpc/v3/proto/broker.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x6b, 0x84, 0x30,
	0x10, 0xc5, 0x1b, 0xdd, 0x55, 0x3a, 0xdb, 0x7f, 0x0c, 0x3d, 0xc8, 0xf6, 0x22, 0x9e, 0x84, 0x52,
	0x03, 0xbb, 0x14, 0xda, 0x3d, 0x16, 0x16, 0x7a, 0x29, 0x14, 0xbf, 0x81, 0xd1, 0x10, 0xc3, 0xea,
	0x26, 0xc4, 0x58, 0xf0, 0x73, 0xf5, 0x0b, 0x16, 0x63, 0x0a, 0xa5, 0xb7, 0xf9, 0xbd, 0x4c, 0xde,
	0x7b, 0x0c, 0x1c, 0x84, 0xb4, 0xed, 0xc8, 0x8a, 0x5a, 0xf5, 0xb4, 0x97, 0xb5, 0x51, 0x54, 0xa8,
	0x27, 0xdd, 0x8d, 0x42, 0x9e, 0x07, 0xca, 0x8c, 0x3a, 0x71, 0x43, 0x85, 0xd1, 0x35, 0xd5, 0x46,
	0x59, 0xe5, 0x95, 0xc2, 0x01, 0x46, 0x0b, 0x65, 0xdf, 0x04, 0xe2, 0x0f, 0x3e, 0x0c, 0x95, 0xe0,
	0x78, 0x0f, 0x6b, 0xab, 0xb4, 0xac, 0x13, 0x92, 0x92, 0xfc, 0xb2, 0x5c, 0x00, 0x6f, 0x20, 0x90,
	0x4d, 0x12, 0x38, 0x29, 0x90, 0x0d, 0xee, 0x21, 0x6a, 0x79, 0xd5, 0x70, 0x93, 0x84, 0x69, 0x98,
	0x6f, 0x76, 0x0f, 0x85, 0x37, 0xf6, 0x36, 0xc5, 0xbb, 0x7b, 0x3d, 0x9e, 0xad, 0x99, 0x4a, 0xbf,
	0x8a, 0x08, 0x2b, 0xa6, 0x9a, 0x29, 0x59, 0xa5, 0x24, 0xbf, 0x2a, 0xdd, 0xbc, 0x7d, 0x85, 0xcd,
	0x9f, 0x55, 0xbc, 0x83, 0xf0, 0xc4, 0x27, 0x9f, 0x3d, 0x8f, 0x73, 0x9f, 0xaf, 0xaa, 0x1b, 0xb9,
	0x0f, 0x5f, 0xe0, 0x10, 0xbc, 0x90, 0x2c, 0x86, 0xf5, 0xb1, 0xd7, 0x76, 0xda, 0x3d, 0x43, 0xf4,
	0xe6, 0xd2, 0xf1, 0x11, 0xe2, 0xcf, 0x91, 0x75, 0x72, 0x68, 0xf1, 0xf6, 0x5f, 0xa3, 0xed, 0xf5,
	0xaf, 0xe0, 0x3e, 0x65, 0x17, 0x2c, 0x72, 0x47, 0xd8, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6f,
	0x55, 0xde, 0x23, 0x42, 0x01, 0x00, 0x00,
}
