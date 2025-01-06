// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: subscription.proto

package pb_subscription

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NotifySubscription_JoinPublication_FullMethodName          = "/NotifySubscription/JoinPublication"
	NotifySubscription_LeavePublication_FullMethodName         = "/NotifySubscription/LeavePublication"
	NotifySubscription_ChangePublicationPlan_FullMethodName    = "/NotifySubscription/ChangePublicationPlan"
	NotifySubscription_SubscribePublisherPlan_FullMethodName   = "/NotifySubscription/SubscribePublisherPlan"
	NotifySubscription_UnsubscribePublisherPlan_FullMethodName = "/NotifySubscription/UnsubscribePublisherPlan"
	NotifySubscription_ChangePublisherPlan_FullMethodName      = "/NotifySubscription/ChangePublisherPlan"
)

// NotifySubscriptionClient is the client API for NotifySubscription service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifySubscriptionClient interface {
	JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error)
	LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error)
	SubscribePublisherPlan(ctx context.Context, in *SubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error)
	UnsubscribePublisherPlan(ctx context.Context, in *UnsubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePublisherPlan(ctx context.Context, in *ChangePublisherPlanRequest, opts ...grpc.CallOption) (*Response, error)
}

type notifySubscriptionClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifySubscriptionClient(cc grpc.ClientConnInterface) NotifySubscriptionClient {
	return &notifySubscriptionClient{cc}
}

func (c *notifySubscriptionClient) JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_JoinPublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifySubscriptionClient) LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_LeavePublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifySubscriptionClient) ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_ChangePublicationPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifySubscriptionClient) SubscribePublisherPlan(ctx context.Context, in *SubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_SubscribePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifySubscriptionClient) UnsubscribePublisherPlan(ctx context.Context, in *UnsubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_UnsubscribePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifySubscriptionClient) ChangePublisherPlan(ctx context.Context, in *ChangePublisherPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifySubscription_ChangePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifySubscriptionServer is the server API for NotifySubscription service.
// All implementations must embed UnimplementedNotifySubscriptionServer
// for forward compatibility.
type NotifySubscriptionServer interface {
	JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error)
	LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error)
	ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error)
	SubscribePublisherPlan(context.Context, *SubscribePublisherRequest) (*Response, error)
	UnsubscribePublisherPlan(context.Context, *UnsubscribePublisherRequest) (*Response, error)
	ChangePublisherPlan(context.Context, *ChangePublisherPlanRequest) (*Response, error)
	mustEmbedUnimplementedNotifySubscriptionServer()
}

// UnimplementedNotifySubscriptionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotifySubscriptionServer struct{}

func (UnimplementedNotifySubscriptionServer) JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (UnimplementedNotifySubscriptionServer) LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (UnimplementedNotifySubscriptionServer) ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (UnimplementedNotifySubscriptionServer) SubscribePublisherPlan(context.Context, *SubscribePublisherRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribePublisherPlan not implemented")
}
func (UnimplementedNotifySubscriptionServer) UnsubscribePublisherPlan(context.Context, *UnsubscribePublisherRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribePublisherPlan not implemented")
}
func (UnimplementedNotifySubscriptionServer) ChangePublisherPlan(context.Context, *ChangePublisherPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherPlan not implemented")
}
func (UnimplementedNotifySubscriptionServer) mustEmbedUnimplementedNotifySubscriptionServer() {}
func (UnimplementedNotifySubscriptionServer) testEmbeddedByValue()                            {}

// UnsafeNotifySubscriptionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifySubscriptionServer will
// result in compilation errors.
type UnsafeNotifySubscriptionServer interface {
	mustEmbedUnimplementedNotifySubscriptionServer()
}

func RegisterNotifySubscriptionServer(s grpc.ServiceRegistrar, srv NotifySubscriptionServer) {
	// If the following call pancis, it indicates UnimplementedNotifySubscriptionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotifySubscription_ServiceDesc, srv)
}

func _NotifySubscription_JoinPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).JoinPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_JoinPublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).JoinPublication(ctx, req.(*JoinPublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifySubscription_LeavePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeavePublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).LeavePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_LeavePublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).LeavePublication(ctx, req.(*LeavePublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifySubscription_ChangePublicationPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublicationPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).ChangePublicationPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_ChangePublicationPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).ChangePublicationPlan(ctx, req.(*ChangePublicationPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifySubscription_SubscribePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).SubscribePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_SubscribePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).SubscribePublisherPlan(ctx, req.(*SubscribePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifySubscription_UnsubscribePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).UnsubscribePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_UnsubscribePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).UnsubscribePublisherPlan(ctx, req.(*UnsubscribePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifySubscription_ChangePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublisherPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifySubscriptionServer).ChangePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifySubscription_ChangePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifySubscriptionServer).ChangePublisherPlan(ctx, req.(*ChangePublisherPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotifySubscription_ServiceDesc is the grpc.ServiceDesc for NotifySubscription service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotifySubscription_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotifySubscription",
	HandlerType: (*NotifySubscriptionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPublication",
			Handler:    _NotifySubscription_JoinPublication_Handler,
		},
		{
			MethodName: "LeavePublication",
			Handler:    _NotifySubscription_LeavePublication_Handler,
		},
		{
			MethodName: "ChangePublicationPlan",
			Handler:    _NotifySubscription_ChangePublicationPlan_Handler,
		},
		{
			MethodName: "SubscribePublisherPlan",
			Handler:    _NotifySubscription_SubscribePublisherPlan_Handler,
		},
		{
			MethodName: "UnsubscribePublisherPlan",
			Handler:    _NotifySubscription_UnsubscribePublisherPlan_Handler,
		},
		{
			MethodName: "ChangePublisherPlan",
			Handler:    _NotifySubscription_ChangePublisherPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
