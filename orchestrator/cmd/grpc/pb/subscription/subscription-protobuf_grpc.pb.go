// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: subscription-protobuf.proto

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
	SubscriptionService_JoinPublication_FullMethodName          = "/SubscriptionService/JoinPublication"
	SubscriptionService_LeavePublication_FullMethodName         = "/SubscriptionService/LeavePublication"
	SubscriptionService_ChangePublicationPlan_FullMethodName    = "/SubscriptionService/ChangePublicationPlan"
	SubscriptionService_SubscribePublisherPlan_FullMethodName   = "/SubscriptionService/SubscribePublisherPlan"
	SubscriptionService_UnsubscribePublisherPlan_FullMethodName = "/SubscriptionService/UnsubscribePublisherPlan"
	SubscriptionService_ChangePublisherPlan_FullMethodName      = "/SubscriptionService/ChangePublisherPlan"
)

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error)
	LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error)
	SubscribePublisherPlan(ctx context.Context, in *SubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error)
	UnsubscribePublisherPlan(ctx context.Context, in *UnsubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error)
	ChangePublisherPlan(ctx context.Context, in *ChangePublisherPlanRequest, opts ...grpc.CallOption) (*Response, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) JoinPublication(ctx context.Context, in *JoinPublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_JoinPublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) LeavePublication(ctx context.Context, in *LeavePublicationRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_LeavePublication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ChangePublicationPlan(ctx context.Context, in *ChangePublicationPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_ChangePublicationPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) SubscribePublisherPlan(ctx context.Context, in *SubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_SubscribePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) UnsubscribePublisherPlan(ctx context.Context, in *UnsubscribePublisherRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_UnsubscribePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ChangePublisherPlan(ctx context.Context, in *ChangePublisherPlanRequest, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, SubscriptionService_ChangePublisherPlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility.
type SubscriptionServiceServer interface {
	JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error)
	LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error)
	ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error)
	SubscribePublisherPlan(context.Context, *SubscribePublisherRequest) (*Response, error)
	UnsubscribePublisherPlan(context.Context, *UnsubscribePublisherRequest) (*Response, error)
	ChangePublisherPlan(context.Context, *ChangePublisherPlanRequest) (*Response, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubscriptionServiceServer struct{}

func (UnimplementedSubscriptionServiceServer) JoinPublication(context.Context, *JoinPublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinPublication not implemented")
}
func (UnimplementedSubscriptionServiceServer) LeavePublication(context.Context, *LeavePublicationRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeavePublication not implemented")
}
func (UnimplementedSubscriptionServiceServer) ChangePublicationPlan(context.Context, *ChangePublicationPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublicationPlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) SubscribePublisherPlan(context.Context, *SubscribePublisherRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribePublisherPlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) UnsubscribePublisherPlan(context.Context, *UnsubscribePublisherRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribePublisherPlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) ChangePublisherPlan(context.Context, *ChangePublisherPlanRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePublisherPlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}
func (UnimplementedSubscriptionServiceServer) testEmbeddedByValue()                             {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubscriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_JoinPublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinPublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).JoinPublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_JoinPublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).JoinPublication(ctx, req.(*JoinPublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_LeavePublication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeavePublicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).LeavePublication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_LeavePublication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).LeavePublication(ctx, req.(*LeavePublicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ChangePublicationPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublicationPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ChangePublicationPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ChangePublicationPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ChangePublicationPlan(ctx, req.(*ChangePublicationPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_SubscribePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).SubscribePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_SubscribePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).SubscribePublisherPlan(ctx, req.(*SubscribePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_UnsubscribePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribePublisherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).UnsubscribePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_UnsubscribePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).UnsubscribePublisherPlan(ctx, req.(*UnsubscribePublisherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ChangePublisherPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePublisherPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ChangePublisherPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_ChangePublisherPlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ChangePublisherPlan(ctx, req.(*ChangePublisherPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinPublication",
			Handler:    _SubscriptionService_JoinPublication_Handler,
		},
		{
			MethodName: "LeavePublication",
			Handler:    _SubscriptionService_LeavePublication_Handler,
		},
		{
			MethodName: "ChangePublicationPlan",
			Handler:    _SubscriptionService_ChangePublicationPlan_Handler,
		},
		{
			MethodName: "SubscribePublisherPlan",
			Handler:    _SubscriptionService_SubscribePublisherPlan_Handler,
		},
		{
			MethodName: "UnsubscribePublisherPlan",
			Handler:    _SubscriptionService_UnsubscribePublisherPlan_Handler,
		},
		{
			MethodName: "ChangePublisherPlan",
			Handler:    _SubscriptionService_ChangePublisherPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription-protobuf.proto",
}
