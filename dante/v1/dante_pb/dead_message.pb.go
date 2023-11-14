// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: o5/dante/v1/dead_message.proto

package dante_pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Urgency int32

const (
	Urgency_URGENCY_UNSPECIFIED Urgency = 0
	Urgency_URGENCY_LOW         Urgency = 1
	Urgency_URGENCY_MEDIUM      Urgency = 2
	Urgency_URGENCY_HIGH        Urgency = 3
)

// Enum value maps for Urgency.
var (
	Urgency_name = map[int32]string{
		0: "URGENCY_UNSPECIFIED",
		1: "URGENCY_LOW",
		2: "URGENCY_MEDIUM",
		3: "URGENCY_HIGH",
	}
	Urgency_value = map[string]int32{
		"URGENCY_UNSPECIFIED": 0,
		"URGENCY_LOW":         1,
		"URGENCY_MEDIUM":      2,
		"URGENCY_HIGH":        3,
	}
)

func (x Urgency) Enum() *Urgency {
	p := new(Urgency)
	*p = x
	return p
}

func (x Urgency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Urgency) Descriptor() protoreflect.EnumDescriptor {
	return file_o5_dante_v1_dead_message_proto_enumTypes[0].Descriptor()
}

func (Urgency) Type() protoreflect.EnumType {
	return &file_o5_dante_v1_dead_message_proto_enumTypes[0]
}

func (x Urgency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Urgency.Descriptor instead.
func (Urgency) EnumDescriptor() ([]byte, []int) {
	return file_o5_dante_v1_dead_message_proto_rawDescGZIP(), []int{0}
}

type DeadMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID used by the underlying messaging infra, can be any format
	InfraMessageId       string                 `protobuf:"bytes,1,opt,name=infra_message_id,json=infraMessageId,proto3" json:"infra_message_id,omitempty"`
	MessageId            string                 `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	QueueName            string                 `protobuf:"bytes,3,opt,name=queue_name,json=queueName,proto3" json:"queue_name,omitempty"`
	GrpcName             string                 `protobuf:"bytes,4,opt,name=grpc_name,json=grpcName,proto3" json:"grpc_name,omitempty"`
	RejectedTimestamp    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=rejected_timestamp,json=rejectedTimestamp,proto3" json:"rejected_timestamp,omitempty"`
	InitialSentTimestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=initial_sent_timestamp,json=initialSentTimestamp,proto3" json:"initial_sent_timestamp,omitempty"`
	Payload              *Any                   `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	// Types that are assignable to Problem:
	//
	//	*DeadMessage_InvariantViolation
	//	*DeadMessage_UnhandledError
	Problem isDeadMessage_Problem `protobuf_oneof:"problem"`
}

func (x *DeadMessage) Reset() {
	*x = DeadMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dante_v1_dead_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadMessage) ProtoMessage() {}

func (x *DeadMessage) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dante_v1_dead_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadMessage.ProtoReflect.Descriptor instead.
func (*DeadMessage) Descriptor() ([]byte, []int) {
	return file_o5_dante_v1_dead_message_proto_rawDescGZIP(), []int{0}
}

func (x *DeadMessage) GetInfraMessageId() string {
	if x != nil {
		return x.InfraMessageId
	}
	return ""
}

func (x *DeadMessage) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *DeadMessage) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *DeadMessage) GetGrpcName() string {
	if x != nil {
		return x.GrpcName
	}
	return ""
}

func (x *DeadMessage) GetRejectedTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.RejectedTimestamp
	}
	return nil
}

func (x *DeadMessage) GetInitialSentTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.InitialSentTimestamp
	}
	return nil
}

func (x *DeadMessage) GetPayload() *Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (m *DeadMessage) GetProblem() isDeadMessage_Problem {
	if m != nil {
		return m.Problem
	}
	return nil
}

func (x *DeadMessage) GetInvariantViolation() *InvariantViolation {
	if x, ok := x.GetProblem().(*DeadMessage_InvariantViolation); ok {
		return x.InvariantViolation
	}
	return nil
}

func (x *DeadMessage) GetUnhandledError() *UnhandledError {
	if x, ok := x.GetProblem().(*DeadMessage_UnhandledError); ok {
		return x.UnhandledError
	}
	return nil
}

type isDeadMessage_Problem interface {
	isDeadMessage_Problem()
}

type DeadMessage_InvariantViolation struct {
	InvariantViolation *InvariantViolation `protobuf:"bytes,10,opt,name=invariant_violation,json=invariantViolation,proto3,oneof"`
}

type DeadMessage_UnhandledError struct {
	UnhandledError *UnhandledError `protobuf:"bytes,11,opt,name=unhandled_error,json=unhandledError,proto3,oneof"`
}

func (*DeadMessage_InvariantViolation) isDeadMessage_Problem() {}

func (*DeadMessage_UnhandledError) isDeadMessage_Problem() {}

type InvariantViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string  `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Error       *Any    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Urgency     Urgency `protobuf:"varint,3,opt,name=urgency,proto3,enum=o5.dante.v1.Urgency" json:"urgency,omitempty"`
}

func (x *InvariantViolation) Reset() {
	*x = InvariantViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dante_v1_dead_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvariantViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvariantViolation) ProtoMessage() {}

func (x *InvariantViolation) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dante_v1_dead_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvariantViolation.ProtoReflect.Descriptor instead.
func (*InvariantViolation) Descriptor() ([]byte, []int) {
	return file_o5_dante_v1_dead_message_proto_rawDescGZIP(), []int{1}
}

func (x *InvariantViolation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *InvariantViolation) GetError() *Any {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *InvariantViolation) GetUrgency() Urgency {
	if x != nil {
		return x.Urgency
	}
	return Urgency_URGENCY_UNSPECIFIED
}

