// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: auth.proto

package pb_auth

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
	NotifyAuth_JoinPublication_FullMethodName                      = "/NotifyAuth/JoinPublication"
	NotifyAuth_RollbackJoinPublication_FullMethodName              = "/NotifyAuth/RollbackJoinPublication"
	NotifyAuth_LeavePublication_FullMethodName                     = "/NotifyAuth/LeavePublication"
	NotifyAuth_RollbackLeavePublication_FullMethodName             = "/NotifyAuth/RollbackLeavePublication"
	NotifyAuth_ChangePublicationPlan_FullMethodName                = "/NotifyAuth/ChangePublicationPlan"
	NotifyAuth_RollbackChangePublicationPlan_FullMethodName        = "/NotifyAuth/RollbackChangePublicationPlan"
	NotifyAuth_ConfirmPublisherSubscription_FullMethodName         = "/NotifyAuth/ConfirmPublisherSubscription"
	NotifyAuth_RollbackConfirmPublisherSubscription_FullMethodName = "/NotifyAuth/RollbackConfirmPublisherSubscription"
	NotifyAuth_RevokePublisherSubscription_FullMethodName          = "/NotifyAuth/RevokePublisherSubscription"
	NotifyAuth_RollbackRevokePublisherSubscription_FullMethodName  = "/NotifyAuth/RollbackRevokePublisherSubscription"
	NotifyAuth_ChangePublisherSubscription_FullMethodName          = "/NotifyAuth/ChangePublisherSubscription"
	NotifyAuth_RollbackChangePublisherSubscription_FullMethodName  = "/NotifyAuth/RollbackChangePublisherSubscription"
)

