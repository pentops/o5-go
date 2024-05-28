// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/messaging/v1/message.proto

package messaging_pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message wraps messaging content for the underlying platform routing and
// delivery.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The 'full_name' of the gRPC method, /<package>.<service>/<method>
	GrpcMethod string     `protobuf:"bytes,2,opt,name=grpc_method,json=grpcMethod,proto3" json:"grpc_method,omitempty"`
	Body       *anypb.Any `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Then o5.application.name which sent the message
	SourceApp string `protobuf:"bytes,4,opt,name=source_app,json=sourceApp,proto3" json:"source_app,omitempty"`
	// The o5.environment.full_name which sent the message
	SourceEnv string `protobuf:"bytes,5,opt,name=source_env,json=sourceEnv,proto3" json:"source_env,omitempty"`
	// compatibility with named topic messaging, specifies the destination name
	DestinationTopic string `protobuf:"bytes,6,opt,name=destination_topic,json=destinationTopic,proto3" json:"destination_topic,omitempty"`
	// In the reply to a request/reply message, Identifies the service or instance
	// which first sent the request, and then to which the reply should be routed.
	ReplyDest *string           `protobuf:"bytes,7,opt,name=reply_dest,json=replyDest,proto3,oneof" json:"reply_dest,omitempty"`
	Headers   map[string]string `protobuf:"bytes,8,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_messaging_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_o5_messaging_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_o5_messaging_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetGrpcMethod() string {
	if x != nil {
		return x.GrpcMethod
	}
	return ""
}

func (x *Message) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Message) GetSourceApp() string {
	if x != nil {
		return x.SourceApp
	}
	return ""
}

func (x *Message) GetSourceEnv() string {
	if x != nil {
		return x.SourceEnv
	}
	return ""
}

func (x *Message) GetDestinationTopic() string {
	if x != nil {
		return x.DestinationTopic
	}
	return ""
}

func (x *Message) GetReplyDest() string {
	if x != nil && x.ReplyDest != nil {
		return *x.ReplyDest
	}
	return ""
}

func (x *Message) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_o5_messaging_v1_message_proto protoreflect.FileDescriptor

var file_o5_messaging_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x35, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x6f, 0x35, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0,
	0x01, 0x01, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x60, 0x0a,
	0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x3f, 0xba, 0x48, 0x3c, 0x72, 0x3a, 0x32, 0x38, 0x5e, 0x5c, 0x2f, 0x28, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2b, 0x5c, 0x2e, 0x29, 0x2b, 0x28, 0x5b, 0x41,
	0x2d, 0x5a, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2b,
	0x29, 0x5c, 0x2f, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5f, 0x5d,
	0x2b, 0x29, 0x24, 0x52, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x28, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x64, 0x65,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x44, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x35, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x64, 0x65, 0x73, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_o5_messaging_v1_message_proto_rawDescOnce sync.Once
	file_o5_messaging_v1_message_proto_rawDescData = file_o5_messaging_v1_message_proto_rawDesc
)

func file_o5_messaging_v1_message_proto_rawDescGZIP() []byte {
	file_o5_messaging_v1_message_proto_rawDescOnce.Do(func() {
		file_o5_messaging_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_messaging_v1_message_proto_rawDescData)
	})
	return file_o5_messaging_v1_message_proto_rawDescData
}

var file_o5_messaging_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_o5_messaging_v1_message_proto_goTypes = []interface{}{
	(*Message)(nil),   // 0: o5.messaging.v1.Message
	nil,               // 1: o5.messaging.v1.Message.HeadersEntry
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_o5_messaging_v1_message_proto_depIdxs = []int32{
	2, // 0: o5.messaging.v1.Message.body:type_name -> google.protobuf.Any
	1, // 1: o5.messaging.v1.Message.headers:type_name -> o5.messaging.v1.Message.HeadersEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_o5_messaging_v1_message_proto_init() }
func file_o5_messaging_v1_message_proto_init() {
	if File_o5_messaging_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_o5_messaging_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_o5_messaging_v1_message_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_messaging_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_o5_messaging_v1_message_proto_goTypes,
		DependencyIndexes: file_o5_messaging_v1_message_proto_depIdxs,
		MessageInfos:      file_o5_messaging_v1_message_proto_msgTypes,
	}.Build()
	File_o5_messaging_v1_message_proto = out.File
	file_o5_messaging_v1_message_proto_rawDesc = nil
	file_o5_messaging_v1_message_proto_goTypes = nil
	file_o5_messaging_v1_message_proto_depIdxs = nil
}
