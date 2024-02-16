// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: o5/dante/v1/service/dead_message_query.proto

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

// DeadMessageQueryServiceClient is the client API for DeadMessageQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeadMessageQueryServiceClient interface {
	GetDeadMessage(ctx context.Context, in *GetDeadMessageRequest, opts ...grpc.CallOption) (*GetDeadMessageResponse, error)
	ListDeadMessages(ctx context.Context, in *ListDeadMessagesRequest, opts ...grpc.CallOption) (*ListDeadMessagesResponse, error)
	ListDeadMessageEvents(ctx context.Context, in *ListDeadMessageEventsRequest, opts ...grpc.CallOption) (*ListDeadMessageEventsResponse, error)
}

type deadMessageQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeadMessageQueryServiceClient(cc grpc.ClientConnInterface) DeadMessageQueryServiceClient {
	return &deadMessageQueryServiceClient{cc}
}

func (c *deadMessageQueryServiceClient) GetDeadMessage(ctx context.Context, in *GetDeadMessageRequest, opts ...grpc.CallOption) (*GetDeadMessageResponse, error) {
	out := new(GetDeadMessageResponse)
	err := c.cc.Invoke(ctx, "/o5.dante.v1.service.DeadMessageQueryService/GetDeadMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadMessageQueryServiceClient) ListDeadMessages(ctx context.Context, in *ListDeadMessagesRequest, opts ...grpc.CallOption) (*ListDeadMessagesResponse, error) {
	out := new(ListDeadMessagesResponse)
	err := c.cc.Invoke(ctx, "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deadMessageQueryServiceClient) ListDeadMessageEvents(ctx context.Context, in *ListDeadMessageEventsRequest, opts ...grpc.CallOption) (*ListDeadMessageEventsResponse, error) {
	out := new(ListDeadMessageEventsResponse)
	err := c.cc.Invoke(ctx, "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessageEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeadMessageQueryServiceServer is the server API for DeadMessageQueryService service.
// All implementations must embed UnimplementedDeadMessageQueryServiceServer
// for forward compatibility
type DeadMessageQueryServiceServer interface {
	GetDeadMessage(context.Context, *GetDeadMessageRequest) (*GetDeadMessageResponse, error)
	ListDeadMessages(context.Context, *ListDeadMessagesRequest) (*ListDeadMessagesResponse, error)
	ListDeadMessageEvents(context.Context, *ListDeadMessageEventsRequest) (*ListDeadMessageEventsResponse, error)
	mustEmbedUnimplementedDeadMessageQueryServiceServer()
}

// UnimplementedDeadMessageQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeadMessageQueryServiceServer struct {
}

func (UnimplementedDeadMessageQueryServiceServer) GetDeadMessage(context.Context, *GetDeadMessageRequest) (*GetDeadMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeadMessage not implemented")
}
func (UnimplementedDeadMessageQueryServiceServer) ListDeadMessages(context.Context, *ListDeadMessagesRequest) (*ListDeadMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeadMessages not implemented")
}
func (UnimplementedDeadMessageQueryServiceServer) ListDeadMessageEvents(context.Context, *ListDeadMessageEventsRequest) (*ListDeadMessageEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDeadMessageEvents not implemented")
}
func (UnimplementedDeadMessageQueryServiceServer) mustEmbedUnimplementedDeadMessageQueryServiceServer() {
}

// UnsafeDeadMessageQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeadMessageQueryServiceServer will
// result in compilation errors.
type UnsafeDeadMessageQueryServiceServer interface {
	mustEmbedUnimplementedDeadMessageQueryServiceServer()
}

func RegisterDeadMessageQueryServiceServer(s grpc.ServiceRegistrar, srv DeadMessageQueryServiceServer) {
	s.RegisterService(&DeadMessageQueryService_ServiceDesc, srv)
}

func _DeadMessageQueryService_GetDeadMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeadMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageQueryServiceServer).GetDeadMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.dante.v1.service.DeadMessageQueryService/GetDeadMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageQueryServiceServer).GetDeadMessage(ctx, req.(*GetDeadMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeadMessageQueryService_ListDeadMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeadMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageQueryServiceServer).ListDeadMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageQueryServiceServer).ListDeadMessages(ctx, req.(*ListDeadMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeadMessageQueryService_ListDeadMessageEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeadMessageEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeadMessageQueryServiceServer).ListDeadMessageEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.dante.v1.service.DeadMessageQueryService/ListDeadMessageEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeadMessageQueryServiceServer).ListDeadMessageEvents(ctx, req.(*ListDeadMessageEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeadMessageQueryService_ServiceDesc is the grpc.ServiceDesc for DeadMessageQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeadMessageQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.dante.v1.service.DeadMessageQueryService",
	HandlerType: (*DeadMessageQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDeadMessage",
			Handler:    _DeadMessageQueryService_GetDeadMessage_Handler,
		},
		{
			MethodName: "ListDeadMessages",
			Handler:    _DeadMessageQueryService_ListDeadMessages_Handler,
		},
		{
			MethodName: "ListDeadMessageEvents",
			Handler:    _DeadMessageQueryService_ListDeadMessageEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/dante/v1/service/dead_message_query.proto",
}
