// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Request
	User
	Response
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

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

type Request struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Id     string  `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Age    int32   `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Height float64 `protobuf:"fixed64,3,opt,name=height" json:"height,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetHeight() float64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Response struct {
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*Response)(nil), "api.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserSerivce service

type UserSerivceClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userSerivceClient struct {
	cc *grpc.ClientConn
}

func NewUserSerivceClient(cc *grpc.ClientConn) UserSerivceClient {
	return &userSerivceClient{cc}
}

func (c *userSerivceClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/api.UserSerivce/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserSerivce service

type UserSerivceServer interface {
	Create(context.Context, *Request) (*Response, error)
}

func RegisterUserSerivceServer(s *grpc.Server, srv UserSerivceServer) {
	s.RegisterService(&_UserSerivce_serviceDesc, srv)
}

func _UserSerivce_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSerivceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserSerivce/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSerivceServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSerivce_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserSerivce",
	HandlerType: (*UserSerivceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserSerivce_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x50, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0xfd, 0x36, 0xed, 0x17, 0xed, 0x54, 0x45, 0x56, 0x94, 0xd8, 0x4b, 0x42, 0x69, 0xa1, 0x82,
	0x49, 0x6b, 0x2b, 0x82, 0xb7, 0x1a, 0x4f, 0x1e, 0x5d, 0xf1, 0xa2, 0x28, 0x6c, 0xdb, 0x21, 0x5d,
	0xb4, 0x6e, 0xdc, 0xdd, 0xb4, 0xa0, 0x78, 0xf6, 0x27, 0xf8, 0xf3, 0x02, 0xf9, 0x23, 0x4a, 0xb6,
	0xd1, 0x77, 0x58, 0x76, 0xe6, 0xbd, 0xc7, 0xcc, 0x1b, 0x68, 0xf0, 0x54, 0x44, 0xa9, 0x92, 0x46,
	0xd2, 0x1a, 0x4f, 0x45, 0xeb, 0x2c, 0x11, 0x66, 0x9e, 0x4d, 0xa2, 0xa9, 0x5c, 0xf4, 0x17, 0x2b,
	0x61, 0x9e, 0xe4, 0xaa, 0x9f, 0xc8, 0xd0, 0x2a, 0xc2, 0x25, 0x7f, 0x16, 0x33, 0x6e, 0xa4, 0xd2,
	0xfd, 0xbf, 0xef, 0xda, 0xdc, 0x1e, 0xc0, 0x06, 0xc3, 0xd7, 0x0c, 0xb5, 0xa1, 0x5d, 0xa8, 0x67,
	0x1a, 0x95, 0x47, 0x02, 0xd2, 0x6b, 0x0e, 0x1b, 0x51, 0x39, 0xe1, 0x56, 0xa3, 0x8a, 0xdd, 0x22,
	0xf7, 0x9d, 0x80, 0x30, 0x4b, 0xb7, 0x3f, 0x09, 0xd4, 0xcb, 0x36, 0x3d, 0x02, 0x47, 0xcc, 0xac,
	0xba, 0x11, 0x1f, 0x16, 0xb9, 0xbf, 0x0f, 0x7b, 0x8f, 0xf7, 0x3c, 0x7c, 0xbb, 0x08, 0xef, 0x06,
	0xe1, 0xf9, 0xc3, 0xfb, 0xe8, 0xf8, 0x64, 0xf0, 0xd1, 0x61, 0x8e, 0x98, 0xd1, 0x0e, 0xd4, 0x78,
	0x82, 0x9e, 0x13, 0x90, 0xde, 0xff, 0x98, 0x16, 0xb9, 0xbf, 0xb3, 0xfb, 0xfd, 0x0b, 0xe2, 0x7d,
	0x11, 0x56, 0xd2, 0x34, 0x02, 0x77, 0x8e, 0x22, 0x99, 0x1b, 0xaf, 0x16, 0x90, 0x1e, 0x89, 0x0f,
	0x8a, 0xdc, 0xa7, 0x57, 0xff, 0x2a, 0x5c, 0xdb, 0x77, 0x2c, 0xc7, 0xac, 0x52, 0xb5, 0x01, 0x36,
	0x19, 0xea, 0x54, 0xbe, 0x68, 0x1c, 0x9e, 0x42, 0xb3, 0x5c, 0xea, 0x06, 0x95, 0x58, 0x4e, 0x91,
	0x76, 0xc1, 0xbd, 0x54, 0xc8, 0x0d, 0xd2, 0x2d, 0x9b, 0xa3, 0xca, 0xd8, 0xda, 0xae, 0xaa, 0xb5,
	0x6b, 0xe2, 0xda, 0x23, 0x8c, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x55, 0x0d, 0xb2, 0x4e,
	0x01, 0x00, 0x00,
}
