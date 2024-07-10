// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.1
// source: AccountProtocol/account.proto

package account

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Register_RegisterRequest_FullMethodName = "/account.Register/RegisterRequest"
)

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegisterClient interface {
	RegisterRequest(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
}

type registerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegisterClient(cc grpc.ClientConnInterface) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) RegisterRequest(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, Register_RegisterRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
// All implementations must embed UnimplementedRegisterServer
// for forward compatibility
type RegisterServer interface {
	RegisterRequest(context.Context, *RegisterReq) (*RegisterResp, error)
	mustEmbedUnimplementedRegisterServer()
}

// UnimplementedRegisterServer must be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (UnimplementedRegisterServer) RegisterRequest(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterRequest not implemented")
}
func (UnimplementedRegisterServer) mustEmbedUnimplementedRegisterServer() {}

// UnsafeRegisterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegisterServer will
// result in compilation errors.
type UnsafeRegisterServer interface {
	mustEmbedUnimplementedRegisterServer()
}

func RegisterRegisterServer(s grpc.ServiceRegistrar, srv RegisterServer) {
	s.RegisterService(&Register_ServiceDesc, srv)
}

func _Register_RegisterRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).RegisterRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Register_RegisterRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).RegisterRequest(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Register_ServiceDesc is the grpc.ServiceDesc for Register service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Register_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterRequest",
			Handler:    _Register_RegisterRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AccountProtocol/account.proto",
}

const (
	Login_LoginRequest_FullMethodName = "/account.Login/LoginRequest"
)

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginClient interface {
	LoginRequest(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) LoginRequest(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Login_LoginRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
// All implementations must embed UnimplementedLoginServer
// for forward compatibility
type LoginServer interface {
	LoginRequest(context.Context, *LoginReq) (*LoginResp, error)
	mustEmbedUnimplementedLoginServer()
}

// UnimplementedLoginServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (UnimplementedLoginServer) LoginRequest(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginRequest not implemented")
}
func (UnimplementedLoginServer) mustEmbedUnimplementedLoginServer() {}

// UnsafeLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServer will
// result in compilation errors.
type UnsafeLoginServer interface {
	mustEmbedUnimplementedLoginServer()
}

func RegisterLoginServer(s grpc.ServiceRegistrar, srv LoginServer) {
	s.RegisterService(&Login_ServiceDesc, srv)
}

func _Login_LoginRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Login_LoginRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginRequest(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Login_ServiceDesc is the grpc.ServiceDesc for Login service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Login_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginRequest",
			Handler:    _Login_LoginRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "AccountProtocol/account.proto",
}
