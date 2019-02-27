// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto

It has these top-level messages:
	VersionResponse
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

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
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

// Server API for Helloworld service

type HelloworldServer interface {
	GetVersion(context.Context, *google_protobuf.Empty) (*VersionResponse, error)
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

var _Helloworld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Helloworld",
	HandlerType: (*HelloworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Helloworld_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto",
}

func init() {
	proto.RegisterFile("github.com/bwoodworthIBLX/helloworld/pkg/pb/service.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0xc9, 0x57, 0x7c, 0xd1, 0x69, 0x94, 0x2d, 0x24, 0xac, 0x16, 0x92, 0x4a, 0xd0, 0xcc,
	0x15, 0xad, 0x24, 0x5d, 0x40, 0x54, 0xb0, 0x90, 0x14, 0x22, 0x16, 0xc2, 0xcc, 0xee, 0xcd, 0xec,
	0xe0, 0x64, 0xee, 0x30, 0x33, 0xd9, 0x25, 0xad, 0xaf, 0xe0, 0xa3, 0xf9, 0x0a, 0x3e, 0x88, 0xb8,
	0x3b, 0x2b, 0xc1, 0x54, 0x76, 0xf7, 0xef, 0x39, 0xfc, 0x0e, 0xbb, 0x52, 0x3a, 0x56, 0x2b, 0xc9,
	0x0b, 0x5a, 0x82, 0x6c, 0x88, 0xca, 0x86, 0x7c, 0xac, 0xee, 0x66, 0xf7, 0x4f, 0x50, 0xa1, 0x31,
	0xd4, 0x90, 0x37, 0x25, 0xb8, 0x57, 0x05, 0x4e, 0x42, 0x40, 0x5f, 0xeb, 0x02, 0xb9, 0xf3, 0x14,
	0x29, 0x1b, 0xa6, 0x36, 0x3f, 0x54, 0x44, 0xca, 0x20, 0xb4, 0x63, 0xb9, 0x5a, 0x00, 0x2e, 0x5d,
	0x5c, 0x77, 0x57, 0xf9, 0x51, 0x5a, 0x0a, 0xa7, 0x41, 0x58, 0x4b, 0x51, 0x44, 0x4d, 0x36, 0xa4,
	0xed, 0x74, 0xc3, 0xde, 0xac, 0x17, 0xb1, 0xd3, 0x28, 0x26, 0x0a, 0xed, 0xa4, 0x16, 0x46, 0x97,
	0x22, 0x22, 0x6c, 0x15, 0xe9, 0xf9, 0x6c, 0xe3, 0x38, 0x34, 0x42, 0x29, 0xf4, 0x40, 0xae, 0x95,
	0xdf, 0xb6, 0x1a, 0x9f, 0xb2, 0xbd, 0x47, 0xf4, 0x41, 0x93, 0x9d, 0x63, 0x70, 0x64, 0x03, 0x66,
	0x23, 0x36, 0xac, 0xbb, 0xd1, 0x68, 0x70, 0x3c, 0x38, 0xd9, 0x9d, 0xf7, 0xed, 0xc5, 0x0b, 0x63,
	0xb7, 0x3f, 0xf8, 0xd9, 0x03, 0x63, 0x37, 0x18, 0xd3, 0x77, 0x76, 0xc0, 0x3b, 0x24, 0xde, 0xf3,
	0xf2, 0xeb, 0x6f, 0xde, 0x7c, 0xc4, 0xfb, 0x7c, 0x7e, 0xf9, 0x8c, 0xf7, 0xdf, 0x3e, 0x3e, 0xdf,
	0xff, 0xb1, 0x6c, 0x07, 0x92, 0xfe, 0xec, 0xfc, 0x99, 0xff, 0x21, 0xf8, 0xa9, 0x93, 0xf2, 0x7f,
	0xeb, 0x76, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x74, 0x8d, 0x19, 0x55, 0xb1, 0x01, 0x00, 0x00,
}
