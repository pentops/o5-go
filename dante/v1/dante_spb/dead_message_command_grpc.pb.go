// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: o5/dante/v1/service/dead_message_command.proto

package dante_spb

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
	DeadMessageCommandService_UpdateDeadMessage_FullMethodName = "/o5.dante.v1.service.DeadMessageCommandService/UpdateDeadMessage"
	DeadMessageCommandService_ReplayDeadMessage_FullMethodName = "/o5.dante.v1.service.DeadMessageCommandService/ReplayDeadMessage"
	DeadMessageCommandService_RejectDeadMessage_FullMethodName = "/o5.dante.v1.service.DeadMessageCommandService/RejectDeadMessage"
)

// DeadMessageCommandServiceClient is the client API for DeadMessageCommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeadMessageCommandServiceClient interface {
	UpdateDeadMessage(ctx context.Context, in *UpdateDeadMessageRequest, opts ...grpc.CallOption) (*UpdateDeadMessageResponse, error)
	ReplayDeadMessage(ctx context.Context, in *ReplayDeadMessageRequest, opts ...grpc.CallOption) (*ReplayDeadMessageResponse, error)
	RejectDeadMessage(ctx context.Context, in *RejectDeadMessageRequest, opts ...grpc.CallOption) (*RejectDeadMessageResponse, error)
}

type deadMessageCommandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeadMessageCommandServiceClient(cc grpc.ClientConnInterface) DeadMessageCommandServiceClient {
	return &deadMessageCommandServiceClient{cc}
}

func (c *deadMessageCommandServiceClient) UpdateDeadMessage(ctx context.Context, in *UpdateDeadMessageRequest, opts ...grpc.CallOption) (*UpdateDeadMessageResponse, error) {
	out := new(UpdateDeadMessageResponse)
	err := c.cc.Invoke(ctx, DeadMessageCommandService_UpdateDeadMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadMessageCommandServiceClient) ReplayDeadMessage(ctx context.Context, in *ReplayDeadMessageRequest, opts ...grpc.CallOption) (*ReplayDeadMessageResponse, error) {
	out := new(ReplayDeadMessageResponse)
	err := c.cc.Invoke(ctx, DeadMessageCommandService_ReplayDeadMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadMessageCommandServiceClient) RejectDeadMessage(ctx context.Context, in *RejectDeadMessageRequest, opts ...grpc.CallOption) (*RejectDeadMessageResponse, error) {
	out := new(RejectDeadMessageResponse)
	err := c.cc.Invoke(ctx, DeadMessageCommandService_RejectDeadMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeadMessageCommandServiceServer is the server API for DeadMessageCommandService service.
// All implementations must embed UnimplementedDeadMessageCommandServiceServer
// for forward compatibility
type DeadMessageCommandServiceServer interface {
	UpdateDeadMessage(context.Context, *UpdateDeadMessageRequest) (*UpdateDeadMessageResponse, error)
	ReplayDeadMessage(context.Context, *ReplayDeadMessageRequest) (*ReplayDeadMessageResponse, error)
	RejectDeadMessage(context.Context, *RejectDeadMessageRequest) (*RejectDeadMessageResponse, error)
	mustEmbedUnimplementedDeadMessageCommandServiceServer()
}

// UnimplementedDeadMessageCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeadMessageCommandServiceServer struct {
}

func (UnimplementedDeadMessageCommandServiceServer) UpdateDeadMessage(context.Context, *UpdateDeadMessageRequest) (*UpdateDeadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeadMessage not implemented")
}
func (UnimplementedDeadMessageCommandServiceServer) ReplayDeadMessage(context.Context, *ReplayDeadMessageRequest) (*ReplayDeadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplayDeadMessage not implemented")
}
func (UnimplementedDeadMessageCommandServiceServer) RejectDeadMessage(context.Context, *RejectDeadMessageRequest) (*RejectDeadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectDeadMessage not implemented")
}
func (UnimplementedDeadMessageCommandServiceServer) mustEmbedUnimplementedDeadMessageCommandServiceServer() {
}

// UnsafeDeadMessageCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeadMessageCommandServiceServer will
// result in compilation errors.
type UnsafeDeadMessageCommandServiceServer interface {
	mustEmbedUnimplementedDeadMessageCommandServiceServer()
}

func RegisterDeadMessageCommandServiceServer(s grpc.ServiceRegistrar, srv DeadMessageCommandServiceServer) {
	s.RegisterService(&DeadMessageCommandService_ServiceDesc, srv)
}

func _DeadMessageCommandService_UpdateDeadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageCommandServiceServer).UpdateDeadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeadMessageCommandService_UpdateDeadMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageCommandServiceServer).UpdateDeadMessage(ctx, req.(*UpdateDeadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeadMessageCommandService_ReplayDeadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplayDeadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageCommandServiceServer).ReplayDeadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeadMessageCommandService_ReplayDeadMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageCommandServiceServer).ReplayDeadMessage(ctx, req.(*ReplayDeadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeadMessageCommandService_RejectDeadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectDeadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageCommandServiceServer).RejectDeadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeadMessageCommandService_RejectDeadMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageCommandServiceServer).RejectDeadMessage(ctx, req.(*RejectDeadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeadMessageCommandService_ServiceDesc is the grpc.ServiceDesc for DeadMessageCommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeadMessageCommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.dante.v1.service.DeadMessageCommandService",
	HandlerType: (*DeadMessageCommandServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDeadMessage",
			Handler:    _DeadMessageCommandService_UpdateDeadMessage_Handler,
		},
		{
			MethodName: "ReplayDeadMessage",
			Handler:    _DeadMessageCommandService_ReplayDeadMessage_Handler,
		},
		{
			MethodName: "RejectDeadMessage",
			Handler:    _DeadMessageCommandService_RejectDeadMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/dante/v1/service/dead_message_command.proto",
}