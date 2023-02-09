// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: proto/userinfo.proto

package userinfo

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

// UserInfoSvrClient is the client API for UserInfoSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInfoSvrClient interface {
	// 获取用户信息
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error)
	// 更新用户信息
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoRsp, error)
}

type userInfoSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoSvrClient(cc grpc.ClientConnInterface) UserInfoSvrClient {
	return &userInfoSvrClient{cc}
}

func (c *userInfoSvrClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoRsp, error) {
	out := new(GetUserInfoRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.userinfo.UserInfoSvr/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoSvrClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoReq, opts ...grpc.CallOption) (*UpdateUserInfoRsp, error) {
	out := new(UpdateUserInfoRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.userinfo.UserInfoSvr/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoSvrServer is the server API for UserInfoSvr service.
// All implementations must embed UnimplementedUserInfoSvrServer
// for forward compatibility
type UserInfoSvrServer interface {
	// 获取用户信息
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error)
	// 更新用户信息
	UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoRsp, error)
	mustEmbedUnimplementedUserInfoSvrServer()
}

// UnimplementedUserInfoSvrServer must be embedded to have forward compatible implementations.
type UnimplementedUserInfoSvrServer struct {
}

func (UnimplementedUserInfoSvrServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserInfoSvrServer) UpdateUserInfo(context.Context, *UpdateUserInfoReq) (*UpdateUserInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserInfoSvrServer) mustEmbedUnimplementedUserInfoSvrServer() {}

// UnsafeUserInfoSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInfoSvrServer will
// result in compilation errors.
type UnsafeUserInfoSvrServer interface {
	mustEmbedUnimplementedUserInfoSvrServer()
}

func RegisterUserInfoSvrServer(s grpc.ServiceRegistrar, srv UserInfoSvrServer) {
	s.RegisterService(&UserInfoSvr_ServiceDesc, srv)
}

func _UserInfoSvr_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoSvrServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.userinfo.UserInfoSvr/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoSvrServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoSvr_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoSvrServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.userinfo.UserInfoSvr/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoSvrServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInfoSvr_ServiceDesc is the grpc.ServiceDesc for UserInfoSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInfoSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "funstar.server.userinfo.UserInfoSvr",
	HandlerType: (*UserInfoSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInfo",
			Handler:    _UserInfoSvr_GetUserInfo_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserInfoSvr_UpdateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/userinfo.proto",
}
