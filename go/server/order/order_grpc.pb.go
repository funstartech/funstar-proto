// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: order.proto

package order

import (
	context "context"
	wxpay "github.com/funstartech/funstar-proto/go/wxpay"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderSvrClient is the client API for OrderSvr service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderSvrClient interface {
	// 创建订单
	CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error)
	// 查询订单信息
	GetOrderInfo(ctx context.Context, in *GetOrderInfoReq, opts ...grpc.CallOption) (*GetOrderInfoRsp, error)
	// 查询用户订单
	GetUserOrders(ctx context.Context, in *GetUserOrdersReq, opts ...grpc.CallOption) (*GetUserOrdersRsp, error)
	// 取消订单
	CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderRsp, error)
	// 查询支付结果
	QueryPaymentResult(ctx context.Context, in *QueryPaymentResultReq, opts ...grpc.CallOption) (*QueryPaymentResultRsp, error)
	// 微信支付付款回调
	WxPayCallback(ctx context.Context, in *wxpay.WxPayCallbackReq, opts ...grpc.CallOption) (*wxpay.WxPayCallbackRsp, error)
}

type orderSvrClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderSvrClient(cc grpc.ClientConnInterface) OrderSvrClient {
	return &orderSvrClient{cc}
}

func (c *orderSvrClient) CreateOrder(ctx context.Context, in *CreateOrderReq, opts ...grpc.CallOption) (*CreateOrderRsp, error) {
	out := new(CreateOrderRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSvrClient) GetOrderInfo(ctx context.Context, in *GetOrderInfoReq, opts ...grpc.CallOption) (*GetOrderInfoRsp, error) {
	out := new(GetOrderInfoRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/GetOrderInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSvrClient) GetUserOrders(ctx context.Context, in *GetUserOrdersReq, opts ...grpc.CallOption) (*GetUserOrdersRsp, error) {
	out := new(GetUserOrdersRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/GetUserOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSvrClient) CancelOrder(ctx context.Context, in *CancelOrderReq, opts ...grpc.CallOption) (*CancelOrderRsp, error) {
	out := new(CancelOrderRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSvrClient) QueryPaymentResult(ctx context.Context, in *QueryPaymentResultReq, opts ...grpc.CallOption) (*QueryPaymentResultRsp, error) {
	out := new(QueryPaymentResultRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/QueryPaymentResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderSvrClient) WxPayCallback(ctx context.Context, in *wxpay.WxPayCallbackReq, opts ...grpc.CallOption) (*wxpay.WxPayCallbackRsp, error) {
	out := new(wxpay.WxPayCallbackRsp)
	err := c.cc.Invoke(ctx, "/funstar.server.order.OrderSvr/WxPayCallback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderSvrServer is the server API for OrderSvr service.
// All implementations must embed UnimplementedOrderSvrServer
// for forward compatibility
type OrderSvrServer interface {
	// 创建订单
	CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error)
	// 查询订单信息
	GetOrderInfo(context.Context, *GetOrderInfoReq) (*GetOrderInfoRsp, error)
	// 查询用户订单
	GetUserOrders(context.Context, *GetUserOrdersReq) (*GetUserOrdersRsp, error)
	// 取消订单
	CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderRsp, error)
	// 查询支付结果
	QueryPaymentResult(context.Context, *QueryPaymentResultReq) (*QueryPaymentResultRsp, error)
	// 微信支付付款回调
	WxPayCallback(context.Context, *wxpay.WxPayCallbackReq) (*wxpay.WxPayCallbackRsp, error)
	mustEmbedUnimplementedOrderSvrServer()
}

// UnimplementedOrderSvrServer must be embedded to have forward compatible implementations.
type UnimplementedOrderSvrServer struct {
}

func (UnimplementedOrderSvrServer) CreateOrder(context.Context, *CreateOrderReq) (*CreateOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderSvrServer) GetOrderInfo(context.Context, *GetOrderInfoReq) (*GetOrderInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderInfo not implemented")
}
func (UnimplementedOrderSvrServer) GetUserOrders(context.Context, *GetUserOrdersReq) (*GetUserOrdersRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserOrders not implemented")
}
func (UnimplementedOrderSvrServer) CancelOrder(context.Context, *CancelOrderReq) (*CancelOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOrder not implemented")
}
func (UnimplementedOrderSvrServer) QueryPaymentResult(context.Context, *QueryPaymentResultReq) (*QueryPaymentResultRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPaymentResult not implemented")
}
func (UnimplementedOrderSvrServer) WxPayCallback(context.Context, *wxpay.WxPayCallbackReq) (*wxpay.WxPayCallbackRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WxPayCallback not implemented")
}
func (UnimplementedOrderSvrServer) mustEmbedUnimplementedOrderSvrServer() {}

// UnsafeOrderSvrServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderSvrServer will
// result in compilation errors.
type UnsafeOrderSvrServer interface {
	mustEmbedUnimplementedOrderSvrServer()
}

func RegisterOrderSvrServer(s grpc.ServiceRegistrar, srv OrderSvrServer) {
	s.RegisterService(&OrderSvr_ServiceDesc, srv)
}

func _OrderSvr_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).CreateOrder(ctx, req.(*CreateOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSvr_GetOrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).GetOrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/GetOrderInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).GetOrderInfo(ctx, req.(*GetOrderInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSvr_GetUserOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserOrdersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).GetUserOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/GetUserOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).GetUserOrders(ctx, req.(*GetUserOrdersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSvr_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).CancelOrder(ctx, req.(*CancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSvr_QueryPaymentResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPaymentResultReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).QueryPaymentResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/QueryPaymentResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).QueryPaymentResult(ctx, req.(*QueryPaymentResultReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderSvr_WxPayCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wxpay.WxPayCallbackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderSvrServer).WxPayCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/funstar.server.order.OrderSvr/WxPayCallback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderSvrServer).WxPayCallback(ctx, req.(*wxpay.WxPayCallbackReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderSvr_ServiceDesc is the grpc.ServiceDesc for OrderSvr service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderSvr_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "funstar.server.order.OrderSvr",
	HandlerType: (*OrderSvrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderSvr_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderInfo",
			Handler:    _OrderSvr_GetOrderInfo_Handler,
		},
		{
			MethodName: "GetUserOrders",
			Handler:    _OrderSvr_GetUserOrders_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _OrderSvr_CancelOrder_Handler,
		},
		{
			MethodName: "QueryPaymentResult",
			Handler:    _OrderSvr_QueryPaymentResult_Handler,
		},
		{
			MethodName: "WxPayCallback",
			Handler:    _OrderSvr_WxPayCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
