// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto

It has these top-level messages:
	VersionResponse
	HelloRequest
	HelloResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

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

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
	proto.RegisterType((*HelloRequest)(nil), "service.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "service.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Helloworld service

type HelloworldClient interface {
	GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type helloworldClient struct {
	cc *grpc.ClientConn
}

func NewHelloworldClient(cc *grpc.ClientConn) HelloworldClient {
	return &helloworldClient{cc}
}

func (c *helloworldClient) GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/service.Helloworld/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloworldClient) HelloWorld(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/service.Helloworld/HelloWorld", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Helloworld service

type HelloworldServer interface {
	GetVersion(context.Context, *google_protobuf.Empty) (*VersionResponse, error)
	HelloWorld(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterHelloworldServer(s *grpc.Server, srv HelloworldServer) {
	s.RegisterService(&_Helloworld_serviceDesc, srv)
}

func _Helloworld_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Helloworld/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).GetVersion(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Helloworld_HelloWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloworldServer).HelloWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Helloworld/HelloWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloworldServer).HelloWorld(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Helloworld_GetVersion_Handler,
		},
		{
			MethodName: "HelloWorld",
			Handler:    _Helloworld_HelloWorld_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto",
}

func init() {
	proto.RegisterFile("github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x4e, 0x2a, 0x31,
	0x14, 0xce, 0xdc, 0xdc, 0x00, 0xb7, 0x57, 0xa3, 0x36, 0x91, 0x90, 0xd1, 0x85, 0x99, 0x95, 0x11,
	0x69, 0x8d, 0xae, 0x94, 0x1d, 0x89, 0x11, 0x13, 0x63, 0x0c, 0x0b, 0x35, 0xee, 0x3a, 0x70, 0x28,
	0x8d, 0x43, 0x4f, 0x6d, 0x0b, 0x13, 0xb6, 0xbe, 0x82, 0xef, 0xe1, 0xcb, 0xf8, 0x0a, 0x3e, 0x88,
	0x61, 0xa6, 0x10, 0x02, 0x2b, 0x77, 0xe7, 0xe7, 0x3b, 0xdf, 0xf9, 0xce, 0xd7, 0x92, 0x4b, 0xa9,
	0xfc, 0x68, 0x92, 0xb2, 0x3e, 0x8e, 0x79, 0x9a, 0x23, 0x0e, 0x72, 0xb4, 0x7e, 0x74, 0xdb, 0xb9,
	0x7b, 0xe6, 0x23, 0xc8, 0x32, 0xcc, 0xd1, 0x66, 0x03, 0x6e, 0x5e, 0x25, 0x37, 0x29, 0x77, 0x60,
	0xa7, 0xaa, 0x0f, 0xcc, 0x58, 0xf4, 0x48, 0xab, 0x21, 0x8d, 0x0f, 0x24, 0xa2, 0xcc, 0x80, 0x17,
	0xe5, 0x74, 0x32, 0xe4, 0x30, 0x36, 0x7e, 0x56, 0xa2, 0xe2, 0xc3, 0xd0, 0x14, 0x46, 0x71, 0xa1,
	0x35, 0x7a, 0xe1, 0x15, 0x6a, 0x17, 0xba, 0xed, 0x95, 0xf5, 0xd9, 0x6c, 0xe8, 0x4b, 0x8e, 0x7e,
	0x4b, 0x82, 0x6e, 0x4d, 0x45, 0xa6, 0x06, 0xc2, 0x03, 0xdf, 0x08, 0xc2, 0xf0, 0xe9, 0x0a, 0xd8,
	0xe5, 0x42, 0x4a, 0xb0, 0x1c, 0x4d, 0x41, 0xbf, 0xb9, 0x2a, 0x69, 0x92, 0x9d, 0x47, 0xb0, 0x4e,
	0xa1, 0xee, 0x81, 0x33, 0xa8, 0x1d, 0xd0, 0x06, 0xa9, 0x4e, 0xcb, 0x52, 0x23, 0x3a, 0x8a, 0x8e,
	0xff, 0xf5, 0x16, 0x69, 0x92, 0x90, 0xad, 0xee, 0xfc, 0xfc, 0x1e, 0xbc, 0x4d, 0xc0, 0x79, 0x4a,
	0xc9, 0x5f, 0x2d, 0xc6, 0x10, 0x60, 0x45, 0x9c, 0x34, 0xc9, 0x76, 0xc0, 0x04, 0xba, 0x98, 0xd4,
	0xa4, 0x05, 0xf0, 0x4a, 0xcb, 0x00, 0x5c, 0xe6, 0xe7, 0x9f, 0x11, 0x21, 0xdd, 0xa5, 0xa1, 0xf4,
	0x81, 0x90, 0x1b, 0xf0, 0x41, 0x0f, 0xad, 0xb3, 0xd2, 0x24, 0xb6, 0x70, 0x90, 0x5d, 0xcf, 0x1d,
	0x8c, 0x1b, 0x6c, 0xe1, 0xf8, 0x9a, 0xf2, 0x64, 0xf7, 0xfd, 0xeb, 0xfb, 0xe3, 0x0f, 0xa1, 0x35,
	0x1e, 0x14, 0xd3, 0xfb, 0xc0, 0xff, 0x54, 0xf0, 0xef, 0x2f, 0x27, 0x57, 0xcf, 0x88, 0xeb, 0xeb,
	0xe5, 0x40, 0xb7, 0x57, 0xd0, 0xfd, 0x4f, 0x2a, 0xe5, 0xa3, 0x5f, 0x45, 0x27, 0x9d, 0xb3, 0x17,
	0xf6, 0x8b, 0xaf, 0xd1, 0x36, 0x69, 0x5a, 0x29, 0xd4, 0x5f, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff,
	0xc5, 0x8f, 0x8f, 0xc2, 0x53, 0x02, 0x00, 0x00,
}