// NotifyAuthClient is the client API for NotifyAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// NotifyAuth service definition
type NotifyAuthClient interface {
	// Join a publication
	JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackJoinPublication(ctx context.Context, in *RollbackJoinPublicationRequest, opts ...grpc.CallOption) (*Response, error)
	// Leave a publication
	LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackLeavePublication(ctx context.Context, in *RollbackLeavePublicationRequest, opts ...grpc.CallOption) (*Response, error)
	// Change the plan of a publication
	ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackChangePublicationPlan(ctx context.Context, in *RollbackChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error)
	// Confirm publisher subscription
	ConfirmPublisherSubscription(ctx context.Context, in *ConfirmPublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackConfirmPublisherSubscription(ctx context.Context, in *RollbackConfirmPublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	// Revoke publisher subscription
	RevokePublisherSubscription(ctx context.Context, in *RevokePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackRevokePublisherSubscription(ctx context.Context, in *RollbackRevokePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	// Change publisher subscription
	ChangePublisherSubscription(ctx context.Context, in *ChangePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
	RollbackChangePublisherSubscription(ctx context.Context, in *RollbackChangePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error)
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

func (c *notifyAuthClient) RollbackJoinPublication(ctx context.Context, in *RollbackJoinPublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackJoinPublication_FullMethodName, in, out, cOpts...)
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

func (c *notifyAuthClient) RollbackLeavePublication(ctx context.Context, in *RollbackLeavePublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackLeavePublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_ChangePublicationPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifyAuthClient) RollbackChangePublicationPlan(ctx context.Context, in *RollbackChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackChangePublicationPlan_FullMethodName, in, out, cOpts...)
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

func (c *notifyAuthClient) RollbackConfirmPublisherSubscription(ctx context.Context, in *RollbackConfirmPublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackConfirmPublisherSubscription_FullMethodName, in, out, cOpts...)
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

func (c *notifyAuthClient) RollbackRevokePublisherSubscription(ctx context.Context, in *RollbackRevokePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackRevokePublisherSubscription_FullMethodName, in, out, cOpts...)
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

func (c *notifyAuthClient) RollbackChangePublisherSubscription(ctx context.Context, in *RollbackChangePublisherSubscriptionRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, NotifyAuth_RollbackChangePublisherSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifyAuthServer is the server API for NotifyAuth service.
// All implementations must embed UnimplementedNotifyAuthServer
// for forward compatibility.
//
// NotifyAuth service definition
type NotifyAuthServer interface {
	// Join a publication
	JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error)
	RollbackJoinPublication(context.Context, *RollbackJoinPublicationRequest) (*Response, error)
	// Leave a publication
	LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error)
	RollbackLeavePublication(context.Context, *RollbackLeavePublicationRequest) (*Response, error)
	// Change the plan of a publication
	ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error)
	RollbackChangePublicationPlan(context.Context, *RollbackChangePublicationPlanRequest) (*Response, error)
	// Confirm publisher subscription
	ConfirmPublisherSubscription(context.Context, *ConfirmPublisherSubscriptionRequest) (*Response, error)
	RollbackConfirmPublisherSubscription(context.Context, *RollbackConfirmPublisherSubscriptionRequest) (*Response, error)
	// Revoke publisher subscription
	RevokePublisherSubscription(context.Context, *RevokePublisherSubscriptionRequest) (*Response, error)
	RollbackRevokePublisherSubscription(context.Context, *RollbackRevokePublisherSubscriptionRequest) (*Response, error)
	// Change publisher subscription
	ChangePublisherSubscription(context.Context, *ChangePublisherSubscriptionRequest) (*Response, error)
	RollbackChangePublisherSubscription(context.Context, *RollbackChangePublisherSubscriptionRequest) (*Response, error)
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
func (UnimplementedNotifyAuthServer) RollbackJoinPublication(context.Context, *RollbackJoinPublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackJoinPublication not implemented")
}
func (UnimplementedNotifyAuthServer) LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (UnimplementedNotifyAuthServer) RollbackLeavePublication(context.Context, *RollbackLeavePublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackLeavePublication not implemented")
}
func (UnimplementedNotifyAuthServer) ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (UnimplementedNotifyAuthServer) RollbackChangePublicationPlan(context.Context, *RollbackChangePublicationPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublicationPlan not implemented")
}
func (UnimplementedNotifyAuthServer) ConfirmPublisherSubscription(context.Context, *ConfirmPublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) RollbackConfirmPublisherSubscription(context.Context, *RollbackConfirmPublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackConfirmPublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) RevokePublisherSubscription(context.Context, *RevokePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokePublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) RollbackRevokePublisherSubscription(context.Context, *RollbackRevokePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackRevokePublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) ChangePublisherSubscription(context.Context, *ChangePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherSubscription not implemented")
}
func (UnimplementedNotifyAuthServer) RollbackChangePublisherSubscription(context.Context, *RollbackChangePublisherSubscriptionRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackChangePublisherSubscription not implemented")
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

func _NotifyAuth_RollbackJoinPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackJoinPublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackJoinPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackJoinPublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackJoinPublication(ctx, req.(*RollbackJoinPublicationRequest))
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

func _NotifyAuth_RollbackLeavePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackLeavePublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackLeavePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackLeavePublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackLeavePublication(ctx, req.(*RollbackLeavePublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_ChangePublicationPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublicationPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).ChangePublicationPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_ChangePublicationPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).ChangePublicationPlan(ctx, req.(*ChangePublicationPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotifyAuth_RollbackChangePublicationPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackChangePublicationPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackChangePublicationPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackChangePublicationPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackChangePublicationPlan(ctx, req.(*RollbackChangePublicationPlanRequest))
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

func _NotifyAuth_RollbackConfirmPublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackConfirmPublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackConfirmPublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackConfirmPublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackConfirmPublisherSubscription(ctx, req.(*RollbackConfirmPublisherSubscriptionRequest))
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

func _NotifyAuth_RollbackRevokePublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackRevokePublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackRevokePublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackRevokePublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackRevokePublisherSubscription(ctx, req.(*RollbackRevokePublisherSubscriptionRequest))
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

func _NotifyAuth_RollbackChangePublisherSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackChangePublisherSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifyAuthServer).RollbackChangePublisherSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NotifyAuth_RollbackChangePublisherSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifyAuthServer).RollbackChangePublisherSubscription(ctx, req.(*RollbackChangePublisherSubscriptionRequest))
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
			MethodName: "RollbackJoinPublication",
			Handler:    _NotifyAuth_RollbackJoinPublication_Handler,
		},
		{
			MethodName: "LeavePublication",
			Handler:    _NotifyAuth_LeavePublication_Handler,
		},
		{
			MethodName: "RollbackLeavePublication",
			Handler:    _NotifyAuth_RollbackLeavePublication_Handler,
		},
		{
			MethodName: "ChangePublicationPlan",
			Handler:    _NotifyAuth_ChangePublicationPlan_Handler,
		},
		{
			MethodName: "RollbackChangePublicationPlan",
			Handler:    _NotifyAuth_RollbackChangePublicationPlan_Handler,
		},
		{
			MethodName: "ConfirmPublisherSubscription",
			Handler:    _NotifyAuth_ConfirmPublisherSubscription_Handler,
		},
		{
			MethodName: "RollbackConfirmPublisherSubscription",
			Handler:    _NotifyAuth_RollbackConfirmPublisherSubscription_Handler,
		},
		{
			MethodName: "RevokePublisherSubscription",
			Handler:    _NotifyAuth_RevokePublisherSubscription_Handler,
		},
		{
			MethodName: "RollbackRevokePublisherSubscription",
			Handler:    _NotifyAuth_RollbackRevokePublisherSubscription_Handler,
		},
		{
			MethodName: "ChangePublisherSubscription",
			Handler:    _NotifyAuth_ChangePublisherSubscription_Handler,
		},
		{
			MethodName: "RollbackChangePublisherSubscription",
			Handler:    _NotifyAuth_RollbackChangePublisherSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