type UnhandledError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *UnhandledError) Reset() {
	*x = UnhandledError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dante_v1_dead_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnhandledError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnhandledError) ProtoMessage() {}

func (x *UnhandledError) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dante_v1_dead_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnhandledError.ProtoReflect.Descriptor instead.
func (*UnhandledError) Descriptor() ([]byte, []int) {
	return file_o5_dante_v1_dead_message_proto_rawDescGZIP(), []int{2}
}

func (x *UnhandledError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Any wraps the well-known any, but encodes messages as JSON in addition
type Any struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto *anypb.Any `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Json  string     `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *Any) Reset() {
	*x = Any{}
	if protoimpl.UnsafeEnabled {
		mi := &file_o5_dante_v1_dead_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Any) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Any) ProtoMessage() {}

func (x *Any) ProtoReflect() protoreflect.Message {
	mi := &file_o5_dante_v1_dead_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Any.ProtoReflect.Descriptor instead.
func (*Any) Descriptor() ([]byte, []int) {
	return file_o5_dante_v1_dead_message_proto_rawDescGZIP(), []int{3}
}

func (x *Any) GetProto() *anypb.Any {
	if x != nil {
		return x.Proto
	}
	return nil
}

func (x *Any) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

var File_o5_dante_v1_dead_message_proto protoreflect.FileDescriptor

var file_o5_dante_v1_dead_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x35, 0x2f, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x61, 0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6f, 0x35, 0x2e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x04, 0x0a, 0x0b, 0x44, 0x65, 0x61, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x70,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x12, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x72,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x50, 0x0a, 0x16, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x14, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x52,
	0x0a, 0x13, 0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x76, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12,
	0x69, 0x6e, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0f, 0x75, 0x6e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x64, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x35,
	0x2e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x75, 0x6e, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6f, 0x35, 0x2e, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x75, 0x72, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6f, 0x35, 0x2e, 0x64, 0x61, 0x6e,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x07, 0x75,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x26, 0x0a, 0x0e, 0x55, 0x6e, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x45,
	0x0a, 0x03, 0x41, 0x6e, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x2a, 0x59, 0x0a, 0x07, 0x55, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x17, 0x0a, 0x13, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x52, 0x47,
	0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x52,
	0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x03,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6f, 0x35, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x6e,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x6e, 0x74, 0x65, 0x5f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_o5_dante_v1_dead_message_proto_rawDescOnce sync.Once
	file_o5_dante_v1_dead_message_proto_rawDescData = file_o5_dante_v1_dead_message_proto_rawDesc
)

func file_o5_dante_v1_dead_message_proto_rawDescGZIP() []byte {
	file_o5_dante_v1_dead_message_proto_rawDescOnce.Do(func() {
		file_o5_dante_v1_dead_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_o5_dante_v1_dead_message_proto_rawDescData)
	})
	return file_o5_dante_v1_dead_message_proto_rawDescData
}

var file_o5_dante_v1_dead_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_o5_dante_v1_dead_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_o5_dante_v1_dead_message_proto_goTypes = []interface{}{
	(Urgency)(0),                  // 0: o5.dante.v1.Urgency
	(*DeadMessage)(nil),           // 1: o5.dante.v1.DeadMessage
	(*InvariantViolation)(nil),    // 2: o5.dante.v1.InvariantViolation
	(*UnhandledError)(nil),        // 3: o5.dante.v1.UnhandledError
	(*Any)(nil),                   // 4: o5.dante.v1.Any
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
}
var file_o5_dante_v1_dead_message_proto_depIdxs = []int32{
	5, // 0: o5.dante.v1.DeadMessage.rejected_timestamp:type_name -> google.protobuf.Timestamp
	5, // 1: o5.dante.v1.DeadMessage.initial_sent_timestamp:type_name -> google.protobuf.Timestamp
	4, // 2: o5.dante.v1.DeadMessage.payload:type_name -> o5.dante.v1.Any
	2, // 3: o5.dante.v1.DeadMessage.invariant_violation:type_name -> o5.dante.v1.InvariantViolation
	3, // 4: o5.dante.v1.DeadMessage.unhandled_error:type_name -> o5.dante.v1.UnhandledError
	4, // 5: o5.dante.v1.InvariantViolation.error:type_name -> o5.dante.v1.Any
	0, // 6: o5.dante.v1.InvariantViolation.urgency:type_name -> o5.dante.v1.Urgency
	6, // 7: o5.dante.v1.Any.proto:type_name -> google.protobuf.Any
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_o5_dante_v1_dead_message_proto_init() }
func file_o5_dante_v1_dead_message_proto_init() {
	if File_o5_dante_v1_dead_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_o5_dante_v1_dead_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadMessage); i {
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
		file_o5_dante_v1_dead_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvariantViolation); i {
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
		file_o5_dante_v1_dead_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnhandledError); i {
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
		file_o5_dante_v1_dead_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Any); i {
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
	file_o5_dante_v1_dead_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DeadMessage_InvariantViolation)(nil),
		(*DeadMessage_UnhandledError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_o5_dante_v1_dead_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_o5_dante_v1_dead_message_proto_goTypes,
		DependencyIndexes: file_o5_dante_v1_dead_message_proto_depIdxs,
		EnumInfos:         file_o5_dante_v1_dead_message_proto_enumTypes,
		MessageInfos:      file_o5_dante_v1_dead_message_proto_msgTypes,
	}.Build()
	File_o5_dante_v1_dead_message_proto = out.File
	file_o5_dante_v1_dead_message_proto_rawDesc = nil
	file_o5_dante_v1_dead_message_proto_goTypes = nil
	file_o5_dante_v1_dead_message_proto_depIdxs = nil
}