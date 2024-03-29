// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: o5/deployer/v1/topic/deployer.proto

package deployer_tpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DeployerTopic_RequestDeployment_FullMethodName  = "/o5.deployer.v1.topic.DeployerTopic/RequestDeployment"
	DeployerTopic_DeploymentComplete_FullMethodName = "/o5.deployer.v1.topic.DeployerTopic/DeploymentComplete"
	DeployerTopic_DeploymentFailed_FullMethodName   = "/o5.deployer.v1.topic.DeployerTopic/DeploymentFailed"
	DeployerTopic_TriggerDeployment_FullMethodName  = "/o5.deployer.v1.topic.DeployerTopic/TriggerDeployment"
)

// DeployerTopicClient is the client API for DeployerTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeployerTopicClient interface {
	// Handled by the Stack side
	RequestDeployment(ctx context.Context, in *RequestDeploymentMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeploymentComplete(ctx context.Context, in *DeploymentCompleteMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeploymentFailed(ctx context.Context, in *DeploymentFailedMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Unlocks the Deployment
	TriggerDeployment(ctx context.Context, in *TriggerDeploymentMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deployerTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewDeployerTopicClient(cc grpc.ClientConnInterface) DeployerTopicClient {
	return &deployerTopicClient{cc}
}

func (c *deployerTopicClient) RequestDeployment(ctx context.Context, in *RequestDeploymentMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeployerTopic_RequestDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerTopicClient) DeploymentComplete(ctx context.Context, in *DeploymentCompleteMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeployerTopic_DeploymentComplete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerTopicClient) DeploymentFailed(ctx context.Context, in *DeploymentFailedMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeployerTopic_DeploymentFailed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deployerTopicClient) TriggerDeployment(ctx context.Context, in *TriggerDeploymentMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeployerTopic_TriggerDeployment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployerTopicServer is the server API for DeployerTopic service.
// All implementations must embed UnimplementedDeployerTopicServer
// for forward compatibility
type DeployerTopicServer interface {
	// Handled by the Stack side
	RequestDeployment(context.Context, *RequestDeploymentMessage) (*emptypb.Empty, error)
	DeploymentComplete(context.Context, *DeploymentCompleteMessage) (*emptypb.Empty, error)
	DeploymentFailed(context.Context, *DeploymentFailedMessage) (*emptypb.Empty, error)
	// Unlocks the Deployment
	TriggerDeployment(context.Context, *TriggerDeploymentMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeployerTopicServer()
}

// UnimplementedDeployerTopicServer must be embedded to have forward compatible implementations.
type UnimplementedDeployerTopicServer struct {
}

func (UnimplementedDeployerTopicServer) RequestDeployment(context.Context, *RequestDeploymentMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDeployment not implemented")
}
func (UnimplementedDeployerTopicServer) DeploymentComplete(context.Context, *DeploymentCompleteMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentComplete not implemented")
}
func (UnimplementedDeployerTopicServer) DeploymentFailed(context.Context, *DeploymentFailedMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentFailed not implemented")
}
func (UnimplementedDeployerTopicServer) TriggerDeployment(context.Context, *TriggerDeploymentMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerDeployment not implemented")
}
func (UnimplementedDeployerTopicServer) mustEmbedUnimplementedDeployerTopicServer() {}

// UnsafeDeployerTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeployerTopicServer will
// result in compilation errors.
type UnsafeDeployerTopicServer interface {
	mustEmbedUnimplementedDeployerTopicServer()
}

func RegisterDeployerTopicServer(s grpc.ServiceRegistrar, srv DeployerTopicServer) {
	s.RegisterService(&DeployerTopic_ServiceDesc, srv)
}

func _DeployerTopic_RequestDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeploymentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerTopicServer).RequestDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployerTopic_RequestDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerTopicServer).RequestDeployment(ctx, req.(*RequestDeploymentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployerTopic_DeploymentComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentCompleteMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerTopicServer).DeploymentComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployerTopic_DeploymentComplete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerTopicServer).DeploymentComplete(ctx, req.(*DeploymentCompleteMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployerTopic_DeploymentFailed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentFailedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerTopicServer).DeploymentFailed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployerTopic_DeploymentFailed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerTopicServer).DeploymentFailed(ctx, req.(*DeploymentFailedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeployerTopic_TriggerDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerDeploymentMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployerTopicServer).TriggerDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeployerTopic_TriggerDeployment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployerTopicServer).TriggerDeployment(ctx, req.(*TriggerDeploymentMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// DeployerTopic_ServiceDesc is the grpc.ServiceDesc for DeployerTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeployerTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.deployer.v1.topic.DeployerTopic",
	HandlerType: (*DeployerTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDeployment",
			Handler:    _DeployerTopic_RequestDeployment_Handler,
		},
		{
			MethodName: "DeploymentComplete",
			Handler:    _DeployerTopic_DeploymentComplete_Handler,
		},
		{
			MethodName: "DeploymentFailed",
			Handler:    _DeployerTopic_DeploymentFailed_Handler,
		},
		{
			MethodName: "TriggerDeployment",
			Handler:    _DeployerTopic_TriggerDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/deployer/v1/topic/deployer.proto",
}
