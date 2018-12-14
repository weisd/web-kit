// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/protobuf/user/api.proto

package user

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/googleapis/google/api"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_f8bb12704ea300e0, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "user.LoginRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServiceClient interface {
	// Login 创建用户
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error)
}

type aPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAPIServiceClient(cc *grpc.ClientConn) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.APIService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
type APIServiceServer interface {
	// Login 创建用户
	Login(context.Context, *LoginRequest) (*User, error)
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.APIService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _APIService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobuf/user/api.proto",
}

func init() { proto.RegisterFile("api/protobuf/user/api.proto", fileDescriptor_api_f8bb12704ea300e0) }

var fileDescriptor_api_f8bb12704ea300e0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0x95, 0x7e, 0x5f, 0xf9, 0x63, 0xb1, 0xf2, 0xa6, 0x55, 0x28, 0x52, 0xd4, 0x15, 0xaa,
	0x48, 0x4c, 0x8b, 0xc4, 0xa2, 0xac, 0x0a, 0x08, 0x04, 0x62, 0x51, 0x15, 0x38, 0x80, 0x93, 0x1a,
	0xd7, 0x6a, 0x9b, 0x31, 0xf6, 0xa4, 0x11, 0x5b, 0x8e, 0x00, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0x65,
	0x5a, 0x4a, 0x25, 0xd8, 0x24, 0xf3, 0xe6, 0xfd, 0xec, 0xe7, 0x19, 0xb6, 0x2f, 0xad, 0x11, 0xd6,
	0x01, 0x42, 0x5a, 0x3c, 0x89, 0xc2, 0x2b, 0x27, 0xa4, 0x35, 0x09, 0x75, 0xf8, 0xff, 0x4a, 0x87,
	0x2d, 0x0d, 0xa0, 0x67, 0xaa, 0xea, 0x0b, 0x99, 0xe7, 0x80, 0x12, 0x0d, 0xe4, 0x7e, 0xc9, 0x84,
	0x47, 0xf4, 0xcb, 0x62, 0xad, 0xf2, 0xd8, 0x97, 0x52, 0x6b, 0xe5, 0x04, 0x58, 0x22, 0xfe, 0xa0,
	0x63, 0x6d, 0x70, 0x52, 0xa4, 0x49, 0x06, 0x73, 0xa1, 0x41, 0xc3, 0x4f, 0x74, 0xa5, 0x48, 0x50,
	0xb5, 0xc2, 0x4f, 0x37, 0xf0, 0x79, 0x69, 0x70, 0x0a, 0xa5, 0xd0, 0x10, 0x93, 0x19, 0x2f, 0xe4,
	0xcc, 0x8c, 0x25, 0x82, 0xf3, 0x62, 0x5d, 0xae, 0xce, 0xb5, 0x7e, 0x4f, 0x55, 0x7d, 0x96, 0x6e,
	0xfb, 0x92, 0xed, 0xdd, 0x81, 0x36, 0xf9, 0x48, 0x3d, 0x17, 0xca, 0x23, 0x6f, 0xb2, 0x6d, 0x99,
	0x65, 0x50, 0xe4, 0xd8, 0x0c, 0xa2, 0xe0, 0x70, 0x77, 0xf4, 0x2d, 0x79, 0xc8, 0x76, 0xac, 0xf4,
	0xbe, 0x04, 0x37, 0x6e, 0xd6, 0xc8, 0x5a, 0xeb, 0xde, 0x03, 0x63, 0x83, 0xe1, 0xcd, 0xbd, 0x72,
	0x0b, 0x93, 0x29, 0x7e, 0xc5, 0xea, 0x74, 0x27, 0xe7, 0x09, 0x25, 0x6d, 0x06, 0x84, 0x6c, 0xd9,
	0x7b, 0xf4, 0xca, 0xb5, 0x0f, 0x5e, 0x3f, 0x3e, 0xdf, 0x6b, 0x8d, 0x36, 0xa7, 0x7d, 0xd2, 0xd3,
	0x16, 0x5d, 0x31, 0xab, 0xf0, 0x7e, 0xd0, 0x39, 0xbf, 0x7d, 0x1b, 0x5c, 0xf3, 0x7a, 0xef, 0x5f,
	0x37, 0x39, 0xee, 0x04, 0x35, 0x77, 0xc6, 0x1a, 0x7a, 0x34, 0xbc, 0x88, 0x4a, 0x95, 0xc6, 0x53,
	0x83, 0x91, 0x53, 0x16, 0xbc, 0x41, 0x70, 0x2f, 0x3c, 0x9a, 0x20, 0x5a, 0xdf, 0x17, 0x62, 0x63,
	0x47, 0xa5, 0x32, 0x7e, 0x2c, 0x56, 0x70, 0xba, 0x45, 0xe3, 0x9e, 0x7c, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0xb4, 0x71, 0xcf, 0xe4, 0x01, 0x00, 0x00,
}
