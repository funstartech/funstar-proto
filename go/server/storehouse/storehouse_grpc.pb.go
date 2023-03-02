// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: storehouse.proto

package storehouse

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

// StorehouseSvrClient is the client API for StorehouseSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorehouseSvrClient interface {
	// 拉取用户仓库
	GetUserStores(ctx context.Context, in *GetUserStoresReq, opts ...grpc.CallOption) (*GetUserStoresRsp, error)
	// 创建提货单
	CreatePickUpOrder(ctx context.Context, in *CreatePickUpOrderReq, opts ...grpc.CallOption) (*CreatePickUpOrderRsp, error)
	// 查询用户提货单
	GetUserPickUpOrders(ctx context.Context, in *GetUserPickUpOrdersReq, opts ...grpc.CallOption) (*GetUserPickUpOrdersRsp, error)
}

type storehouseSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewStorehouseSvrClient(cc grpc.ClientConnInterface) StorehouseSvrClient {
	return &storehouseSvrClient{cc}
}

func (c *storehouseSvrClient) GetUserStores(ctx context.Context, in *GetUserStoresReq, opts ...grpc.CallOption) (*GetUserStoresRsp, error) {
	out := new(GetUserStoresRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.storehouse.StorehouseSvr/GetUserStores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storehouseSvrClient) CreatePickUpOrder(ctx context.Context, in *CreatePickUpOrderReq, opts ...grpc.CallOption) (*CreatePickUpOrderRsp, error) {
	out := new(CreatePickUpOrderRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.storehouse.StorehouseSvr/CreatePickUpOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storehouseSvrClient) GetUserPickUpOrders(ctx context.Context, in *GetUserPickUpOrdersReq, opts ...grpc.CallOption) (*GetUserPickUpOrdersRsp, error) {
	out := new(GetUserPickUpOrdersRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.storehouse.StorehouseSvr/GetUserPickUpOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorehouseSvrServer is the server API for StorehouseSvr service.
// All implementations must embed UnimplementedStorehouseSvrServer
// for forward compatibility
type StorehouseSvrServer interface {
	// 拉取用户仓库
	GetUserStores(context.Context, *GetUserStoresReq) (*GetUserStoresRsp, error)
	// 创建提货单
	CreatePickUpOrder(context.Context, *CreatePickUpOrderReq) (*CreatePickUpOrderRsp, error)
	// 查询用户提货单
	GetUserPickUpOrders(context.Context, *GetUserPickUpOrdersReq) (*GetUserPickUpOrdersRsp, error)
	mustEmbedUnimplementedStorehouseSvrServer()
}

// UnimplementedStorehouseSvrServer must be embedded to have forward compatible implementations.
type UnimplementedStorehouseSvrServer struct {
}

func (UnimplementedStorehouseSvrServer) GetUserStores(context.Context, *GetUserStoresReq) (*GetUserStoresRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserStores not implemented")
}
func (UnimplementedStorehouseSvrServer) CreatePickUpOrder(context.Context, *CreatePickUpOrderReq) (*CreatePickUpOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePickUpOrder not implemented")
}
func (UnimplementedStorehouseSvrServer) GetUserPickUpOrders(context.Context, *GetUserPickUpOrdersReq) (*GetUserPickUpOrdersRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPickUpOrders not implemented")
}
func (UnimplementedStorehouseSvrServer) mustEmbedUnimplementedStorehouseSvrServer() {}

// UnsafeStorehouseSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorehouseSvrServer will
// result in compilation errors.
type UnsafeStorehouseSvrServer interface {
	mustEmbedUnimplementedStorehouseSvrServer()
}

func RegisterStorehouseSvrServer(s grpc.ServiceRegistrar, srv StorehouseSvrServer) {
	s.RegisterService(&StorehouseSvr_ServiceDesc, srv)
}

func _StorehouseSvr_GetUserStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserStoresReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorehouseSvrServer).GetUserStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.storehouse.StorehouseSvr/GetUserStores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorehouseSvrServer).GetUserStores(ctx, req.(*GetUserStoresReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorehouseSvr_CreatePickUpOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePickUpOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorehouseSvrServer).CreatePickUpOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.storehouse.StorehouseSvr/CreatePickUpOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorehouseSvrServer).CreatePickUpOrder(ctx, req.(*CreatePickUpOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorehouseSvr_GetUserPickUpOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPickUpOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorehouseSvrServer).GetUserPickUpOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.storehouse.StorehouseSvr/GetUserPickUpOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorehouseSvrServer).GetUserPickUpOrders(ctx, req.(*GetUserPickUpOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StorehouseSvr_ServiceDesc is the grpc.ServiceDesc for StorehouseSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorehouseSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "funstar.server.storehouse.StorehouseSvr",
	HandlerType: (*StorehouseSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserStores",
			Handler:    _StorehouseSvr_GetUserStores_Handler,
		},
		{
			MethodName: "CreatePickUpOrder",
			Handler:    _StorehouseSvr_CreatePickUpOrder_Handler,
		},
		{
			MethodName: "GetUserPickUpOrders",
			Handler:    _StorehouseSvr_GetUserPickUpOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storehouse.proto",
}
