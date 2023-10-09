// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: o5/github/v1/webhook.proto

package github_pb

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

// WebhookTopicClient is the client API for WebhookTopic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebhookTopicClient interface {
	Push(ctx context.Context, in *PushMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type webhookTopicClient struct {
	cc grpc.ClientConnInterface
}

func NewWebhookTopicClient(cc grpc.ClientConnInterface) WebhookTopicClient {
	return &webhookTopicClient{cc}
}

func (c *webhookTopicClient) Push(ctx context.Context, in *PushMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/o5.github.v1.WebhookTopic/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebhookTopicServer is the server API for WebhookTopic service.
// All implementations must embed UnimplementedWebhookTopicServer
// for forward compatibility
type WebhookTopicServer interface {
	Push(context.Context, *PushMessage) (*emptypb.Empty, error)
	mustEmbedUnimplementedWebhookTopicServer()
}

// UnimplementedWebhookTopicServer must be embedded to have forward compatible implementations.
type UnimplementedWebhookTopicServer struct {
}

func (UnimplementedWebhookTopicServer) Push(context.Context, *PushMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedWebhookTopicServer) mustEmbedUnimplementedWebhookTopicServer() {}

// UnsafeWebhookTopicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebhookTopicServer will
// result in compilation errors.
type UnsafeWebhookTopicServer interface {
	mustEmbedUnimplementedWebhookTopicServer()
}

func RegisterWebhookTopicServer(s grpc.ServiceRegistrar, srv WebhookTopicServer) {
	s.RegisterService(&WebhookTopic_ServiceDesc, srv)
}

func _WebhookTopic_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebhookTopicServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/o5.github.v1.WebhookTopic/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebhookTopicServer).Push(ctx, req.(*PushMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// WebhookTopic_ServiceDesc is the grpc.ServiceDesc for WebhookTopic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebhookTopic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "o5.github.v1.WebhookTopic",
	HandlerType: (*WebhookTopicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _WebhookTopic_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "o5/github/v1/webhook.proto",
}
