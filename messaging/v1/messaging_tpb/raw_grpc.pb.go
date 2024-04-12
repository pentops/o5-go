// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: o5/messaging/v1/topic/raw.proto

package messaging_tpb

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

// RawMessageTopicClient is the client API for RawMessageTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RawMessageTopicClient interface {
	Raw(ctx context.Context, in *RawMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type rawMessageTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewRawMessageTopicClient(cc grpc.ClientConnInterface) RawMessageTopicClient {
	return &rawMessageTopicClient{cc}
}

func (c *rawMessageTopicClient) Raw(ctx context.Context, in *RawMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.messaging.v1.topic.RawMessageTopic/Raw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RawMessageTopicServer is the server API for RawMessageTopic service.
// All implementations must embed UnimplementedRawMessageTopicServer
// for forward compatibility
type RawMessageTopicServer interface {
	Raw(context.Context, *RawMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedRawMessageTopicServer()
}

// UnimplementedRawMessageTopicServer must be embedded to have forward compatible implementations.
type UnimplementedRawMessageTopicServer struct {
}

func (UnimplementedRawMessageTopicServer) Raw(context.Context, *RawMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Raw not implemented")
}
func (UnimplementedRawMessageTopicServer) mustEmbedUnimplementedRawMessageTopicServer() {}

// UnsafeRawMessageTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RawMessageTopicServer will
// result in compilation errors.
type UnsafeRawMessageTopicServer interface {
	mustEmbedUnimplementedRawMessageTopicServer()
}

func RegisterRawMessageTopicServer(s grpc.ServiceRegistrar, srv RawMessageTopicServer) {
	s.RegisterService(&RawMessageTopic_ServiceDesc, srv)
}

func _RawMessageTopic_Raw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RawMessageTopicServer).Raw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.messaging.v1.topic.RawMessageTopic/Raw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RawMessageTopicServer).Raw(ctx, req.(*RawMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// RawMessageTopic_ServiceDesc is the grpc.ServiceDesc for RawMessageTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RawMessageTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.messaging.v1.topic.RawMessageTopic",
	HandlerType: (*RawMessageTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Raw",
			Handler:    _RawMessageTopic_Raw_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/messaging/v1/topic/raw.proto",
}
