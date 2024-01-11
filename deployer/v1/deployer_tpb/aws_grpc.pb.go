// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: o5/deployer/v1/topic/aws.proto

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

// AWSCommandTopicClient is the client API for AWSCommandTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AWSCommandTopicClient interface {
	CreateNewStack(ctx context.Context, in *CreateNewStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateStack(ctx context.Context, in *UpdateStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteStack(ctx context.Context, in *DeleteStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ScaleStack(ctx context.Context, in *ScaleStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CancelStackUpdate(ctx context.Context, in *CancelStackUpdateMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StabalizeStack(ctx context.Context, in *StabalizeStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RunDatabaseMigration(ctx context.Context, in *RunDatabaseMigrationMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aWSCommandTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewAWSCommandTopicClient(cc grpc.ClientConnInterface) AWSCommandTopicClient {
	return &aWSCommandTopicClient{cc}
}

func (c *aWSCommandTopicClient) CreateNewStack(ctx context.Context, in *CreateNewStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/CreateNewStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) UpdateStack(ctx context.Context, in *UpdateStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/UpdateStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) DeleteStack(ctx context.Context, in *DeleteStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/DeleteStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) ScaleStack(ctx context.Context, in *ScaleStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/ScaleStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) CancelStackUpdate(ctx context.Context, in *CancelStackUpdateMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/CancelStackUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) StabalizeStack(ctx context.Context, in *StabalizeStackMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/StabalizeStack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCommandTopicClient) RunDatabaseMigration(ctx context.Context, in *RunDatabaseMigrationMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.deployer.v1.topic.AWSCommandTopic/RunDatabaseMigration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AWSCommandTopicServer is the server API for AWSCommandTopic service.
// All implementations must embed UnimplementedAWSCommandTopicServer
// for forward compatibility
type AWSCommandTopicServer interface {
	CreateNewStack(context.Context, *CreateNewStackMessage) (*emptypb.Empty, error)
	UpdateStack(context.Context, *UpdateStackMessage) (*emptypb.Empty, error)
	DeleteStack(context.Context, *DeleteStackMessage) (*emptypb.Empty, error)
	ScaleStack(context.Context, *ScaleStackMessage) (*emptypb.Empty, error)
	CancelStackUpdate(context.Context, *CancelStackUpdateMessage) (*emptypb.Empty, error)
	StabalizeStack(context.Context, *StabalizeStackMessage) (*emptypb.Empty, error)
	RunDatabaseMigration(context.Context, *RunDatabaseMigrationMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedAWSCommandTopicServer()
}

// UnimplementedAWSCommandTopicServer must be embedded to have forward compatible implementations.
type UnimplementedAWSCommandTopicServer struct {
}

func (UnimplementedAWSCommandTopicServer) CreateNewStack(context.Context, *CreateNewStackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewStack not implemented")
}
func (UnimplementedAWSCommandTopicServer) UpdateStack(context.Context, *UpdateStackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStack not implemented")
}
func (UnimplementedAWSCommandTopicServer) DeleteStack(context.Context, *DeleteStackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStack not implemented")
}
func (UnimplementedAWSCommandTopicServer) ScaleStack(context.Context, *ScaleStackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScaleStack not implemented")
}
func (UnimplementedAWSCommandTopicServer) CancelStackUpdate(context.Context, *CancelStackUpdateMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelStackUpdate not implemented")
}
func (UnimplementedAWSCommandTopicServer) StabalizeStack(context.Context, *StabalizeStackMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StabalizeStack not implemented")
}
func (UnimplementedAWSCommandTopicServer) RunDatabaseMigration(context.Context, *RunDatabaseMigrationMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunDatabaseMigration not implemented")
}
func (UnimplementedAWSCommandTopicServer) mustEmbedUnimplementedAWSCommandTopicServer() {}

// UnsafeAWSCommandTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AWSCommandTopicServer will
// result in compilation errors.
type UnsafeAWSCommandTopicServer interface {
	mustEmbedUnimplementedAWSCommandTopicServer()
}

func RegisterAWSCommandTopicServer(s grpc.ServiceRegistrar, srv AWSCommandTopicServer) {
	s.RegisterService(&AWSCommandTopic_ServiceDesc, srv)
}

func _AWSCommandTopic_CreateNewStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewStackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).CreateNewStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/CreateNewStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).CreateNewStack(ctx, req.(*CreateNewStackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_UpdateStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).UpdateStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/UpdateStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).UpdateStack(ctx, req.(*UpdateStackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_DeleteStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).DeleteStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/DeleteStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).DeleteStack(ctx, req.(*DeleteStackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_ScaleStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScaleStackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).ScaleStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/ScaleStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).ScaleStack(ctx, req.(*ScaleStackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_CancelStackUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelStackUpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).CancelStackUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/CancelStackUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).CancelStackUpdate(ctx, req.(*CancelStackUpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_StabalizeStack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StabalizeStackMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).StabalizeStack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/StabalizeStack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).StabalizeStack(ctx, req.(*StabalizeStackMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCommandTopic_RunDatabaseMigration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunDatabaseMigrationMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCommandTopicServer).RunDatabaseMigration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.deployer.v1.topic.AWSCommandTopic/RunDatabaseMigration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCommandTopicServer).RunDatabaseMigration(ctx, req.(*RunDatabaseMigrationMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// AWSCommandTopic_ServiceDesc is the grpc.ServiceDesc for AWSCommandTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AWSCommandTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.deployer.v1.topic.AWSCommandTopic",
	HandlerType: (*AWSCommandTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewStack",
			Handler:    _AWSCommandTopic_CreateNewStack_Handler,
		},
		{
			MethodName: "UpdateStack",
			Handler:    _AWSCommandTopic_UpdateStack_Handler,
		},
		{
			MethodName: "DeleteStack",
			Handler:    _AWSCommandTopic_DeleteStack_Handler,
		},
		{
			MethodName: "ScaleStack",
			Handler:    _AWSCommandTopic_ScaleStack_Handler,
		},
		{
			MethodName: "CancelStackUpdate",
			Handler:    _AWSCommandTopic_CancelStackUpdate_Handler,
		},
		{
			MethodName: "StabalizeStack",
			Handler:    _AWSCommandTopic_StabalizeStack_Handler,
		},
		{
			MethodName: "RunDatabaseMigration",
			Handler:    _AWSCommandTopic_RunDatabaseMigration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/deployer/v1/topic/aws.proto",
}
