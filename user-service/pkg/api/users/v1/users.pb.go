// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package v1

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

type LoginRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignupRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignupRequest) Reset()         { *m = SignupRequest{} }
func (m *SignupRequest) String() string { return proto.CompactTextString(m) }
func (*SignupRequest) ProtoMessage()    {}
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *SignupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignupRequest.Unmarshal(m, b)
}
func (m *SignupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignupRequest.Marshal(b, m, deterministic)
}
func (m *SignupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignupRequest.Merge(m, src)
}
func (m *SignupRequest) XXX_Size() int {
	return xxx_messageInfo_SignupRequest.Size(m)
}
func (m *SignupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignupRequest proto.InternalMessageInfo

func (m *SignupRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignupRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserNameUsedRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNameUsedRequest) Reset()         { *m = UserNameUsedRequest{} }
func (m *UserNameUsedRequest) String() string { return proto.CompactTextString(m) }
func (*UserNameUsedRequest) ProtoMessage()    {}
func (*UserNameUsedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *UserNameUsedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNameUsedRequest.Unmarshal(m, b)
}
func (m *UserNameUsedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNameUsedRequest.Marshal(b, m, deterministic)
}
func (m *UserNameUsedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNameUsedRequest.Merge(m, src)
}
func (m *UserNameUsedRequest) XXX_Size() int {
	return xxx_messageInfo_UserNameUsedRequest.Size(m)
}
func (m *UserNameUsedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNameUsedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserNameUsedRequest proto.InternalMessageInfo

func (m *UserNameUsedRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserNameUsedResponse struct {
	Used                 bool     `protobuf:"varint,1,opt,name=Used,proto3" json:"Used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserNameUsedResponse) Reset()         { *m = UserNameUsedResponse{} }
func (m *UserNameUsedResponse) String() string { return proto.CompactTextString(m) }
func (*UserNameUsedResponse) ProtoMessage()    {}
func (*UserNameUsedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{3}
}

func (m *UserNameUsedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserNameUsedResponse.Unmarshal(m, b)
}
func (m *UserNameUsedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserNameUsedResponse.Marshal(b, m, deterministic)
}
func (m *UserNameUsedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserNameUsedResponse.Merge(m, src)
}
func (m *UserNameUsedResponse) XXX_Size() int {
	return xxx_messageInfo_UserNameUsedResponse.Size(m)
}
func (m *UserNameUsedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserNameUsedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserNameUsedResponse proto.InternalMessageInfo

func (m *UserNameUsedResponse) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

type EmailUsedRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailUsedRequest) Reset()         { *m = EmailUsedRequest{} }
func (m *EmailUsedRequest) String() string { return proto.CompactTextString(m) }
func (*EmailUsedRequest) ProtoMessage()    {}
func (*EmailUsedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{4}
}

func (m *EmailUsedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailUsedRequest.Unmarshal(m, b)
}
func (m *EmailUsedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailUsedRequest.Marshal(b, m, deterministic)
}
func (m *EmailUsedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailUsedRequest.Merge(m, src)
}
func (m *EmailUsedRequest) XXX_Size() int {
	return xxx_messageInfo_EmailUsedRequest.Size(m)
}
func (m *EmailUsedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailUsedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailUsedRequest proto.InternalMessageInfo

func (m *EmailUsedRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type EmailUsedResponse struct {
	Used                 bool     `protobuf:"varint,1,opt,name=Used,proto3" json:"Used,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailUsedResponse) Reset()         { *m = EmailUsedResponse{} }
func (m *EmailUsedResponse) String() string { return proto.CompactTextString(m) }
func (*EmailUsedResponse) ProtoMessage()    {}
func (*EmailUsedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{5}
}

func (m *EmailUsedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailUsedResponse.Unmarshal(m, b)
}
func (m *EmailUsedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailUsedResponse.Marshal(b, m, deterministic)
}
func (m *EmailUsedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailUsedResponse.Merge(m, src)
}
func (m *EmailUsedResponse) XXX_Size() int {
	return xxx_messageInfo_EmailUsedResponse.Size(m)
}
func (m *EmailUsedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailUsedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmailUsedResponse proto.InternalMessageInfo

func (m *EmailUsedResponse) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

type AuthUserRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserRequest) Reset()         { *m = AuthUserRequest{} }
func (m *AuthUserRequest) String() string { return proto.CompactTextString(m) }
func (*AuthUserRequest) ProtoMessage()    {}
func (*AuthUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{6}
}

func (m *AuthUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUserRequest.Unmarshal(m, b)
}
func (m *AuthUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUserRequest.Marshal(b, m, deterministic)
}
func (m *AuthUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserRequest.Merge(m, src)
}
func (m *AuthUserRequest) XXX_Size() int {
	return xxx_messageInfo_AuthUserRequest.Size(m)
}
func (m *AuthUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserRequest proto.InternalMessageInfo

func (m *AuthUserRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthUserResponse struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthUserResponse) Reset()         { *m = AuthUserResponse{} }
func (m *AuthUserResponse) String() string { return proto.CompactTextString(m) }
func (*AuthUserResponse) ProtoMessage()    {}
func (*AuthUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{7}
}

func (m *AuthUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthUserResponse.Unmarshal(m, b)
}
func (m *AuthUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthUserResponse.Marshal(b, m, deterministic)
}
func (m *AuthUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthUserResponse.Merge(m, src)
}
func (m *AuthUserResponse) XXX_Size() int {
	return xxx_messageInfo_AuthUserResponse.Size(m)
}
func (m *AuthUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthUserResponse proto.InternalMessageInfo

func (m *AuthUserResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *AuthUserResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AuthResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{8}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "v1.LoginRequest")
	proto.RegisterType((*SignupRequest)(nil), "v1.SignupRequest")
	proto.RegisterType((*UserNameUsedRequest)(nil), "v1.UserNameUsedRequest")
	proto.RegisterType((*UserNameUsedResponse)(nil), "v1.UserNameUsedResponse")
	proto.RegisterType((*EmailUsedRequest)(nil), "v1.EmailUsedRequest")
	proto.RegisterType((*EmailUsedResponse)(nil), "v1.EmailUsedResponse")
	proto.RegisterType((*AuthUserRequest)(nil), "v1.AuthUserRequest")
	proto.RegisterType((*AuthUserResponse)(nil), "v1.AuthUserResponse")
	proto.RegisterType((*AuthResponse)(nil), "v1.AuthResponse")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x59, 0xa6, 0xa3, 0xbb, 0x9b, 0xba, 0x65, 0x15, 0x47, 0x9e, 0x24, 0x08, 0x1b, 0x8a,
	0x83, 0x29, 0x82, 0x8f, 0x0e, 0xe6, 0xc3, 0x40, 0x44, 0xf6, 0xe7, 0xd1, 0x87, 0xe9, 0xc2, 0x2c,
	0xba, 0xa5, 0x36, 0x6d, 0xfd, 0x34, 0x7e, 0x57, 0x49, 0x9a, 0xc6, 0xa4, 0x56, 0xf1, 0xad, 0x27,
	0xf7, 0xdc, 0x5f, 0x6e, 0xef, 0x69, 0xa1, 0x91, 0x08, 0x16, 0x89, 0x41, 0x18, 0xf1, 0x98, 0x63,
	0x94, 0x0e, 0xe9, 0x0d, 0x34, 0xef, 0xf8, 0x3a, 0xd8, 0x4e, 0xd9, 0x7b, 0xc2, 0x44, 0x8c, 0x7d,
	0xd8, 0x55, 0xba, 0x5b, 0x39, 0xae, 0xf4, 0xeb, 0xd3, 0x4c, 0x60, 0x02, 0xde, 0xc3, 0x52, 0x88,
	0x0f, 0x1e, 0xad, 0xba, 0x48, 0x15, 0x8c, 0xa6, 0x8f, 0xb0, 0x37, 0x0b, 0xd6, 0xdb, 0x24, 0xcc,
	0x11, 0x04, 0xbc, 0x85, 0x60, 0xd1, 0xfd, 0x72, 0xc3, 0x34, 0xc5, 0x68, 0x89, 0xbf, 0xdd, 0x2c,
	0x83, 0x37, 0x4d, 0xc9, 0x84, 0x83, 0xaf, 0x16, 0xf0, 0x43, 0xe8, 0xe4, 0xdd, 0x0b, 0xc1, 0x56,
	0xff, 0xb8, 0x84, 0x9e, 0x82, 0xef, 0xb6, 0x88, 0x90, 0x6f, 0x05, 0xc3, 0x18, 0x76, 0xa4, 0x56,
	0x7e, 0x6f, 0xaa, 0x9e, 0x69, 0x1f, 0x5a, 0x6a, 0x06, 0x9b, 0x6d, 0x86, 0xac, 0x58, 0x43, 0xd2,
	0x1e, 0xb4, 0x2d, 0xe7, 0x1f, 0xc8, 0x1e, 0x1c, 0x8c, 0x92, 0xf8, 0x45, 0x8e, 0x60, 0x11, 0xe7,
	0xfc, 0x95, 0x99, 0xad, 0x2a, 0x41, 0xe7, 0xd0, 0xfa, 0x36, 0x6a, 0xe0, 0x3e, 0xa0, 0xc9, 0x58,
	0xdb, 0xd0, 0x64, 0xec, 0xbc, 0x27, 0xfa, 0x6d, 0x99, 0x55, 0x7b, 0xce, 0x13, 0x68, 0x4a, 0xaa,
	0x21, 0x96, 0xde, 0x7d, 0xf1, 0x89, 0xa0, 0x21, 0x6d, 0x33, 0x16, 0xa5, 0xc1, 0x33, 0xc3, 0x67,
	0x3a, 0x77, 0xdc, 0x1a, 0xa4, 0xc3, 0x81, 0xfd, 0x49, 0x10, 0x75, 0xe2, 0x20, 0xcf, 0xa1, 0x96,
	0x45, 0x8e, 0xdb, 0xb2, 0xe6, 0xc4, 0x5f, 0x62, 0x1f, 0x41, 0xd3, 0xce, 0x03, 0x1f, 0x49, 0x47,
	0x49, 0xa8, 0xa4, 0xfb, 0xb3, 0xa0, 0x11, 0xd7, 0x50, 0x37, 0xcb, 0xc7, 0xbe, 0xb4, 0x15, 0x53,
	0x23, 0x87, 0x85, 0x53, 0xdd, 0x79, 0x05, 0x5e, 0xbe, 0x64, 0xdc, 0xc9, 0x47, 0xb3, 0xb2, 0x21,
	0xbe, 0x7b, 0x98, 0xb5, 0x3d, 0xd5, 0xd4, 0x2f, 0x72, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x50, 0x02, 0x61, 0x31, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	UserNameUsed(ctx context.Context, in *UserNameUsedRequest, opts ...grpc.CallOption) (*UserNameUsedResponse, error)
	EmailUsed(ctx context.Context, in *EmailUsedRequest, opts ...grpc.CallOption) (*EmailUsedResponse, error)
	AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/v1.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/v1.AuthService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserNameUsed(ctx context.Context, in *UserNameUsedRequest, opts ...grpc.CallOption) (*UserNameUsedResponse, error) {
	out := new(UserNameUsedResponse)
	err := c.cc.Invoke(ctx, "/v1.AuthService/UserNameUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EmailUsed(ctx context.Context, in *EmailUsedRequest, opts ...grpc.CallOption) (*EmailUsedResponse, error) {
	out := new(EmailUsedResponse)
	err := c.cc.Invoke(ctx, "/v1.AuthService/EmailUsed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AuthUser(ctx context.Context, in *AuthUserRequest, opts ...grpc.CallOption) (*AuthUserResponse, error) {
	out := new(AuthUserResponse)
	err := c.cc.Invoke(ctx, "/v1.AuthService/AuthUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*AuthResponse, error)
	Signup(context.Context, *SignupRequest) (*AuthResponse, error)
	UserNameUsed(context.Context, *UserNameUsedRequest) (*UserNameUsedResponse, error)
	EmailUsed(context.Context, *EmailUsedRequest) (*EmailUsedResponse, error)
	AuthUser(context.Context, *AuthUserRequest) (*AuthUserResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthServiceServer) Signup(ctx context.Context, req *SignupRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (*UnimplementedAuthServiceServer) UserNameUsed(ctx context.Context, req *UserNameUsedRequest) (*UserNameUsedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserNameUsed not implemented")
}
func (*UnimplementedAuthServiceServer) EmailUsed(ctx context.Context, req *EmailUsedRequest) (*EmailUsedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailUsed not implemented")
}
func (*UnimplementedAuthServiceServer) AuthUser(ctx context.Context, req *AuthUserRequest) (*AuthUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUser not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserNameUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNameUsedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserNameUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/UserNameUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserNameUsed(ctx, req.(*UserNameUsedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EmailUsed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailUsedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EmailUsed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/EmailUsed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EmailUsed(ctx, req.(*EmailUsedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AuthUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AuthUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.AuthService/AuthUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AuthUser(ctx, req.(*AuthUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
		{
			MethodName: "Signup",
			Handler:    _AuthService_Signup_Handler,
		},
		{
			MethodName: "UserNameUsed",
			Handler:    _AuthService_UserNameUsed_Handler,
		},
		{
			MethodName: "EmailUsed",
			Handler:    _AuthService_EmailUsed_Handler,
		},
		{
			MethodName: "AuthUser",
			Handler:    _AuthService_AuthUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
