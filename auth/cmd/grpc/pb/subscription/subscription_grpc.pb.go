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
	NotifyAuth_JoinPublication_FullMethodName              = "/NotifyAuth/JoinPublication"
	NotifyAuth_LeavePublication_FullMethodName             = "/NotifyAuth/LeavePublication"
	NotifyAuth_ChangeSubscriberSubscription_FullMethodName = "/NotifyAuth/ChangeSubscriberSubscription"
	NotifyAuth_ConfirmPublisherSubscription_FullMethodName = "/NotifyAuth/ConfirmPublisherSubscription"
	NotifyAuth_RevokePublisherSubscription_FullMethodName  = "/NotifyAuth/RevokePublisherSubscription"
	NotifyAuth_ChangePublisherSubscription_FullMethodName  = "/NotifyAuth/ChangePublisherSubscription"
)

// NotifyAuthClient is the client API for NotifyAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyAuthClient interface {
	JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error)
	LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error)
	ChangeSubscriberSubscription(ctx context.Context, in *ChangeSubscriberSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	ConfirmPublisherSubscription(ctx context.Context, in *ConfirmPublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	RevokePublisherSubscription(ctx context.Context, in *RevokePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePublisherSubscription(ctx context.Context, in *ChangePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
}

type notifyAuthClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyAuthClient(cc grpc.ClientConnInterface) NotifyAuthClient {
	return &notifyAuthClient{cc}
}

func (c *notifyAuthClient) JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_JoinPublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_LeavePublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) ChangeSubscriberSubscription(ctx context.Context, in *ChangeSubscriberSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_ChangeSubscriberSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) ConfirmPublisherSubscription(ctx context.Context, in *ConfirmPublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_ConfirmPublisherSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) RevokePublisherSubscription(ctx context.Context, in *RevokePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RevokePublisherSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) ChangePublisherSubscription(ctx context.Context, in *ChangePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_ChangePublisherSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyAuthServer is the server API for NotifyAuth service.
// All implementations must embed UnimplementedNotifyAuthServer
// for forward compatibility.
type NotifyAuthServer interface {
	JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error)
	LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error)
	ChangeSubscriberSubscription(context.Context, *ChangeSubscriberSubscriptionRequest) (*Response, error)
	ConfirmPublisherSubscription(context.Context, *ConfirmPublisherSubscriptionRequest) (*Response, error)
	RevokePublisherSubscription(context.Context, *RevokePublisherSubscriptionRequest) (*Response, error)
	ChangePublisherSubscription(context.Context, *ChangePublisherSubscriptionRequest) (*Response, error)
	mustEmbedUnimplementedNotifyAuthServer()
}

// UnimplementedNotifyAuthServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotifyAuthServer struct{}

func (UnimplementedNotifyAuthServer) JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (UnimplementedNotifyAuthServer) LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (UnimplementedNotifyAuthServer) ChangeSubscriberSubscription(context.Context, *ChangeSubscriberSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeSubscriberSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) ConfirmPublisherSubscription(context.Context, *ConfirmPublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) RevokePublisherSubscription(context.Context, *RevokePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) ChangePublisherSubscription(context.Context, *ChangePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) mustEmbedUnimplementedNotifyAuthServer() {}
func (UnimplementedNotifyAuthServer) testEmbeddedByValue()                    {}

// UnsafeNotifyAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyAuthServer will
// result in compilation errors.
type UnsafeNotifyAuthServer interface {
	mustEmbedUnimplementedNotifyAuthServer()
}

func RegisterNotifyAuthServer(s grpc.ServiceRegistrar, srv NotifyAuthServer) {
	// If the following call pancis, it indicates UnimplementedNotifyAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NotifyAuth_ServiceDesc, srv)
}

func _NotifyAuth_JoinPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).JoinPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_JoinPublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).JoinPublication(ctx, req.(*JoinPublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_LeavePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeavePublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).LeavePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_LeavePublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).LeavePublication(ctx, req.(*LeavePublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_ChangeSubscriberSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeSubscriberSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).ChangeSubscriberSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_ChangeSubscriberSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).ChangeSubscriberSubscription(ctx, req.(*ChangeSubscriberSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_ConfirmPublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).ConfirmPublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_ConfirmPublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).ConfirmPublisherSubscription(ctx, req.(*ConfirmPublisherSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_RevokePublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RevokePublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RevokePublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RevokePublisherSubscription(ctx, req.(*RevokePublisherSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_ChangePublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).ChangePublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_ChangePublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).ChangePublisherSubscription(ctx, req.(*ChangePublisherSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NotifyAuth_ServiceDesc is the grpc.ServiceDesc for NotifyAuth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotifyAuth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NotifyAuth",
	HandlerType: (*NotifyAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPublication",
			Handler:    _NotifyAuth_JoinPublication_Handler,
		},
		{
			MethodName: "LeavePublication",
			Handler:    _NotifyAuth_LeavePublication_Handler,
		},
		{
			MethodName: "ChangeSubscriberSubscription",
			Handler:    _NotifyAuth_ChangeSubscriberSubscription_Handler,
		},
		{
			MethodName: "ConfirmPublisherSubscription",
			Handler:    _NotifyAuth_ConfirmPublisherSubscription_Handler,
		},
		{
			MethodName: "RevokePublisherSubscription",
			Handler:    _NotifyAuth_RevokePublisherSubscription_Handler,
		},
		{
			MethodName: "ChangePublisherSubscription",
			Handler:    _NotifyAuth_ChangePublisherSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
